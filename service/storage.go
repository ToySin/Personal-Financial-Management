package service

import (
	"time"

	"github.com/ToySin/finance/portfolio"
)

type PortfolioStorage interface {
	SavePortfolio(p *portfolio.Portfolio) error

	// GetPortfolio returns the portfolio of the given date.
	// - The date should include the year and month. Other fields are ignored.
	// - If the portfolio does not exist, it returns an error.
	GetPortfolio(date time.Time) (*portfolio.Portfolio, error)
}
