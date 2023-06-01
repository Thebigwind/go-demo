package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	// PEM-encoded certificate
	certPEM := []byte(`-----BEGIN CERTIFICATE-----
MIIC+jCCAeKgAwIBAgIQQg4C4jJQq8wz4xZJvZ0F6TAKBggqhkjOPQQDAzBMMQsw
CQYDVQQGEwJDTjERMA8GA1UECAwITmV3IFlvcmsxEDAOBgNVBAcMB0JhbmRvbjEY
MBYGA1UECgwPTmV0d29ya3MgSW5jLjEXMBUGA1UECwwOTmV0d29ya3MgSW5jLjEW
MBQGA1UEAwwNKi5uZXR3b3Jrcy5jb20wHhcNMjEwMzA2MDUxNzE4WhcNMjIwMzA2
MDUxNzE4WjBMMQswCQYDVQQGEwJDTjERMA8GA1UECAwITmV3IFlvcmsxEDAOBgNV
BAcMB0JhbmRvbjEYMBYGA1UECgwPTmV0d29ya3MgSW5jLjEXMBUGA1UECwwOTmV0
d29ya3MgSW5jLjEWMBQGA1UEAwwNKi5uZXR3b3Jrcy5jb20wWTATBgcqhkjOPQIB
BggqhkjOPQMBBwNCAAS+X0ENkGK2iOaUJhO6PZi7hW0nY6K0vU7PihH2s9w9sZkS
W1v0x5+o8lK7+8vF9IhRt9rS+M5+q6y5+8f+Z4yo4GOMIGLMA4GA1UdDwEB/wQE
AwIHgDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBTbIb4+e5vKfLpYyUoGQ4c+7pDp
gzAfBgNVHSMEGDAWgBTbIb4+e5vKfLpYyUoGQ4c+7pDpgzAKBggqhkjOPQQDAwNo
ADBlAjEAj8xPOWnXsGK5zQ9n6L7n9YwKx+1RZ8bKq6tHwRt1tH5uXz2rAgkGEGI7
vFpXtXHc6l8CMQCkFz2sHlqH9Vn6xZz9aRgTtVXa2ybSbZz4gZyDjF+LW60=
-----END CERTIFICATE-----`)

	// decode the certificate from PEM format
	block, _ := pem.Decode(certPEM)
	if block == nil {
		panic("failed to decode certificate")
	}

	// parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	// get the key encryption algorithm
	alg := cert.PublicKeyAlgorithm.String()

	fmt.Println(alg) // output: ECDSA
}
