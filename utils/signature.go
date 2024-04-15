package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type ISignature interface {
	CreateSignature() string
}

type Signature struct {
	Amount       int64
	PrivateKey   string
	MerchantCode string
	MerchanReff  string
	Channel      string
}

func (s *Signature) CreateSignature() string {
	var signStr string
	if s.Amount != 0 {
		signStr = s.MerchantCode + s.MerchanReff + fmt.Sprint(s.Amount)
	} else {
		signStr = s.MerchantCode + s.Channel + s.MerchanReff
	}

	key := []byte(s.PrivateKey)
	message := []byte(signStr)

	hash := hmac.New(sha256.New, key)
	hash.Write(message)
	signature := hex.EncodeToString(hash.Sum(nil))

	return signature
}
