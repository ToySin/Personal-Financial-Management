package storage

import (
	"database/sql"
	"fmt"
	"time"

	env "github.com/caarlos0/env/v11"
	"gorm.io/driver/mariadb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/ToySin/finance/portfolio"
	"github.com/ToySin/finance/utils"
)

const (
	DBTypeMySQL   = "mysql"
	DBTypeMariaDB = "mariadb"
)

// Config represents the configuration for the database connection
type Config struct {
	DBType        string `envconfig:"DB_TYPE" envDefault:"mysql"`
	FinanceDBHost string `envconfig:"FINANCE_DB_HOST" envDefault:"localhost"`
	FinanceDBPort string `envconfig:"FINANCE_DB_PORT" envDefault:"5432"`
	FinanceDBUser string `envconfig:"FINANCE_DB_USER" envDefault:""`
	FinanceDBPass string `envconfig:"FINANCE_DB_PASS" envDefault:""`
	FinanceDBName string `envconfig:"FINANCE_DB_NAME" envDefault:"finance"`
}

// CreateConfigFromEnv creates a new Config from the environment variables.
func CreateConfigFromEnv() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// CreateDB creates a new database connection.
func (c *Config) createDB() (*sql.DB, error) {
	var dsn string
	switch c.DBType {
	case DBTypeMySQL, DBTypeMariaDB:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
			c.FinanceDBUser,
			c.FinanceDBPass,
			c.FinanceDBHost,
			c.FinanceDBPort,
			c.FinanceDBName)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", c.DBType)
	}
	return sql.Open(c.DBType, dsn)
}

// SQLClient is a database client for the finance service.
type SQLClient struct {
	db *gorm.DB
}

// NewSQLClient creates a new SQLClient.
func NewSQLClient(cfg *Config) (*SQLClient, error) {
	var err error
	db, err := cfg.createDB()
	if err != nil {
		return nil, err
	}

	var gdb *gorm.DB
	switch cfg.DBType {
	case DBTypeMySQL:
		gdb, err = gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	case DBTypeMariaDB:
		gdb, err = gorm.Open(mariadb.New(mariadb.Config{Conn: db}), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.DBType)
	}
	if err != nil {
		return nil, err
	}

	gdb.AutoMigrate(&TransactionTable{})
	return &SQLClient{
		db: gdb,
	}, nil
}

// SaveTransaction upserts the transaction to the database.
func (c *SQLClient) SaveTransaction(t *portfolio.Transaction) error {
	transactionTable := TransactionTable{
		UUID:     t.UUID,
		Date:     t.Date,
		Name:     t.Name,
		Category: t.Category.ToCode(),
		Amount:   int(t.Amount),
		Note:     t.Note,
	}
	return c.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&transactionTable).Error
}

// SavePortfolio upserts the transactions of the portfolio to the database.
func (c *SQLClient) SavePortfolio(p *portfolio.Portfolio) error {
	for category, transactions := range p.Transactions {
		for _, t := range transactions {
			transactionTable := TransactionTable{
				UUID:     t.UUID,
				Date:     t.Date,
				Name:     t.Name,
				Category: category.ToCode(),
				Amount:   int(t.Amount),
				Note:     t.Note,
			}

			if err := c.db.Save(&transactionTable).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// GetPortfolio retrieves the portfolio of the given month from the database.
func (c *SQLClient) GetPortfolio(month time.Time) (*portfolio.Portfolio, error) {
	firstDayOfMonth := utils.GetFirstDayOfMonth(month.Year(), month.Month())
	nextMonth := firstDayOfMonth.AddDate(0, 1, 0)
	var transactionTables []TransactionTable
	if err := c.db.Where("date >= ? AND date < ?", firstDayOfMonth, nextMonth).Find(&transactionTables).Error; err != nil {
		return nil, err
	}

	p := portfolio.NewPortfolio(month)
	for _, t := range transactionTables {
		transaction := &portfolio.Transaction{
			UUID:     t.UUID,
			Date:     t.Date,
			Name:     t.Name,
			Category: portfolio.Category.FromCode(portfolio.UnknownCategory, t.Category),
			Amount:   portfolio.Amount(t.Amount),
			Note:     t.Note,
		}
		p.AddTransaction(transaction.Category, transaction)
	}
	return p, nil
}
