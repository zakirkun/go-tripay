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

/*
used to make a payment limit within hours. Example:

	hour := 24 // one day
	SetTripayExpiredTime(hour)
*/
func SetTripayExpiredTime(hour int) TripayExpiredTime {
	return TripayExpiredTime(int(time.Now().Unix()) + (hour * 60 * 60))
}

/*
used to create a new transaction or generate a payment code. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	c.SetSignature(utils.Signature{Amount: 50000, PrivateKey: "your_private_key", MerchantCode: "T14302",MerchanReff: "INV345675"})

	req := ClosePaymentBodyRequest{
		Method:        utils.CHANNEL_BCAVA,
		MerchantRef:   "INV345675",
		Amount:        50000,
		CustomerName:  "John Doe",
		CustomerEmail: "johndoe@gmail.com",
		CustomerPhone: "62891829828",
		ReturnURL:     "https://thisisreturnurl.com/redirect",
		ExpiredTime:   SetTripayExpiredTime(24), // 24 Hour
		Signature:     client.GetSignature(),
		OrderItems: []OrderItemClosePaymentRequest{
			{
				SKU:        "Produk1",
				Name:       "nama produk 1",
				Price:      50000,
				Quantity:   1,
				ProductURL: "https://producturl.com",
				ImageURL:   "https://imageurl.com",
			},
		},
	}

	response, err := c.ClosePaymentRequestTransaction(req)
	if err != nil {
		// do something
	}
	// do something
*/
func (c Client) ClosePaymentRequestTransaction(req ClosePaymentBodyRequest) (tripayResponses[closePaymentTransactionOrderResponse], error) {
	return closePaymentRequestTransaction(c, nil, req)
}

/*
used to create a new transaction or generate a payment code. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	c.SetSignature(utils.Signature{ Amount: 50000, PrivateKey: "your_private_key", MerchantCode: "T14302",MerchanReff: "INV345675" })

	req := ClosePaymentBodyRequest{
		Method:        utils.CHANNEL_BCAVA,
		MerchantRef:   "INV345675",
		Amount:        50000,
		CustomerName:  "John Doe",
		CustomerEmail: "johndoe@gmail.com",
		CustomerPhone: "62891829828",
		ReturnURL:     "https://thisisreturnurl.com/redirect",
		ExpiredTime:   SetTripayExpiredTime(24), // 24 Hour
		Signature:     client.GetSignature(),
		OrderItems: []OrderItemClosePaymentRequest{
			{
				SKU:        "Produk1",
				Name:       "nama produk 1",
				Price:      50000,
				Quantity:   1,
				ProductURL: "https://producturl.com",
				ImageURL:   "https://imageurl.com",
			},
		},
	}

	response, err := c.ClosePaymentRequestTransactionWithContext(context.Background,req)
	if err != nil {
		// do something
	}
	// do something
*/
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

/*
Used to retrieve details of transactions that have been made. Can also be used to check payment status. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	referenceId := "reference_id"
	response, err := c.ClosePaymentTransactionGetDetail(referenceId)
	if err != nil {
		// do something
	}
	// do something
*/
func (c Client) ClosePaymentTransactionGetDetail(reference string) (tripayResponses[closePaymentTransactionOrderResponse], error) {
	return closePaymentTransactionGetDetail(c, nil, reference)
}

/*
Used to retrieve details of transactions that have been made. Can also be used to check payment status. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	referenceId := "reference_id"
	response, err := c.ClosePaymentTransactionGetDetailWithContext(referenceId)
	if err != nil {
		// do something
	}
	// do something
*/
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
