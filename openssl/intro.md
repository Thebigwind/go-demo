https://www.cnblogs.com/wangao1236/p/11609429.html
https://www.cnblogs.com/bandaoyu/p/14625120.html
https://linux.cn/article-13368-1.html

OpenSSL 简介
OpenSSL 是一个开源项目，其组成主要包括三个组件：
openssl：多用途的命令行工具
libcrypto：加密算法库
libssl：加密模块应用库，实现了ssl及tls
OpenSSL 主要用于秘钥证书管理、对称加密和非对称加密


1.1 指令
常用指令包括：genrsa、req、x509

1.1.1 genrsa
主要用于生成私钥，选择算法、加密私钥使用的对称加密密码和秘钥长度
基本用法：openssl genrsa [args] [numbits]

[args]:
    args1 对生成的私钥文件是否要使用加密算法进行对称加密:
        -des : CBC模式的DES加密
            -des3 : CBC模式的3DES加密
            -aes128 : CBC模式的AES128加密
        -aes192 : CBC模式的AES192加密
        -aes256 : CBC模式的AES256加密
    args2 对称加密密码
        -passout passwords
        其中passwords为对称加密(des、3des、aes)的密码(使用这个参数就省去了console交互提示输入密码的环节)
    args3 输出文件
        -out file : 输出证书私钥文件
[numbits]: 密钥长度，理解为私钥长度 

生成一个 2048 位的 RSA 私钥，并用 des3 加密（密码为 123456），保存为 server.key 文件：
 openssl genrsa -des3 -passout pass:123456 -out server.key 1024



1.1.2 req
req 的基本功能主要有两个：生成证书请求和生成自签名证书（当然这并不是其全部功能，但是这两个最为常见）
基本用法：openssl req [args] outfile

[args]
    args1 是输入输入文件格式：-inform arg
        -inform DER 使用输入文件格式为 DER
        -inform PEM 使用输入文件格式为 PEM
    args2 输出文件格式：-outform arg   
        -outform DER 使用输出文件格式为 DER
        -outform PEM 使用输出文件格式为 PEM
    args3 是待处理文件
        -in inputfilepath
    args4 待输出文件
        -out outputfilepath
    args5 用于签名待生成的请求证书的私钥文件的解密密码
        -passin passwords       
    args6 用于签名待生成的请求证书的私钥文件
        -key file
    args7 指定输入密钥的编码格式 -keyform arg  
        -keyform DER
        -keyform NET
        -keyform PEM
    args8 生成新的证书请求
        -new
    args9 输出一个 X509 格式的证书,签名证书时使用
        -x509
    args10 使用 X509 签名证书的有效时间  
        -days // -days 3650 有效期 10 年
    args11 生成一个 bits 长度的 RSA 私钥文件，用于签发，与-key互斥，生成证书请求或者自签名证书时自动生成密钥，然后生成的密钥名称由 -keyout 参数指定
        -newkey rsa:bits
    args12 设置 HASH 算法-[digest]，指定对创建请求时提供的申请者信息进行数字签名时指定的 hash 算法
        -md5
        -sha1 // 高版本浏览器开始不信任这种算法
        -md2
        -mdc2
        -md4
    args13 指定 openssl 配置文件,很多内容不容易通过参数配置，可以指定配置文件
        -config filepath   
    args14 显示格式 txt（用于查看证书、私钥信息）
        -text


利用私钥生成证书请求 CSR：
openssl req -new -key server.key -out server.csr

利用私钥生成自签名证书 CRT：
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt


1.1.3 x509
x509 可以实现显示证书的内容、转换其格式、给 CSR 签名等X.509证书的管理工作
基本用法：openssl x509 [args]

