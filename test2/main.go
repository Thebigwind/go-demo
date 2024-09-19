package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func main() {

	var extmap = map[string]interface{}{
		"node_ip":         "10.10.10.215",
		"node_id":         "10.10.10.215",
		"expiration_date": "2030-11-11 01:00:00",
		"tenant_max_num":  100,
		"cert_max_num":    10000,
	}
	bytes, err := json.Marshal(extmap)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//
	extBase64 := base64.StdEncoding.EncodeToString(bytes)
	fmt.Printf("extBase64:%v\n", extBase64)

	//
	var str = `
[
{
"index": 1,
"info": {
"desc": "PQC证书公钥",
"oid": "1.3.6.1.4.1.60473.01"
}
},
{
"index": 2,
"info": {
"desc": "PQC证书申请公钥",
"oid": "1.3.6.1.4.1.60473.11"
}
},
{
"index": 3,
"info": {
"desc": "PQC签名",
"oid": "1.3.6.1.4.1.60473.21"
}
},
{
"index": 4,
"info": {
"desc": "QSKM扩展项",
"oid": "1.3.6.1.4.1.60473.31"
}
},
{
"index": 5,
"info": {
"desc": "QSCS扩展项",
"oid": "1.3.6.1.4.1.60473.41"
}
}
]
`

	var Arr = []OidInfo{}
	err = json.Unmarshal([]byte(str), &Arr)
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	fmt.Printf("Arr:%v\n", Arr)
}

type OidInfo struct {
	Index int  `json:"index"`
	Info  Info `json:"info"`
}
type Info struct {
	Desc string `json:"desc"`
	Oid  string `json:"oid"`
}
