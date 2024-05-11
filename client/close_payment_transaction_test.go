package client

import (
	"testing"

	"github.com/zakirkun/go-tripay/utils"
)

func TestClosePaymentTransactionRequest(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	client.SetSignature(utils.Signature{
		Amount:       50000,
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		MerchantCode: "T14302",
		MerchanReff:  "INV345675",
	})

	bodyReq := ClosePaymentBodyRequest{
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
	response, err := client.ClosePaymentRequestTransaction(bodyReq)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	t.Log("Success: ", response)
}

func TestClosePaymentTransactionGetTransaction(t *testing.T) {
	referenceId := "DEV-T14302154794FZ2ZT"

	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	response, err := client.ClosePaymentTransactionGetDetail(referenceId)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	t.Log("Success: ", response)
}
