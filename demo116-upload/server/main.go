package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	RootPath      = "uploads/"
	ChunkRootPath = "uploads_tmp/"
)

var (
	// FilesMax 限制上传文件的大小为7 MB
	FilesMax int64 = 7 << 20
	// ValuesMax 限制POST字段内容的大小
	ValuesMax int64 = 512
)

/*
1. 接收客户端发送的文件分片；
2. 根据分片序号将分片保存到指定的位置；
3. 检查所有分片是否
已经上传完毕，如果已经上传完毕，则合并分片成完整的文件并将文件保存到指定位置。

服务端的实现流程如下：
1. 接收客户端发送的 POST 请求，检查是否是上传文件分片的请求；
2. 如果是文件分片上传的请求，则读取分片序号，并创建文件，将分片内容写入文件；
3. 如果是最后一个文件分片上传完成的请求，则合并所有分片，生成完整的文件，将所有分片删除。
*/

func main() {
	// 使服务器监听在任何可用的端口上“：0”
	l, err := net.Listen("tcp", ":50116")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Using port:", l.Addr())
	panic(http.Serve(l, nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			// 不支持的 Content-Type 类型
			fmt.Println("Invalid Content-Type: ", r.Header.Get("Content-Type"))
			http.Error(w, " 不支持的 Content-Type 类型", http.StatusBadRequest)
			return
		}

		// 整个请求的主体大小设置为7.5Mb
		r.Body = http.MaxBytesReader(w, r.Body, FilesMax+ValuesMax)
		reader, err := r.MultipartReader()
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for {
			// A Part represents a single part in a multipart body.
			part, err := reader.NextPart()
			if err != nil {
				if err == io.EOF {
					break
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fileName := part.FileName()
			formName := part.FormName()
			var buf = &bytes.Buffer{}
			// 非文件字段部分大小限制验证（非文件字段，go中filename会是空）
			if fileName == "" {
				var limitError = "请求主体中非文件字段" + formName + "超出大小限制"
				err = uploadSizeLimit(buf, part, ValuesMax, limitError)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				continue
			}

			// 文件字段部分大小限制验证
			var limitError = "请求主体中文件字段" + fileName + "超出大小限制"
			err = uploadSizeLimit(buf, part, FilesMax, limitError)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// 文件创建部分
			if err := uploadFileHandle(r.Header, fileName, buf); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// 非逻辑内容，仅为测试使用
			var chunkNumber = r.Header.Get("chunk-number")
			if chunkNumber == "" {
				http.Error(w, "文件"+fileName+"上传成功", http.StatusOK)
			} else {
				http.Error(w, "分片文件"+fileName+chunkNumber+"上传成功", http.StatusOK)
			}
		}
		return
	}
	http.NotFound(w, r)
}

// 上传内容大小限制
func uploadSizeLimit(buf *bytes.Buffer, part *multipart.Part, maxLimit int64, limitError string) error {
	n, err := io.CopyN(buf, part, maxLimit+1)
	if err != nil && err != io.EOF {
		return err
	}
	maxLimit -= n
	if maxLimit < 0 {
		return errors.New(limitError)
	}
	return nil
}

// 文件上传处理函数
func uploadFileHandle(header http.Header, fileName string, buf *bytes.Buffer) error {
	var chunkNumberStr = header.Get("chunk-number")
	// 1.普通文件上传处理
	if chunkNumberStr == "" {
		//创建文件并写入文件内容
		return createFile(RootPath+fileName, buf.Bytes())
	}
	// 2.分片文件上传处理
	//2.1读取分片编号
	chunkNumber, err := strconv.Atoi(chunkNumberStr)
	if err != nil {
		return err
	}
	//2.2创建分片文件并写入分片内容
	if err := createFile(fmt.Sprintf(ChunkRootPath+fileName+"%d.chunk", chunkNumber), buf.Bytes()); err != nil {
		return err
	}
	//2.3确认是否上传完毕
	if header.Get("chunk-final") == "true" {
		//2.4合并文件
		if err := mergeChunkFiles(fileName); err != nil {
			return err
		}
		//2.5删除分片
		for i := 0; ; i++ {
			chunFileName := fmt.Sprintf(ChunkRootPath+fileName+"%d.chunk", i)
			err := os.Remove(chunFileName)
			if err != nil {
				if os.IsNotExist(err) {
					break
				}
				return err
			}
		}
	}
	return nil
}

// 创建文件并写入内容
func createFile(fileName string, res []byte) error {
	newFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = newFile.Close()
	}()
	bufferedWriter := bufio.NewWriter(newFile)
	_, err = bufferedWriter.Write(res)
	if err != nil && err != io.EOF {
		return err
	}
	return bufferedWriter.Flush()
}

// 合并分片文件
func mergeChunkFiles(fileName string) error {
	var (
		n   int64
		err error
	)
	fmt.Println("来这里了")
	finalFile, err := os.Create(RootPath + fileName)
	if err != nil {
		return err
	}
	defer finalFile.Close()
	// 将分片内容写入最终文件
	for i := 0; ; i++ {
		chunFile, err := os.Open(fmt.Sprintf(ChunkRootPath+fileName+"%d.chunk", i))
		if err != nil {
			if os.IsNotExist(err) {
				break
			}
			return err
		}
		n, err = io.Copy(finalFile, chunFile)
		if err != nil {
			return err
		}
		err = chunFile.Close()
		if err != nil {
			return err
		}
		if n < 1 {
			break
		}
	}
	return nil
}
