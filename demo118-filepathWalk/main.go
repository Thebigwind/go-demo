package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// 获取目录下创建日期最新的文件名
func getLatestFile(dirPath string) (string, error) {
	var latestFileName string
	var latestModTime time.Time

	// 遍历目录
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 判断是否为文件
		if !info.IsDir() {
			// 获取文件修改时间
			modTime := info.ModTime()

			// 检查是否为创建日期最新的文件
			if modTime.After(latestModTime) {
				latestModTime = modTime
				latestFileName = info.Name()
			}
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return latestFileName, nil
}

func main() {
	// 目录路径
	dirPath := "/data/upload"

	// 获取目录下创建日期最新的文件名
	latestFileName, err := getLatestFile(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印结果
	fmt.Println("最新的文件名:", latestFileName)
	//test()
}

func test() {

	// 获取当前系统的架构信息
	architecture := runtime.GOARCH

	// 打印结果
	fmt.Println("当前系统架构:", architecture)

}

//func test2() {
//	command := "cat /etc/system-release | awk '{print $1}'"
//	sysType, err := common.Command(command)
//	if err != nil {
//		fmt.Errorf("查询系统类型err:%s", err.Error())
//	}
//
//	fmt.Printf("sysType:%s", sysType)
//}
