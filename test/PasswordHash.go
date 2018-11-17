package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func passwordhash(password string) string {
	passbyte := []byte(password)
	passhash := sha256.New()
	passhash.Write(passbyte)
	phencoded := base64.URLEncoding.EncodeToString(passhash.Sum(nil))
	return phencoded
}

func main() {
	fmt.Println(passwordhash("tester"))
}

