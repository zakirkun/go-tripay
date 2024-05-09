package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/zakirkun/go-tripay/internal/requester"
)

type OpenPaymentRequest struct {
	Method       string `json:"method"`
	MerchatReff  string `json:"merchant_ref"`
	CustomerName string `json:"customer_name"`
	Signature    string `json:"signature"`
}

type OpenPaymentResponse struct {
	Success         bool   `json:"success"`
	Message         string `json:"message"`
	OpenPaymentData struct {
		UUID          string `json:"uuid"`
		MerchantRef   string `json:"merchant_ref"`
		CustomerName  string `json:"customer_name"`
		PaymentName   string `json:"payment_name"`
		PaymentMethod string `json:"payment_method"`
		PayCode       string `json:"pay_code"`
		QRString      string `json:"qr_string"`
		QRURL         string `json:"qr_url"`
	} `json:"data"`
}

type OpenPaymentDetailResponse struct {
	Success               bool   `json:"success"`
	Message               string `json:"message"`
	OpenPaymentDetailData struct {
		UUID          string `json:"uuid"`
		MerchantRef   string `json:"merchant_ref"`
		CustomerName  string `json:"customer_name"`
		PaymentName   string `json:"payment_name"`
		PaymentMethod string `json:"payment_method"`
		PayCode       string `json:"pay_code"`
		QRString      string `json:"qr_string,omitempty"`
		QRURL         string `json:"qr_url,omitempty"`
	} `json:"data"`
}

type OpenPaymentListParams struct {
	Page        int
	PerPage     int
	Sort        string
	Reference   string
	MerchantRef string
	Method      string
	Status      string
}

type OpenPaymentListResponse struct {
	Success                bool   `json:"success"`
	Message                string `json:"message"`
	OpenPaymentTransaction []struct {
		Reference      string `json:"reference"`
		MerchantRef    string `json:"merchant_ref"`
		PaymentMethod  string `json:"payment_method"`
		PaymentName    string `json:"payment_name"`
		CustomerName   string `json:"customer_name"`
		Amount         int    `json:"amount"`
		FeeMerchant    int    `json:"fee_merchant"`
		FeeCustomer    int    `json:"fee_customer"`
		TotalFee       int    `json:"total_fee"`
		AmountReceived int    `json:"amount_received"`
		CheckoutURL    string `json:"checkout_url"`
		Status         string `json:"status"`
		PaidAt         int64  `json:"paid_at"`
	} `json:"data"`
	Pagination struct {
		Total       int  `json:"total"`
		DataFrom    int  `json:"data_from"`
		DataTo      int  `json:"data_to"`
		PerPage     int  `json:"per_page"`
		CurrentPage int  `json:"current_page"`
		LastPage    int  `json:"last_page"`
		NextPage    *int `json:"next_page"`
	} `json:"pagination"`
}

func (c Client) OpenPaymentDetail(uuid string) (*OpenPaymentDetailResponse, error) {
	return openPaymentDetail(c, uuid, nil)
}

func (c Client) OpenPaymentDetailWithContext(uuid string, ctx context.Context) (*OpenPaymentDetailResponse, error) {
	return openPaymentDetail(c, uuid, ctx)
}

func (c Client) OpenPaymentTransaction(p OpenPaymentRequest) (*OpenPaymentResponse, error) {
	return openPayment(c, p, nil)
}

func (c Client) OpenPaymentTransactionWithContext(p OpenPaymentRequest, ctx context.Context) (*OpenPaymentResponse, error) {
	return openPayment(c, p, ctx)
}

func (c Client) OpenPaymentList(uuid string, p ...OpenPaymentListParams) (*OpenPaymentListResponse, error) {
	return openPaymentList(c, uuid, nil, p...)
}

func (c Client) OpenPaymentListWithContext(uuid string, ctx context.Context, p ...OpenPaymentListParams) (*OpenPaymentListResponse, error) {
	return openPaymentList(c, uuid, ctx, p...)
}

func openPayment(c Client, p OpenPaymentRequest, ctx context.Context) (*OpenPaymentResponse, error) {

	payloadBody, err := json.Marshal(&p)
	if err != nil {
		return nil, err
	}

	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "open-payment/create",
		Method: "POST",
		Body:   payloadBody,
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

	var successResponse OpenPaymentResponse
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return &successResponse, nil
}

func openPaymentDetail(c Client, uuid string, ctx context.Context) (*OpenPaymentDetailResponse, error) {

	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "open-payment/" + uuid + "/detail",
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

	var successResponse OpenPaymentDetailResponse
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return &successResponse, nil
}

func openPaymentList(c Client, uuid string, ctx context.Context, p ...OpenPaymentListParams) (*OpenPaymentListResponse, error) {

	var paymentParams OpenPaymentListParams
	for _, m := range p {
		paymentParams = m
	}

	urlParam := url.Values{}
	urlParam.Set("page", fmt.Sprintf("%d", paymentParams.Page))
	urlParam.Set("per_page", fmt.Sprintf("%d", paymentParams.PerPage))
	urlParam.Set("sort", paymentParams.Sort)
	urlParam.Set("reference", paymentParams.Reference)
	urlParam.Set("merchant_ref", paymentParams.MerchantRef)
	urlParam.Set("method", paymentParams.Method)
	urlParam.Set("status", paymentParams.Status)

	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "open-payment/" + uuid + "/transactions?" + urlParam.Encode(),
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

	var response OpenPaymentListResponse
	json.Unmarshal(bodyReq.ResponseBody, &response)
	return &response, nil

}
