package api

import "github.com/filipeFit/account-service/domain"

type CreateAccountRequest struct {
	CustomerId     uint64 `json:"customerId"`
	AllowOverdraft bool   `json:"allowOverdraft"`
}

type CreateAccountResponse struct {
	ID             uint64  `json:"id"`
	CustomerId     uint64  `json:"customerId"`
	CustomerName   string  `json:"customerName"`
	CustomerEmail  string  `json:"customerEmail"`
	Balance        float64 `json:"balance"`
	Active         bool    `json:"active"`
	AllowOverdraft bool    `json:"allowOverdraft"`
}

type PaymentRequest struct {
	AccountID uint64  `json:"accountId"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
}
type PaymentResponse struct {
	AccountID uint64  `json:"accountId"`
	Balance   float64 `json:"balance"`
}

type Statement struct {
	ID             uint64                   `json:"id"`
	CustomerId     uint64                   `json:"customerId"`
	CustomerName   string                   `json:"customerName"`
	CustomerEmail  string                   `json:"customer_email"`
	Balance        float64                  `json:"balance"`
	Active         bool                     `json:"active"`
	AllowOverdraft bool                     `json:"allowOverdraft"`
	Payments       []PaymentServiceResponse `json:"payments,omitempty"`
}

func ToAccountResponse(account *domain.Account, customerEmail string, customerName string) *CreateAccountResponse {
	accountResponse := CreateAccountResponse{ID: account.ID,
		Balance:        account.Balance,
		CustomerId:     account.CustomerId,
		Active:         account.Active,
		AllowOverdraft: account.AllowOverdraft,
		CustomerEmail:  customerEmail,
		CustomerName:   customerName}
	return &accountResponse
}
