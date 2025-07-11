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

func NewHandler(s *storage.Storage) *Handler {
	return &Handler{storage: s}
}

func (*Handler) HandlePayments(w http.ResponseWriter, r *http.Request) {
	data := types.Payment{
		Success:       true,
		CorrelationId: "1234-5678",
		Amount:        29.90,
		Status:        http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(data)
}

func (*Handler) HandleSummary(w http.ResponseWriter, r *http.Request) {
	data := types.SummaryResponse{
		Default: types.Summary{
			TotalRequests: 123456,
			TotalAmount:   123456,
		},
		Fallback: types.Summary{
			TotalRequests: 123456,
			TotalAmount:   123456,
		},
	}

	json.NewEncoder(w).Encode(data)
}
