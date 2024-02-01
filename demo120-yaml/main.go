package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Pwd struct {
	Password int `yaml:"password"`
}
type Env struct {
	GinMode string   `yaml:"gin_mode"`
	Authors []string `yaml:"authors"`
	Age     int      `yaml:"age"`
	Dev     Pwd      `yaml:"dev"`
	Test    Pwd      `yaml:"test"`
}

func main() {
	path := "yaml/test.yaml"

	// 读取yaml
	env := readYaml(path)

	// 修改值
	env.Age = 18
	env.Test.Password = 88

	// 写入yaml
	writeYaml(path, env)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readYaml(path string) (env Env) {
	content, err := ioutil.ReadFile(path)
	checkError(err)

	err = yaml.Unmarshal(content, &env)
	checkError(err)

	fmt.Println(err, env)
	return env
}

func writeYaml(path string, env Env) {
	data, err := yaml.Marshal(env)
	checkError(err)
	fmt.Printf("data:%v\n", data)
	err = ioutil.WriteFile(path, data, 0777)
	checkError(err)
}
