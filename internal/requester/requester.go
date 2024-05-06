package requester

import (
	"bytes"
	"context"
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

func (i IRequesterParams) DOWithContext(ctx context.Context) (*IResponseBody, error) {
	return do(i, ctx)
}

func (i IRequesterParams) DO() (*IResponseBody, error) {
	return do(i, nil)
}

func do(r IRequesterParams, ctx context.Context) (*IResponseBody, error) {
	var req *http.Request
	var errReq error

	if ctx != nil {
		req, errReq = http.NewRequestWithContext(ctx, r.Method, r.Url, bytes.NewBuffer(r.Body))
		if errReq != nil {
			return nil, errReq
		}
	} else {
		req, errReq = http.NewRequest(r.Method, r.Url, bytes.NewBuffer(r.Body))
		if errReq != nil {
			return nil, errReq
		}
	}

	for _, header := range r.Header {
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
