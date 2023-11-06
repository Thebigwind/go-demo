package main

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	gmx509 "github.com/tjfoc/gmsm/x509"
	"reflect"
	"strings"
	//gmx509 "github.com/tjfoc/gmsm/x509"
)

func main() {
	//test2()
	a := strings.TrimSuffix("zdlz-id_public_key", "_public_key")

	fmt.Printf("a:%s", a)
}
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

func test2() {
	publicKey2PEM := `
-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEpxRo3g8ZHcYfpIlcmGzec2/VEjbH
JTn0q1kxCCgQTDbqtCWOs7QU5b9Gu2pBWQi9AGJuCfZUPn14Ve0DI+QuEg==
-----END PUBLIC KEY-----

`

	sm2PublicKey, err := gmx509.ReadPublicKeyFromPem([]byte(DelStringTailNull(publicKey2PEM)))
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	ecdPubKeyBytes, err := gmx509.MarshalPKIXPublicKey(sm2PublicKey)
	if err != nil {
		fmt.Printf("err2:%v", err)
		return
	}
	inputPubKeyB64 := base64.StdEncoding.EncodeToString(ecdPubKeyBytes)
	fmt.Println(inputPubKeyB64)
}

//请注意，这里使用了 crypto 和 encoding/pem 包来解码和比较两个公钥。你需要替换示例中的 publicKey1PEM 和 publicKey2PEM 字符串为你要比较的两个公钥的 PEM 编码字符串。然后，程序会输出是否这两个公钥相同。
func test1() {
	// 两个公钥的 PEM 编码字符串示例
	publicKey1PEM := `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxZDVIEJlE1j4yJKG1x6H
OJrI7A6LXa1V3YR8BrsqBdHto6DCv2aR6qtSYPvYtWzRP4xweoy6HlK4Mzf3OL7N
xI8ytr9tT9lV0WsYudJWt3QfnCNS4NTm1iGAE70qGd/N9s8hs1nQ9Tx+jEXvVdGy
7lAsyZ4DgKmwU9n5K1RBr0s/SPAd0wa6NT0zzd4cxd2/GPKU+TlZ5uhfxsZ3vgs9
6SwAK89g80WJc2k6cnJYQDpYGiD6ikq0YAR9gprnncalzQW0TbTjC9qTRpAsXUQs
x+xzrzt2BcQ6GvXvNnIzRyY+Vwnj1F2/jkLlR18xIq+azrIqQFl2k9sdMRE4OjwH
JwIDAQAB
-----END PUBLIC KEY-----
`

	publicKey2PEM := `
-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEpxRo3g8ZHcYfpIlcmGzec2/VEjbH
JTn0q1kxCCgQTDbqtCWOs7QU5b9Gu2pBWQi9AGJuCfZUPn14Ve0DI+QuEg==
-----END PUBLIC KEY-----

`

	// 解码 PEM 编码的公钥
	block1, _ := pem.Decode([]byte(publicKey1PEM))
	block2, _ := pem.Decode([]byte(publicKey2PEM))

	// 解析公钥
	pubKey1, err1 := x509.ParsePKIXPublicKey(block1.Bytes)
	pubKey2, err2 := x509.ParsePKIXPublicKey(block2.Bytes)

	if err1 != nil || err2 != nil {
		fmt.Println("公钥解析失败")
		fmt.Printf("err2:%v", err2)
		return
	}

	// 比较两个公钥是否相同
	if reflect.DeepEqual(pubKey1, pubKey2) {
		fmt.Println("两个公钥相同")
	} else {
		fmt.Println("两个公钥不同")
	}
}
