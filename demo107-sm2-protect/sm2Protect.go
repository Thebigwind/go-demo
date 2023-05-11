package main

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/base64"
	"errors"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
	"math/big"
	"math/rand"
)

//本文主要介绍ASN1对密钥数据格式SM2PrivateKey、SM2PublicKey；加密数据格式SM2Cipher；签名数据格式SM2Signature； 密钥对保护数据格式SM2EnvelopedKey等四中类型。
//前三种都比较容易实现，本文主要介绍密钥对保护数据的封装和解析！！！

var oidEncryptionAlgorithmSM4ECB = asn1.ObjectIdentifier{1, 2, 156, 10197, 1, 104, 1}

// 不能使用代码库中的结构，库中对外不可见，asn1编码会失败
type Sm2Cipher struct {
	XCoordinate *big.Int
	YCoordinate *big.Int
	HASH        []byte
	CipherText  []byte
}

type SM2EnvelopedKey struct {
	SymAlgID               pkix.AlgorithmIdentifier
	Sm2CipherSymKey        Sm2Cipher
	Sm2PublicKey           asn1.BitString
	Sm2EncryptedPrivateKey asn1.BitString
}

func SM2ExchangePubKeyToBytes(pKey *sm2.PublicKey) []byte {
	xBuf := pKey.X.Bytes()
	yBuf := pKey.Y.Bytes()
	if n := len(xBuf); n < 32 {
		xBuf = append(zeroByteSlice()[:32-n], xBuf...)
	}
	if n := len(yBuf); n < 32 {
		yBuf = append(zeroByteSlice()[:32-n], yBuf...)
	}
	return append(append([]byte{0x04}, xBuf...), yBuf...)
}

// 私钥结构转切片
func SM2ExchangePriKeyToBytes(pKey *sm2.PrivateKey) []byte {
	dBuf := pKey.D.Bytes()
	if n := len(dBuf); n < 32 {
		dBuf = append(zeroByteSlice()[:32-n], dBuf...)
	}
	return dBuf
}

type Sm2Key interface{}

