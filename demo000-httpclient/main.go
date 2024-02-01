package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Replace with your actual server IP and port
	serverIP := "10.10.10.129"
	OpsAPIHttpPort := "9803"

	server := fmt.Sprintf("http://%s:%s/v1/ops/upgrade_location", serverIP, OpsAPIHttpPort)

	// Your JSON payload
	js := []byte(`{"key": "value"}`)

	// Your token value
	token := "your_token_value"

	// Create a new request
	req, err := http.NewRequest("POST", server, bytes.NewBuffer(js))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Add custom header for token
	req.Header.Set("token", token)

	// Create HTTP client
	client := &http.Client{}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Printf("out:%s\n", string(out))
	// Process the response as needed
	fmt.Println("Response Status:", resp.Status)
}
