package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ToySin/finance/portfolio"
)

type fakePortfolioStorage struct {
	portfolio *portfolio.Portfolio
}

func (s *fakePortfolioStorage) SavePortfolio(p *portfolio.Portfolio) error {
	s.portfolio = p
	return nil
}

func (s *fakePortfolioStorage) GetPortfolio(date time.Time) (*portfolio.Portfolio, error) {
	if s.portfolio == nil {
		return nil, portfolio.ErrPortfolioNotFound
	}
	return s.portfolio, nil
}

func TestService_CreatePortfolio(t *testing.T) {
	s := New()
	s.storage = &fakePortfolioStorage{}

	// Create a new portfolio for the given date.
	date := time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local)
	err := s.CreatePortfolio(date)
	assert.NoError(t, err)

	// Check if the portfolio is created correctly.
	p, err := s.storage.GetPortfolio(date)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, date.Year(), p.Month.Year())
	assert.Equal(t, date.Month(), p.Month.Month())

	s.CreateTransaction(date.AddDate(0, 0, 1), portfolio.FixedExpenseCategory, "test", 1000, "test")
	s.CreateTransaction(date.AddDate(0, 0, 1), portfolio.VariableExpenseCategory, "test", 2000, "test")

	nextMonth := date.AddDate(0, 1, 0)
	err = s.CreatePortfolio(nextMonth)
	assert.NoError(t, err)

	// Check if the last month's fixed expenses are copied to the new portfolio.
	p, err = s.storage.GetPortfolio(nextMonth)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, 1, len(p.Transactions[portfolio.FixedExpenseCategory]))
	assert.Equal(t, 0, len(p.Transactions[portfolio.VariableExpenseCategory]))
}
