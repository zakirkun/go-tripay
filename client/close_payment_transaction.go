package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zakirkun/go-tripay/internal/requester"
	"github.com/zakirkun/go-tripay/utils"
)

type TripayExpiredTime int

type ClosePaymentRequestTransactionResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Data    ClosePaymentTransactionOrder `json:"data"`
}

type ClosePaymentTransactionOrder struct {
	Reference            string                             `json:"reference"`
	MerchantRef          string                             `json:"merchant_ref"`
	PaymentSelectionType string                             `json:"payment_selection_type"`
	PaymentMethod        string                             `json:"payment_method"`
	PaymentName          string                             `json:"payment_name"`
	CustomerName         string                             `json:"customer_name"`
	CustomerEmail        string                             `json:"customer_email"`
	CustomerPhone        string                             `json:"customer_phone"`
	CallbackURL          string                             `json:"callback_url"`
	ReturnURL            string                             `json:"return_url"`
	Amount               int                                `json:"amount"`
	FeeMerchant          int                                `json:"fee_merchant"`
	FeeCustomer          int                                `json:"fee_customer"`
	TotalFee             int                                `json:"total_fee"`
	AmountReceived       int                                `json:"amount_received"`
	PayCode              string                             `json:"pay_code"`
	PayURL               interface{}                        `json:"pay_url"`
	CheckoutURL          string                             `json:"checkout_url"`
	Status               string                             `json:"status"`
	ExpiredTime          int                                `json:"expired_time"`
	OrderItems           []ClosePaymentTransactionOrderItem `json:"order_items"`
	Instructions         []Instruction                      `json:"instructions"`
	QRString             interface{}                        `json:"qr_string"`
	QRURL                interface{}                        `json:"qr_url"`
}

type ClosePaymentTransactionOrderItem struct {
	SKU        string `json:"sku"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
	Subtotal   int    `json:"subtotal"`
	ProductURL string `json:"product_url"`
	ImageURL   string `json:"image_url"`
}

type Instruction struct {
	Title string   `json:"title"`
	Steps []string `json:"steps"`
}

type ClosePaymentBodyRequest struct {
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

type OrderItemClosePaymentRequest struct {
	SKU        string `json:"sku"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
	ProductURL string `json:"product_url"`
	ImageURL   string `json:"image_url"`
}

func (c Client) ClosePaymentRequestTransaction(ctx context.Context, req ClosePaymentBodyRequest) (*ClosePaymentRequestTransactionResponse, error) {
	return closePaymentRequestTransaction(c, nil, req)
}

func (c Client) ClosePaymentRequestTransactionWithContext(ctx context.Context, req ClosePaymentBodyRequest) (*ClosePaymentRequestTransactionResponse, error) {
	return closePaymentRequestTransaction(c, ctx, req)
}

func closePaymentRequestTransaction(c Client, ctx context.Context, reqBody ClosePaymentBodyRequest) (*ClosePaymentRequestTransactionResponse, error) {
	reqBodyByte, _ := json.Marshal(&reqBody)
	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "transaction/create",
		Method: "POST",
		Body:   reqBodyByte,
		Header: c.HeaderRequest(),
	}
	fmt.Printf("reqBodyByte: %v\n", string(reqBodyByte))

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

	var successResponse ClosePaymentRequestTransactionResponse
	json.Unmarshal(bodyReq.ResponseBody, &successResponse)
	return &successResponse, nil
}

func SetTripayExpiredTime(hour int) TripayExpiredTime {
	return TripayExpiredTime(int(time.Now().Unix()) + (hour * 60 * 60))
}
