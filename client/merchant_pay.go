package client

import (
	"context"
	"encoding/json"

	"github.com/zakirkun/go-tripay/internal/requester"
)

type MerchantResponse struct {
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

func (c Client) MerchantPay() (*MerchantResponse, error) {
	return merchantPay(c, nil)
}

func (c Client) MerchantPayWithContext(ctx context.Context) (*MerchantResponse, error) {
	return merchantPay(c, ctx)
}

func merchantPay(c Client, ctx context.Context) (*MerchantResponse, error) {
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
		return nil, errReq
	}

	var successResponse MerchantResponse
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return &successResponse, nil
}
