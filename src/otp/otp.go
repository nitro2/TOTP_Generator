package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
)

type Author struct {
	Secret string
	Epoch  int64
}

func GoogleAuthenticatorCode(author Author) (string, error) {
	// fmt.Println("GoogleAuthenticatorCode ")

	// Padding "=" into base32 string
	if n := len(author.Secret); n%8 != 0 {
		author.Secret += strings.Repeat("=", 8-(n%8))
	}

	// Decode the secret from base32 to original
	decodedKey, err := base32.StdEncoding.DecodeString(strings.ToUpper(author.Secret))
	if err != nil {
		fmt.Errorf("Cannot decode secret %v", err)
		return "", err
	}
	// fmt.Println(decodedKey)

	// hashed message is 8 bytes of (epoch/30)
	rawMes := make([]byte, 8)
	binary.BigEndian.PutUint64(rawMes, uint64(author.Epoch/30))
	// fmt.Printf("rawMes=%v \n", rawMes)

	//Signing the value using HMAC-SHA1 Algorithm
	hash := hmac.New(sha1.New, decodedKey)
	hash.Write(rawMes)
	hashMes := hash.Sum(nil)
	// Get last nibble of hash message
	offset := int(hashMes[19] & 0x0f)
	// fmt.Printf("hashMes=%v offset=%v\n", hashMes, offset)
	// Get 4 bytes offset
	startingMes := hashMes[offset : offset+4]
	// fmt.Println(startingMes)
	// Remove the most significant bit
	num := binary.BigEndian.Uint32(startingMes)
	code := (num & 0x7fffffff) % 1000000

	// fmt.Printf("num=%v code=%v\n", num, code)
	return fmt.Sprintf("%06d", code), nil

}
