package main

import (
	"fmt"
	"reflect"
)

type data1 struct {
	Name string
	Id   int
	//m map[string]string
}
type data2 struct {
	Name string
	Id   int
	m    map[string]string
}

func main() {

	v1 := data1{Name: "a", Id: 1}
	v2 := data1{Name: "a", Id: 1}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	//prints: v1 == v2: true
	fmt.Println("v1 == v2:", v1 == v2)

	v3 := data2{Name: "a", Id: 1}
	v4 := data2{Name: "a", Id: 1}
	fmt.Println("v3 == v4:", reflect.DeepEqual(v3, v4))
	//prints: v3 == v4: true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true
}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	for index := range a {
		if a[index] < a[right] { //以最后侧为分区点
			a[index], a[left] = a[left], a[index] //如果小于分区点，则放在左侧
			left++                                //左侧（比分区点小的）的长度加1
		}
	}
	a[left], a[right] = a[right], a[left] //把分区点 a[right] 放在正确的分区点位置，a[left]正好是右侧的第一个元素
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

func test(m map[string]string) {
	m["1"] = "111"
	m["2"] = "222"
}

func test2(sli []string) {
	sli = append(sli, "aaa")
	sli = append(sli, "bbb")
}
