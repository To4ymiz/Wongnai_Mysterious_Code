package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	fmt.Println(secret)
	decoded, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
}
