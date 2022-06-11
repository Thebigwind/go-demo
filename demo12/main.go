package demo12

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//按行读
func ReadByLine() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
}

//----------------------------------------------
//按行读，指定bufio大小
func ReadBylineBuffio() {

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReaderSize(f, 4096) //指定buffer大小
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
}

//------------------------------------------------
//按行读，ReadBytes

func ReadBytes() {

	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		fmt.Println(string(b))
	}
}

//---------------------------------------------------
//按行读，scanner
func HandleText(textfile string) error {
	file, err := os.Open(textfile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // or
		//line := scanner.Bytes()

		//do_your_function(line)
		fmt.Printf("%s\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
		return err
	}

	return nil
}

//-----------------------------------------------------------
//读整个文件
func ReadFile() {
	b, err := ioutil.ReadFile("app-2019-06-01.log") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'
}
