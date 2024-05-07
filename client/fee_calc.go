package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/zakirkun/go-tripay/internal/requester"
	"github.com/zakirkun/go-tripay/utils"
)

type (
	FeeCalcSuccessResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    []struct {
			Code string
			Name string
			Fee  struct {
				Flat    int         `json:"flat"`
				Percent string      `json:"percent"`
				Min     interface{} `json:"min"`
				Max     interface{} `json:"max"`
			} `json:"fee"`
			TotalFee struct {
				Merchant int `json:"merchant"`
				Customer int `json:"customer"`
			} `json:"total_fee"`
		} `json:"data"`
	}

	FeeCalcFailResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	FeeCalcParam struct {
		Amount int
		Code   utils.TRIPAY_CHANNEL
	}
)

func (c Client) FeeCalc(p FeeCalcParam) (*FeeCalcSuccessResponse, error) {
	return feeCalc(c, p, nil)
}

func (c Client) FeeCalcWithContext(ctx context.Context, p FeeCalcParam) (*FeeCalcSuccessResponse, error) {
	return feeCalc(c, p, ctx)
}

func feeCalc(c Client, p FeeCalcParam, ctx context.Context) (*FeeCalcSuccessResponse, error) {
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

	if errReq != nil || strings.Contains(string(paramReq.Body), "false") {
		var failResponse FeeCalcFailResponse
		json.Unmarshal(bodyReq.ResponseBody, &failResponse)
		return nil, errReq
	}

	var successResponse FeeCalcSuccessResponse
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return &successResponse, nil
}
