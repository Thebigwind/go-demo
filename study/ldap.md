https://wiki.shileizcc.com/confluence/pages/viewpage.action?pageId=39223494

OpenLDAP是开源的轻型目录访问协议（Lightweight Directory Access Protocol，LDAP）。它本身是一个小型文件数据库。Ldap是树形结构的，能够通过server + client(服务端+客户端)的方式。进行统一的用户(账号)管理。

举个栗子：如果有100台机器，一个用户需要登录这100台机器。传统的做法就是每台机器中，都需要创建登录账号，操作100次。想想都会疯掉。如果使用ldap来管理，就只需要在ldap服务中创建一次就可以了。账号清理也是类似的道理。
![image-20230424150318261](/Users/me/Library/Application Support/typora-user-images/image-20230424150318261.png)



当前登录用户为: cn=admin

![image-20230424151050813](/Users/me/Library/Application Support/typora-user-images/image-20230424151050813.png)



## LDAP简介

LDAP是一个数据库，但是又不是一个数据库。说他是数据库，因为他是一个数据存储的东西。但是说他不是数据库，是因为他的作用没有数据库这么强大，而是一个目录

从概念上说，LDAP分成了DN, OU等。OU就是一个树，DN就可以理解为是叶子，叶子还可以有更小的叶子。但是LDAP最大的分层按照IBM的文档是4层

LDAP的目的不是为了写，主要是为了查找，但并不是说LDAP不能写，只是说强项不是写。

LDAP作为一个统一认证的解决方案，主要的优点就在能够快速响应用户的查找需求。比如用户的认证，这可能会有大量的并发。如果用数据库来实现，由于数据库结构分成了各个表，要满足认证这个非常简单的需求，每次都需要去搜索数据库，合成过滤，效率慢也没有好处。虽然可以有Cache，但是还是有点浪费。LDAP就是一张表，只需要用户名和口令，加上一些其他的东西，非常简单。从效率和结构上都可以满足认证的需求。这就是为什么LDAP成为现在很人们的统一认证的解决方案的优势所在。

### LDAP的优势

- **读写效率高：**LDAP也是对读操作进行优化的一种数据库，在读写比例大于7比1的情况下，LDAP会体现出极高的性能。这个特性正适合了身份认证的需要
- **开放的标准协议：**不同于SQL数据库，LDAP的客户端是跨平台的，并且对几乎所有的程序语言都有标准的API接口。即使是改变了LDAP数据库产品的提供厂商，开发人员也不用担心需要修改程序才能适应新的数据库产品。这个优势是使用SQL语言进行查询的关系型数据库难以达到的。
- **强认证方式 ：**可以达到很高的安全级别。在国际化方面，LDAP使用了UTF-8编码来存储各种语言的字符。
- **OpenLDAP开源实现：**OpenLDAP还包含了很多有创造性的新功能，能满足大多数使用者的要求。笔者曾使用过许多商用LDAP产品，OpenLDAP是其中最轻便且消耗系统资源最少的一个。OpenLDAP是开源软件，近年国内很多公司开发的LDAP产品都是基于OpenLDAP开发的。
- **灵活添加数据类型：**LDAP是根据schema的内容定义各种属性之间的从属关系及匹配模式的。例如在关系型数据库中如果要为用户增加一个属性，就要在用户表中增加一个字段，在拥有庞大数量用户的情况下是十分困难的，需要改变表结构。但LDAP只需要在schema中加入新的属性，不会由于用户的属性增多而影响查询性能
- **数据存储是树结构：**整棵树的任何一个分支都可以单独放在一个服务器中进行分布式管理，不仅有利于做服务器的负载均衡，还方便了跨地域的服务器部署。这个优势在查询负载大或企业在不同地域都设有分公司的时候体现尤为明显

### LDAP的特点

1. LDAP 是一种网络协议而不是数据库，而且LDAP的目录不是关系型的，没有RDBMS那么复杂
2. LDAP不支持数据库的Transaction机制，纯粹的无状态、请求-响应的工作模式。
3. LDAP不能存储BLOB，LDAP的读写操作是非对称的，读非常方便，写比较麻烦，
4. LDAP支持复杂的查询过滤器(filter)，可以完成很多类似数据库的查询功能。
5. LDAP使用树状结构，接近于公司组织结构、文件目录结构、域名结构等我们耳熟能详的东
6. LDAP使用简单、接口标准，并支持SSL访问。

### LDAP的主要应用场景

1. 网络服务：DNS服务
2. 统一认证服务：
3. Linux PAM (ssh, login, cvs. . . )
4. Apache访问控制
5. 各种服务登录(ftpd, php based, perl based, python based. . . )
6. 个人信息类，如地址簿
7. 服务器信息，如帐号管理、邮件服务等



