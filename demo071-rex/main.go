package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	//"strings"
	"time"
)

func main() {
	res := "Certificate:\n    Data:\n        Version: 3 (0x2)\n        Serial Number:\n            d5:eb:39:7f:db:d1:4f:d4\n    Signature Algorithm: sha256WithRSAEncryption\n        Issuer: C = CN, ST = Beijing, L = Beijing, O = zdlz, OU = zdlz\n        Validity\n            Not Before: Nov 10 08:11:41 2022 GMT\n            Not After : Nov  7 08:11:41 2032 GMT\n        Subject: C = cn, O = mycomp, OU = mygroup, CN = 10.10.10.180\n        Subject Public Key Info:\n            Public Key Algorithm: rsaEncryption\n                Public-Key: (2048 bit)\n                Modulus:\n                    00:d5:71:90:5c:25:00:2e:2a:b4:57:3c:38:4e:e6:\n                    9f:5d:13:ec:17:0c:2e:0f:93:73:b3:da:f7:92:fe:\n                    13:54:00:2f:26:7e:2d:b0:02:b5:73:74:cc:27:2b:\n                    f7:48:2b:41:e1:ad:58:72:f3:d4:30:5d:48:a6:dc:\n                    95:af:65:73:88:69:eb:62:47:96:96:9f:62:06:09:\n                    ab:cb:48:9a:79:e1:f4:80:53:fa:9f:5e:c6:d2:43:\n                    a3:0a:7a:55:75:ba:2b:88:87:49:7a:18:f1:83:d9:\n                    b9:0a:9e:20:a2:c9:01:0e:22:93:42:02:6d:40:91:\n                    64:86:12:a5:fb:55:9b:ba:4c:4a:7c:bb:c3:ec:dd:\n                    2b:7f:d3:ce:94:f7:96:77:d5:a0:4c:61:00:a4:4c:\n                    1d:31:c6:2d:01:71:cb:50:0b:09:40:ca:78:88:2f:\n                    71:15:ff:e2:8c:74:b4:08:27:ac:0b:3d:12:a3:a3:\n                    3e:eb:4d:67:c2:c2:73:f7:8b:ad:4b:cb:46:97:b4:\n                    43:ed:fe:c6:b2:a5:20:b1:86:d4:05:da:f1:ef:d6:\n                    e4:a9:fc:23:71:36:22:38:3f:9b:fe:9b:dc:d5:2a:\n                    9e:44:c5:e4:35:c6:4b:c0:fb:24:5e:9f:8b:ac:23:\n                    18:18:dc:a0:08:5b:18:c6:9f:d9:6a:07:62:0e:cd:\n                    89:e7\n                Exponent: 65537 (0x10001)\n        X509v3 extensions:\n            1.2.3.4.5.6.7.8:\n                .=ParentSerialNumber=d8:56:c5:7b:86:c0:55:88,MaxDeviceNum=10000\n    Signature Algorithm: sha256WithRSAEncryption\n         b4:53:b2:2b:47:ec:81:95:2e:9e:88:7e:50:c6:00:dc:28:8c:\n         9b:11:a2:b3:c3:e5:10:7d:bd:8a:6a:d2:b8:d9:3d:53:dd:65:\n         f9:84:c0:52:12:24:1d:13:b5:91:b7:dc:22:4b:30:b4:8e:18:\n         e9:56:eb:23:b9:75:7b:3d:00:3b:c7:0a:14:53:0d:d8:6e:41:\n         44:11:90:e1:dc:64:00:c3:fe:2c:07:62:62:98:46:55:31:93:\n         fb:8c:a9:0d:3b:67:d3:18:d0:84:b1:49:d2:16:80:a8:b7:bc:\n         2e:ce:14:7a:91:77:32:93:f1:68:56:01:86:e2:54:77:d6:fe:\n         b2:6b:59:61:33:20:85:37:5d:9e:2a:cb:8e:a6:03:c3:76:6b:\n         f4:35:a1:23:e4:d5:82:f1:d9:52:71:1c:f1:04:d3:27:4f:82:\n         2c:e0:96:60:98:aa:ba:c2:20:46:ce:0f:f7:f1:ec:bf:a6:22:\n         d3:8a:db:5d:ee:bc:50:96:2f:fb:5c:bc:23:fe:1a:59:94:3a:\n         1d:9a:28:f9:49:00:1a:3f:92:70:b1:67:c1:85:9f:26:7e:5f:\n         ba:c9:3f:d9:99:8c:b2:a8:52:30:e9:27:62:36:20:41:bd:c8:\n         80:db:14:15:f7:21:e0:f4:d9:56:c5:31:a8:ba:b7:2f:ef:02:\n         e1:8b:56:9a"
	compileRegex := regexp.MustCompile("Not After : (.+)") // 提取过期时间
	matchArr := compileRegex.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return
	}
	afterStr := matchArr[len(matchArr)-1]
	fmt.Printf("afterStr:%v\n", afterStr)
	//afterStr = strings.TrimPrefix(afterStr,"Not After : ")
	//fmt.Printf("afterStr:%v\n", afterStr)
	afterTime, err := GetGmtTime(afterStr)
	if err != nil {
		return
	}
	fmt.Printf("afterTime:%s\n", afterTime)

	data, err := GetCertificateMaxDeviceNum("aa")
	fmt.Println(data)

	id, pid, err := GetCertificateId("xxx")
	fmt.Println(id)
	fmt.Println(pid)
}

