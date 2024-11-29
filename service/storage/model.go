package storage

import (
	"time"
)

// TransactionTable represents the transaction table in the database.
type TransactionTable struct {
	UUID     string `gorm:"primaryKey"`
	Date     time.Time
	Name     string
	Category string
	Amount   int
	Note     string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName returns the table name of the TransactionTable.
func (t *TransactionTable) TableName() string {
	return "transaction"
}
