package client

import (
	"testing"

	"github.com/zakirkun/go-tripay/utils"
)

func TestOpenPaymentSuccess(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	client.SetSignature(utils.Signature{
		MerchantCode: "T14302",
		Channel:      "BRIVA",
		MerchanReff:  "INV345678",
	})

	payment := OpenPaymentRequest{
		Method:       "BRIVA",
		MerchatReff:  "INV345678",
		CustomerName: "Fulan",
	}

	responseOk, responseBad := client.OpenPaymentTransaction(payment)
	if responseBad != nil {
		t.Errorf("ERROR: %v", responseBad)
	}

	t.Log("Success: ", responseOk)
}
