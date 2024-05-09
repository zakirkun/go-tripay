package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/zakirkun/go-tripay/internal/requester"
)

type (
	OrderItem struct {
		SKU      interface{} `json:"sku"`
		Name     string      `json:"name"`
		Price    int         `json:"price"`
		Quantity int         `json:"quantity"`
		Subtotal int         `json:"subtotal"`
	}

	Data struct {
		Reference        string      `json:"reference"`
		MerchantRef      string      `json:"merchant_ref"`
		PaymentSelection string      `json:"payment_selection_"`
		PaymentMethod    string      `json:"payment_method"`
		PaymentName      string      `json:"payment_name"`
		CustomerName     string      `json:"customer_name"`
		CustomerEmail    string      `json:"customer_email"`
		CustomerPhone    interface{} `json:"customer_phone"`
		CallbackURL      interface{} `json:"callback_url"`
		ReturnURL        interface{} `json:"return_url"`
		Amount           int         `json:"amount"`
		FeeMerchant      int         `json:"fee_merchant"`
		FeeCustomer      int         `json:"fee_customer"`
		TotalFee         int         `json:"total_fee"`
		AmountReceived   int         `json:"amount_received"`
		PayCode          int64       `json:"pay_code"`
		PayURL           interface{} `json:"pay_url"`
		CheckoutURL      string      `json:"checkout_url"`
		OrderItems       []OrderItem `json:"order_items"`
		Status           string      `json:"status"`
		Note             interface{} `json:"note"`
		CreatedAt        int64       `json:"created_at"`
		ExpiredAt        int64       `json:"expired_at"`
		PaidAt           interface{} `json:"paid_at"`
	}

	Pagination struct {
		Sort   string `json:"sort"`
		Offset struct {
			From int `json:"from"`
			To   int `json:"to"`
		} `json:"offset"`
		CurrentPage  int         `json:"current_page"`
		PreviousPage interface{} `json:"previous_page"`
		NextPage     interface{} `json:"next_page"`
		LastPage     int         `json:"last_page"`
		PerPage      int         `json:"per_page"`
		TotalRecords int         `json:"total_records"`
	}

	MerchantTransactionsResponse struct {
		Success    bool       `json:"success"`
		Message    string     `json:"message"`
		Data       []Data     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	MerchantTransactionsParam struct {
		Page        int
		PerPage     int
		Sort        string
		Reference   string
		MerchantRef string
		Method      string
		Status      string
	}
)

func (c Client) MerchantTransactions(p ...MerchantTransactionsParam) (*MerchantTransactionsResponse, error) {
	return merchantTransactions(c, nil, p...)
}

func (c Client) MerchantTransactionsWithContext(ctx context.Context, p ...MerchantTransactionsParam) (*MerchantTransactionsResponse, error) {
	return merchantTransactions(c, ctx, p...)
}

func merchantTransactions(c Client, ctx context.Context, p ...MerchantTransactionsParam) (*MerchantTransactionsResponse, error) {
	var merchatsParams MerchantTransactionsParam

	for _, m := range p {
		merchatsParams = m
	}

	urlParam := url.Values{}
	urlParam.Set("page", fmt.Sprintf("%d", merchatsParams.Page))
	urlParam.Set("per_page", fmt.Sprintf("%d", merchatsParams.PerPage))
	urlParam.Set("sort", merchatsParams.Sort)
	urlParam.Set("reference", merchatsParams.Reference)
	urlParam.Set("merchant_ref", merchatsParams.MerchantRef)
	urlParam.Set("method", merchatsParams.Method)
	urlParam.Set("status", merchatsParams.Status)

	paramReq := requester.IRequesterParams{
		Url:    c.BaseUrl() + "merchant/transactions?" + urlParam.Encode(),
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

	var response MerchantTransactionsResponse
	json.Unmarshal(bodyReq.ResponseBody, &response)
	return &response, nil
}
