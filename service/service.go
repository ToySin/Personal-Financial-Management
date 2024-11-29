package service

import (
	"time"

	"github.com/ToySin/finance/portfolio"
)

type PortfolioStorage interface {
	GetPortfolio(date time.Time) (*portfolio.Portfolio, error)
	SaveTransaction(transaction *portfolio.Transaction) error
}

// Service is a service for managing finance data.
type Service struct {
	dbClient PortfolioStorage
}

// New creates a new Service.
func New(dbClient PortfolioStorage) *Service {
	return &Service{dbClient: dbClient}
}

// GetPortfolio returns the portfolio of the given date.
func (s *Service) GetPortfolio(date time.Time) (*portfolio.Portfolio, error) {
	return s.dbClient.GetPortfolio(date)
}

// CreateTransaction creates a transaction to the portfolio of the given date.
func (s *Service) CreateTransaction(date time.Time, c portfolio.Category, name string, amount portfolio.Amount, note string) error {
	transaction := &portfolio.Transaction{
		Date:     date,
		Name:     name,
		Category: c,
		Amount:   amount,
		Note:     note,
	}
	return s.dbClient.SaveTransaction(transaction)
}
