package service

import (
	"time"

	"github.com/ToySin/finance/portfolio"
)

// Service is a service for managing finance data.
type Service struct {
	storage PortfolioStorage
}

// New creates a new Service.
func New() *Service {
	return &Service{}
}

// CreatePortfolio creates a new portfolio for the given date. (year and month)
// When creating a new portfolio, the last month's fixed expenses are copied to the new portfolio in default.
// TODO(#4): Copy from genral fixed expenses setting instead of the last month's fixed expenses.
func (s *Service) CreatePortfolio(date time.Time) error {
	// If there is no portfolio for the last month,
	// create a new portfolio without copying the last month's fixed expenses.
	var fixedExpense []portfolio.Transaction
	lastPortfolio, err := s.storage.GetPortfolio(date.AddDate(0, -1, 0))
	if err == nil {
		fixedExpense = lastPortfolio.Transactions[portfolio.FixedExpenseCategory]
	}

	p := portfolio.NewPortfolio(date)
	p.Transactions[portfolio.FixedExpenseCategory] = fixedExpense
	return s.storage.SavePortfolio(p)
}

// CreateTransaction creates a transaction to the portfolio of the given date.
func (s *Service) CreateTransaction(date time.Time, c portfolio.Category, name string, amount portfolio.Amount, note string) error {
	p, err := s.storage.GetPortfolio(date)
	if err != nil {
		return err
	}

	transaction := &portfolio.Transaction{
		Date:   date,
		Name:   name,
		Amount: amount,
		Note:   note,
	}
	p.AddTransaction(c, transaction)
	return s.storage.SavePortfolio(p)
}
