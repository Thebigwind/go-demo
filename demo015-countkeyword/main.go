package main

import (
	"bytes"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"strings"
	"time"
	"unsafe"
)

func main() {
	start := time.Now().UnixNano()
	Command()
	end := time.Now().UnixNano()
	fmt.Printf("耗时：%d ms\n", (end-start)/1e6)
}
func Command() {
	count := kingpin.Command("count", "统计命令.")
	count_keyword := count.Command("keyword", "统计关键字.")
	count_keyword_key := count_keyword.Flag("word", "关键字.").
		Short('w').
		String()
	count_keyword_filepath := count_keyword.Flag("path", "文件路径.").
		Short('p').
		String()
	count_keyword_goroutinenum := count_keyword.Flag("num", "线程数.").
		Short('n').
		Int()

	kingpin.CommandLine.HelpFlag.Short('h')
	osParse := kingpin.Parse()
	cmds := strings.Split(osParse, " ")
	switch cmds[0] {

	case "count":

		switch cmds[1] {

		case "keyword":

			filePath := *count_keyword_filepath
			keyWord := *count_keyword_key
			if *count_keyword_goroutinenum <= 1 {

				err, _ := CountHELLO(filePath, keyWord)
				if err != nil {
					fmt.Printf("err:%s\n", err.Error())
					return
				}
			} else {

			}
			return

		}
	}
}

func CountHELLO(filePath string, keyWord string) (error, int) {
	s1, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return err, 0
	}
	count := bytes.Count(s1, str2bytes(keyWord))
	fmt.Printf("%s count:%d\n", keyWord, count)
	return nil, count
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
