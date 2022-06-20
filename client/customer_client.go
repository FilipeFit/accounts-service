package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/filipeFit/account-service/client/restclient"
	"github.com/filipeFit/account-service/config"
	"github.com/filipeFit/account-service/domain/api"
	"io/ioutil"
)

func GetCustomer(customerID uint64, authorization string) (*api.GetCustomerResponse, error) {
	url := fmt.Sprintf("%s/%d", config.Config.CustomerServiceUrl, customerID)
	var customer api.GetCustomerResponse
	response, err := restclient.Get(url, authorization)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &customer); err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("the customer does not exists")
	}

	return &customer, nil
}
