package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	DoIndex()
}

func DoIndex() {

	var a int
	var b int
	var err error

	var wg = sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		a, err = A()
		if err != nil {
			fmt.Printf("err B")
		}
	}()
	go func() {
		defer wg.Done()
		b, err = B()
		if err != nil {
			fmt.Printf("err B")
		}
	}()

	wg.Wait()

	fmt.Printf("A:%v,B:%v\n", a, b)
}

func A() (int, error) {

	req, _ := http.NewRequest(http.MethodGet, "http://baidu.com", nil)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	req = req.WithContext(ctx)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("errB:", err.Error())
		return 0, err
	}

	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("errB:", err.Error())
		return 0, err
	}
	log.Println(string(out))
	return 1, nil
}

func B() (int, error) {
	req, _ := http.NewRequest(http.MethodGet, "http://baidu.com", nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
	defer cancel()
	req = req.WithContext(ctx)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("errB:", err.Error())
		return 0, err
	}

	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("errB:", err.Error())
		return 0, err
	}
	log.Println(string(out))
	return 2, nil
}
