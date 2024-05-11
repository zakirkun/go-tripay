package examples

import (
	"fmt"

	"github.com/zakirkun/go-tripay/client"
	"github.com/zakirkun/go-tripay/utils"
)

func Example_merchant_fee_calculator() {
	c := client.Client{
		MerchantCode: "T14302",
		ApiKey:       "your_api_key_here",
		PrivateKey:   "your_private_key_here",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	feeCalcParam := client.FeeCalcParam{
		Amount: 100000,
		Code:   utils.CHANNEL_ALFAMIDI,
	}

	response, err := c.FeeCalc(feeCalcParam)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %v\n", response)
	/*
		response:
		&{true  [{ALFAMIDI Alfamidi {3500 0.00 <nil> <nil>} {3500 0}}]}
	*/
}
