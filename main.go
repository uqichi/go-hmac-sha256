package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
)

func main() {
	var (
		message   = flag.String("msg", "", "")
		secretKey = flag.String("key", "", "")
	)
	flag.Parse()
	if len(*message) == 0 {
		panic("message is empty")
	}
	if len(*secretKey) == 0 {
		panic("secret key is empty")
	}
	fmt.Println(makeHMAC(*message, *secretKey))
}

func makeHMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}