[args]
    args1 是输入输入文件格式：-inform arg
        -inform DER 使用输入文件格式为 DER
        -inform PEM 使用输入文件格式为 PEM
    args2 输出文件格式:-outform arg
        -outform DER 使用输出文件格式为 DER
        -outform PEM 使用输出文件格式为 PEM
    args3 是待处理 X509 证书文件
        -in inputfilepath
    args4 待输出 X509 证书文件
        -out outputfilepath
    args5 表明输入文件是一个“请求签发证书文件（CSR）”，等待进行签发
        -req
    args6 签名证书的有效时间  
        -days // -days 3650 有效期 10 年
    args7 指定用于签发请求证书的根 CA 证书
        -CA arg
    args8 根 CA 证书格式（默认是 PEM）
        -CAform arg
    args9 指定用于签发请求证书的 CA 私钥证书文件
        -CAkey arg
    args10 指定根 CA 私钥证书文件格式（默认为 PEM 格式）
        -CAkeyform arg
    args11 指定序列号文件（serial number file）
        -CAserial arg
    args12 如果序列号文件（serial number file）没有指定，则自动创建它
        -CAcreateserial
    args12 设置 HASH 算法-[digest]，指定对创建请求时提供的申请者信息进行数字签名时指定的 hash 算法
        -md5
        -sha1 // 高版本浏览器开始不信任这种算法
        -md2
        -mdc2
        -md4

使用根 CA 证书 ca.crt 和私钥 ca.key 对“请求签发证书” server.csr 进行签发，生成 x509 格式证书：
openssl x509 -req -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out serverx509.crt


------------------ 具体使用 ---------------------

2. 具体使用
   2.1 生成 RSA 秘钥对
   使用 genrsa 生成 RSA 秘钥对：openssl genrsa -out server.key 2048

   2.2 生成身份证申请
   使用 req 命令，以之前的 server.key 为输入，生成一个身份证申请（CSR）文件：openssl req -nodes -new -key server.key -subj "/CN=localhost" -out server.csr
   CSR 的公钥从 server.key 中提取，域名是 localhost
   如果启动 https 服务，使用这个 CSR 签署的 CRT，客户端必须访问 localhost 才能访问到这个 HTTPS 服务
   关于配置多个域名和 IP 的 CSR，后面会介绍
   
   2.3 生成数字证书
   使用 x509 使用指定私钥 server,key 签署 server.csr，输出数字证书（ CRT）：openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt
   此处使用自身的私钥签署 CSR

   2.4 HTTPS 验证
   生成证书后，我们可以编写一个 Golang 的 https 服务验证刚刚生成的证书
   
服务端：
    `package main
    
    import (
    "io"
    "log"
    "net/http"
    )
    
    func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
    })
    if e := http.ListenAndServeTLS("0.0.0.0:5200", "/home/ao/Documents/certs/review/server.crt",
    "/home/ao/Documents/certs/review/server.key", nil); e != nil {
    log.Fatal("ListenAndServe: ", e)
    }
    //if e := http.ListenAndServe("0.0.0.0:5200", nil); e != nil {
    //    log.Fatal("ListenAndServe: ", e)
    //}
    }`

