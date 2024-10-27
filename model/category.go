package model

type Category string

const (
	Unknown Category = ""

	// Income
	Salary Category = "월급"
	Bonus  Category = "상여금"

	// FixedExpense
	MonthlyRent   Category = "월세"
	HousingLoan   Category = "전세대출이자"
	EducationLoan Category = "학자금대출이자"
	Insurance     Category = "보험료"
	PhoneBill     Category = "휴대폰통신비"
	InternetBill  Category = "인터넷통신비"

	// Saving
	YouthSaving      Category = "청년도약계좌"
	YouthHouseSaving Category = "청년주택드림청약"

	// Investment
	KoreanStock  Category = "국내주식"
	ForeignStock Category = "해외주식"
)
