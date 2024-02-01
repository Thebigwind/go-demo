package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//)

//func main() {
//	// 读取.yaml文件
//	data, err := ioutil.ReadFile("/Users/me/Thebigwind/go-demo/demo120-yaml2/yaml/default.yml")
//	if err != nil {
//		panic(err)
//	}
//	// 将.yaml文件解析为map
//	m := make(map[interface{}]interface{})
//	err = yaml.Unmarshal(data, &m)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("m:%v", m["codebook_purposes"])
//
//	//// 更新map中的键值对
//	//m["key"] = "new value"
//	//// 将map重新编码为.yaml文件
//	//out, err := yaml.Marshal(m)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//err = ioutil.WriteFile("config.yaml", out, 0644)
//	//if err != nil {
//	//	panic(err)
//	//}
//	fmt.Println("Successfully updated config.yaml")
//}

import (
	"gopkg.in/yaml.v2"
)

// Config 结构体用于定义 YAML 文件的结构
type Config struct {
	Server struct {
		HTTP string `yaml:"http_endpoint"`
		GRPC string `yaml:"grpc_endpoint"`
		File string `yaml:"file_endpoint"`
	} `yaml:"server"`

	// ... 其他字段

	Deploy struct {
		Architecture string   `yaml:"architecture"`
		Node         []string `yaml:"node"`
	} `yaml:"deploy"`
}

func main() {
	// 读取 YAML 文件内容
	filePath := "path/to/your/default.yml"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 解析 YAML 文件内容
	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}

	// 更新配置
	config.Deploy.Architecture = "ha"
	config.Deploy.Node = []string{"10.10.80.11", "10.10.80.12", "10.10.80.13"}
	// 其他更新...

	// 将更新后的配置转换为 YAML 格式
	updatedYAML, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	// 写回文件
	err = ioutil.WriteFile(filePath, updatedYAML, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("YAML file updated successfully")
}

func test() {

}