// SM2密钥结构转字节串（统一入口）
func SM2ExchangeKeyToBytes(key Sm2Key) []byte {
	switch pkey := key.(type) {
	case *sm2.PublicKey:
		return SM2ExchangePubKeyToBytes(pkey)
	case *sm2.PrivateKey:
		return SM2ExchangePriKeyToBytes(pkey)
	default:
		return []byte(nil)
	}
}
func zeroByteSlice() []byte {
	return []byte{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
}

// SM2字节串转密钥结构
// 在调用结束后需要使用类型转换，转成相对应的结构，如data为65字节公钥
// sm2Key := SM2ExchangeBytesToKey(data)
// pubKey := sm2Key.(*sm2.PublicKey)
func SM2ExchangeBytesToKey(data []byte) Sm2Key {
	dataLen := len(data)
	if dataLen != 32 && dataLen != 64 && dataLen != 65 {
		return nil
	}
	switch dataLen {
	case 32:
		key := new(sm2.PrivateKey)
		c := sm2.P256Sm2()
		key.D = new(big.Int).SetBytes(data)
		key.PublicKey.Curve = c
		key.PublicKey.X, key.PublicKey.Y = c.ScalarBaseMult(key.D.Bytes())
		return key
	case 65:
		if data[0] != 0x04 {
			return nil
		}
		data = data[1:]
		fallthrough
	case 64:
		key := new(sm2.PublicKey)
		key.Curve = sm2.P256Sm2()
		var x, y big.Int
		x.SetBytes(data[0:32])
		y.SetBytes(data[32:64])
		key.X = &x
		key.Y = &y
		return key
	default:
		return nil
	}
}

// 封装SM2密钥保护数据
// GMT 0009-2012
func SM2EnvelopedKeyPairMarshal(pubKey *sm2.PublicKey, projectCert *x509.Certificate, projectPkey *sm2.PrivateKey) ([]byte, error) {
	var key = make([]byte, 16)
	var sm2Enveloped SM2EnvelopedKey
	// 1、产生对称密钥
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	// 2、将要保护的SM2私钥转为bytes
	projectPkeyBytes := SM2ExchangeKeyToBytes(projectPkey)
	if projectPkeyBytes == nil || len(projectPkeyBytes) != 32 {
		return nil, errors.New("project private key error")
	}
	// 3、获得私钥的对称密文
	block, err := sm4.NewCipher(key)
	bs := block.BlockSize()
	projectEncryptedKey := make([]byte, len(projectPkeyBytes))
	for i := 0; i < len(projectPkeyBytes)/bs; i++ {
		block.Encrypt(projectEncryptedKey[bs*i:bs*(i+1)], projectPkeyBytes[bs*i:bs*(i+1)])
	}
	// 4、外部公钥加密对称密钥
	keyEncrypted, err := sm2.Encrypt(pubKey, key, nil, sm2.C1C3C2)
	if err != nil {
		return nil, err
	}
	// 5、设置密钥对保护数据中的公钥
	projectPubKey, err := CERTGeneratePubKeyFromObjs(projectCert)
	if err != nil {
		return nil, err
	}
	// 6、设置数据到结构中
	sm2Enveloped.SymAlgID.Algorithm = oidEncryptionAlgorithmSM4ECB
	sm2Enveloped.Sm2CipherSymKey = *sm2CipherMarshal(keyEncrypted)
	sm2Enveloped.Sm2PublicKey.Bytes = SM2ExchangePubKeyToBytes(projectPubKey.(*sm2.PublicKey))
	sm2Enveloped.Sm2EncryptedPrivateKey.Bytes = projectEncryptedKey

	return asn1.Marshal(sm2Enveloped)
}

func SM2EnvelopedKeyPairMarshal2(projectCert *x509.Certificate, projectPkey *sm2.PrivateKey) ([]byte, error) {
	var key = make([]byte, 16)
	var sm2Enveloped SM2EnvelopedKey
	// 1、产生对称密钥
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	// 2、将要保护的SM2私钥转为bytes
	projectPkeyBytes := SM2ExchangeKeyToBytes(projectPkey)
	if projectPkeyBytes == nil || len(projectPkeyBytes) != 32 {
		return nil, errors.New("project private key error")
	}
	// 3、获得私钥的对称密文
	block, err := sm4.NewCipher(key)
	bs := block.BlockSize()
	projectEncryptedKey := make([]byte, len(projectPkeyBytes))
	for i := 0; i < len(projectPkeyBytes)/bs; i++ {
		block.Encrypt(projectEncryptedKey[bs*i:bs*(i+1)], projectPkeyBytes[bs*i:bs*(i+1)])
	}
	// 4、外部公钥加密对称密钥
	//keyEncrypted, err := sm2.Encrypt(pubKey, key, nil, sm2.C1C3C2)
	//if err != nil {
	//	return nil, err
	//}
	// 5、设置密钥对保护数据中的公钥
	projectPubKey, err := CERTGeneratePubKeyFromObjs(projectCert)
	if err != nil {
		return nil, err
	}
	// 6、设置数据到结构中
	sm2Enveloped.SymAlgID.Algorithm = oidEncryptionAlgorithmSM4ECB
	//sm2Enveloped.Sm2CipherSymKey = *sm2CipherMarshal(keyEncrypted)
	sm2Enveloped.Sm2PublicKey.Bytes = SM2ExchangePubKeyToBytes(projectPubKey.(*sm2.PublicKey))
	sm2Enveloped.Sm2EncryptedPrivateKey.Bytes = projectEncryptedKey

	return asn1.Marshal(sm2Enveloped)
}

func sm2CipherMarshal(cipherData []byte) *Sm2Cipher {
	var cipher Sm2Cipher
	cipher.XCoordinate = new(big.Int)
	cipher.XCoordinate.SetBytes(cipherData[1:33])
	cipher.YCoordinate = new(big.Int)
	cipher.YCoordinate.SetBytes(cipherData[33:65])
	cipher.HASH = cipherData[65:97]
	cipher.CipherText = cipherData[97:]
	return &cipher
}

func CERTGeneratePubKeyFromObjs(cert *x509.Certificate) (crypto.PublicKey, error) {
	switch cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		ecdSa := cert.PublicKey.(*ecdsa.PublicKey)
		return &sm2.PublicKey{
			Curve: ecdSa.Curve,
			X:     ecdSa.X,
			Y:     ecdSa.Y,
		}, nil
	case *rsa.PublicKey:
		return cert.PublicKey.(*rsa.PublicKey), nil
	default:
		return nil, errors.New("unknown this cert")
	}
}

