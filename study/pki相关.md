**1.**生成根证书私钥

openssl genrsa -out root.key 2048



**2.** 创建证书签名请求配置文件



cat > root_csr.conf <<EOF

[ req ]

default_bits = 2048

prompt = no

default_md = sha256

distinguished_name = dn

req_extensions = v3_req



[ dn ]

C = CN

ST = Beijing

L = Beijing

O = zdlz

OU = zdlz

CN = root



[v3_req]

basicConstraints=critical,CA:TRUE

subjectAltName = @alt_names



[ alt_names ]

IP.1 = 127.0.0.1



EOF



**3.**使用根私钥生成证书签名请求 **(CSR)**

openssl req -new -key root.key -out root.csr -config root_csr.conf



**4.**自签名生成根证书，此处是自签**root**证书，其他情况则是找证书签发机构签发：

openssl x509 -req -in root.csr -out root.crt -signkey root.key -CAcreateserial -days 3650