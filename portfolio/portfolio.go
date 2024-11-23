package portfolio

import (
	"sync"
	"time"
)

// Ammount is a type for the amount of money.
type Amount int64

// Category is a type for the category of the transaction.
type Category string

const (
	UnknownCategory          Category = ""
	IncomeCategory           Category = "수입"
	FixedExpenseCategory     Category = "고정지출"
	VariableExpenseCategory  Category = "변동지출"
	SavingCategory           Category = "저축"
	InvestmentCategory       Category = "투자"
	InvestmentIncomeCategory Category = "투자수익"
)

// Transaction is a financial transaction.
// It represents a single income or expense item.
type Transaction struct {
	Date     time.Time // date of the transaction
	Name     string    // name of the transaction
	Category Category  // category of the transaction
	Amount   Amount    // amount of the transaction
	Note     string    // note of the transaction
}

// Portfolio is a collection of monthly financial transactions.
// It should be created with NewPortfolio().
type Portfolio struct {
	Month        time.Time                  // a year and month of the portfolio
	Transactions map[Category][]Transaction // a list of transactions
	Balance      Amount                     // total balance of the month
}

// NewPortfolio creates a new portfolio for the given year-month.
// The day, hour, minute, second, and nsecond of the time is ignored.
func NewPortfolio(yearAndMonth time.Time) *Portfolio {
	// Set the day to the first day of the month
	yearAndMonth = yearAndMonth.AddDate(0, 0, -yearAndMonth.Day()+1)
	return &Portfolio{
		Month:        yearAndMonth,
		Transactions: map[Category][]Transaction{},
		Balance:      0,
	}
}

// TotalAmount returns the total amount of the given category.
func (p *Portfolio) TotalAmount(c Category) Amount {
	var total Amount
	for _, b := range p.Transactions[c] {
		total += b.Amount
	}
	return total
}

// AddTransaction adds a transaction to the portfolio.
//
// The category of the transaction are set automatically based on the given category.
// The day of the transaction is only required.
// The year and month of the transaction are set to the year and month of the portfolio.
func (p *Portfolio) AddTransaction(c Category, transaction *Transaction) {
	if p.Transactions[c] == nil {
		p.Transactions[c] = make([]Transaction, 0)
	}
	categoryList := p.Transactions[c]

	transaction.Date = time.Date(
		p.Month.Year(),
		p.Month.Month(),
		transaction.Date.Day(),
		0, 0, 0, 0, time.Local)
	transaction.Category = c
	categoryList = append(categoryList, *transaction)
	p.Transactions[c] = categoryList
}

func (p *Portfolio) updateBalance() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	p.Balance = 0
	for c, categoryList := range p.Transactions {
		wg.Add(1)
		go func(c Category, categoryList []Transaction) {
			defer wg.Done()
			var totalAmount Amount
			for _, t := range categoryList {
				totalAmount += t.Amount
			}
			mu.Lock()
			switch c {
			// Incoming categories
			case IncomeCategory,
				InvestmentIncomeCategory:

				p.Balance += totalAmount

			// Outgoing categories
			case FixedExpenseCategory,
				VariableExpenseCategory,
				SavingCategory,
				InvestmentCategory:

				p.Balance -= totalAmount
			}
			mu.Unlock()

		}(c, categoryList)
	}
	wg.Wait()
}

// GetBalance returns the total balance of the portfolio.
func (p *Portfolio) GetBalance() Amount {
	p.updateBalance()
	return p.Balance
}
