# Personal-Financial-Management
For my personal purpose

# Requirements
- 매월 가계 비율 작성하기
- 연 / 월간 이율 계산하기

# Terminology
- Portfolio
    - 월 별 종합 가계부 단위
- Category
    - 근로수익, 고정지출, 변동지출, 저축, 투자 등의 단위
- Transaction
    - 각 Category의 항목들. 고정지출 (월세, 대출이자, 보험료, 통신비 등)

# Required Env
```go
type Config struct {
	FinanceDBHost string `envconfig:"FINANCE_DB_HOST" default:"localhost"`
	FinanceDBPort string `envconfig:"FINANCE_DB_PORT" default:"5432"`
	FinanceDBUser string `envconfig:"FINANCE_DB_USER" default:""`
	FinanceDBPass string `envconfig:"FINANCE_DB_PASS" default:""`
	FinanceDBName string `envconfig:"FINANCE_DB_NAME" default:"finance"`
}

type config struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	PlatformEnv string `env:"PLATFORM_ENV" envDefault:"local"`
}
```

```javascript
const CONFIG = {
    API_BASE_URL: "http://api.example.com",
};
  
export default CONFIG;
```