| 属性            | 描述                                                         |
| --------------- | :----------------------------------------------------------- |
| dn              | 唯一标识名类似于linux文件系统中的绝对路径，每个对象都有唯一标识名：uid=dpgdy,ou=people,dc=gdy,dc=com |
| rdn             | 通常指相对标识名，类似于linux系统中的相对路径，例如uid=dpgdy |
| uid             | 通常指一个用户的登录名称，例如uid=dpgdy，与系统中的uid不是一个概念 |
| sn              | 通常指一个人的姓氏，例如：sn:Guo                             |
| giceName        | 通常指一个人的名字，例如，giveName：Guodayyong，但是不能指姓氏 |
| objectClass     | 特殊属性，包括数据存储的方式及相关属性信息                   |
| dc              | 通常指一个域名：例如dc=example、dc=com                       |
| ou              | 通常指一个组织单元的名称。 例如ou=people，dcexample，dc=com  |
| cn              | 通常指一个对象的名称，如果是人，则是全名                     |
| mail            | 通常指登录账号的邮箱地址，例如 mail：dayong@126.com          |
| telephoneNumber | 通常指登录账号的手机号码，例如 telephoneNumber：XXXXXXXX     |
| c               | 通常指一个而为国家的名称，比如CN,US等国家代号，比如c:CN      |
| I               | 通常指一个地方的地名，例如 I：shanghai                       |
|                 |                                                              |

DN 的英文名称是(distinguished name)，直接翻译过来就是专有名称。简单的就可以理解为一个路径。

这个路径可以指到 OU ，也可以指到 CN。

其中 DN 有三个属性，分别是CN，OU，DC。

DC (Domain Component)

CN (Common Name)

OU (Organizational Unit)

O (Organization Name)  ，O 是可选项，有时候你不一定能够看得到。



如：dn = "ou=dev,dc=xyz,dc=com"

 DC 是最高的，叫做域名，基本上所有的 ldap 服务都会至少有一个 DC。

在 DC 下一级就会有一个 OU，OU 可以理解为一个组织单元，我们可以有多个组织单元。你可以在组织单元中组织用户组，也可以在组织单元中组织用户，你还可以在组织单元中组织组织单元。

在 OU 下面就是 CN 了，可以理解是 CN 就是一个具体的实例了，比如说一个具体的用户。

所以要定位一个实例，那么路径就是 CN - OU - DC

你可能会有多个 OU，多个 DC，但是最后都会定位到最高一级的 DC

这长串字符串放到一起，就是 DN 了。



## 组织顺序 这个组织顺序是逆序的。

举例来说，下面的 DN：

CN=cwikius,ou=Users,dc=jumpcloud,dc=com

实际的查找顺序是

DC=COM

DC=jumpcloud

OU=Users

CN=cwikius

最高一级的路径在最后面，如果理解为文件系统的查找路径的话就是：

COM/jumpcloud/Users/cwikius

最后的 CN=cwikius 可以理解为一个具体的文件，只是这个文件没有后缀罢了。



**下面列出部分常用objectClass要求必设的属性。**
● account：userid。
● organization：o。
● person：cn和sn。
● organizationalPerson：与person相同。
● organizationalRole：cn。
● organizationUnit：ou。
● posixGroup：cn、gidNumber。
● posixAccount：cn、gidNumber、homeDirectory、uid、uidNumber。



### LDAP协议的操作分为三大类：

查询操作：search, compare
更新操作：add, delete, modify, modify DN(rename)
认证和访问控制：bind, unbind, abandon



### LDAP信息模型

LDAP存储信息的基本单位是Entry，一个节点为一个Entry；
每个Entry有一套Attributes；
每个Attribute有Type和一个或者多个Values；
Type有语法规则(哪些值才能赋给这种类型的属性)和匹配规则；
匹配规则由比较规则和排序规则组成，比如caseIgnoreMatch和integerMatch；
Entry的属性是由Schema来定义的。

### LDAP功能模型

这里只介绍搜索，其他部分参考后文。在LDAP的搜索中，共有8个选项：

Base Object，搜索的起始根路径；
Search Scope，分三类：

​		    base，只检索Base Object; 

​            onelevel，检索Base Object下面的第一层目录；

​            sub，检索从Base Object开始的所有下层目录；
Dereferencing选项，是否解除别名节点的引用；
Size Limit，返回的Entries的数目，0为不限制；
Time Limit，0为不限制；
Attribute Only参数，true指示只返回属性类型，否则类型和值都返回；
Search Filter，搜索过滤条件；
要求搜索返回的属性列表，默认为都返回。

### LDIF文件及格式语法

LDIF为轻量级目录访问协议数据交换格式，是存储LDAP配置信息及目录内容的标准文本文件格式。

### LDIF文件存取OpenLDAP条目标准格式:

```yaml
# 注释，用于对条目进行解释  
dn：条目名称  
objectClass（对象类）： 属性值  
objectClass（对象类）： 属性值  
……
```

### LDIF格式范例：

