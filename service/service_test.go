package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ToySin/finance/portfolio"
)

type fakePortfolioStorage struct {
	transactions []*portfolio.Transaction
}

func (s *fakePortfolioStorage) GetPortfolio(date time.Time) (*portfolio.Portfolio, error) {
	p := &portfolio.Portfolio{
		Month:        date,
		Transactions: make(map[portfolio.Category][]*portfolio.Transaction),
	}
	for _, t := range s.transactions {
		if t.Date.Year() == date.Year() && t.Date.Month() == date.Month() {
			p.Transactions[t.Category] = append(p.Transactions[t.Category], t)
		}
	}
	return p, nil
}

func (s *fakePortfolioStorage) SaveTransaction(transaction *portfolio.Transaction) error {
	s.transactions = append(s.transactions, transaction)
	return nil
}

func TestService_CreatePortfolio(t *testing.T) {
	s := New(&fakePortfolioStorage{})

	// Create a new portfolio for the given date.
	date := time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local)
	s.CreateTransaction(date.AddDate(0, 0, 1), portfolio.FixedExpenseCategory, "test", 1000, "test")
	s.CreateTransaction(date.AddDate(0, 0, 2), portfolio.VariableExpenseCategory, "test", 2000, "test")

	// Check if the last month's fixed expenses are copied to the new portfolio.
	p, err := s.GetPortfolio(date)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, 1, len(p.Transactions[portfolio.FixedExpenseCategory]))
	assert.Equal(t, 1, len(p.Transactions[portfolio.VariableExpenseCategory]))
	assert.Equal(t, "test", p.Transactions[portfolio.FixedExpenseCategory][0].Name)
	assert.Equal(t, "test", p.Transactions[portfolio.VariableExpenseCategory][0].Name)
}
