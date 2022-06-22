package main

import "net/http"
import "log"
import "io/ioutil"
import "context"
import "time"

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
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
