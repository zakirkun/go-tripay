package client

import (
	"context"
	"encoding/json"
	"net/http"

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

func (c Client) OpenPaymentTransaction(p OpenPaymentRequest) (*OpenPaymentResponse, error) {
	return openPayment(c, p, nil)
}

func (c Client) OpenPaymentTransactionWithConext(p OpenPaymentRequest, ctx context.Context) (*OpenPaymentResponse, error) {
	return openPayment(c, p, ctx)
}

func openPayment(c Client, p OpenPaymentRequest, ctx context.Context) (*OpenPaymentResponse, error) {

	p.Signature = c.GetSignature()

	payloadBody, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "open-payment/create",
		Method: http.MethodPost,
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
