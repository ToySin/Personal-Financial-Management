package portfolio

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ToySin/finance/utils"
)

var myTestPortfolio = &Portfolio{
	Month: time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
	Transactions: map[Category][]Transaction{
		IncomeCategory: {
			{
				Date:     utils.GetLastBusinessDay(2024, 10),
				Name:     string(utils.Salary),
				Amount:   399,
				Category: IncomeCategory,
			},
			{
				Date:     utils.GetLastBusinessDay(2024, 10),
				Name:     string(utils.Bonus),
				Amount:   1,
				Category: IncomeCategory,
			},
		},
		FixedExpenseCategory: {
			{
				Date:     time.Date(2024, 10, 8, 0, 0, 0, 0, time.Local),
				Name:     string(utils.MonthlyRent),
				Amount:   15,
				Category: FixedExpenseCategory,
			},
			{
				Date:     time.Date(2024, 10, 8, 0, 0, 0, 0, time.Local),
				Name:     string(utils.HousingLoan),
				Amount:   20,
				Category: FixedExpenseCategory,
			},
			{
				Date:     time.Date(2024, 10, 15, 0, 0, 0, 0, time.Local),
				Name:     string(utils.EducationLoan),
				Amount:   15,
				Category: FixedExpenseCategory,
			},
			{
				Date:     time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
				Name:     string(utils.Insurance),
				Amount:   17,
				Category: FixedExpenseCategory,
			},
			{
				Date:     time.Date(2024, 10, 27, 0, 0, 0, 0, time.Local),
				Name:     string(utils.PhoneBill),
				Amount:   17,
				Category: FixedExpenseCategory,
			},
			{
				Date:     time.Date(2024, 10, 15, 0, 0, 0, 0, time.Local),
				Name:     string(utils.InternetBill),
				Amount:   22,
				Category: FixedExpenseCategory,
			},
		},
		SavingCategory: {
			{
				Date:     time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
				Name:     string(utils.YouthSaving),
				Amount:   70,
				Category: SavingCategory,
			},
			{
				Date:     time.Date(2024, 10, 15, 0, 0, 0, 0, time.Local),
				Name:     string(utils.YouthHouseSaving),
				Amount:   5,
				Category: SavingCategory,
			},
		},
		InvestmentCategory: {
			{
				Date:     time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
				Name:     string(utils.KoreanStock),
				Amount:   30,
				Category: InvestmentCategory,
			},
			{
				Date:     time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
				Name:     string(utils.ForeignStock),
				Amount:   60,
				Category: InvestmentCategory,
			},
		},
	},
	Balance: 0,
}

func TestGetBalance(t *testing.T) {
	t.Logf("myTestPortfolio.Balance: %v", myTestPortfolio.GetBalance())
	expectedBalance := myTestPortfolio.TotalAmount(IncomeCategory) -
		myTestPortfolio.TotalAmount(FixedExpenseCategory) -
		myTestPortfolio.TotalAmount(SavingCategory) -
		myTestPortfolio.TotalAmount(InvestmentCategory)
	assert.Equal(t, expectedBalance, myTestPortfolio.GetBalance())
}
