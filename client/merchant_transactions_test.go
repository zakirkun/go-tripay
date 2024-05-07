package client

import (
	"context"
	"testing"
	"time"

	"github.com/zakirkun/go-tripay/utils"
)

func TestMerchantTransactionsWithContextSuccess(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	ctx, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	reponseOk, err := client.MerchantTransactionsWithContext(ctx)
	if err != nil {
		t.FailNow()
	}

	t.Log(reponseOk)
}

func TestMerchantTransactionsWithContextFailed(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSz",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	ctx, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	reponseOk, err := client.MerchantTransactionsWithContext(ctx)
	if err != nil || reponseOk.Success {
		t.Log(reponseOk)
		t.FailNow()
	}
}
