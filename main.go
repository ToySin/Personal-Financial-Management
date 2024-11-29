package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	env "github.com/caarlos0/env/v11"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/mysql"

	"github.com/ToySin/finance/api/handler"
	"github.com/ToySin/finance/service"
	"github.com/ToySin/finance/service/storage"
)

type config struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	PlatformEnv string `env:"PLATFORM_ENV" envDefault:"local"`
}

func main() {
	// Parse the environment variables.
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		slog.Error("failed to parse the environment variables", "reason", err)
		return
	}

	// Create a new service with a storage.
	storageCfg, err := storage.CreateConfigFromEnv()
	if err != nil {
		slog.Error("failed to create a storage config", "reason", err)
		return
	}

	if cfg.PlatformEnv == "local" {
		p := mysql.Preset(
			mysql.WithUser("toy", "sin"),
			mysql.WithDatabase("finance"),
		)
		container, err := gnomock.Start(p)
		if err != nil {
			slog.Error("failed to start a gnomock container", "reason", err)
			return
		}
		defer func() { _ = gnomock.Stop(container) }()

		addr := container.DefaultAddress()
		storageCfg.FinanceDBHost = strings.Split(addr, ":")[0]
		storageCfg.FinanceDBPort = strings.Split(addr, ":")[1]
		storageCfg.FinanceDBUser = "toy"
		storageCfg.FinanceDBPass = "sin"
	}

	db, err := storageCfg.CreateDB()
	if err != nil {
		slog.Error("failed to create a database connection", "reason", err)
		return
	}
	client, err := storage.NewSQLClient(db)
	if err != nil {
		slog.Error("failed to create a storage client", "reason", err)
		return
	}
	s := service.New(client)

	// Create a new API handler with the service.
	h := handler.NewAPIHandler(s)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Start the http server.
	http.HandleFunc("/api/portfolio", h.GetPortfolioHandler)
	http.HandleFunc("/api/transaction", h.CreateTransactionHandler)

	slog.Info("starting the server", "port", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil); err != nil {
		slog.Error("failed to start the server", "reason", err)
	}
}
