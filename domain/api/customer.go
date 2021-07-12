package api

type GetCustomerResponse struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Nick       string `json:"nick"`
	Document   string `json:"document"`
	Active     bool   `json:"active"`
	LeiCompany string `json:"leiCompany"`
}