```yaml
dn: uid=Guodayong,ou=people,dc=gdy,dc=com  //DN描述项，在整个目录树上为***的  
objectClass: top  
objectClass: posixAccount  
objectClass: shadowAccount  
objectClass: person  
objectClass: inetOrgPerson  
objectClass: hostObject  
sn: wang  
cn: wangxiaomei  
telephoneNumber：157****8900  
mail: wangxiaomei@126.com
```

以下为一个典型的添加数据类的格式：

```
dn: dc=example,dc=com
objectclass: dcObject
objectclass: organization
o: example
dc: example

dn: cn=Manager,dc=example,dc=com
objectclass: organizationalRole
cn: Manager

dn: ou=People,dc=example,dc=com
objectclass: top
objectclass: organizationalUnit
ou: People

dn: uid=yingyuan,ou=People,dc=example,dc=com
objectClass: Top
objectClass: Person
objectClass: OrganizationalPerson
objectClass: InetOrgPerson
uid: yingyuan
cn: Yingyuan Cheng
sn: Cheng
userPassword: yingyuan 
mail: yingyuan@staff.example.com.cn
description: A little little boy living in the big big 
 world.
jpegPhoto:: /9j/4AAQSkZJRgABAAAAAQABAAD/2wBDABALDA4MChAODQ4
 SERATGCgaGBYWGDEjJR0oOjM9PDkzODdASFxOQERXRTc4UG1RV19iZ2hnP
```

数据更新有好几种情况:

数据添加：

```
dn: uid=bjensen, ou=people, dc=example, dc=com
changetype: add
objectclass: top
objectclass: person
```

数据删除：

```
dn: uid=bjensen, ou=people, dc=example, dc=com
changetype: delete
```

数据修改：

```
dn: uid=bjensen, ou=people, dc=example, dc=com
changetype: modify
add: telephoneNumber
telephoneNumber: +1 216 555 1212
```

```
dn: uid=bjensen, ou=people, dc=example, dc=com
changetype: modify
delete: telephoneNumber
telephoneNumber: +1 216 555 1212
```

```
dn: uid=bjensen, ou=people, dc=example, dc=com
changetype: modify
replace: telephoneNumber
telephoneNumber: +1 216 555 1212
telephoneNumber: +1 405 555 1212
```

```
dn: uid=bjensen, ou=people, dc=example, dc=com
changetype: modify
add: mail
mail: bjensen@example.com
-
delete: telephoneNumber
telephoneNumber: +1 216 555 1212
-
delete: description
-
```

目录的移动和重命名：

```
dn: uid=bjensen, ou=People, dc=example, dc=com
changetype: moddn
newsuperior: ou=Terminated Employees, dc=example, dc=com
dn: uid=bjensen, ou=People, dc=example, dc=com
changetype: moddn
newrdn: uid=babsj
deleteoldrdn: 0
```

命令实战：

```
# 搜索主机ldap.example.com，范围为dc=example,dc=com以及其子目录，
# 过滤条件为cn=Barbara Jensen
$ ldapsearch -h ldap.example.com -s sub -b "dc=example,dc=com" "(cn=Barbara Jensen)“

# 仅搜索基目录，过滤条件为所有类
$ ldapsearch -h ldap.example.com -s base -b "uid=bjensen,ou=people,dc=example,dc=com" "(objectclass=*)"

# 搜索用户为uid=bjensen,ou=people,dc=exampe,dc=com
# 口令为hifalutin
$ ldapsearch -h localhost -D "uid=bjensen,ou=people,dc=example,dc=com" -w hifalutin -s sub -b "dc=example,dc=com" "(cn=Barbara Jensen)“

# 仅返回mail, roomNumber属性 
$ ldapsearch -h localhost -s sub -b "dc=example,dc=com" "(cn=Barbara Jensen)" mail roomNumber

# 返回所有属性和操作属性
$ ldapsearch -h localhost -s sub -b "dc=example,dc=com" "(cn=Barbara Jensen)" /
"*" modifiersName modifyTimeStamp

# 过滤条件为或连接
$ ldapsearch -h localhost -s sub -b "dc=example,dc=com" "(|(L=cupertino)(L=sunnyvale))"

# 过滤条件为复合条件
$ ldapsearch -h localhost -s sub -b "dc=example,dc=com" /
"(&(|(L=cupertino)(L=sunnyvale))(objectclass=person))"

# 从ldif文件中更新数据（updates.ldif含changetype）
$ ldapmodify –h ldap.example.com –D "cn=directory manager" –w secret < updates.ldif

# 从ldif文件中添加数据(不含changetype)
$ ldapmodify –h ldap.example.com –D "cn=directory manager" –w secret –a < updates.ldif

# 更新数据的过程中如果遇到错误则继续(-c)
# 并把错误写入rejects.ldif文件(-e rejects.ldif)
$ ldapmodify –h ldap.example.com –D "cn=directory manager" –w secret /
–c –e rejects.ldif < updates.ldif

```