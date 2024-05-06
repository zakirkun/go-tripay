package client

import (
	"context"
	"encoding/json"

	"github.com/zakirkun/go-tripay/internal/requester"
)

type MerchantResponseOK struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []struct {
		Group       string `json:"group"`
		Code        string `json:"code"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		FeeMerchant struct {
			Flat    uint `json:"flat"`
			Percent uint `json:"percent"`
		} `json:"fee_merchant"`
		FeeCustomer struct {
			Flat    uint `json:"flat"`
			Percent uint `json:"percent"`
		} `json:"fee_customer"`
		TotalFee struct {
			Flat    uint    `json:"flat"`
			Percent float64 `json:"percent"`
		} `json:"total_fee"`
		MinimumFee uint   `json:"minimum_fee"`
		MaximumFee uint   `json:"maximum_fee"`
		IconURL    string `json:"icon_url"`
		Active     bool   `json:"active"`
	} `json:"data"`
}

type MerchantResponseFail struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (c Client) MerchantPay() (*MerchantResponseOK, *MerchantResponseFail) {
	return merchantPay(c, nil)
}

func (c Client) MerchantPayWithContext(ctx context.Context) (*MerchantResponseOK, *MerchantResponseFail) {
	return merchantPay(c, ctx)
}

func merchantPay(c Client, ctx context.Context) (*MerchantResponseOK, *MerchantResponseFail) {
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
		var failResponse MerchantResponseFail
		_ = json.Unmarshal(bodyReq.ResponseBody, &failResponse)
		return nil, &failResponse
	}

	var successResponse MerchantResponseOK
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return &successResponse, nil
}
