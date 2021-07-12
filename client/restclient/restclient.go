package restclient

import "net/http"

func Get(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	return client.Do(request)
}
