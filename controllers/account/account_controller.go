package account

import (
	"net/http"
	"strconv"

	"github.com/filipeFit/account-service/domain/api"
	"github.com/filipeFit/account-service/handlers"
	"github.com/filipeFit/account-service/services"
	"github.com/gin-gonic/gin"
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
	authorization := c.Request.Header.Get("Authorization")
	response, err := accountService.CreateAccount(&request, authorization)
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
	authorization := c.Request.Header.Get("Authorization")
	response, apiErr := accountService.FindByAccountID(accountId, authorization)
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}

func FindAll(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	response, apiErr := accountService.FindAll(authorization)
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

	authorization := c.Request.Header.Get("Authorization")

	response, apiErr := accountService.AccountStatement(accountId, authorization)
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}
