package client

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/zakirkun/go-tripay/internal/requester"
)

func (c Client) Instruction(channelCode string, payCode string, amount string, allow_html string) (tripayResponses[[]instructionResponse], error) {

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
		return tripayResponses[[]instructionResponse]{}, err
	}

	var response tripayResponses[[]instructionResponse]
	_ = json.Unmarshal(body.ResponseBody, &response)

	return response, nil
}
