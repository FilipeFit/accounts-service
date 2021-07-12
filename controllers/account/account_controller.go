package account

import (
	"github.com/filipeFit/account-service/domain/api"
	"github.com/filipeFit/account-service/handlers"
	"github.com/filipeFit/account-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	accountService = services.AccountService
)

func CreateAccount(c *gin.Context) {
	var request api.CreateAccountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := handlers.NewBadRequestError("invalid json body")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}

	response, err := accountService.CreateAccount(&request)
	if err != nil {
		c.JSON(err.ResponseStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func FindByAccountId(c *gin.Context) {
	accountId, err := strconv.ParseUint(c.Param("accountId"), 10, 64)
	if err != nil {
		apiErr := handlers.NewBadRequestError("invalid account id")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}

	response, apiErr := accountService.FindByAccountID(accountId)
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}

func FindAll(c *gin.Context) {
	response, apiErr := accountService.FindAll()
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}

func Payment(c *gin.Context) {
	var request api.PaymentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := handlers.NewBadRequestError("invalid json body")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	response, err := accountService.PerformPayment(&request)
	if err != nil {
		c.JSON(err.ResponseStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func Statement(c *gin.Context) {
	accountId, err := strconv.ParseUint(c.Param("accountId"), 10, 64)
	if err != nil {
		apiErr := handlers.NewBadRequestError("invalid account id")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}

	response, apiErr := accountService.AccountStatement(accountId)
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}
