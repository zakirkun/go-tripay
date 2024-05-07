package client

import (
	"context"
	"testing"

	"github.com/zakirkun/go-tripay/utils"
)

func TestFeeCalcWithContextSuccess(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	feeCalcParam := FeeCalcParam{
		Code:   utils.CHANNEL_ALFAMIDI,
		Amount: 100000,
	}
	response, err := client.FeeCalcWithContext(context.Background(), feeCalcParam)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	t.Log("Success: ", response)
}

func TestFeeCalcWithContextFail(t *testing.T) {
	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	feeCalcParam := FeeCalcParam{
		Code:   "salah",
		Amount: 100000,
	}
	response, err := client.FeeCalcWithContext(context.Background(), feeCalcParam)
	if err != nil {
		t.FailNow()
	}

	if response.Success != false {
		t.FailNow()
	}

}