// 解封装SM2密钥对保护数据格式
// GMT 0009-2012
func SM2EnvelopedKeyPairUnmarshal(pkey *sm2.PrivateKey, encCert, encData []byte) (*sm2.PrivateKey, error) {

	// 1、解码ANS.1编码的SM2密钥对的保护数据
	var sm2Enveloped SM2EnvelopedKey
	_, err := asn1.Unmarshal(encData, &sm2Enveloped)
	if err != nil {
		return nil, err
	}
	// 2、 校验OID
	if !sm2Enveloped.SymAlgID.Algorithm.Equal(oidEncryptionAlgorithmSM4ECB) {
		return nil, errors.New("enveloped data oid error")
	}
	// 3、 解码ASN.1编码的SM2信封的对称密钥的密文
	cipherData := sm2CipherUnMarshal(&sm2Enveloped.Sm2CipherSymKey)
	// 4、 SM2解密对称密钥的密文
	symKey, err := sm2.Decrypt(pkey, cipherData, sm2.C1C3C2)
	if err != nil {
		return nil, err
	}
	// 5、 SM4解密加密私钥
	if sm2Enveloped.Sm2EncryptedPrivateKey.BitLength != 32*8 {
		return nil, errors.New("enveloped encrypted private key length error")
	}
	pKeyEncData := sm2Enveloped.Sm2EncryptedPrivateKey.Bytes
	block, err := sm4.NewCipher(symKey)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	decKeyBytes := make([]byte, len(pKeyEncData))
	for i := 0; i < len(pKeyEncData)/bs; i++ {
		block.Decrypt(decKeyBytes[bs*i:bs*(i+1)], pKeyEncData[bs*i:bs*(i+1)])
	}

	if len(decKeyBytes) != 32 {
		return nil, errors.New("decrypt enc private key length error")
	}
	// 6、获得私钥对象
	decKey := SM2ExchangeBytesToKey(decKeyBytes)
	if decKey == nil {
		return nil, errors.New("sm2 key bytes to exchange error")
	}
	priKey := decKey.(*sm2.PrivateKey)
	// 7、验证保护密钥对数据中的密钥对是否匹配
	if sm2Enveloped.Sm2PublicKey.BitLength != 65*8 {
		return nil, errors.New("enveloped public key bytes length error")
	}
	pubKey := SM2ExchangeBytesToKey(sm2Enveloped.Sm2PublicKey.Bytes)
	_, ok := pubKey.(*sm2.PublicKey)
	if !ok {
		return nil, errors.New("sm2 key pair public key error")
	}
	ok = SM2KeyPairCheck(pubKey.(*sm2.PublicKey), priKey)
	if !ok {
		return nil, errors.New("check key pair error")
	}
	// 8、验证生成的证书和保护数据中的公钥是否一致
	if encCert != nil {
		cert, err := CERTGeneratePubKeyFromBytes(encCert)
		if err != nil {
			return nil, err
		}
		ok = SM2KeyPairCheck(cert, priKey)
		if !ok {
			return nil, errors.New("check key pair error")
		}
	}

	return priKey, nil
}

func sm2CipherUnMarshal(cipher *Sm2Cipher) []byte {

	x := cipher.XCoordinate.Bytes()
	y := cipher.YCoordinate.Bytes()
	hash := cipher.HASH

	cipherText := cipher.CipherText

	if n := len(x); n < 32 {
		x = append(zeroByteSlice()[:32-n], x...)
	}
	if n := len(y); n < 32 {
		y = append(zeroByteSlice()[:32-n], y...)
	}
	var c []byte
	c = append(c, x...)          // x分量
	c = append(c, y...)          // y分量
	c = append(c, hash...)       // 哈希
	c = append(c, cipherText...) // 密文
	return append([]byte{0x04}, c...)
}

func SM2KeyPairCheck(pub Sm2Key, pri *sm2.PrivateKey) bool {
	var err error
	var pkey *sm2.PublicKey
	var plain = []byte("HelloWorld")

	switch pub.(type) {
	case *x509.Certificate:
		pkey = pub.(*x509.Certificate).PublicKey.(*sm2.PublicKey)
	case *sm2.PublicKey:
		pkey = pub.(*sm2.PublicKey)
	default:
		return false
	}

	encData, err := sm2.Encrypt(pkey, plain, nil, sm2.C1C3C2)
	if err != nil {
		return false
	}
	plainData, err := sm2.Decrypt(pri, encData, sm2.C1C3C2)
	if err != nil {
		return false
	}
	if bytes.Compare(plainData, plain) != 0 {
		return false
	}
	r, s, err := sm2.Sm2Sign(pri, plain, nil, nil)
	if err != nil {
		return false
	}
	if !sm2.Sm2Verify(pkey, plain, nil, r, s) {
		return false
	}
	return true
}

func CERTGeneratePubKeyFromBytes(data []byte) (crypto.PublicKey, error) {
	cert, err := CERTGenerateObjectFromBytes(data)
	if err != nil {
		return nil, err
	}
	switch cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		ecdSa := cert.PublicKey.(*ecdsa.PublicKey)
		return &sm2.PublicKey{
			Curve: ecdSa.Curve,
			X:     ecdSa.X,
			Y:     ecdSa.Y,
		}, nil
	case *rsa.PublicKey:
		return cert.PublicKey.(*rsa.PublicKey), nil
	default:
		return nil, errors.New("unknown this cert")
	}
}

// 从证书中获取证书对象
func CERTGenerateObjectFromBytes(data []byte) (*x509.Certificate, error) {
	c, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		c = data
	}
	return x509.ParseCertificate(c)
}
