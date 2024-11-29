package storage

import (
	"strings"
	"testing"
	"time"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/ToySin/finance/portfolio"
)

const (
	dbUser = "toy"
	dbPass = "sin"
	dbName = "finance"
)

func GetTestTransaction1() *portfolio.Transaction {
	return &portfolio.Transaction{
		UUID:     "test1-uuid",
		Date:     time.Date(2024, 10, 2, 14, 12, 0, 0, time.Local),
		Name:     "test1",
		Category: portfolio.FixedExpenseCategory,
		Amount:   1000,
		Note:     "test1",
	}
}

func GetTestTransaction2() *portfolio.Transaction {
	return &portfolio.Transaction{
		UUID:     "test2-uuid",
		Date:     time.Date(2024, 10, 3, 7, 55, 12, 0, time.Local),
		Name:     "test2",
		Category: portfolio.VariableExpenseCategory,
		Amount:   2000,
		Note:     "test2",
	}
}

func GetTestPortfolio() *portfolio.Portfolio {
	return &portfolio.Portfolio{
		Month: time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
		Transactions: map[portfolio.Category][]*portfolio.Transaction{
			portfolio.FixedExpenseCategory: {
				GetTestTransaction1(),
			},
			portfolio.VariableExpenseCategory: {
				GetTestTransaction2(),
			},
		},
	}
}

func TestSavePortfolio(t *testing.T) {
	p := mysql.Preset(
		mysql.WithUser(dbUser, dbPass),
		mysql.WithDatabase(dbName),
	)
	container, err := gnomock.Start(p)
	assert.NoError(t, err)

	defer func() { _ = gnomock.Stop(container) }()

	addr := container.DefaultAddress()
	config := &Config{
		FinanceDBHost: strings.Split(addr, ":")[0],
		FinanceDBPort: strings.Split(addr, ":")[1],
		FinanceDBUser: dbUser,
		FinanceDBPass: dbPass,
		FinanceDBName: dbName,
	}
	db, err := config.CreateDB()
	assert.NoError(t, err)

	client, err := NewSQLClient(db)
	assert.NoError(t, err)

	// Save the portfolio to the database.
	p1 := GetTestPortfolio()
	err = client.SavePortfolio(p1)
	assert.NoError(t, err)

	// Check if the portfolio is saved correctly.
	p2, err := client.GetPortfolio(p1.Month)
	assert.NoError(t, err)
	assert.NotNil(t, p2)
	assert.Equal(t, p1.Month.Year(), p2.Month.Year())
	assert.Equal(t, p1.Month.Month(), p2.Month.Month())

	t1 := GetTestTransaction1()
	p2T1 := p2.Transactions[t1.Category][0]
	assert.Equal(t, t1.UUID, p2T1.UUID)
	assert.Equal(t, t1.Category, p2T1.Category)
	assert.Equal(t, t1.Name, p2T1.Name)
	assert.Equal(t, t1.Amount, p2T1.Amount)
	assert.Equal(t, t1.Note, p2T1.Note)

	t2 := GetTestTransaction2()
	p2T2 := p2.Transactions[t2.Category][0]
	assert.Equal(t, t2.UUID, p2T2.UUID)
	assert.Equal(t, t2.Category, p2T2.Category)
	assert.Equal(t, t2.Name, p2T2.Name)
	assert.Equal(t, t2.Amount, p2T2.Amount)
	assert.Equal(t, t2.Note, p2T2.Note)
}

func TestSaveTransaction(t *testing.T) {
	p := mysql.Preset(
		mysql.WithUser(dbUser, dbPass),
		mysql.WithDatabase(dbName),
	)
	container, err := gnomock.Start(p)
	assert.NoError(t, err)

	defer func() { _ = gnomock.Stop(container) }()

	addr := container.DefaultAddress()
	config := &Config{
		FinanceDBHost: strings.Split(addr, ":")[0],
		FinanceDBPort: strings.Split(addr, ":")[1],
		FinanceDBUser: dbUser,
		FinanceDBPass: dbPass,
		FinanceDBName: dbName,
	}
	db, err := config.CreateDB()
	assert.NoError(t, err)

	client, err := NewSQLClient(db)
	assert.NoError(t, err)

	// Save the transaction to the database.
	t1 := GetTestTransaction1()
	err = client.SaveTransaction(t1)
	assert.NoError(t, err)

	// Save the transaction to the database.
	t2 := GetTestTransaction2()
	err = client.SaveTransaction(t2)
	assert.NoError(t, err)

	// Check if the transaction is saved correctly.
	p1, err := client.GetPortfolio(t1.Date)
	assert.NoError(t, err)
	assert.NotNil(t, p1)
	assert.Equal(t, t1.Date.Year(), p1.Month.Year())
	assert.Equal(t, t1.Date.Month(), p1.Month.Month())
	assert.Equal(t, 1, len(p1.Transactions[t1.Category]))
	assert.Equal(t, t1.UUID, p1.Transactions[t1.Category][0].UUID)
	assert.Equal(t, t1.Category, p1.Transactions[t1.Category][0].Category)
	assert.Equal(t, t1.Name, p1.Transactions[t1.Category][0].Name)
	assert.Equal(t, t1.Amount, p1.Transactions[t1.Category][0].Amount)
	assert.Equal(t, t1.Note, p1.Transactions[t1.Category][0].Note)
	assert.Equal(t, 1, len(p1.Transactions[t2.Category]))
	assert.Equal(t, t2.UUID, p1.Transactions[t2.Category][0].UUID)
	assert.Equal(t, t2.Category, p1.Transactions[t2.Category][0].Category)
	assert.Equal(t, t2.Name, p1.Transactions[t2.Category][0].Name)
	assert.Equal(t, t2.Amount, p1.Transactions[t2.Category][0].Amount)
	assert.Equal(t, t2.Note, p1.Transactions[t2.Category][0].Note)
}
