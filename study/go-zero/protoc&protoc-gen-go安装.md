

要将证书发布到OpenLDAP，您需要确保OpenLDAP中适当的属性已经设置。以下是一些您需要设置的属性：

1. userCertificate：这是用于存储用户证书的属性。您可以在用户的LDAP条目中添加此属性，并将其设置为用户的证书。
2. caCertificate：这是用于存储CA证书的属性。如果您使用的是自签名证书，那么您需要将CA证书添加到此属性中。
3. certificateRevocationList：这是用于存储撤销列表的属性。如果您需要撤销某个证书，您可以在此属性中添加相应的撤销条目。
4. userCertificate;binary：这是用于存储二进制证书的属性。如果您使用的是DER编码的证书，那么您需要将其添加到此属性中。
5. sslcertificatemode：这是用于控制OpenLDAP是否强制要求使用SSL / TLS进行通信的属性。如果您想要强制使用SSL / TLS，则应将此属性设置为“require”.

这些是在OpenLDAP中设置证书时需要考虑的一些属性。您需要确保这些属性已正确设置并与您的证书一起使用。





LDAP Result Code 17 "Undefined Attribute Type"通常意味着您正在尝试使用未定义的属性类型。在LDAP中，如果您使用未知的属性类型，服务器将返回此错误代码。

在您的情况下，您正在尝试使用 userCertificate 属性，但它需要使用二进制传输（binary transfer）。因此，您需要在属性名称后面添加“;binary”来指示此属性需要使用二进制传输。例如：

```
phpCopy code
userCertificate;binary
```

请注意，在使用二进制传输时，您需要确保您的证书已经以正确的方式编码为二进制格式。如果您的证书不是二进制格式，则您需要先将其转换为二进制格式，然后再将其添加到 userCertificate;binary 属性中。

如果您仍然遇到问题，请确保您的LDAP服务器支持使用 userCertificate;binary 属性，或者尝试使用其他属性来存储您的证书。



LDAP Result Code 17 "Undefined Attribute Type"通常意味着您正在尝试使用未定义的属性类型。在LDAP中，如果您使用未知的属性类型，服务器将返回此错误代码。

在您的情况下，您正在尝试使用 userCertificate 属性，但它需要使用二进制传输（binary transfer）。因此，您需要在属性名称后面添加“;binary”来指示此属性需要使用二进制传输。例如：

```
phpCopy code
userCertificate;binary
```

请注意，在使用二进制传输时，您需要确保您的证书已经以正确的方式编码为二进制格式。如果您的证书不是二进制格式，则您需要先将其转换为二进制格式，然后再将其添加到 userCertificate;binary 属性中。

如果您仍然遇到问题，请确保您的LDAP服务器支持使用 userCertificate;binary 属性，或者尝试使用其他属性来存储您的证书。







如果要在 LDAP 中存储证书，您需要确保将其添加到正确的对象类中。通常，存储证书的对象类是 `inetOrgPerson` 或 `user`。

在 `inetOrgPerson` 或 `user` 对象类中，您需要设置以下属性：

- `userCertificate;binary`: 存储证书的二进制数据，必须使用 `;binary` 后缀来指示它是二进制数据。
- `userCertificate;binary;certificateRevocationList`: 存储证书撤销列表（CRL）的二进制数据，同样需要使用 `;binary` 后缀。

以下是一个示例 LDIF 条目，它演示了如何在 `inetOrgPerson` 对象类中存储证书和 CRL：

```
dn: uid=testuser,ou=users,dc=example,dc=com
objectClass: inetOrgPerson
objectClass: user
cn: Test User
sn: User
userCertificate;binary: <binary_data_here>
userCertificate;binary;certificateRevocationList: <binary_data_here>
uid: testuser
```



golang添加吊销的证书到吊销列表，发布到openldap

要将吊销的证书添加到 OpenLDAP 的证书吊销列表（CRL）中，您需要执行以下步骤：

1. 将要吊销的证书转换为 DER 格式。您可以使用 OpenSSL 或其他工具将证书转换为 DER 格式，并将其保存到文件中。
2. 创建包含 CRL 分发点 URL 和更新日期的 `cRLDistributionPoint` 属性。这是 CRL 中必需的属性，它指定了 CRL 在何处可以获得以及何时更新。在此示例中，我们使用 `http://crl.example.com/my-ca.crl` 作为 CRL 分发点 URL，并使用当前日期作为更新日期。
3. 创建包含 CRL 中吊销的证书的序列号的 `certificateRevocationList;binary` 属性。在此示例中，我们将序列号设置为 `12345678`。
4. 创建新的 `ldap.Entry`，并将上述两个属性添加到条目中。设置条目的 `dn` 属性以指定 CRL 条目在目录中的位置。
5. 将新的 `ldap.Entry` 添加到 LDAP 目录中。

下面是一个示例代码片段，它演示了如何使用 Go LDAP 库将吊销的证书添加到 OpenLDAP 的证书吊销列表中：

```go
package main

import (
    "encoding/base64"
    "fmt"
    "time"

    "gopkg.in/ldap.v3"
)

func main() {
    // 读取要吊销的证书文件
    certData, err := ioutil.ReadFile("path/to/revoked_certificate.crt")
    if err != nil {
        panic(err)
    }

    // 将证书转换为DER格式
    derData, err := base64.StdEncoding.DecodeString(string(certData))
    if err != nil {
        panic(err)
    }

    // 创建LDAP客户端连接
    conn, err := ldap.Dial("tcp", "ldap.example.com:389")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // 进行身份验证
    err = conn.Bind("cn=admin,dc=example,dc=com", "password")
    if err != nil {
        panic(err)
    }

    // 创建CRL分发点属性
    cRLDistPointAttr := ldap.NewEntryAttribute("cRLDistributionPoint", []string{"http://crl.example.com/my-ca.crl"})
    cRLDistPointAttr.AddValue(fmt.Sprintf("20130510000000Z"))

    // 创建CRL条目属性
    crlAttr := ldap.NewEntryAttribute("certificateRevocationList;binary", []string{string(derData)})
    crlAttr.AddValue("12345678")

    // 创建新的LDAP条目并添加属性
    entry := ldap.NewEntry("cn=my-ca-crl,ou=crls,dc=example,dc=com", []ldap.EntryAttribute{cRLDistPointAttr, crlAttr})

    // 将条目添加到LDAP目录
    err = conn.Add(entry)
    if err != nil {
        panic(err)
    }

    fmt.Println("Revoked certificate added to CRL successfully")
}

```