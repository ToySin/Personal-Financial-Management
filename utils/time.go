package utils

import (
	"time"
)

// TODO: Get the holiday list from the government API

// Holiday list of 2024 in Korea
var holidays = map[string]struct{}{
	"2024-01-01": {}, // 새해
	"2024-02-09": {}, // 설날 전날
	"2024-02-10": {}, // 설날
	"2024-02-11": {}, // 설날 다음날
	"2024-03-01": {}, // 삼일절
	"2024-05-05": {}, // 어린이날
	"2024-06-06": {}, // 현충일
	"2024-08-15": {}, // 광복절
	"2024-09-17": {}, // 추석 전날
	"2024-09-18": {}, // 추석
	"2024-09-19": {}, // 추석 다음날
	"2024-10-03": {}, // 개천절
	"2024-10-09": {}, // 한글날
	"2024-12-25": {}, // 크리스마스
}

func isHoliday(date time.Time) bool {
	_, exists := holidays[date.Format("1999-12-03")]
	return exists
}

func isBusinessDay(date time.Time) bool {
	weekday := date.Weekday()
	return weekday != time.Saturday && weekday != time.Sunday && !isHoliday(date)
}

// GetLastBusinessDay returns the last business day of the given year and month.
func GetLastBusinessDay(year int, month time.Month) time.Time {
	// Get the last day of the month
	lastDay := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)

	// Find the last business day
	for !isBusinessDay(lastDay) {
		lastDay = lastDay.AddDate(0, 0, -1)
	}

	return lastDay
}
