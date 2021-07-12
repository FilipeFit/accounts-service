package api

type PaymentServiceResponse struct {
	ID          uint64  `json:"id"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Channel     string  `json:"channel"`
	Description string  `json:"description"`
	Destination string  `json:"destination"`
}
