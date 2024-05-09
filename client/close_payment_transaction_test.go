package client

import (
	"context"
	"testing"

	"github.com/zakirkun/go-tripay/utils"
)

func TestClosePaymentTransactionRequestSuccess(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	s := utils.Signature{
		Amount:       10,
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		MerchantCode: "T0001",
		MerchanReff:  "INV345675",
	}

	bodyReq := ClosePaymentBodyRequest{
		Method:        utils.CHANNEL_BCAVA,
		MerchantRef:   "INV345675",
		Amount:        50000,
		CustomerName:  "John Doe",
		CustomerEmail: "johndoe@gmail.com",
		CustomerPhone: "62891829828",
		ReturnURL:     "https://thisisreturnurl.com/redirect",
		ExpiredTime:   SetTripayExpiredTime(24), // 24 Hour
		Signature:     s.CreateSignature(),
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
	reponseOk, responseBad := client.ClosePaymentRequestTransaction(context.Background(), bodyReq)
	if responseBad != nil {
		t.Errorf("ERROR: %v", responseBad)
	}

	t.Log("Success: ", reponseOk)
}
