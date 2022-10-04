package main

import (
	"fmt"
	"os"
	"time"

	"otp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Missing secret? Need exactly 1 argument.")
		os.Exit(1)
	}
	secret := os.Args[1]
	if o, err := otp.GoogleAuthenticatorCode(otp.Author{Secret: secret, Epoch: time.Now().Unix()}); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	} else {
		fmt.Println(o)
	}
}
