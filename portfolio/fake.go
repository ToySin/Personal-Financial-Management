package portfolio

import (
	"time"

	"github.com/ToySin/finance/utils"
)

func GetTestPortfolio() *Portfolio {
	return &Portfolio{
		Month: time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
		Transactions: map[Category][]Transaction{
			IncomeCategory: {
				{
					Date:     utils.GetLastBusinessDay(2024, 10),
					Name:     string(utils.Salary),
					Amount:   5000000,
					Category: IncomeCategory,
				},
			},
			FixedExpenseCategory: {
				{
					Date:     time.Date(2024, 10, 8, 0, 0, 0, 0, time.Local),
					Name:     string(utils.MonthlyRent),
					Amount:   500000,
					Category: FixedExpenseCategory,
				},
				{
					Date:     time.Date(2024, 10, 8, 0, 0, 0, 0, time.Local),
					Name:     string(utils.HousingLoan),
					Amount:   200000,
					Category: FixedExpenseCategory,
				},
				{
					Date:     time.Date(2024, 10, 15, 0, 0, 0, 0, time.Local),
					Name:     string(utils.EducationLoan),
					Amount:   300000,
					Category: FixedExpenseCategory,
				},
				{
					Date:     time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
					Name:     string(utils.Insurance),
					Amount:   200000,
					Category: FixedExpenseCategory,
				},
				{
					Date:     time.Date(2024, 10, 27, 0, 0, 0, 0, time.Local),
					Name:     string(utils.PhoneBill),
					Amount:   30000,
					Category: FixedExpenseCategory,
				},
				{
					Date:     time.Date(2024, 10, 15, 0, 0, 0, 0, time.Local),
					Name:     string(utils.InternetBill),
					Amount:   30000,
					Category: FixedExpenseCategory,
				},
			},
			SavingCategory: {
				{
					Date:     time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
					Name:     string(utils.YouthSaving),
					Amount:   700000,
					Category: SavingCategory,
				},
				{
					Date:     time.Date(2024, 10, 15, 0, 0, 0, 0, time.Local),
					Name:     string(utils.YouthHouseSaving),
					Amount:   100000,
					Category: SavingCategory,
				},
			},
			InvestmentCategory: {
				{
					Date:     time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
					Name:     string(utils.KoreanStock),
					Amount:   300000,
					Category: InvestmentCategory,
				},
				{
					Date:     time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local),
					Name:     string(utils.ForeignStock),
					Amount:   600000,
					Category: InvestmentCategory,
				},
			},
		},
		Balance: 0,
	}
}
