package utils

type TransactionName string

const (
	// Income
	Salary TransactionName = "월급"
	Bonus  TransactionName = "상여금"

	// FixedExpense
	MonthlyRent   TransactionName = "월세"
	HousingLoan   TransactionName = "전세대출이자"
	EducationLoan TransactionName = "학자금대출이자"
	Insurance     TransactionName = "보험료"
	PhoneBill     TransactionName = "휴대폰통신비"
	InternetBill  TransactionName = "인터넷통신비"

	// Saving
	YouthSaving      TransactionName = "청년도약계좌"
	YouthHouseSaving TransactionName = "청년주택드림청약"

	// Investment
	KoreanStock  TransactionName = "국내주식"
	ForeignStock TransactionName = "해외주식"
)
