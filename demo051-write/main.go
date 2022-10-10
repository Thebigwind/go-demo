package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func testWrite1(filepath string) {
	file, err := os.Create(filepath)
	checkError(err)
	defer file.Close()

	file.Write([]byte("hello world\n"))
	file.Write([]byte("测试中文\n"))
	file.WriteString("使用 Write String写入文件\n")
}

//使用fmt.Fprintln格式化写入文件
func testWrite2(filepath string) {
	file, err := os.Create(filepath)
	checkError(err)
	defer file.Close()

	fmt.Fprintln(file, "hello world")
	fmt.Fprint(file, "测试中文\n")
	fmt.Fprintln(file, 1, 2, 3)
	fmt.Fprintf(file, "name=%s, age=%d\n", "admin", 20)
	//hello world
	//测试中文
	//1 2 3
	//name=admin, age=20
}

func testWrite3(filepath string) {
	file, err := os.Create(filepath)
	checkError(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("hello world\n")
	writer.Write([]byte("测试中文\n"))

	//在执行完 Write 操作之后，应该要调用 Flush，将内容全部打到文件中去
	//否则文件中不会有内容
	writer.Flush()
}

func testWrite4(filepath string) {
	file, err := os.Create(filepath)
	checkError(err)
	defer file.Close()

	content := `这个是中文
hello world
天下无敌
	`

	err = ioutil.WriteFile(filepath, []byte(content), os.ModePerm)
	checkError(err)
}

//使用文件追加模式写入文件
func testWrite5(filepath string) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("杭州 浙江\n")
	}
	writer.Flush()
}

func testCopy() {
	src, err := os.Open("1.txt")
	checkError(err)
	defer src.Close()

	reader := bufio.NewReader(src)

	dest, err := os.Create("1_copy.txt")
	checkError(err)
	defer dest.Close()
	writer := bufio.NewWriter(dest)

	_, err = io.Copy(writer, reader)
	checkError(err)
}

func main() {
	testWrite1("1.txt")
	testWrite2("2.txt")
	testWrite3("3.txt")
	testWrite4("4.txt")
	testWrite5("1.txt")
	testCopy()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
