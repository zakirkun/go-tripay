package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func MakeSignature(amount int64, merchantCode, privateKey, merchantRef string) string {
	key := []byte(privateKey)
	message := []byte(merchantCode + merchantRef + fmt.Sprint(amount))

	hasher := hmac.New(sha256.New, key)
	hasher.Write(message)
	signature := hex.EncodeToString(hasher.Sum(nil))

	return signature
}

func MakeSignatureOpen(merchantCode, privateKey, merchantRef, channel string) string {
	key := []byte(privateKey)
	message := []byte(merchantCode + channel + merchantRef)

	hasher := hmac.New(sha256.New, key)
	hasher.Write(message)
	signature := hex.EncodeToString(hasher.Sum(nil))

	return signature
}

func JsonToStruct(data []byte, to any) {
	err := json.Unmarshal(data, &to)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
