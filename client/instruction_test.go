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
		Mode:         utils.MODE_DEVELOPMENT,
	}

	ip := InstructionRequestParam{
		ChannelCode: utils.CHANNEL_BRIVA,
		PayCode:     "",
		Amount:      "10000",
		AllowHtml:   "",
	}
	response, err := client.Instruction(ip)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	t.Log("Success: ", response)
}
