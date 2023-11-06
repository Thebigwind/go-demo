package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	//const (
	//	zdlz_crt       = "-----BEGIN CERTIFICATE-----\nMIIDUTCCAjmgAwIBAgIUfLLH7x/V3rlQtLFWrZyaCwQ85YswDQYJKoZIhvcNAQEL\nBQAwNzEWMBQGA1UEAwwNZGVtby56ZGx6LmNvbTELMAkGA1UEBhMCY24xEDAOBgNV\nBAcMB2JlaWppbmcwIBcNMjIwNDI3MTU0MzM4WhgPMjA3MjA0MTQxNTQzMzhaMDcx\nFjAUBgNVBAMMDWRlbW8uemRsei5jb20xCzAJBgNVBAYTAmNuMRAwDgYDVQQHDAdi\nZWlqaW5nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxxYLRUhTKpbc\nMnfoYEbnPAXNsTHn4nx0NA3m56byTT5HdnRdTAyendWSuTzw23wKe1qAmg63Xwc3\nrzlzdLmlUHgtasLnbp6+jao4+PUBkOsr1WZHQpupQppUR96vVFd+67PRM+PJKYCA\n7gzWVDDArdHXQ5Wq+duXue8TYPECC3R7EJvX1MQNcRMNQEUejcOxW5jzVyUf7SSG\nitCsTeP+YOeRU9DRRtseqwT0UfqQBQqyRJKXQ6tY+t8SDYEAAIzhA5zdHA3UJNXX\nPgzXVhSGeQev+5VK/2eoGCOfcMG04fZM72+U4hH2toqU1Vjgc2B8Mrzyy8/nLm6x\n6RfFPR4ZEQIDAQABo1MwUTAdBgNVHQ4EFgQUlV74d0MO1gJZDlHBg5lHeVnMwM0w\nHwYDVR0jBBgwFoAUlV74d0MO1gJZDlHBg5lHeVnMwM0wDwYDVR0TAQH/BAUwAwEB\n/zANBgkqhkiG9w0BAQsFAAOCAQEAapiAJ43Ru5NUTnpkYPsEkAbbw9Q9fkT1uHOP\nyeqHLWMQwzMJYvMdX5MFUeMG1e7rcfB3Q8uVtAZbDCvjuM6VMcD/fNSNXRCIOkZX\nfDCYaYv/v/H7b5Kq3z39CUBMWAzeygLYRHrIKR9O+SaVAjPsDrWF3xVSgif0sUT3\njnF78+tjvBnLcVD95dXY9R+SxosTB0dT640Yenp7e+73CKNyMxvFFuEKiCW9MHCj\nOleG/a6c7WND/nXe9mkOUf4uGuzYcNFWSI7spotQoTA5FiJZSlx9fZqZNI2WELeG\nDAVcW1Xu9RQ3sBhlutJ1SMs3SbGNzjask38VIYFsib/zXw9UXg==\n-----END CERTIFICATE-----"
	//	zdlz_gmssl_crt = "-----BEGIN CERTIFICATE-----\nMIIBvzCCAWQCCQDkS6BFLMnMKjAKBggqgRzPVQGDdTBmMQswCQYDVQQGEwJDTjEQ\nMA4GA1UECAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRs\nejENMAsGA1UECwwEemRsejEVMBMGA1UEAwwMc3RhdGVfc2VjcmV0MCAXDTIyMDgy\nMTAyNTUzNFoYDzIwNzIwODA4MDI1NTM0WjBmMQswCQYDVQQGEwJDTjEQMA4GA1UE\nCAwHQmVpamluZzEQMA4GA1UEBwwHQmVpamluZzENMAsGA1UECgwEemRsejENMAsG\nA1UECwwEemRsejEVMBMGA1UEAwwMc3RhdGVfc2VjcmV0MFkwEwYHKoZIzj0CAQYI\nKoEcz1UBgi0DQgAEhy+9gmPC0vYUH8kz9zg+9XRDIuELOj7Et1L5DcDRbSGwjFkA\n0tswBlj5xpycpgXeAUE3d/sUi7+TqRFXCTnNCzAKBggqgRzPVQGDdQNJADBGAiEA\n+lxh4L4pZT6rnsYiDI0gEtPluvokf/t1fCeFNUl2M7gCIQDX930NvpVUhl5XMdey\nJeAysqdSpQaA4Kxd9SALokAriQ==\n-----END CERTIFICATE-----"
	//)
	//
	//err := ioutil.WriteFile("aa.crt", []byte(zdlz_crt), 0755)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//
	//}
	test()
}

func test() {
	bytes, err := ioutil.ReadFile("admin.p12")
	if err != nil {
		fmt.Println(err)
	} else {
		data := base64.StdEncoding.EncodeToString(bytes)
		fmt.Printf("data:%s\n", data)
	}
}
