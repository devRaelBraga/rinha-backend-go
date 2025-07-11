package types

type Payment struct {
	Success       bool    `json:"success"`
	CorrelationId string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
	Status        int16   `json:"status"`
}

type Summary struct {
	TotalRequests int64   `json:"totalRequests"`
	TotalAmount   float64 `json:"totalAmount"`
}

type SummaryResponse struct {
	Default  Summary `json:"default"`
	Fallback Summary `json:"fallback"`
}
