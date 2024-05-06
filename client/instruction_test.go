package client

import (
	"testing"

	"github.com/zakirkun/go-tripay/utils"
)

func TestInstruction(t *testing.T) {

	client := Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         utils.URL_DEVELOPMENT,
	}

	reponseOk, responseBad := client.Instruction("BRIVA", "", "10000", "")
	if responseBad != nil {
		t.Errorf("ERROR: %v", responseBad)
	}

	t.Log("Success: ", reponseOk)
}
