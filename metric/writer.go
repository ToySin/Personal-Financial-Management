package metric

import (
	"fmt"
	"io"

	humanize "github.com/dustin/go-humanize"

	"github.com/ToySin/finance/portfolio"
)

const (
	dateFormat = "2006-01-02"
)

var categorySequence = []portfolio.Category{
	portfolio.IncomeCategory,
	portfolio.FixedExpenseCategory,
	portfolio.VariableExpenseCategory,
	portfolio.SavingCategory,
	portfolio.InvestmentCategory,
	portfolio.InvestmentIncomeCategory,
}

func WritePortfolio(w io.Writer, portfolio portfolio.Portfolio) error {
	fmt.Fprintf(w, "# %d년 %d월 가계부 요약\n\n", portfolio.Month.Year(), int(portfolio.Month.Month()))

	for _, category := range categorySequence {
		if transactions, exist := portfolio.Transactions[category]; exist && len(transactions) > 0 {
			fmt.Fprintf(w, "## %s\n", category)
			for _, t := range transactions {
				amount := humanize.Comma(int64(t.Amount))
				fmt.Fprintf(w, "- [%s] %s %v원\n",
					t.Date.Format(dateFormat), t.Name, amount)
			}
			fmt.Fprintf(w, "\n총 %s: %v원\n\n", category, humanize.Comma(int64(portfolio.TotalAmount(category))))
		}
	}

	fmt.Fprintf(w, "## 🔄 잔액 (Balance)\n")
	fmt.Fprintf(w, "- 전 월 제외 이번 달 잔액: %v원\n", humanize.Comma(int64(portfolio.GetBalance())))

	return nil
}
