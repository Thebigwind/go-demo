# go-zero model生成

2022-04-25 17:47 更新

## model生成

首先，下载好[演示工程](https://raw.githubusercontent.com/zeromicro/go-zero-pages/master/cn/resource/book.zip)后，我们以user的model来进行代码生成演示。

## 前言

model是服务访问持久化数据层的桥梁，业务的持久化数据常存在于mysql，mongo等数据库中，我们都知道，对于一个数据库的操作莫过于CURD， 而这些工作也会占用一部分时间来进行开发，我曾经在编写一个业务时写了40个model文件，根据不同业务需求的复杂性，平均每个model文件差不多需要 10分钟，对于40个文件来说，400分钟的工作时间，差不多一天的工作量，而goctl工具可以在10秒钟来完成这400分钟的工作。

## 准备工作

进入演示工程book，找到user/model下的user.sql文件，将其在你自己的数据库中执行建表。

## 代码生成(带缓存)

### 方式一(ddl)

进入service/user/model目录，执行命令

```
$ cd service/user/model
$ goctl model mysql ddl -src user.sql -dir . -c
Done.
```

### 方式二(datasource)

```
$ goctl model mysql datasource -url="$datasource" -table="user" -c -dir .
Done.
```

> $datasource为数据库连接地址
