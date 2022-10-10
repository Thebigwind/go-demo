package main

import "net/http"
import "log"
import "io/ioutil"
import "context"
import "time"

func main() {
	server := "http://10.10.10.162:9801/v1/get_auth_random_num?type=0&id=emRsemRldkB6ZGx6LnRlY2g=" //http://www.baidu.com
	req, err := http.NewRequest(http.MethodGet, server, nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	defer cancel()
	req = req.WithContext(ctx)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
