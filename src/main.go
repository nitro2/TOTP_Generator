package main

import (
	"fmt"

	"otp"
)

func main() {
	o, _ := otp.GoogleAuthenticatorCode(otp.Author{Secret: "dummySECRETdummy", Epoch: 1523822557})
	fmt.Println(o)
}
