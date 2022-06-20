package services

import (
	"fmt"
	"net/http"

	"github.com/filipeFit/account-service/client"
	"github.com/filipeFit/account-service/config"
	"github.com/filipeFit/account-service/domain/api"
	"github.com/filipeFit/account-service/handlers"
	"github.com/filipeFit/account-service/repositories"
)

type accountService struct{}

type accountServiceInterface interface {
	CreateAccount(request *api.CreateAccountRequest, authorization string) (*api.CreateAccountResponse, handlers.ApiError)
	FindByAccountID(id uint64, authorization string) (*api.CreateAccountResponse, handlers.ApiError)
	FindAll(authorization string) ([]api.CreateAccountResponse, handlers.ApiError)
	PerformPayment(payment *api.PaymentRequest) (*api.PaymentResponse, handlers.ApiError)
	AccountStatement(accountID uint64, authorization string) (*api.Statement, handlers.ApiError)
}

var (
	AccountService    accountServiceInterface
	accountRepository = repositories.NewAccountRepository(config.DB)
)

func init() {
	AccountService = &accountService{}
}
func (s *accountService) CreateAccount(request *api.CreateAccountRequest, authorization string) (*api.CreateAccountResponse, handlers.ApiError) {
	customer, err := client.GetCustomer(request.CustomerId, authorization)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError,
			fmt.Sprintf("The informed customer %d does not exists", request.CustomerId))
	}

	if isCustomerHaveAccount(request.CustomerId) == true {
		return nil, handlers.NewApiError(http.StatusBadRequest, "customer already have a valid account")
	}

	account, err := accountRepository.Create(request)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError, "error in saving the account")
	}
	response := api.ToAccountResponse(account, customer.Email, customer.Name)
	return response, nil
}

func (s *accountService) FindByAccountID(id uint64, authorization string) (*api.CreateAccountResponse, handlers.ApiError) {
	account, err := accountRepository.FindByID(id)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusNotFound, "account not found")
	}

	customer, err := client.GetCustomer(account.CustomerId, authorization)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError,
			fmt.Sprintf("The informed customer %d does not exists", account.CustomerId))
	}

	response := api.ToAccountResponse(account, customer.Email, customer.Name)
	return response, nil
}

func (s *accountService) FindAll(authorization string) ([]api.CreateAccountResponse, handlers.ApiError) {
	accounts, err := accountRepository.FindAll()
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError, "error searching for accounts")
	}
	var accountsResponse []api.CreateAccountResponse
	for _, account := range accounts {
		customer, err := client.GetCustomer(account.CustomerId, authorization)
		if err != nil {
			return nil, handlers.NewApiError(http.StatusBadRequest,
				fmt.Sprintf("The informed customer %d does not exists", account.CustomerId))
		}
		var accountResponse = api.ToAccountResponse(&account, customer.Email, customer.Name)
		accountsResponse = append(accountsResponse, *accountResponse)
	}

	return accountsResponse, nil
}

func (s *accountService) PerformPayment(payment *api.PaymentRequest) (*api.PaymentResponse, handlers.ApiError) {
	account, err := accountRepository.FindByID(payment.AccountID)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusBadRequest, "account not found")
	}

	balance := account.Balance
	if payment.Type == "C" {
		balance = account.Balance + payment.Amount
	}
	if payment.Type == "D" {
		if payment.Amount > account.Balance && !account.AllowOverdraft {
			return nil, handlers.NewApiError(http.StatusBadRequest, "not enough balance")
		}
		balance = account.Balance - payment.Amount
	}

	updatedBalance, err := accountRepository.UpdateBalance(payment.AccountID, balance)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusUnprocessableEntity, "not able to perform the payment")
	}
	paymentResponse := api.PaymentResponse{
		AccountID: payment.AccountID,
		Balance:   updatedBalance,
	}

	return &paymentResponse, nil
}

func (s *accountService) AccountStatement(accountID uint64, authorization string) (*api.Statement, handlers.ApiError) {
	account, err := accountRepository.FindByID(accountID)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusNotFound, "account not found")
	}

	customer, err := client.GetCustomer(account.CustomerId, authorization)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError,
			fmt.Sprintf("The informed customer %d does not exists", account.CustomerId))
	}
	payments, err := client.GetPayments(accountID, authorization)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError,
			fmt.Sprintf("not possible to retrieve payments for the account %d", account.ID))
	}

	statement := api.Statement{ID: accountID,
		Balance:        account.Balance,
		CustomerId:     customer.ID,
		Active:         account.Active,
		AllowOverdraft: account.AllowOverdraft,
		CustomerEmail:  customer.Email,
		CustomerName:   customer.Name,
		Payments:       payments}

	return &statement, nil

}

func isCustomerHaveAccount(customerId uint64) bool {
	_, err := accountRepository.FindByCustomerID(customerId)
	if err != nil && err.Error() == "record not found" {
		return false
	}
	return true
}
