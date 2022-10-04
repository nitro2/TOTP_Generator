package otp

import (
	"encoding/base32"
	"fmt"
)

type Author struct {
	Secret string
	Epoch  int
}

func GoogleAuthenticatorCode(author Author) string {
	decodedKey := base32.StdEncoding.EncodeToString([]byte(author.Secret))
	fmt.Println(decodedKey)
	return decodedKey
}
