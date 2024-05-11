package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/zakirkun/go-tripay/internal/requester"
	"github.com/zakirkun/go-tripay/utils"
)

type (
	MerchantTransactionsParam struct {
		Page        int
		PerPage     int
		Sort        string
		Reference   string
		MerchantRef string
		Method      utils.TRIPAY_CHANNEL
		Status      string
	}
)

/*
used to get a list of merchant transactions. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	response, err := c.MerchantTransactions()
	if err != nil{
		// do something
	}
	// do something
*/
func (c Client) MerchantTransactions(p ...MerchantTransactionsParam) (*merchantTransactionsResponse, error) {
	return merchantTransactions(c, nil, p...)
}

/*
used to get a list of merchant transactions. Example:

	c := Client{ MerchantCode: "T14302", ApiKey: "your_api_key", PrivateKey: "your_private_key", Mode: utils.MODE_DEVELOPMENT }
	response, err := c.MerchantTransactionsWithContext(context.Background())
	if err != nil{
		// do something
	}
	// do something
*/
func (c Client) MerchantTransactionsWithContext(ctx context.Context, p ...MerchantTransactionsParam) (*merchantTransactionsResponse, error) {
	return merchantTransactions(c, ctx, p...)
}

func merchantTransactions(c Client, ctx context.Context, p ...MerchantTransactionsParam) (*merchantTransactionsResponse, error) {
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
	urlParam.Set("method", string(merchatsParams.Method))
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

	var response merchantTransactionsResponse
	json.Unmarshal(bodyReq.ResponseBody, &response)
	return &response, nil
}
