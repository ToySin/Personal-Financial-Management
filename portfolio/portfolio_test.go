package portfolio

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
