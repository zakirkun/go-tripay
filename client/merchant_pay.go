package client

import (
	"context"
	"encoding/json"

	"github.com/zakirkun/go-tripay/internal/requester"
)

/*
used to get a list of payment channels that are active on your Merchant account along with complete information including transaction fees from each channel. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	response, err := c.MerchantPay()
	if err != nil{
		// do something
	}
	// do something
*/
func (c Client) MerchantPay() (tripayResponses[[]merchantResponse], error) {
	return merchantPay(c, nil)
}

/*
used to get a list of payment channels that are active on your Merchant account along with complete information including transaction fees from each channel. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	response, err := c.MerchantPayWithContext(context.Background())
	if err != nil{
		// do something
	}
	// do something
*/
func (c Client) MerchantPayWithContext(ctx context.Context) (tripayResponses[[]merchantResponse], error) {
	return merchantPay(c, ctx)
}

func merchantPay(c Client, ctx context.Context) (tripayResponses[[]merchantResponse], error) {
	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "merchant/payment-channel",
		Method: "GET",
		Body:   nil,
		Header: c.HeaderRequest(),
	}
	req := requester.NewRequester(paramReq)

	bodyReq := new(requester.IResponseBody)
	var errReq error
	if ctx != nil {
		bodyReq, errReq = req.DOWithContext(ctx)
	} else {
		bodyReq, errReq = req.DO()
	}

	if errReq != nil {
		return tripayResponses[[]merchantResponse]{}, errReq
	}

	var successResponse tripayResponses[[]merchantResponse]
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return successResponse, nil
}
