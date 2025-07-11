package main

import (
	"net/http"

	"github.com/devRaelBraga/rinha-backend-go/internal/handler"
	"github.com/devRaelBraga/rinha-backend-go/internal/storage"
)

func main() {
	s := storage.NewStorage()
	h := handler.NewHandler(s)

	http.HandleFunc("/payments", h.HandlePayments)
	http.HandleFunc("/payment-summary", h.HandleSummary)
	http.ListenAndServe(":9999", nil)
}
