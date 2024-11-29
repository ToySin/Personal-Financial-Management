package handler

// CreateTransactionRequest represents a request for creating a transaction.
type CreateTransactionRequest struct {
	Date     string `json:"date"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Note     string `json:"note"`
}
