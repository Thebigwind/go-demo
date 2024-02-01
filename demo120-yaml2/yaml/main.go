package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config 结构体用于定义 new_default.yml 的结构
type Config struct {
	MongoDB map[string]interface{} `yaml:"mongodb"`
}

//type NewYamlConfig struct {
//	Mysql map[string]interface{} `yaml:"mysql"`
//}

type NewYamlConfig map[string]interface{}

func main() {
	// 读取 new_default.yml 文件内容
	filePath := "/Users/me/Thebigwind/go-demo/demo120-yaml2/yaml/default.yml"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 解析 new_default.yml 文件内容
	var config = map[string]interface{}{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("config:%+v\n", config)
	fmt.Printf("config:%+v", config["mysql"])
	//
	//// 更新指定字段
	//config.MongoDB["addrs"] = []string{"10.10.10.125:27017", "10.10.10.126:27017", "10.10.10.127:27017"}
	//
	//// 将更新后的配置转换为 YAML 格式
	//updatedYAML, err := yaml.Marshal(&config)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 写回 new_default.yml 文件
	//err = ioutil.WriteFile(filePath, updatedYAML, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("new_default.yml file updated successfully")
}
