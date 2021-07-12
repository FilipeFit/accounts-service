package app

import (
	"github.com/filipeFit/account-service/controllers/account"
	"github.com/filipeFit/account-service/controllers/heath"
)

func mapRoutes() {
	router.GET("/health", heath.HealthCheck)
	router.POST("/accounts", account.CreateAccount)
	router.POST("/accounts/balances", account.Payment)
	router.GET("/accounts", account.FindAll)
	router.GET("/accounts/:accountId", account.FindByAccountId)
	router.GET("/accounts/:accountId/statements", account.Statement)
}
