```
mkdir /opt/data/qscs-ldap/database -p
mkdir /opt/data/qscs-ldap/config -p

docker run -itd -p 389:389 --name qscs-openldap \
-v /opt/data/qscs-ldap/database:/var/lib/ldap \
-v /opt/data/qscs-ldap/config:/etc/openldap/slapd.d \
--env LDAP_ORGANISATION="zdlz" \
--env LDAP_DOMAIN="zdlz.com" \
--env LDAP_ADMIN_PASSWORD="123456" \
--detach osixia/openldap:1.4.0
```

login DN ：cn=admin,dc=zdlz,dc=com
Password: 123456



```
docker run \
-p 8080:80 \
--privileged \
--name phpldapadmin \
--env PHPLDAPADMIN_HTTPS=false \
--env PHPLDAPADMIN_LDAP_HOSTS=10.10.10.125  \
--detach osixia/phpldapadmin
```

![](/Users/me/Library/Application Support/typora-user-images/image-20230425142028236.png)







qscs-add.ldif

```
dn: ou=People,dc=zdlz,dc=com
objectClass: organizationalUnit
ou: People

dn: ou=dev,dc=zdlz,dc=com
objectClass: organizationalUnit
ou: dev


dn: ou=golang,ou=dev,dc=zdlz,dc=com
objectClass: organizationalUnit
ou: golang

dn: ou=clang,ou=dev,dc=zdlz,dc=com
objectClass: organizationalUnit
ou: clang

dn: ou=java,ou=dev,dc=zdlz,dc=com
objectClass: organizationalUnit
ou: java
```



ldapadd -x -D cn=admin,dc=zdlz,dc=com -W -f qscs-add.ldif

![image-20230425133542355](/Users/me/Library/Application Support/typora-user-images/image-20230425133542355.png)



cert-base.ldif

```bash
dn: ou=certs,dc=zdlz,dc=com
objectClass: top
objectClass: organizationalUnit
ou: certs
```



root@90a1deb38ce6:~#  ldapadd -x -D cn=admin,dc=zdlz,dc=com -W -f qscs-base.ldif
Enter LDAP Password:
adding new entry "ou=certs,dc=zdlz,dc=com"

root@90a1deb38ce6:~#

![image-20230425134359316](/Users/me/Library/Application Support/typora-user-images/image-20230425134359316.png)





#### 添加一个 CA 证书
dn: cn=ca,ou=certs,dc=example,dc=com
objectClass: top
objectClass: device
cn: ca
userCertificate;binary:<CA_CERTIFICATE_DER_FORMAT>



#### 添加普通证书
dn: cn=user,ou=certs,dc=example,dc=com
objectClass: top
objectClass: person
cn: user
userCertificate;binary:<USER_CERTIFICATE_DER_FORMAT>



请注意，`userCertificate;binary` 属性的值应该是一个 DER 格式的证书，您需要使用 OpenSSL 命令将证书转换为 DER 格式。

例如，如果您有一个名为 `ca.crt` 的 PEM 格式证书文件，您可以使用以下命令将其转换为 DER 格式：

```

openssl x509 -in ca.crt -outform der -out ca.der
```

然后，您需要将 `userCertificate;binary` 属性的值更新为转换后的 DER 编码。在您的 `base.ldif` 文件中，该属性应该类似于以下内容：

```

userCertificate;binary:<CA_CERTIFICATE_DER_FORMAT>
```

将 `<CA_CERTIFICATE_DER_FORMAT>` 替换为实际证书的 DER 编码，例如：

```

userCertificate;binary:<0x30 0x82 0x01 0x23 ...>
```

确保在 `ldapadd` 命令中指定正确的 `base.ldif` 文件和管理员凭据，以便成功导入证书和 CRL。









