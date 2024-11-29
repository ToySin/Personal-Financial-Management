package portfolio

import (
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/ToySin/finance/utils"
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

// ToCode returns the code of the category.
func (c Category) ToCode() string {
	switch c {
	case IncomeCategory:
		return "CATEGORY_INCOME"
	case FixedExpenseCategory:
		return "CATEGORY_FIXED_EXPENSE"
	case VariableExpenseCategory:
		return "CATEGORY_VARIABLE_EXPENSE"
	case SavingCategory:
		return "CATEGORY_SAVING"
	case InvestmentCategory:
		return "CATEGORY_INVESTMENT"
	case InvestmentIncomeCategory:
		return "CATEGORY_INVESTMENT_INCOME"
	default:
		return "CATEGORY_UNKNOWN"
	}
}

// FromCode returns the category from the code.
func (c Category) FromCode(code string) Category {
	switch code {
	case "CATEGORY_INCOME":
		return IncomeCategory
	case "CATEGORY_FIXED_EXPENSE":
		return FixedExpenseCategory
	case "CATEGORY_VARIABLE_EXPENSE":
		return VariableExpenseCategory
	case "CATEGORY_SAVING":
		return SavingCategory
	case "CATEGORY_INVESTMENT":
		return InvestmentCategory
	case "CATEGORY_INVESTMENT_INCOME":
		return InvestmentIncomeCategory
	default:
		return UnknownCategory
	}
}

// Transaction is a financial transaction.
// It represents a single income or expense item.
type Transaction struct {
	UUID     string    // unique identifier of the transaction
	Date     time.Time // date of the transaction
	Name     string    // name of the transaction
	Category Category  // category of the transaction
	Amount   Amount    // amount of the transaction
	Note     string    // note of the transaction
}

// NewTransaction creates a new transaction with the given values.
// Since the UUID is generated automatically, Transaction instances should be created with this function.
func NewTransaction(date time.Time, name string, category Category, amount Amount, note string) *Transaction {
	return &Transaction{
		UUID:     uuid.NewString(),
		Date:     date,
		Name:     name,
		Category: category,
		Amount:   amount,
		Note:     note,
	}
}

// Portfolio is a collection of monthly financial transactions.
// It should be created with NewPortfolio().
type Portfolio struct {
	Month        time.Time                   // a year and month of the portfolio
	Transactions map[Category][]*Transaction // a list of transactions
	Balance      Amount                      // total balance of the month
}

// NewPortfolio creates a new portfolio for the given year-month.
// The day, hour, minute, second, and nsecond of the time is ignored.
func NewPortfolio(month time.Time) *Portfolio {
	return &Portfolio{
		Month:        utils.GetFirstDayOfMonth(month.Year(), month.Month()),
		Transactions: map[Category][]*Transaction{},
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
		p.Transactions[c] = make([]*Transaction, 0)
	}
	categoryList := p.Transactions[c]

	transaction.Date = time.Date(
		p.Month.Year(),
		p.Month.Month(),
		transaction.Date.Day(),
		0, 0, 0, 0, time.Local)
	transaction.Category = c
	categoryList = append(categoryList, transaction)
	p.Transactions[c] = categoryList
}

func (p *Portfolio) updateBalance() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	p.Balance = 0
	for c, categoryList := range p.Transactions {
		wg.Add(1)
		go func(c Category, categoryList []*Transaction) {
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
