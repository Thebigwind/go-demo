package common

import (
	"encoding/base64"
	"fmt"
	gmx509 "github.com/tjfoc/gmsm/x509"
)

func DelStringTailNull(in string) string {
	inBytes := []byte(in)
	var outBytes []byte
	for i := len(inBytes) - 1; i >= 0; i-- {
		if inBytes[i] == 0 {
			outBytes = inBytes[0:i]
		} else if i == len(inBytes)-1 {
			outBytes = inBytes
			break
		} else {
			break
		}
	}
	return string(outBytes)
}

func CompareSm2Pubkey(pub1 string, pub2 string) bool {
	pub1 = GetSm2PubkeyB64(pub1)
	pub2 = GetSm2PubkeyB64(pub2)

	return pub1 == pub2
}

func GetSm2PubkeyB64(pubkey string) string {
	fmt.Printf("len(pubkey) ori:%v\n", len(pubkey))
	pubkey = DelStringTailNull(pubkey)
	fmt.Printf("len(pubkey):%v\n", len(pubkey))
	sm2PublicKey, err := gmx509.ReadPublicKeyFromPem([]byte(pubkey))
	if err != nil {
		fmt.Printf("gmx509.ReadPublicKeyFromPem err:%v", err)
		return pubkey
	}
	ecdPubKeyBytes, err := gmx509.MarshalPKIXPublicKey(sm2PublicKey)
	if err != nil {
		fmt.Printf("gmx509.MarshalPKIXPublicKey err2:%v", err)
		return pubkey
	}
	inputPubKeyB64 := base64.StdEncoding.EncodeToString(ecdPubKeyBytes)
	fmt.Println(inputPubKeyB64)
	return inputPubKeyB64
}
