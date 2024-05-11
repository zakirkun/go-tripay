# Tripay Go (SDK tripay written in go)
## installation
```sh
go get github.com/zakirkun/go-tripay
```
## How To Use?

### register client
```go
c := client.Client{
    MerchantCode: "T14302",
    ApiKey:       "your_api_key_here",
    PrivateKey:   "your_private_key_here",
    Mode:         utils.MODE_DEVELOPMENT,
}
```

### payment instruction
[docs](https://tripay.co.id/developer?tab=payment-instruction)
```go
	req := client.InstructionRequestParam{
		ChannelCode: utils.CHANNEL_BRIVA,
		PayCode:     "",
		Amount:      "10000",
		AllowHtml:   "",
	}

	response, err := c.Instruction(req)
	if err != nil {
		// do something
	}
        // do something
```

### payment channel
[docs](https://tripay.co.id/developer?tab=merchant-payment-channel)
```go
	response, err := c.MerchantPay()
	if err != nil {
        // do something
	}
        // do something
```

### fee calculation
[docs](https://tripay.co.id/developer?tab=merchant-fee-calculator)
```go

	feeCalcParam := client.FeeCalcParam{
		Amount: 100000,
		Code:   utils.CHANNEL_ALFAMIDI,
	}

	response, err := c.FeeCalc(feeCalcParam)
	if err != nil {
        // do something
	}
        // do something
```

### list transactions
[docs](https://tripay.co.id/developer?tab=merchant-transactions)
```go
	merchanTransactionParam := client.MerchantTransactionsParam{
        Page:        1,
		PerPage:     10,
		Sort:        "asc", // asc or desc
		Reference:   "reference",
		MerchantRef: "merchant_ref",
		Method:      utils.CHANNEL_BCAVA,
		Status: "status",
	}

	response, err := c.MerchantTransactions(merchanTransactionParam)
	if err != nil {
        // do something
	}
        // do something
```

## close payment
### request transaction
[docs](https://tripay.co.id/developer?tab=transaction-create)
```go
	signStr := utils.Signature{
		Amount:       50000,
		PrivateKey:   "your_private_key_here",
		MerchantCode: "T14302",
		MerchanReff:  "INV345675",
	}

	c.SetSignature(signStr)

	bodyReq := client.ClosePaymentBodyRequest{
		Method:        utils.CHANNEL_QRIS_SHOPEEPAY,
		MerchantRef:   "INV345675",
		Amount:        50000,
		CustomerName:  "Farda Ayu Nurfatika",
		CustomerEmail: "fardaayu@gmail.com",
		CustomerPhone: "6285111990223",
		ReturnURL:     "https://thisisreturnurl.com/redirect",
		ExpiredTime:   client.SetTripayExpiredTime(24), // 24 Hour
		Signature:     c.GetSignature(),
		OrderItems: []client.OrderItemClosePaymentRequest{
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

	response, err := c.ClosePaymentRequestTransaction(bodyReq)
	if err != nil {
		// do something
	}
        // do something
```
### get detail transaction
[docs](https://tripay.co.id/developer?tab=transaction-detail)
```go
	signStr := utils.Signature{
		Amount:       50000,
		PrivateKey:   "your_private_key_here",
		MerchantCode: "T14302",
		MerchanReff:  "INV345675",
	}

	c.SetSignature(signStr)

	referenceId := "your_reference_id"
	response, err := c.ClosePaymentTransactionGetDetail(referenceId)
	if err != nil {
        // do something
	}
        // do something
```

## open payment
```go
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

	response, err := c.OpenPaymentTransaction(payment)
	if err != nil {
		// do something
	}
        // do something
```