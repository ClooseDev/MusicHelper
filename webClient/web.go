package webClient

import (
	"io"
	"net/http"
	"time"
)

func MakeGetRequest(path string, params map[string]string, headers map[string]string) (io.ReadCloser, error)  {
	return call(path, params, headers, "GET", nil)
}

func MakePostRequest(path string, params map[string]string, headers map[string]string, body io.Reader) (io.ReadCloser, error) {
	return call(path, params, headers, "POST", body)
}

func call(path string, params map[string]string, headers map[string]string, requestType string, body io.Reader) (io.ReadCloser, error) {
	timeout := 3 * time.Second
	client := http.Client{Timeout: timeout}
	request, err := http.NewRequest(requestType, path, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	return response.Body, err
}
