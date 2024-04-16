package utils

import (
	"encoding/json"
	"fmt"
)

func JsonToStruct(data []byte, to any) {
	err := json.Unmarshal(data, &to)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
