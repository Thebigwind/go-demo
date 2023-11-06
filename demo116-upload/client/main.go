package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"time"
)

// 分片大小设置为10M
const chunkSize = 10 * 1024 * 1024

//请求体
type Request struct {
	Method          string
	Url             string
	Timeout         time.Duration
	Headers         map[string]string //设置请求header
	Values          map[string]string //普通post字段内容
	Files           map[string]string //文件post字段内容
	body            *bytes.Buffer
	multipartWriter *multipart.Writer
	IsChunk         bool //是否分片上传
}

/*
1. 将待上传的大文件分割成多个小文件，每个小文件大小相同；
2. 将小文件依次上传到服务端，服务端会将这些小文件合并成完整的文件。

客户端的实现流程如下：
1. 打开待上传的大文件；
2. 将文件分割成多个小文件，每个小文件大小相同（最后一个小文件大小可能会不同）；
3. 将每个小文件依次上传到服务端；
4. 服务端会将这些小文件合并成完整的文件。
*/
func main() {
	//1.文件切片上传
	var r1 = &Request{
		Method:  "POST",
		Url:     "http://localhost:50116/upload",
		Timeout: 0,
		Headers: make(map[string]string),
		Files:   make(map[string]string),
		IsChunk: true,
	}
	r1.Files["file1"] = "./cc.txt"
	r1.Upload()

	//// 2.文件上传
	//var r2 = &Request{
	//	Method:  "POST",
	//	Url:     "http://localhost:50116/upload",
	//	Timeout: 3 * time.Second,
	//	Headers: make(map[string]string),
	//	Values:  make(map[string]string),
	//	Files:   make(map[string]string),
	//}
	//r2.Files = map[string]string{"file1": "./aa.txt", "file2": "./bb.txt"}
	//r2.Values["n1"] = "100"
	//r2.Values["n2"] = "200"
	//r2.Upload()
}

func (r *Request) Upload() {
	if r.IsChunk {
		r.MultipartUpload()
	} else {
		r.FileUpload()
	}
}

//分片上传
func (r *Request) MultipartUpload() {
	if !r.IsChunk {
		return
	}
	// 多文件
	for formName, filename := range r.Files {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		for i := 0; ; i++ {
			// 每次都需要重新初始化，防止文件内容重复（需要优化）
			r.body = &bytes.Buffer{}
			r.multipartWriter = multipart.NewWriter(r.body)
			var buf = make([]byte, chunkSize) //此处chunkSize是读取分片的核心设置
			n, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			// 创建表单数据
			if err := r.addFileToRequset(formName, file.Name()); err != nil {
				panic(err)
			}
			if err := r.multipartWriter.Close(); err != nil {
				panic(err)
			}

			//组装header信息
			r.Headers = map[string]string{
				"Content-Type": r.multipartWriter.FormDataContentType(),
				"chunk-number": fmt.Sprintf("%d", i),
			}
			//最后一片
			if n < chunkSize {
				r.Headers["chunk-final"] = "true"
			}
			// 发送 POST 请求上传文件
			r.doHttp()
		}
	}

}

func (r *Request) FileUpload() {
	// 创建多部分上传请求
	r.body = &bytes.Buffer{}
	//通过包multipart实现了MIME的multipart解析
	r.multipartWriter = multipart.NewWriter(r.body)
	if err := r.multiTypeUpload(); err != nil {
		panic(err)
	}
	// 发送 POST 请求上传文件
	r.Headers = map[string]string{
		"Content-Type": r.multipartWriter.FormDataContentType(),
	}
	r.doHttp()
}

// 创建http请求
func (r *Request) doHttp() {
	req, err := http.NewRequest(r.Method, r.Url, r.body)
	if err != nil {
		panic(err)
	}
	// 设置头信息为：multipart/form-data
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	// 实例化一个http客户端对象
	client := &http.Client{
		Timeout: r.Timeout,
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	resBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		reqDump, _ := httputil.DumpRequest(req, false)
		fmt.Println(string(reqDump))
	} else {
		fmt.Println("Upload error:")
	}
	fmt.Println(resp.Status, string(resBody))
}

// 创建多种类型的表单数据
func (r *Request) multiTypeUpload() error {
	var err error = nil
	if r.Values != nil {
		for k, v := range r.Values {
			err = r.multipartWriter.WriteField(k, v)
		}
		if err != nil {
			return err
		}
	}
	if r.Files != nil {
		for k, v := range r.Files {
			err = r.addFileToRequset(k, v)
		}
		if err != nil {
			return err
		}
	}
	if err := r.multipartWriter.Close(); err != nil {
		return err
	}
	return err
}

// 创建表单数据
func (r *Request) addFileToRequset(fileName, file string) error {
	_, err := r.multipartWriter.CreateFormFile(fileName, filepath.Base(file))
	if err != nil {
		return err
	}
	// 打开文件
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// 拷贝文件内容到新建FormFile中
	if r.IsChunk {
		_, err = io.CopyN(r.body, f, chunkSize)
	} else {
		_, err = io.Copy(r.body, f)
	}
	if err != nil {
		return err
	}
	return nil
}
