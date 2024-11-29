package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ToySin/finance/metric"
	"github.com/ToySin/finance/portfolio"
	"github.com/ToySin/finance/service"
)

// APIHandler is a handler for the API.
type APIHandler struct {
	service *service.Service
}

// NewAPIHandler creates a new APIHandler.
func NewAPIHandler(s *service.Service) *APIHandler {
	return &APIHandler{service: s}
}

// GetPortfolioHandler is a handler for getting the portfolio.
func (h *APIHandler) GetPortfolioHandler(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-01", r.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	p, err := h.service.GetPortfolio(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if err := metric.WritePortfolio(w, *p); err != nil {
		http.Error(w, "failed to write portfolio", http.StatusInternalServerError)
		return
	}
}

// CraeteTransactionHandler is a handler for creating a transaction.
func (h *APIHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	category := portfolio.Category.FromCode(portfolio.UnknownCategory, req.Category)
	if category == portfolio.UnknownCategory {
		http.Error(w, "invalid category", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTransaction(date, category, req.Name, portfolio.Amount(req.Amount), req.Note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
