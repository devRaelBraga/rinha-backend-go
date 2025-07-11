package storage

import (
	"log"
	"time"

	"github.com/devRaelBraga/rinha-backend-go/internal/types"
)

type Storage struct {
	summaries   map[string]types.Summary
	defaultURL  string
	fallbackURL string
}

func NewStorage() *Storage {
	s := &Storage{
		summaries:   make(map[string]types.Summary),
		defaultURL:  "default",
		fallbackURL: "fallback",
	}
	s.summaries[s.defaultURL] = types.Summary{TotalRequests: 0, TotalAmount: 0}
	s.summaries[s.fallbackURL] = types.Summary{TotalRequests: 0, TotalAmount: 0}
	log.Printf("Initialized storage with defaultURL=%s, fallbackURL=%s", s.defaultURL, s.fallbackURL)
	return s
}

func Increment(s *Storage, processorURL string, amount float64) {
	summary := s.summaries[processorURL]
	summary.TotalRequests++
	summary.TotalAmount += amount
	s.summaries[processorURL] = summary
	log.Printf("Incremented summary for %s: requests=%d, amount=%.2f", processorURL, summary.TotalRequests, summary.TotalAmount)
}

func GetSummary(s *Storage, from, to *time.Time) (types.SummaryResponse, error) {
	summary := types.SummaryResponse{
		Default:  s.summaries[s.defaultURL],
		Fallback: s.summaries[s.fallbackURL],
	}
	log.Printf("Retrieved summary: default={requests=%d, amount=%.2f}, fallback={requests=%d, amount=%.2f}",
		summary.Default.TotalRequests, summary.Default.TotalAmount,
		summary.Fallback.TotalRequests, summary.Fallback.TotalAmount)
	return summary, nil
}
