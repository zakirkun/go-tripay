package client

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/zakirkun/go-tripay/internal/requester"
)

type InstructionResponseOk struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []struct {
		Title string   `json:"title"`
		Steps []string `json:"steps"`
	} `json:"data"`
}

type InstructionResponseBad struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (c Client) Instruction(channelCode string, payCode string, amount string, allow_html string) (*InstructionResponseOk, *InstructionResponseBad) {

	params := url.Values{}
	params.Set("code", channelCode)

	if payCode != "" {
		params.Set("pay_code", payCode)
	}

	if amount != "" {
		params.Set("amount", amount)
	}

	if allow_html != "" {
		params.Set("allow_html", allow_html)
	}

	queryString := params.Encode()

	requester := requester.NewRequester(requester.IRequesterParams{
		c.BaseUrl() + "payment/instruction?" + queryString,
		http.MethodGet,
		nil,
		c.HeaderRequest(),
	})

	body, err := requester.DO()

	if err != nil {
		var responseBad InstructionResponseBad
		_ = json.Unmarshal(body.ResponseBody, &responseBad)

		return nil, &responseBad
	}

	var responseOk InstructionResponseOk
	_ = json.Unmarshal(body.ResponseBody, &responseOk)

	return &responseOk, nil
}
