# Docker部署

## [#](https://mszlu.com/docker/04/04.html#_1-镜像分层原理)1. 镜像分层原理

> 为什么centos镜像只有200多MB？

因为centos镜像使用了宿主机的内核。

![image-20220312122525847](https://mszlu.com/assets/image-20220312122525847.2f48ad6c.png)

**base镜像的定义**

1. 不依赖其他镜像，从scratch构建
2. 其他镜像可以在其基础进行扩展

base镜像一般都是各种Linux发行版本的Docker镜像，比如：Ubuntu，Debian或者CentOS等。

Linux操作系统由用户空间和内核空间构成。

内核空间是`kernel`，用户空间是`rootfs`，不同发行版的区别主要是rootfs。比如Ubuntu 14.04使用 upstart 管理服务，apt 管理软件包；而 CentOS 7 使用 systemd 和 yum。这些都是用户空间的不同，Kernel差别不大。

所以Docker可以同时支持多种 Linux 镜像，模拟出不同的操作系统环境。

> 为什么jdk的镜像有500多MB？

![image-20220312123335633](https://mszlu.com/assets/image-20220312123335633.beaf40e0.png)

jdk镜像包含了rootfs和jdk本身，所以jdk的镜像要加上rootfs的大小，才是jdk镜像的大小。

> 为什么tomcat正常下载几十兆，镜像却要几百兆？

![image-20220312133435172](https://mszlu.com/assets/image-20220312133435172.8a2e6aac.png)

**分层说明：**

![image-20220312133514003](https://mszlu.com/assets/image-20220312133514003.388ca98c.png)

**修改时复制策略（copy-on-write）**

Docker通过一个修改时复制策略来保证base镜像的安全性，以及更高的性能和空间利用率。

- 当容器需要读取文件的时候

从最上层的镜像层开始往下找，找到后读取到内存中，若已经在内存中，可以直接使用。换句话说，运行在同一台机器上的Docker容器共享运行时相同的文件。

- 当容器需要修改文件的时候

从上往下查找，找到后复制到容器层，对于容器来说，可以看到的是容器层的这个文件，看不到镜像层里的文件，然后直接修改容器层的文件。

- 当容器需要删除文件的时候

从上往下查找，找到后在容器中记录删除，并不是真正的删除，而是软删除。这导致镜像体积只会增加，不会减少。

当容器需要增加文件的时候 直接在最上层的容器可写层增加，不会影响镜像层。

## [#](https://mszlu.com/docker/04/04.html#_2-部署)2. 部署

### [#](https://mszlu.com/docker/04/04.html#_2-1-tomcat部署)2.1 tomcat部署

1. 拉取镜像

```bash
docker pull tomcat
```

1. 创建容器

   ```bash
   docker run -di --name tomcat -v /mnt/docker/tomcat/webapps:/usr/local/tomcat/webapps -p 8080:8080 tomcat
   ```

2. 在webapps下添加ROOT文件夹，其中添加一个index.html

3. 访问 http://ip:8080/index.html 即可

### [#](https://mszlu.com/docker/04/04.html#_2-2-nginx部署)2.2 nginx部署

1. 拉取镜像

   ```bash
   docker pull nginx
   ```

2. 创建容器

   ```bash
   docker run -di --name nginx -p 80:80 -v /mnt/docker/nginx:/etc/nginx nginx
   ```

3. 访问 http://ip 即可

4. 通过nginx访问tomcat

   nginx.conf

   ```bash
   user  nginx;
   worker_processes  1;
   
   error_log  /var/log/nginx/error.log warn;
   pid        /var/run/nginx.pid;
   
   
   events {
       worker_connections  1024;
   }
   
   
   http {
       include       /etc/nginx/mime.types;
       default_type  application/octet-stream;
   
       log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                         '$status $body_bytes_sent "$http_referer" '
                         '"$http_user_agent" "$http_x_forwarded_for"';
   
       access_log  /var/log/nginx/access.log  main;
   
       sendfile        on;
       #tcp_nopush     on;
   
       keepalive_timeout  65;
   
       #gzip  on;
   
       server {
       	listen 80;
       	server_name localhost;
       	location / {
       		proxy_pass http://172.17.0.2:8080;
       	}
       }
   }
   ```

   

5. 访问http://ip/index.html

### [#](https://mszlu.com/docker/04/04.html#_2-3-mysql部署)2.3 mysql部署

1. 拉取镜像

   ```bash
   docker pull mysql:5.7
   ```

2. 创建镜像

   ```bash
   docker run -id \
   -p 3307:3306 \
   --name=c_mysql \
   -v /mnt/docker/mysql/conf:/etc/mysql/conf.d \
   -v /mnt/docker/mysql/logs:/logs \
   -v /mnt/docker/mysql/data:/var/lib/mysql \
   -e MYSQL_ROOT_PASSWORD=root \
   mysql:5.7
   ```

3. 可以进入容器访问，也可以通过3307 外部访问

### [#](https://mszlu.com/docker/04/04.html#_2-4-redis部署)2.4 Redis部署

1. 拉取镜像。

```bash
docker pull redis
```



1. 创建容器

```bash
docker run -di --name redis -p 6379:6379 redis
```



连接容器中的 Redis 时，只需要连接宿主机的 IP + 指定的映射端口即可。

### [#](https://mszlu.com/docker/04/04.html#_2-5-mongodb部署)2.5 MongoDB部署

1. 拉取镜像

```bash
docker pull mongo
```



1. 创建容器

```bash
docker run -di --name mongo -p 27017:27017 mongo
```

连接容器中的 MongoDB 时，只需要连接宿主机的 IP + 指定的映射端口即可。

### [#](https://mszlu.com/docker/04/04.html#_2-6-elasticsearch部署)2.6 Elasticsearch部署

1. 拉取镜像

```bash
docker pull elasticsearch:7.8.1
```



1. 创建容器，为了方便演示，修改 ES 启动占用内存大小。

```bash
docker run -e ES_JAVA_OPTS="-Xms256m -Xmx512m" -e "discovery.type=single-node" -di --name es -p 9200:9200 -p 9300:9300 -p 5601:5601 -v /mnt/docker/es/plugins:/usr/share/elasticsearch/plugins elasticsearch:7.8.1
```



1. 安装中文分词器

```bash
# 进入容器
docker exec -it es /bin/bash
# 安装中文分词器 
elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.8.1/elasticsearch-analysis-ik-7.8.1.zip
# 重启 es
docker restart es
```



1. 访问：http://ip:9200/

### [#](https://mszlu.com/docker/04/04.html#_2-7-solr部署)2.7 Solr部署

1. 拉取镜像

```bash
docker pull solr
```



1. 创建容器

```bash
docker run -di --name=solr -p 8983:8983 solr
```



1. 访问：http://ip:8983/solr/#/ 结果如下

### [#](https://mszlu.com/docker/04/04.html#_2-8-rabbitmq部署)2.8 RabbitMQ部署

1. 拉取镜像

```bash
docker pull rabbitmq
```



1. 创建容器

```bash
docker run -di --name rabbitmq -p 4369:4369 -p 5671:5671 -p 5672:5672 -p 15671:15671 -p 15672:15672 -p 25672:25672 rabbitmq
```



1. 进入容器并开启管理功能

```bash
# 进入容器
docker exec -it rabbitmq /bin/bash
# 开启 RabbitMQ 管理功能
rabbitmq-plugins enable rabbitmq_management
```



1. 访问：http://ip:15672/ 使用 `guest` 登录账号密码