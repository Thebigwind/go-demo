package main

import (
	"crypto/rand"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"math/big"
	"net"
	"os"
	"time"
)

func GenerateCertKeySM2(host, commonName string, alternateIPs []net.IP, alternateDNS []string) {
	priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Issuer:       pkix.Name{},
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{"Test"},
			Country:      []string{"CN"},
			ExtraNames: []pkix.AttributeTypeAndValue{
				{
					Type:  []int{2, 5, 4, 42},
					Value: "Gopher",
				},
				{
					Type:  []int{2, 5, 4, 6},
					Value: "NL",
				},
			},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		SignatureAlgorithm:    x509.SM2WithSM3,
	}

	if ip := net.ParseIP(host); ip != nil {
		template.IPAddresses = append(template.IPAddresses, ip)
	} else {
		template.DNSNames = append(template.DNSNames, host)
	}

	template.IPAddresses = append(template.IPAddresses, alternateIPs...)
	template.DNSNames = append(template.DNSNames, alternateDNS...)

	publicKey := priv.Public().(*sm2.PublicKey)
	certificateToPem, err := x509.CreateCertificateToPem(&template, &template, publicKey, priv)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("caSM2.crt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(certificateToPem)

	privateKey, err := x509.MarshalSm2PrivateKey(priv, []byte("jjjjjj"))
	if err != nil {
		panic(err)
	}
	block := pem.Block{
		Type:    "SM2 PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKey,
	}
	file, err = os.Create("caSM2.key")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = pem.Encode(file, &block)
	if err != nil {
		panic(err)
	}
}

func main() {
	ip := []byte("127.0.0.1")
	alternateDNS := []string{"localhost"}
	GenerateCertKeySM2("192.168.0.1", "www.example.com", []net.IP{net.ParseIP(string(ip))}, alternateDNS)
}
