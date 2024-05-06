package client

import "github.com/zakirkun/go-tripay/utils"

type Client struct {
	MerchantCode string
	ApiKey       string
	PrivateKey   string
	Mode         utils.TRIPAY_MODE
}

func (c Client) HeaderRequest() []map[string]string {
	// Define your headers as a slice of maps
	headers := []map[string]string{
		{"Content-Type": "application/json"},
		{"Authorization": "Bearer " + c.ApiKey},
	}
	return headers
}

func (c Client) BaseUrl() string {
	if c.Mode == utils.MODE_DEVELOPMENT {
		return string(utils.URL_DEVELOPMENT)
	}

	return string(utils.URL_PRODUCTION)
}