客户端：
    `
    package main
    
    import (
    "crypto/tls"
    "crypto/x509"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    )
    
    func loadCA(caFile string) *x509.CertPool {
    pool := x509.NewCertPool()
    
        ca, err := ioutil.ReadFile(caFile)
        if err != nil {
            log.Fatal("ReadFile: ", err)
        }
        pool.AppendCertsFromPEM(ca)
        return pool
    }
    
    func main() {
    //c := &http.Client{
    //    Transport: &http.Transport{
    //        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    //    }}
    c := &http.Client{
    Transport: &http.Transport{
    TLSClientConfig: &tls.Config{RootCAs: loadCA("/home/ao/Documents/certs/review/server.crt")},
    }}
    
        resp, err := c.Get("https://localhost:5200")
        if err != nil {
            log.Fatal("http.Client.Get: ", err)
        }
    
        defer resp.Body.Close()
        io.Copy(os.Stdout, resp.Body)
    }
`

    改变客户端请求域名为 127.0.0.1 时，客户端结果如下：
    $ go run client.go
    2019/09/30 15:11:41 http.Client.Get: Get https://127.0.0.1:5200: x509: cannot validate certificate for 127.0.0.1 because it doesn't contain any IP SANs
    exit status 1
    
    服务端输出如下：
    2019/09/30 15:11:41 http: TLS handshake error from 127.0.0.1:33596: remote error: tls: bad certificate
    
    若客户端保持 127.0.0.1 不变，改变 http.Transport 的 TLSClientConfig 为 insecure 配置时，可以正常返回：
    $ go run client.go
    2019/09/30 15:11:41 http.Client.Get: Get https://127.0.0.1:5200: x509: cannot validate certificate for 127.0.0.1 because it doesn't contain any IP SANs
    exit status 1
    此时就像浏览器保持信任网站证书，继续选择浏览的动作是一样的

   2.5 配置域名为 IP
    上述我们使用 localhost，作为 CSR 里的域名，导致请求时必须使用域名
    若我们使用本地 IP 作为域名呢：openssl req -nodes -new -key server.key -subj "/CN=127.0.0.1" -out server.csr
    此时客户端请求 https://127.0.0.1:5200 返回：
   $ go run client.go
   2019/09/30 15:19:24 http.Client.Get: Get https://127.0.0.1:5200: x509: cannot validate certificate for 127.0.0.1 because it doesn't contain any IP SANs
   exit status 1

   启用 insecure 时正常返回
   此处说明，我们不能简单通过 -subj "/CN=[IP]" 实现 CSR 中包含域名
   
   重新生成 CSR 和 CRT
   openssl genrsa -out server.key 2048
   $ echo subjectAltName = IP:127.0.0.1 > extfile.cnf
   $ openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -extfile extfile.cnf -out server.crt
   再次测试发现，请求 127.0.0.1 时可以了


   2.6 不使用自签名证书
    上述我们使用自签名证书，下面我们尝试模拟一个 CA 签署证书：
    首先生成 CA 的秘钥和自签名证书，中间不生成 CSR：
    $ openssl genrsa -out ca.key 2048
    $ openssl req -x509 -new -nodes -key ca.key -days 10000 -out ca.crt -subj "/CN=localhost.ca.com"
    生成私钥、证书，并使用 CA 签名：
    $ openssl genrsa -out server.key 2048
    $ openssl req -new -key server.key -subj "/CN=127.0.0.1" -out server.csr
    $ openssl x509 -req -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -extfile extfile.cnf
    经测试，正确


   2.7 CSR 配置多个域名和 IP
    上述提到 -subj 配置域名只能指定单个域名或者 IP
    下面介绍一下多域名 CSR 配置
    新增一个配置文件：

[req]
distinguished_name = req_distinguished_name
req_extensions = v3_req

[req_distinguished_name]
countryName = CN
countryName_default = CN
stateOrProvinceName = Beijing
stateOrProvinceName_default = Beijing
localityName = Beijing
localityName_default = Beijing
organizationName = WangAo
organizationName_default = WangAo
organizationalUnitName = Dev
organizationalUnitName_default = Dev
commonName = test.openssl.com
commonName_default = test.openssl.com
commonName_max = 64

[v3_req]
basicConstraints = CA:TRUE
subjectAltName = @alt_names

[alt_names]
DNS.1 = test.openssl.com
IP.1 = 127.0.0.1
IP.2 = 10.0.2.15   

我们为服务制定了一个域名和两个IP地址
下面利用之前的 CA，重新签署数字证书
openssl genrsa -out server.key 2048
$  openssl req -nodes -new -key server.key -out server.csr -subj "/CN=test.openssl.com"
$ openssl x509 -req -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -extensions v3_req -extfile openssl.cnf

CA 签署数字证书时制定了 -extensions 和 -extfile
分别向 127.0.0.1、10.0.2.15 和 test.openssl.com 请求，均可以成功







