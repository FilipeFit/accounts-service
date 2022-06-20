package restclient

import (
	"net/http"
)

func Get(url, authorization string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", authorization)
	client := http.Client{}

	return client.Do(request)
}