func GetGmtTime(secStr string) (*time.Time, error) {
	// May 29 08:00:17 2023 GMT
	formatTimeStr := "Jan 2 15:04:05 2006 GMT"
	start, err := time.Parse(formatTimeStr, secStr)
	if err == nil {
		realTime := start.Add(8 * time.Hour)
		return &realTime, nil
	} else {
		return nil, err
	}
}

func GetCertificateMaxDeviceNum(certFile string) (int, error) {
	if len(certFile) == 0 {
		return 0, errors.New("cert is null")
	}

	//获取证书text信息
	res := "Certificate:\n    Data:\n        Version: 3 (0x2)\n        Serial Number:\n            d5:eb:39:7f:db:d1:4f:d4\n    Signature Algorithm: sha256WithRSAEncryption\n        Issuer: C = CN, ST = Beijing, L = Beijing, O = zdlz, OU = zdlz\n        Validity\n            Not Before: Nov 10 08:11:41 2022 GMT\n            Not After : Nov  7 08:11:41 2032 GMT\n        Subject: C = cn, O = mycomp, OU = mygroup, CN = 10.10.10.180\n        Subject Public Key Info:\n            Public Key Algorithm: rsaEncryption\n                Public-Key: (2048 bit)\n                Modulus:\n                    00:d5:71:90:5c:25:00:2e:2a:b4:57:3c:38:4e:e6:\n                    9f:5d:13:ec:17:0c:2e:0f:93:73:b3:da:f7:92:fe:\n                    13:54:00:2f:26:7e:2d:b0:02:b5:73:74:cc:27:2b:\n                    f7:48:2b:41:e1:ad:58:72:f3:d4:30:5d:48:a6:dc:\n                    95:af:65:73:88:69:eb:62:47:96:96:9f:62:06:09:\n                    ab:cb:48:9a:79:e1:f4:80:53:fa:9f:5e:c6:d2:43:\n                    a3:0a:7a:55:75:ba:2b:88:87:49:7a:18:f1:83:d9:\n                    b9:0a:9e:20:a2:c9:01:0e:22:93:42:02:6d:40:91:\n                    64:86:12:a5:fb:55:9b:ba:4c:4a:7c:bb:c3:ec:dd:\n                    2b:7f:d3:ce:94:f7:96:77:d5:a0:4c:61:00:a4:4c:\n                    1d:31:c6:2d:01:71:cb:50:0b:09:40:ca:78:88:2f:\n                    71:15:ff:e2:8c:74:b4:08:27:ac:0b:3d:12:a3:a3:\n                    3e:eb:4d:67:c2:c2:73:f7:8b:ad:4b:cb:46:97:b4:\n                    43:ed:fe:c6:b2:a5:20:b1:86:d4:05:da:f1:ef:d6:\n                    e4:a9:fc:23:71:36:22:38:3f:9b:fe:9b:dc:d5:2a:\n                    9e:44:c5:e4:35:c6:4b:c0:fb:24:5e:9f:8b:ac:23:\n                    18:18:dc:a0:08:5b:18:c6:9f:d9:6a:07:62:0e:cd:\n                    89:e7\n                Exponent: 65537 (0x10001)\n        X509v3 extensions:\n            1.2.3.4.5.6.7.8:\n                .=ParentSerialNumber=d8:56:c5:7b:86:c0:55:88,MaxDeviceNum=10000\n    Signature Algorithm: sha256WithRSAEncryption\n         b4:53:b2:2b:47:ec:81:95:2e:9e:88:7e:50:c6:00:dc:28:8c:\n         9b:11:a2:b3:c3:e5:10:7d:bd:8a:6a:d2:b8:d9:3d:53:dd:65:\n         f9:84:c0:52:12:24:1d:13:b5:91:b7:dc:22:4b:30:b4:8e:18:\n         e9:56:eb:23:b9:75:7b:3d:00:3b:c7:0a:14:53:0d:d8:6e:41:\n         44:11:90:e1:dc:64:00:c3:fe:2c:07:62:62:98:46:55:31:93:\n         fb:8c:a9:0d:3b:67:d3:18:d0:84:b1:49:d2:16:80:a8:b7:bc:\n         2e:ce:14:7a:91:77:32:93:f1:68:56:01:86:e2:54:77:d6:fe:\n         b2:6b:59:61:33:20:85:37:5d:9e:2a:cb:8e:a6:03:c3:76:6b:\n         f4:35:a1:23:e4:d5:82:f1:d9:52:71:1c:f1:04:d3:27:4f:82:\n         2c:e0:96:60:98:aa:ba:c2:20:46:ce:0f:f7:f1:ec:bf:a6:22:\n         d3:8a:db:5d:ee:bc:50:96:2f:fb:5c:bc:23:fe:1a:59:94:3a:\n         1d:9a:28:f9:49:00:1a:3f:92:70:b1:67:c1:85:9f:26:7e:5f:\n         ba:c9:3f:d9:99:8c:b2:a8:52:30:e9:27:62:36:20:41:bd:c8:\n         80:db:14:15:f7:21:e0:f4:d9:56:c5:31:a8:ba:b7:2f:ef:02:\n         e1:8b:56:9a"
	res = "Certificate:\n    Data:\n        Version: 3 (0x2)\n        Serial Number:\n            d5:eb:39:7f:db:d1:4f:d9\n    Signature Algorithm: sha256WithRSAEncryption\n        Issuer: C = CN, ST = Beijing, L = Beijing, O = zdlz, OU = zdlz\n        Validity\n            Not Before: Nov 11 15:28:04 2022 GMT\n            Not After : Nov  8 15:28:04 2032 GMT\n        Subject: C = CN, ST = Beijing, L = Beijing, O = zdlz, OU = zdlz\n        Subject Public Key Info:\n            Public Key Algorithm: rsaEncryption\n                Public-Key: (2048 bit)\n                Modulus:\n                    00:c3:3e:2a:11:72:47:a9:01:2e:dc:c9:a1:bf:49:\n                    3f:1d:c8:0b:af:d9:2f:62:44:0f:08:fe:c5:10:36:\n                    51:27:3f:e8:48:83:d9:82:22:ed:73:3d:57:e7:ae:\n                    15:66:e0:43:3d:06:d1:48:9f:c7:3c:2b:59:a0:3e:\n                    ea:4a:ea:b5:92:7b:d2:10:3b:b4:3b:68:85:7e:0b:\n                    ec:54:65:0d:6d:76:91:f3:3f:0e:79:64:64:96:56:\n                    71:d9:12:18:18:38:aa:d8:a7:eb:9a:72:fa:90:52:\n                    9e:a8:ea:82:11:f6:d4:f8:70:92:fa:64:af:b2:68:\n                    43:a0:d7:2a:e6:11:bd:9f:99:aa:1c:c6:fe:c0:45:\n                    9f:80:f3:13:7b:57:d4:6d:1a:42:b7:f9:09:cc:db:\n                    9d:9b:4c:03:a8:cc:aa:03:8b:00:84:71:73:ec:39:\n                    3a:d3:39:54:4e:9d:ad:c6:2e:3b:03:d1:79:1f:ec:\n                    6f:46:cd:74:ba:03:97:7f:17:fd:73:39:fc:2b:7f:\n                    90:e1:47:cb:81:26:48:80:a3:3c:2b:e0:7b:7c:05:\n                    aa:72:e2:b8:1c:20:bd:5a:7a:9f:ac:d0:34:00:0c:\n                    5a:5d:80:4b:b9:42:64:af:37:b3:b0:0b:49:fb:d2:\n                    71:45:27:9d:06:d3:b6:13:3f:9c:9b:3a:7d:86:6d:\n                    9d:65\n                Exponent: 65537 (0x10001)\n        X509v3 extensions:\n            1.2.3.4:\n                .6ParentSerialNumber=D856C57B86C0558A,MaxDeviceNum=10000\n    Signature Algorithm: sha256WithRSAEncryption\n         63:c2:ed:a7:e9:30:e3:31:d5:9d:3b:2a:d4:51:04:a8:2f:bd:\n         52:ca:d1:c1:cf:64:dc:93:75:71:08:c0:de:0f:cc:7e:26:74:\n         74:f9:30:1a:f8:1a:49:e7:08:bb:f7:e3:ee:b4:47:da:3e:bd:\n         00:cd:c3:8a:fa:54:b6:f8:99:73:af:ac:94:4b:b6:2b:ed:df:\n         3b:12:a9:0e:d8:fc:28:0d:cd:00:62:53:61:14:b1:cd:83:58:\n         5e:b8:27:c0:86:35:bd:aa:7d:7a:fb:6d:89:81:b1:f1:62:4c:\n         cd:69:9a:fd:b7:49:a7:e1:1b:cf:b9:19:19:21:97:c2:05:10:\n         d2:63:b8:3d:48:a9:2a:25:27:21:51:d7:bc:84:69:68:14:75:\n         b0:b2:4e:93:70:81:4e:b0:f7:71:92:a1:5f:d7:d9:e9:b5:76:\n         26:ab:8a:92:33:60:87:d8:00:75:f4:ed:1c:1e:70:f3:06:af:\n         f5:29:91:24:7f:f3:d8:fa:b2:fb:bf:f3:6d:26:93:f2:ff:d3:\n         d5:61:5d:73:b5:7d:1a:e0:ea:eb:1a:03:59:89:b7:77:d0:e4:\n         34:20:0b:8b:84:81:b2:10:2c:b0:4e:b2:bf:8c:81:56:b8:03:\n         db:8a:f1:dc:d3:3f:05:3c:57:05:cd:1d:af:3b:42:a2:03:2c:\n         2a:28:fd:6e"
	// 获取MaxDeviceNum
	compileRegex := regexp.MustCompile(`MaxDeviceNum=(.+)`)
	matchArr := compileRegex.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return 0, errors.New("get certificate MaxDeviceNum error")
	}
	infoStr := matchArr[0]

	infoStr = strings.TrimPrefix(infoStr, "MaxDeviceNum=")
	num := ""
	if strings.Contains(infoStr, ",") {
		arr := strings.Split(infoStr, ",")
		num = arr[0]
	} else {
		num = infoStr
	}

	numInt, err := strconv.Atoi(num)
	if err != nil {
		return 0, errors.New("get certificate MaxDeviceNum error")
	}
	return numInt, nil
}

