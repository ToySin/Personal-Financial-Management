package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLastBusinessDay(t *testing.T) {
	// Given
	// 2024-06-29 is Saturday
	// 2024-06-30 is Sunday
	year := 2024
	month := 6

	// When
	lastBusinessDay := GetLastBusinessDay(year, time.Month(month))

	// Then
	expected := time.Date(year, time.Month(month), 28, 0, 0, 0, 0, time.Local)
	assert.Equal(t, expected, lastBusinessDay)
}
