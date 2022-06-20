package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/filipeFit/account-service/client/restclient"
	"github.com/filipeFit/account-service/config"
	"github.com/filipeFit/account-service/domain/api"
)

func GetPayments(accountID uint64, authorization string) ([]api.PaymentServiceResponse, error) {
	url := fmt.Sprintf("%s/%d", config.Config.PaymentsServiceUrl, accountID)
	var payments []api.PaymentServiceResponse
	response, err := restclient.Get(url, authorization)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &payments); err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("the customer does not exists")
	}

	return payments, nil
}
