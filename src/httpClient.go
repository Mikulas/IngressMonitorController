package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpClient struct {
	url string
}

type HttpResponse struct {
	statusCode int
	bytes      []byte
}

func createHttpClient(url string) *HttpClient {
	client := HttpClient{url: url}
	return &client
}

func (client *HttpClient) post(body string) (*HttpResponse, error) {
	return client.postWithHeaders(body, nil)
}

func (client *HttpClient) postWithHeaders(body string, headers map[string]string) (*HttpResponse, error) {
	payload := strings.NewReader(body)

	request, err := http.NewRequest("POST", client.url, payload)
	if err != nil {
		return nil, err
	}

	client.addHeaders(request, headers)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	httpResponse := &HttpResponse{statusCode: response.StatusCode}

	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	httpResponse.bytes = responseBytes

	return httpResponse, err
}

func (client *HttpClient) postUrlEncodedFormBody(body string) (*HttpResponse, error) {
	requestHeaders := make(map[string]string)
	requestHeaders["content-type"] = "application/x-www-form-urlencoded"
	requestHeaders["cache-control"] = "no-cache"

	return client.postWithHeaders(body, requestHeaders)
}

func (client *HttpClient) addHeaders(request *http.Request, headers map[string]string) {
	if headers != nil {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}
}
