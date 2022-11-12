package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

type QskmClient struct {
	Ctx context.Context
}

func NewQskmClient(ctx context.Context) *QskmClient {
	qskm := &QskmClient{}
	qskm.Ctx = ctx
	return qskm
}

func (q *QskmClient) GetRandom() (string, error) {

	server := "http://10.10.10.162:9801/v1/get_auth_random_num?type=0&id=emRsemRldkB6ZGx6LnRlY2g=" //http://www.baidu.com
	req, err := http.NewRequest(http.MethodGet, server, nil)
	if err != nil {
		log.Fatal(err)
	}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	//defer cancel()
	req = req.WithContext(q.Ctx)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	result := string(out)
	log.Println(result)
	return result, nil
}
