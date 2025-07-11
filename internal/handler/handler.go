package handler

import (
	"encoding/json"
	"net/http"

	"github.com/devRaelBraga/rinha-backend-go/internal/storage"
	"github.com/devRaelBraga/rinha-backend-go/internal/types"
)

type Handler struct {
	storage *storage.Storage
}

type PaymentRequest struct {
	CorrelationID string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
}

func NewHandler(s *storage.Storage) *Handler {
	return &Handler{storage: s}
}

func (h *Handler) HandlePayments(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.storage.Increment("default", req.Amount)

	data := types.Payment{
		Success:       true,
		CorrelationId: req.CorrelationID,
		Amount:        req.Amount,
		Status:        http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(data)
}

func (h *Handler) HandleSummary(w http.ResponseWriter, r *http.Request) {
	// data := types.SummaryResponse{
	// 	Default: types.Summary{
	// 		TotalRequests: 123456,
	// 		TotalAmount:   123456,
	// 	},
	// 	Fallback: types.Summary{
	// 		TotalRequests: 123456,
	// 		TotalAmount:   123456,
	// 	},
	// }

	// h.storage.Increment("default", 123.123)
	response, _ := h.storage.GetSummary()

	json.NewEncoder(w).Encode(response)
}
