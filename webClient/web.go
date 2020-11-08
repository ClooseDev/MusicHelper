package webClient

import (
	"io"
	"net/http"
	"time"
)

func MakeGetRequest(path string, params map[string]string) (io.ReadCloser, error)  {
	timeout := 3 * time.Second
	client := http.Client{Timeout: timeout}
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	return response.Body, err
}
