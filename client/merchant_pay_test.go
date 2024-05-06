package client

import (
	"context"
	"testing"
	"time"

	"github.com/zakirkun/go-tripay/utils"
)

func TestMerchantPay(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.URL_DEVELOPMENT,
	}

	reponseOk, responseBad := client.MerchantPay()
	if responseBad != nil {
		t.Errorf("ERROR: %v", responseBad)
	}

	t.Log("Success: ", reponseOk)
}

func TestMerchantPayWithCtx(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.URL_DEVELOPMENT,
	}

	ctx, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	reponseOk, responseBad := client.MerchantPayWithContext(ctx)
	if responseBad != nil {
		t.Errorf("ERROR: %v", responseBad)
	}

	t.Log("Success: ", reponseOk)
}
