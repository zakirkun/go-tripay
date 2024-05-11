package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zakirkun/go-tripay/internal/requester"
	"github.com/zakirkun/go-tripay/utils"
)

type (
	TripayExpiredTime       int
	ClosePaymentBodyRequest struct {
		Method        utils.TRIPAY_CHANNEL           `json:"method"`
		MerchantRef   string                         `json:"merchant_ref"`
		Amount        int                            `json:"amount"`
		CustomerName  string                         `json:"customer_name"`
		CustomerEmail string                         `json:"customer_email"`
		CustomerPhone string                         `json:"customer_phone"`
		OrderItems    []OrderItemClosePaymentRequest `json:"order_items"`
		ReturnURL     string                         `json:"return_url"`
		ExpiredTime   TripayExpiredTime              `json:"expired_time"`
		Signature     string                         `json:"signature"`
	}

	OrderItemClosePaymentRequest struct {
		SKU        string `json:"sku"`
		Name       string `json:"name"`
		Price      int    `json:"price"`
		Quantity   int    `json:"quantity"`
		ProductURL string `json:"product_url"`
		ImageURL   string `json:"image_url"`
	}

	TransactionDetailBodyRequest struct {
		Reference string
	}
)

func SetTripayExpiredTime(hour int) TripayExpiredTime {
	return TripayExpiredTime(int(time.Now().Unix()) + (hour * 60 * 60))
}

func (c Client) ClosePaymentRequestTransaction(req ClosePaymentBodyRequest) (tripayResponses[closePaymentTransactionOrderResponse], error) {
	return closePaymentRequestTransaction(c, nil, req)
}

func (c Client) ClosePaymentRequestTransactionWithContext(ctx context.Context, req ClosePaymentBodyRequest) (tripayResponses[closePaymentTransactionOrderResponse], error) {
	return closePaymentRequestTransaction(c, ctx, req)
}

func closePaymentRequestTransaction(c Client, ctx context.Context, reqBody ClosePaymentBodyRequest) (tripayResponses[closePaymentTransactionOrderResponse], error) {
	reqBodyByte, _ := json.Marshal(&reqBody)
	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "transaction/create",
		Method: "POST",
		Body:   reqBodyByte,
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
		return tripayResponses[closePaymentTransactionOrderResponse]{}, errReq
	}

	var successResponse tripayResponses[closePaymentTransactionOrderResponse]
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return successResponse, nil
}

func (c Client) ClosePaymentTransactionGetDetail(reference string) (tripayResponses[closePaymentTransactionOrderResponse], error) {
	return closePaymentTransactionGetDetail(c, nil, reference)
}

func (c Client) ClosePaymentTransactionGetDetailWithContext(ctx context.Context, reference string) (tripayResponses[closePaymentTransactionOrderResponse], error) {
	return closePaymentTransactionGetDetail(c, ctx, reference)
}

func closePaymentTransactionGetDetail(c Client, ctx context.Context, reference string) (tripayResponses[closePaymentTransactionOrderResponse], error) {
	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "transaction/detail?" + fmt.Sprintf("reference=%s", reference),
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
		return tripayResponses[closePaymentTransactionOrderResponse]{}, errReq
	}

	var successResponse tripayResponses[closePaymentTransactionOrderResponse]
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return successResponse, nil
}
