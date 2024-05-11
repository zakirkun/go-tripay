package examples

import (
	"fmt"

	"github.com/zakirkun/go-tripay/client"
	"github.com/zakirkun/go-tripay/utils"
)

func Example_open_payment_create() {

	c := client.Client{
		MerchantCode: "T14302",
		ApiKey:       "your_api_key_here",
		PrivateKey:   "your_private_key_here",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	c.SetSignature(utils.Signature{
		MerchantCode: "T14302",
		Channel:      "BCAVA",
		MerchanReff:  "INV345675",
	})

	payment := client.OpenPaymentRequest{
		Method:       "BCAVA",
		MerchatReff:  "INV345675",
		CustomerName: "Fulan Fulan",
		Signature:    c.GetSignature(),
	}

	responseOk, responseBad := c.OpenPaymentTransaction(payment)
	if responseBad != nil {
		fmt.Errorf("ERROR: %v", responseBad)
	}

	fmt.Printf("Success: ", responseOk)
}
