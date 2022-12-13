package main

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type Config struct {
	Addr           string `json:"addr"`
	AesKey         string `json:"aes_key"`
	HTTPS          bool   `json:"https"`
	Secret         string `json:"secret"`
	PrivateKeyPath string `json:"private_key_path"`
	CertFilePath   string `json:"cert_file_path"`
}

var (
	appConfig  Config
	configPath = `/configs/remote_config.json`
	client     *clientv3.Client
	err        error
)

func init() {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}

	initConfig()
}

func watchAndUpdate() {
	watch := client.Watch(context.Background(), configPath)

	for wresp := range watch {
		for _, ev := range wresp.Events {
			log.Println("new values is ", string(ev.Kv.Value))
			err = json.Unmarshal(ev.Kv.Value, &appConfig)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

}

func initConfig() {
	var resp *clientv3.GetResponse
	resp, err = client.Get(context.Background(), configPath)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(resp.Kvs[0].Value, &appConfig)
}

func getConfig() Config {
	return appConfig
}

func main() {
	defer client.Close()
	watchAndUpdate()
}
