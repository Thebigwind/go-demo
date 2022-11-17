package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	err := httpPost()
	if err != nil {
		return
	}
}
func httpPost() error {
	type bodyData struct {
		Check string `json:"check"`
	}
	ActiveCheck := "AYTAkNOMRAwDgYDVQQIDAdCZWlqaW5nMRAwDgYDVQQHDAdCZWlqaW5nMQ0wCwYo2gAwIBAgIJAKq+j7FgdZWAMA0GCSqGSIb3DQEBCw"
	data := bodyData{Check: ActiveCheck}
	fmt.Println(data)
	js, _ := json.Marshal(data)
	fmt.Printf(string(js))

	resp, err := http.Post("http://127.0.0.1:9801/v1/active_update_status", "application/json", strings.NewReader(string(js)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