func GetCertificateId(certFile string) (string, string, error) {
	if len(certFile) == 0 {
		return "", "", errors.New("cert is null")
	}

	//获取证书subject信息
	//cmdStr := fmt.Sprintf("gmssl x509 -in %s -noout -serial", certFile)
	//res, err := Command(cmdStr)
	//if err != nil {
	//	return "", "", err
	//}
	res := "serial=D5EB397FDBD14FD4"
	// 获取id
	compileRegex := regexp.MustCompile(`serial=(.+)`)
	matchArr := compileRegex.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", "", errors.New("get certificate id error")
	}
	id := matchArr[len(matchArr)-1]

	//获取证书text信息
	res = "Certificate:\n    Data:\n        Version: 3 (0x2)\n        Serial Number:\n            d5:eb:39:7f:db:d1:4f:d4\n    Signature Algorithm: sha256WithRSAEncryption\n        Issuer: C = CN, ST = Beijing, L = Beijing, O = zdlz, OU = zdlz\n        Validity\n            Not Before: Nov 10 08:11:41 2022 GMT\n            Not After : Nov  7 08:11:41 2032 GMT\n        Subject: C = cn, O = mycomp, OU = mygroup, CN = 10.10.10.180\n        Subject Public Key Info:\n            Public Key Algorithm: rsaEncryption\n                Public-Key: (2048 bit)\n                Modulus:\n                    00:d5:71:90:5c:25:00:2e:2a:b4:57:3c:38:4e:e6:\n                    9f:5d:13:ec:17:0c:2e:0f:93:73:b3:da:f7:92:fe:\n                    13:54:00:2f:26:7e:2d:b0:02:b5:73:74:cc:27:2b:\n                    f7:48:2b:41:e1:ad:58:72:f3:d4:30:5d:48:a6:dc:\n                    95:af:65:73:88:69:eb:62:47:96:96:9f:62:06:09:\n                    ab:cb:48:9a:79:e1:f4:80:53:fa:9f:5e:c6:d2:43:\n                    a3:0a:7a:55:75:ba:2b:88:87:49:7a:18:f1:83:d9:\n                    b9:0a:9e:20:a2:c9:01:0e:22:93:42:02:6d:40:91:\n                    64:86:12:a5:fb:55:9b:ba:4c:4a:7c:bb:c3:ec:dd:\n                    2b:7f:d3:ce:94:f7:96:77:d5:a0:4c:61:00:a4:4c:\n                    1d:31:c6:2d:01:71:cb:50:0b:09:40:ca:78:88:2f:\n                    71:15:ff:e2:8c:74:b4:08:27:ac:0b:3d:12:a3:a3:\n                    3e:eb:4d:67:c2:c2:73:f7:8b:ad:4b:cb:46:97:b4:\n                    43:ed:fe:c6:b2:a5:20:b1:86:d4:05:da:f1:ef:d6:\n                    e4:a9:fc:23:71:36:22:38:3f:9b:fe:9b:dc:d5:2a:\n                    9e:44:c5:e4:35:c6:4b:c0:fb:24:5e:9f:8b:ac:23:\n                    18:18:dc:a0:08:5b:18:c6:9f:d9:6a:07:62:0e:cd:\n                    89:e7\n                Exponent: 65537 (0x10001)\n        X509v3 extensions:\n            1.2.3.4.5.6.7.8:\n                .=ParentSerialNumber=d8:56:c5:7b:86:c0:55:88,MaxDeviceNum=10000\n    Signature Algorithm: sha256WithRSAEncryption\n         b4:53:b2:2b:47:ec:81:95:2e:9e:88:7e:50:c6:00:dc:28:8c:\n         9b:11:a2:b3:c3:e5:10:7d:bd:8a:6a:d2:b8:d9:3d:53:dd:65:\n         f9:84:c0:52:12:24:1d:13:b5:91:b7:dc:22:4b:30:b4:8e:18:\n         e9:56:eb:23:b9:75:7b:3d:00:3b:c7:0a:14:53:0d:d8:6e:41:\n         44:11:90:e1:dc:64:00:c3:fe:2c:07:62:62:98:46:55:31:93:\n         fb:8c:a9:0d:3b:67:d3:18:d0:84:b1:49:d2:16:80:a8:b7:bc:\n         2e:ce:14:7a:91:77:32:93:f1:68:56:01:86:e2:54:77:d6:fe:\n         b2:6b:59:61:33:20:85:37:5d:9e:2a:cb:8e:a6:03:c3:76:6b:\n         f4:35:a1:23:e4:d5:82:f1:d9:52:71:1c:f1:04:d3:27:4f:82:\n         2c:e0:96:60:98:aa:ba:c2:20:46:ce:0f:f7:f1:ec:bf:a6:22:\n         d3:8a:db:5d:ee:bc:50:96:2f:fb:5c:bc:23:fe:1a:59:94:3a:\n         1d:9a:28:f9:49:00:1a:3f:92:70:b1:67:c1:85:9f:26:7e:5f:\n         ba:c9:3f:d9:99:8c:b2:a8:52:30:e9:27:62:36:20:41:bd:c8:\n         80:db:14:15:f7:21:e0:f4:d9:56:c5:31:a8:ba:b7:2f:ef:02:\n         e1:8b:56:9a"
	res = "Certificate:\n    Data:\n        Version: 3 (0x2)\n        Serial Number:\n            d5:eb:39:7f:db:d1:4f:d9\n    Signature Algorithm: sha256WithRSAEncryption\n        Issuer: C = CN, ST = Beijing, L = Beijing, O = zdlz, OU = zdlz\n        Validity\n            Not Before: Nov 11 15:28:04 2022 GMT\n            Not After : Nov  8 15:28:04 2032 GMT\n        Subject: C = CN, ST = Beijing, L = Beijing, O = zdlz, OU = zdlz\n        Subject Public Key Info:\n            Public Key Algorithm: rsaEncryption\n                Public-Key: (2048 bit)\n                Modulus:\n                    00:c3:3e:2a:11:72:47:a9:01:2e:dc:c9:a1:bf:49:\n                    3f:1d:c8:0b:af:d9:2f:62:44:0f:08:fe:c5:10:36:\n                    51:27:3f:e8:48:83:d9:82:22:ed:73:3d:57:e7:ae:\n                    15:66:e0:43:3d:06:d1:48:9f:c7:3c:2b:59:a0:3e:\n                    ea:4a:ea:b5:92:7b:d2:10:3b:b4:3b:68:85:7e:0b:\n                    ec:54:65:0d:6d:76:91:f3:3f:0e:79:64:64:96:56:\n                    71:d9:12:18:18:38:aa:d8:a7:eb:9a:72:fa:90:52:\n                    9e:a8:ea:82:11:f6:d4:f8:70:92:fa:64:af:b2:68:\n                    43:a0:d7:2a:e6:11:bd:9f:99:aa:1c:c6:fe:c0:45:\n                    9f:80:f3:13:7b:57:d4:6d:1a:42:b7:f9:09:cc:db:\n                    9d:9b:4c:03:a8:cc:aa:03:8b:00:84:71:73:ec:39:\n                    3a:d3:39:54:4e:9d:ad:c6:2e:3b:03:d1:79:1f:ec:\n                    6f:46:cd:74:ba:03:97:7f:17:fd:73:39:fc:2b:7f:\n                    90:e1:47:cb:81:26:48:80:a3:3c:2b:e0:7b:7c:05:\n                    aa:72:e2:b8:1c:20:bd:5a:7a:9f:ac:d0:34:00:0c:\n                    5a:5d:80:4b:b9:42:64:af:37:b3:b0:0b:49:fb:d2:\n                    71:45:27:9d:06:d3:b6:13:3f:9c:9b:3a:7d:86:6d:\n                    9d:65\n                Exponent: 65537 (0x10001)\n        X509v3 extensions:\n            1.2.3.4:\n                .6ParentSerialNumber=D856C57B86C0558A,MaxDeviceNum=10000\n    Signature Algorithm: sha256WithRSAEncryption\n         63:c2:ed:a7:e9:30:e3:31:d5:9d:3b:2a:d4:51:04:a8:2f:bd:\n         52:ca:d1:c1:cf:64:dc:93:75:71:08:c0:de:0f:cc:7e:26:74:\n         74:f9:30:1a:f8:1a:49:e7:08:bb:f7:e3:ee:b4:47:da:3e:bd:\n         00:cd:c3:8a:fa:54:b6:f8:99:73:af:ac:94:4b:b6:2b:ed:df:\n         3b:12:a9:0e:d8:fc:28:0d:cd:00:62:53:61:14:b1:cd:83:58:\n         5e:b8:27:c0:86:35:bd:aa:7d:7a:fb:6d:89:81:b1:f1:62:4c:\n         cd:69:9a:fd:b7:49:a7:e1:1b:cf:b9:19:19:21:97:c2:05:10:\n         d2:63:b8:3d:48:a9:2a:25:27:21:51:d7:bc:84:69:68:14:75:\n         b0:b2:4e:93:70:81:4e:b0:f7:71:92:a1:5f:d7:d9:e9:b5:76:\n         26:ab:8a:92:33:60:87:d8:00:75:f4:ed:1c:1e:70:f3:06:af:\n         f5:29:91:24:7f:f3:d8:fa:b2:fb:bf:f3:6d:26:93:f2:ff:d3:\n         d5:61:5d:73:b5:7d:1a:e0:ea:eb:1a:03:59:89:b7:77:d0:e4:\n         34:20:0b:8b:84:81:b2:10:2c:b0:4e:b2:bf:8c:81:56:b8:03:\n         db:8a:f1:dc:d3:3f:05:3c:57:05:cd:1d:af:3b:42:a2:03:2c:\n         2a:28:fd:6e"
	//获取父级id
	compileRegex = regexp.MustCompile(`ParentSerialNumber=(.+)`)
	matchArr = compileRegex.FindStringSubmatch(res)
	if len(matchArr) == 0 {
		return "", "", errors.New("get certificate id error")
	}
	pid := ""

	infoStr := strings.TrimPrefix(matchArr[len(matchArr)-1], "ParentSerialNumber:")
	if strings.Contains(infoStr, ",") {
		sli := strings.Split(infoStr, ",")
		pid = sli[0]
	}

	return id, pid, nil
}
