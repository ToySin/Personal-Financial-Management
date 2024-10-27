package portfolio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	testPortfolio := GetTestPortfolio()

	expectedBalance := testPortfolio.TotalAmount(IncomeCategory) -
		testPortfolio.TotalAmount(FixedExpenseCategory) -
		testPortfolio.TotalAmount(SavingCategory) -
		testPortfolio.TotalAmount(InvestmentCategory)
	assert.Equal(t, expectedBalance, testPortfolio.GetBalance())

	t.Logf("testPortfolio.Balance: %v", testPortfolio.GetBalance())
}
