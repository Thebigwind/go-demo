package common

import (
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm3"
)

//sm3摘要
func EncryptSM3(msg string) string {
	hash := sm3.New()
	hash.Write([]byte(msg))
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
