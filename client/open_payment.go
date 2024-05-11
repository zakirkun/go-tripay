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

type OpenPaymentListParams struct {
	Page        int
	PerPage     int
	Sort        string
	Reference   string
	MerchantRef string
	Method      string
	Status      string
}

/*
is used to retrieve the list of payments entered in the open payment. this method only work in production!
*/
func (c Client) OpenPaymentDetail(uuid string) (tripayResponses[openPaymentDetailResponse], error) {
	return openPaymentDetail(c, uuid, nil)
}

/*
is used to retrieve the list of payments entered in the open payment. this method only work in production!
*/
func (c Client) OpenPaymentDetailWithContext(uuid string, ctx context.Context) (tripayResponses[openPaymentDetailResponse], error) {
	return openPaymentDetail(c, uuid, ctx)
}

/*
Used to create a new transaction or generate a payment code for Open Payment type. this method only work in production!
*/
func (c Client) OpenPaymentTransaction(p OpenPaymentRequest) (tripayResponses[openPaymentDataResponse], error) {
	return openPayment(c, p, nil)
}

/*
Used to create a new transaction or generate a payment code for Open Payment type. this method only work in production!
*/
func (c Client) OpenPaymentTransactionWithContext(p OpenPaymentRequest, ctx context.Context) (tripayResponses[openPaymentDataResponse], error) {
	return openPayment(c, p, ctx)
}

/*
used to retrieve details of open payment transactions that have been made. this method only work in production!
*/
func (c Client) OpenPaymentList(uuid string, p ...OpenPaymentListParams) (*openPaymentListResponse, error) {
	return openPaymentList(c, uuid, nil, p...)
}

/*
used to retrieve details of open payment transactions that have been made. this method only work in production!
*/
func (c Client) OpenPaymentListWithContext(uuid string, ctx context.Context, p ...OpenPaymentListParams) (*openPaymentListResponse, error) {
	return openPaymentList(c, uuid, ctx, p...)
}

func openPayment(c Client, p OpenPaymentRequest, ctx context.Context) (tripayResponses[openPaymentDataResponse], error) {

	payloadBody, err := json.Marshal(&p)
	if err != nil {
		return tripayResponses[openPaymentDataResponse]{}, err
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
		return tripayResponses[openPaymentDataResponse]{}, errReq
	}

	var successResponse tripayResponses[openPaymentDataResponse]
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return successResponse, nil
}

func openPaymentDetail(c Client, uuid string, ctx context.Context) (tripayResponses[openPaymentDetailResponse], error) {

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
		return tripayResponses[openPaymentDetailResponse]{}, errReq
	}

	var successResponse tripayResponses[openPaymentDetailResponse]
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return successResponse, nil
}

func openPaymentList(c Client, uuid string, ctx context.Context, p ...OpenPaymentListParams) (*openPaymentListResponse, error) {

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

	var response openPaymentListResponse
	json.Unmarshal(bodyReq.ResponseBody, &response)
	return &response, nil

}
