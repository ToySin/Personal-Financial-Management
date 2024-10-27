package metric

import (
	"fmt"
	"io"
	"os"

	humanize "github.com/dustin/go-humanize"

	"github.com/ToySin/finance/portfolio"
)

const (
	dateFormat = "1999-12-03"
)

type Writer interface {
	Portfolio(portfolio.Portfolio) error
}

type FileWriter struct{}

func writePortfolio(w io.Writer, portfolio portfolio.Portfolio) error {
	fmt.Fprintf(w, "# %d년 %d월 가계부 요약\n\n", portfolio.Month.Year(), int(portfolio.Month.Month()))

	for category, transactions := range portfolio.Transactions {
		fmt.Fprintf(w, "## %s\n", category)
		for _, t := range transactions {
			fmt.Fprintf(w, "- [%s] %v원 | 메모: %s\n",
				t.Date.Format(dateFormat), humanize.Comma(int64(t.Amount)), t.Note)
		}
		fmt.Fprintf(w, "\n총 %s: %v원\n\n", category, humanize.Comma(int64(portfolio.TotalAmount(category))))
	}

	fmt.Fprintf(w, "## 🔄 잔액 (Balance)\n")
	fmt.Fprintf(w, "- 전 월 제외 이번 달 잔액: %v원\n", humanize.Comma(int64(portfolio.GetBalance())))

	return nil
}

func (w *FileWriter) Portfolio(filename string, p portfolio.Portfolio) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return writePortfolio(f, p)
}
