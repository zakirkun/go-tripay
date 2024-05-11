package examples

import (
	"fmt"

	"github.com/zakirkun/go-tripay/client"
	"github.com/zakirkun/go-tripay/utils"
)

func Example_payment_instructions() {
	c := client.Client{
		MerchantCode: "T14302",
		ApiKey:       "your_api_key_here",
		PrivateKey:   "your_private_key_here",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	ip := client.InstructionRequestParam{
		ChannelCode: utils.CHANNEL_BRIVA,
		PayCode:     "",
		Amount:      "10000",
		AllowHtml:   "",
	}

	response, err := c.Instruction(ip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v\n", response)
}
