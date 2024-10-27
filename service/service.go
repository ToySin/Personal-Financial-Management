package service

import (
	"time"

	"github.com/ToySin/finance/portfolio"
)

type storage interface {
	SavePortfolio(p *portfolio.Portfolio) error

	// GetPortfolio returns the portfolio of the given date.
	// - The date should include the year and month. Other fields are ignored.
	// - If the portfolio does not exist, it returns an error.
	GetPortfolio(date time.Time) (*portfolio.Portfolio, error)
}

type Service struct {
	storage storage
}

func New() *Service {
	return &Service{}
}

func (s *Service) AddMeal(date time.Time, name string, price int) error {
	p, err := s.storage.GetPortfolio(date)
	if err != nil {
		return err
	}

	newTransaction := &portfolio.Transaction{
		Date:   date,
		Name:   name,
		Amount: portfolio.Amount(price),
	}
	p.AddTransaction(portfolio.VariableExpenseCategory, newTransaction)

	return s.storage.SavePortfolio(p)
}
