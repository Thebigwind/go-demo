package demo041

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

//hmac

func GenerateHMAC(text string, key string) string {

	textBytes := []byte(text)
	keyBytes := []byte(key)

	hash := hmac.New(sha256.New, keyBytes)

	hash.Write(textBytes)

	result := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)

}

func VerifyHMAC(HMAC string, text string, key string) bool {

	HMACBytes := []byte(HMAC)

	nowHMAC := GenerateHMAC(text, key)
	nowHMACBytes := []byte(nowHMAC)

	return hmac.Equal(HMACBytes, nowHMACBytes)
}
