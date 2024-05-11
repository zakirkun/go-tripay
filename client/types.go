package client

type (
	tripayDataResponse interface {
		closePaymentTransactionOrderResponse | []feeCalcResponse | []merchantResponse | []instructionResponse | openPaymentDetailResponse | openPaymentDataResponse
	}

	tripayResponses[X tripayDataResponse] struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    X      `json:"data"`
	}

	closePaymentTransactionOrderResponse struct {
		Reference            string                                     `json:"reference"`
		MerchantRef          string                                     `json:"merchant_ref"`
		PaymentSelectionType string                                     `json:"payment_selection_type"`
		PaymentMethod        string                                     `json:"payment_method"`
		PaymentName          string                                     `json:"payment_name"`
		CustomerName         string                                     `json:"customer_name"`
		CustomerEmail        string                                     `json:"customer_email"`
		CustomerPhone        string                                     `json:"customer_phone"`
		CallbackURL          string                                     `json:"callback_url"`
		ReturnURL            string                                     `json:"return_url"`
		Amount               int                                        `json:"amount"`
		FeeMerchant          int                                        `json:"fee_merchant"`
		FeeCustomer          int                                        `json:"fee_customer"`
		TotalFee             int                                        `json:"total_fee"`
		AmountReceived       int                                        `json:"amount_received"`
		PayCode              string                                     `json:"pay_code"`
		PayURL               interface{}                                `json:"pay_url"`
		CheckoutURL          string                                     `json:"checkout_url"`
		Status               string                                     `json:"status"`
		ExpiredTime          int                                        `json:"expired_time"`
		OrderItems           []closePaymentTransactionOrderItemResponse `json:"order_items"`
		Instructions         []instructionResponse                      `json:"instructions"`
		QRString             interface{}                                `json:"qr_string"`
		QRURL                interface{}                                `json:"qr_url"`
	}

	closePaymentTransactionOrderItemResponse struct {
		SKU        string `json:"sku"`
		Name       string `json:"name"`
		Price      int    `json:"price"`
		Quantity   int    `json:"quantity"`
		Subtotal   int    `json:"subtotal"`
		ProductURL string `json:"product_url"`
		ImageURL   string `json:"image_url"`
	}

	instructionResponse struct {
		Title string   `json:"title"`
		Steps []string `json:"steps"`
	}

	feeCalcResponse struct {
		Code string
		Name string
		Fee  struct {
			Flat    int         `json:"flat"`
			Percent string      `json:"percent"`
			Min     interface{} `json:"min"`
			Max     interface{} `json:"max"`
		} `json:"fee"`
		TotalFee struct {
			Merchant int `json:"merchant"`
			Customer int `json:"customer"`
		} `json:"total_fee"`
	}

	merchantResponse struct {
		Group       string `json:"group"`
		Code        string `json:"code"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		FeeMerchant struct {
			Flat    uint `json:"flat"`
			Percent uint `json:"percent"`
		} `json:"fee_merchant"`
		FeeCustomer struct {
			Flat    uint `json:"flat"`
			Percent uint `json:"percent"`
		} `json:"fee_customer"`
		TotalFee struct {
			Flat    uint    `json:"flat"`
			Percent float64 `json:"percent"`
		} `json:"total_fee"`
		MinimumFee uint   `json:"minimum_fee"`
		MaximumFee uint   `json:"maximum_fee"`
		IconURL    string `json:"icon_url"`
		Active     bool   `json:"active"`
	}

	merchantTransactionsOrderItemResponse struct {
		SKU      interface{} `json:"sku"`
		Name     string      `json:"name"`
		Price    int         `json:"price"`
		Quantity int         `json:"quantity"`
		Subtotal int         `json:"subtotal"`
	}

	merchantTransactionsDataResponse struct {
		Reference        string                                  `json:"reference"`
		MerchantRef      string                                  `json:"merchant_ref"`
		PaymentSelection string                                  `json:"payment_selection_"`
		PaymentMethod    string                                  `json:"payment_method"`
		PaymentName      string                                  `json:"payment_name"`
		CustomerName     string                                  `json:"customer_name"`
		CustomerEmail    string                                  `json:"customer_email"`
		CustomerPhone    interface{}                             `json:"customer_phone"`
		CallbackURL      interface{}                             `json:"callback_url"`
		ReturnURL        interface{}                             `json:"return_url"`
		Amount           int                                     `json:"amount"`
		FeeMerchant      int                                     `json:"fee_merchant"`
		FeeCustomer      int                                     `json:"fee_customer"`
		TotalFee         int                                     `json:"total_fee"`
		AmountReceived   int                                     `json:"amount_received"`
		PayCode          int64                                   `json:"pay_code"`
		PayURL           interface{}                             `json:"pay_url"`
		CheckoutURL      string                                  `json:"checkout_url"`
		OrderItems       []merchantTransactionsOrderItemResponse `json:"order_items"`
		Status           string                                  `json:"status"`
		Note             interface{}                             `json:"note"`
		CreatedAt        int64                                   `json:"created_at"`
		ExpiredAt        int64                                   `json:"expired_at"`
		PaidAt           interface{}                             `json:"paid_at"`
	}

	merchantTransactionsPaginationResponse struct {
		Sort   string `json:"sort"`
		Offset struct {
			From int `json:"from"`
			To   int `json:"to"`
		} `json:"offset"`
		CurrentPage  int         `json:"current_page"`
		PreviousPage interface{} `json:"previous_page"`
		NextPage     interface{} `json:"next_page"`
		LastPage     int         `json:"last_page"`
		PerPage      int         `json:"per_page"`
		TotalRecords int         `json:"total_records"`
	}

	merchantTransactionsResponse struct {
		Success    bool                                   `json:"success"`
		Message    string                                 `json:"message"`
		Data       []merchantTransactionsDataResponse     `json:"data"`
		Pagination merchantTransactionsPaginationResponse `json:"pagination"`
	}

	openPaymentDetailResponse struct {
		UUID          string `json:"uuid"`
		MerchantRef   string `json:"merchant_ref"`
		CustomerName  string `json:"customer_name"`
		PaymentName   string `json:"payment_name"`
		PaymentMethod string `json:"payment_method"`
		PayCode       string `json:"pay_code"`
		QRString      string `json:"qr_string,omitempty"`
		QRURL         string `json:"qr_url,omitempty"`
	}

	openPaymentDataResponse struct {
		UUID          string `json:"uuid"`
		MerchantRef   string `json:"merchant_ref"`
		CustomerName  string `json:"customer_name"`
		PaymentName   string `json:"payment_name"`
		PaymentMethod string `json:"payment_method"`
		PayCode       string `json:"pay_code"`
		QRString      string `json:"qr_string"`
		QRURL         string `json:"qr_url"`
	}

	openPaymentListResponse struct {
		Success                bool   `json:"success"`
		Message                string `json:"message"`
		OpenPaymentTransaction []struct {
			Reference      string `json:"reference"`
			MerchantRef    string `json:"merchant_ref"`
			PaymentMethod  string `json:"payment_method"`
			PaymentName    string `json:"payment_name"`
			CustomerName   string `json:"customer_name"`
			Amount         int    `json:"amount"`
			FeeMerchant    int    `json:"fee_merchant"`
			FeeCustomer    int    `json:"fee_customer"`
			TotalFee       int    `json:"total_fee"`
			AmountReceived int    `json:"amount_received"`
			CheckoutURL    string `json:"checkout_url"`
			Status         string `json:"status"`
			PaidAt         int64  `json:"paid_at"`
		} `json:"data"`
		Pagination struct {
			Total       int  `json:"total"`
			DataFrom    int  `json:"data_from"`
			DataTo      int  `json:"data_to"`
			PerPage     int  `json:"per_page"`
			CurrentPage int  `json:"current_page"`
			LastPage    int  `json:"last_page"`
			NextPage    *int `json:"next_page"`
		} `json:"pagination"`
	}
)
