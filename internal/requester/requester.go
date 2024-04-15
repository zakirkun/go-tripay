package requester

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func NewRequester(params IRequesterParams) IRequester {
	return IRequesterParams{
		Url:    params.Url,
		Method: params.Method,
		Body:   params.Body,
		Header: params.Header,
	}
}

func (i IRequesterParams) DO() (*IResponseBody, error) {

	// Create a new HTTP request with the POST method and the request body
	req, err := http.NewRequest(i.Method, i.Url, bytes.NewBuffer(i.Body))
	if err != nil {
		return nil, err
	}

	for _, header := range i.Header {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &IResponseBody{
		HttpCode:     resp.StatusCode,
		ResponseBody: responseBody,
	}, nil
}
