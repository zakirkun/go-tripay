package client

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/zakirkun/go-tripay/internal/requester"
	"github.com/zakirkun/go-tripay/utils"
)

type InstructionRequestParam struct {
	ChannelCode utils.TRIPAY_CHANNEL
	PayCode     string
	Amount      string
	AllowHtml   string
}

/*
used to retrieve payment instructions from each channel. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	param := InstructionRequestParam{ ChannelCode: utils.CHANNEL_BRIVA, PayCode: "", Amount: "10000", AllowHtml: "" }
	response, err := c.Instruction(param)
	if err != nil{
		// do something
	}
	// do something
*/
func (c Client) Instruction(ip InstructionRequestParam) (tripayResponses[[]instructionResponse], error) {

	params := url.Values{}
	params.Set("code", string(ip.ChannelCode))

	if ip.PayCode != "" {
		params.Set("pay_code", ip.PayCode)
	}

	if ip.Amount != "" {
		params.Set("amount", ip.Amount)
	}

	if ip.AllowHtml != "" {
		params.Set("ip.AllowHtml", ip.AllowHtml)
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
