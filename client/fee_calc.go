package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/zakirkun/go-tripay/internal/requester"
	"github.com/zakirkun/go-tripay/utils"
)

type FeeCalcParam struct {
	Amount int
	Code   utils.TRIPAY_CHANNEL
}

/*
used to get the details of the transaction fee calculation for each channel based on the nominal specified. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	fcParam := FeeCalcParam{Code: utils.CHANNEL_ALFAMIDI,Amount: 100000}
	response, err := client.FeeCalc(fcParam)
	if err != nil {
		// do something
	}
	// do something
*/
func (c Client) FeeCalc(p FeeCalcParam) (tripayResponses[[]feeCalcResponse], error) {
	return feeCalc(c, p, nil)
}

/*
used to get the details of the transaction fee calculation for each channel based on the nominal specified. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	fcParam := FeeCalcParam{Code: utils.CHANNEL_ALFAMIDI,Amount: 100000}
	response, err := client.FeeCalcWithContext(context.Background(), fcParam)
	if err != nil {
		// do something
	}
	// do something
*/
func (c Client) FeeCalcWithContext(ctx context.Context, p FeeCalcParam) (tripayResponses[[]feeCalcResponse], error) {
	return feeCalc(c, p, ctx)
}

func feeCalc(c Client, p FeeCalcParam, ctx context.Context) (tripayResponses[[]feeCalcResponse], error) {
	param := url.Values{}
	param.Set("code", string(p.Code))
	param.Set("amount", fmt.Sprintf("%d", p.Amount))

	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "merchant/fee-calculator?" + param.Encode(),
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
		return tripayResponses[[]feeCalcResponse]{}, errReq
	}

	var successResponse tripayResponses[[]feeCalcResponse]
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return successResponse, nil
}
