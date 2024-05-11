package client

import (
	"context"
	"encoding/json"

	"github.com/zakirkun/go-tripay/internal/requester"
)

func (c Client) MerchantPay() (tripayResponses[[]merchantResponse], error) {
	return merchantPay(c, nil)
}

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
