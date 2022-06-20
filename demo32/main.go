package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {

}

//把 JSON 数据从 byt 变量反序列化（如解析、解码等等）成名为 dat 的 map/字典对象
// 当我们处理一个多层嵌套的 JSON 对象时，这些类型断言会让处理变得非常繁琐
func test1() {
	byt := []byte(`{
        "num":6.13,
        "strs":["a","b"],
        "obj":{"foo":{"bar":"zip","zap":6}}
    }`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	obj := dat["obj"].(map[string]interface{})
	obj2 := obj["foo"].(map[string]interface{})
	fmt.Println(obj2)
}

// 利用 Go struct 的标签功能把 byt 变量中的字节反序列化成一个具体的结构 ourData
// 在 struct 的定义中，标签不是必需的。如果你的 struct 中包含了标签，那么它意味着 Go 的 反射 API[4] 可以访问标签的值
func test2() {
	type ourData struct {
		Num  float64                      `json:"num"`
		Strs []string                     `json:"strs"`
		Obj  map[string]map[string]string `json:"obj"`
	}

	byt := []byte(`{
        "num":6.13,
        "strs":["a","b"],
        "obj":{"foo":{"bar":"zip","zap":6}}
    }`)

	res := ourData{}
	json.Unmarshal(byt, &res)
	fmt.Println(res.Num)
	fmt.Println(res.Strs)
	fmt.Println(res.Obj)
}

const JSON = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func test3() {
	value := gjson.Get(JSON, "name.last")
	println(value.String())
}

func test4() {
	value, _ := sjson.Set(JSON, "name.last", "Anderson")
	println(value)
}
