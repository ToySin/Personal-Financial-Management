package portfolio

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testTransaction1 = &Transaction{
		Date:     time.Date(2024, 10, 2, 14, 12, 0, 0, time.Local),
		Name:     "test1",
		Category: FixedExpenseCategory,
		Amount:   1000,
	}

	testTransaction2 = &Transaction{
		Date:     time.Date(2024, 10, 3, 7, 55, 12, 0, time.Local),
		Name:     "test2",
		Category: VariableExpenseCategory,
		Amount:   2000,
	}
)

func TestGetBalance(t *testing.T) {
	// Given
	testPortfolio := GetTestPortfolio()

	// When
	expectedBalance := testPortfolio.TotalAmount(IncomeCategory) -
		testPortfolio.TotalAmount(FixedExpenseCategory) -
		testPortfolio.TotalAmount(SavingCategory) -
		testPortfolio.TotalAmount(InvestmentCategory)

	// Then
	assert.Equal(t, expectedBalance, testPortfolio.GetBalance())

	t.Logf("testPortfolio.Balance: %v", testPortfolio.GetBalance())
}

func TestTransactionDate(t *testing.T) {
	// Given
	now := time.Now()
	p := NewPortfolio(now)

	// When
	p.AddTransaction(IncomeCategory, &Transaction{
		Date:   time.Date(1999, 12, 3, 0, 0, 0, 0, time.Local),
		Name:   "Test",
		Amount: 1000,
	})

	// Then
	transaction := p.Transactions[IncomeCategory][0]
	assert.Equal(t, now.Year(), transaction.Date.Year())
	assert.Equal(t, now.Month(), transaction.Date.Month())
	assert.Equal(t, 3, transaction.Date.Day())
}

func TestAddTransaction(t *testing.T) {
	// Given
	now := time.Now()
	p := NewPortfolio(now)

	// When
	p.AddTransaction(FixedExpenseCategory, testTransaction1)
	p.AddTransaction(VariableExpenseCategory, testTransaction2)

	// Then
	assert.Equal(t, 1, len(p.Transactions[FixedExpenseCategory]))
	assert.Equal(t, 1000, int(p.Transactions[FixedExpenseCategory][0].Amount))
	assert.Equal(t, 1, len(p.Transactions[VariableExpenseCategory]))
	assert.Equal(t, 2000, int(p.Transactions[VariableExpenseCategory][0].Amount))
	assert.Equal(t, 0, len(p.Transactions[IncomeCategory]))
	assert.Equal(t, 0, len(p.Transactions[SavingCategory]))
}
