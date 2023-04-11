

## ngin x安装：

###  rpm方式：

https://www.bbsmax.com/A/lk5aPMelJ1/

```bash
查看 yum 源是否存在
yum list | grep nginx
如果不存在 或者 不是自己想要的版本 可以自己设置Nginx的源
用vim 打开 (没有会自己创建)
vim /etc/yum.repos.d/nginx.repo
写入如下代码 (官方提供的放心用)
[nginx]
name=nginx repo
baseurl=http://nginx.org/packages/OS/OSRELEASE/$basearch/
gpgcheck=0
enabled=1
完成后，你需要修改一下对应的操作系统和版本号，因为我的是centos和7的版本，所以改为这样
baseurl=http://nginx.org/packages/centos/7/$basearch/
一切就绪 安装 Nginx
yum install nginx
查看 Nginx 的版本
Nginx -v
启动服务
systemctl start nginx.service
关闭服务
systemctl stop nginx.service
重启
systemctl restart nginx.service
查看服务的启动状态
ps aux | grep nginx
查看端口占用情况
netstat -tlnp
```



### 源码方式：

  安装脚本：  

```bash
#!/bin/bash
#https://www.cnblogs.com/lixiangang/p/16512468.html

result(){
  if [ $? -eq 0 ]; then
    echo "$1成功。"
  else
    echo "$1失败"
    exit 1
  fi
}


if [ -d "/opt/nginx-1.21.4" ]; then
  echo "/opt/nginx-1.21.4 已存在"
  exit 1;
fi

cd /opt && wget http://nginx.org/download/nginx-1.21.4.tar.gz
result "下载包"

#mkdir -p /opt/nginx
tar -vxzf nginx-1.21.4.tar.gz
result "解压包"

#编译
cd /opt/nginx-1.21.4 && ./configure
result "configure"
#
#执行make命令(要是执行不成功请检查最开始安装的四个依赖有没有安装成功)
#yum -y install openssl openssl-devel make zlib zlib-devel gcc gcc-c++ libtool    pcre pcre-devel
echo "开始make"
make
result "make编译"

echo "开始make install"
make install
result "make install"

echo "指定配置文件启动"
 /usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
result "启动nginx"

#安装成功：浏览器输入ip.index.html (如http://192.168.14.110/index.html) 显示欢迎Welcome to nginx!
```



## 常用命令

普通启动服务：/usr/local/nginx/sbin/nginx
配置文件启动：/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf

暴力停止服务：/usr/local/nginx/sbin/nginx -s stop

优雅停止服务：/usr/local/nginx/sbin/nginx -s quit

检查配置文件：/usr/local/nginx/sbin/nginx -t

重新加载配置：/usr/local/nginx/sbin/nginx -s reload

查看相关进程：ps -ef | grep nginx



## 配置文件：

#### 整体结构图

 ![图片](https://mmbiz.qpic.cn/mmbiz_png/fEsWkVrSk56QLvAACZRLNEWPf4BcfKoichjDPI9oZ4RsLH4Lde39kvKIgxQSfoDDDiaUDPF5ic9KiaFVvUUbaBERkQ/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

#### 配置演示图；

![图片](https://mmbiz.qpic.cn/mmbiz_png/fEsWkVrSk56QLvAACZRLNEWPf4BcfKoicCtjqfAibuOyz2TBvghgHDevHvwZAaHWOiaj0AHjXRZlQ8ffzQwDMhRew/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)



**全局块**

配置影响Nginx全局的指令。主要包括：

- 配置运行Nginx服务器用户（组）
- worker process数
- Nginx进程
- PID存放路径错误日志的存放路径
- 一个Nginx进程打开的最多文件描述符数目

```bash
#配置worker进程运行用户（和用户组），nobody也是一个Linux用户，一般用于启动程序，没有密码
user nobody;
#user www www;

#配置工作进程数目，根据硬件调整，通常等于CPU数量或者2倍于CPU数量
worker_processes 1;

#配置全局错误日志及类型，[debug | info | notice | warn | error | crit]，默认是error
error_log logs/error.log;
#error_log logs/error.log notice;
#error_log logs/error.log info;

#配置进程pid文件
pid logs/nginx.pid;

#一个nginx进程打开的最多文件描述符数目，理论值应该是最多打开文件数（系统的值ulimit -n）与Nginx进程数相除，但是Nginx分配请求并不均匀，所以建议与ulimit -n的值保持一致。
worker_rlimit_nofile 65535;
```



**events块**

配置影响Nginx服务器或与用户的网络连接。主要包括：

- 事件驱动模型的选择
- 最大连接数的配置

```bash
#参考事件模型，use [ kqueue | rtsig | epoll | /dev/poll | select | poll ]; 
#epoll模型是Linux 2.6以上版本内核中的高性能网络I/O模型，如果跑在FreeBSD上面，就用kqueue模型。
use epoll;

#单个进程最大连接数（最大连接数=连接数*进程数）
worker_connections 65535;
```



**http块**

可以嵌套多个server，配置代理，缓存，日志定义等绝大多数功能和第三方模块的配置。主要包括：

- 定义MIMI-Type
- 自定义服务日志
- 允许sendfile方式传输文件
- 连接超时时间
- 单连接请求数上限

```bash
#常见的一些基础配置
include mime.types; #文件扩展名与文件类型映射表
default_type application/octet-stream; #默认文件类型
charset utf-8; #默认编码
server_names_hash_bucket_size 128; #服务器名字的hash表大小
client_header_buffer_size 32k; #上传文件大小限制
large_client_header_buffers 4 64k; #设定请求缓冲
client_max_body_size 8m; #设定请求缓冲
sendfile on; #开启高效文件传输模式，对于普通应用设为on，如果用来进行下载等应用磁盘IO重负载应用，可设置为off，以平衡磁盘与网络I/O处理速度，降低系统的负载。注意：如果图片显示不正常把这个改成off。
autoindex on; #开启目录列表访问，合适下载服务器，默认关闭。
tcp_nopush on; #防止网络阻塞
tcp_nodelay on; #防止网络阻塞
keepalive_timeout 120; #长连接超时时间，单位是秒

#FastCGI相关参数是为了改善网站的性能：减少资源占用，提高访问速度。
fastcgi_connect_timeout 300;
fastcgi_send_timeout 300;
fastcgi_read_timeout 300;
fastcgi_buffer_size 64k;
fastcgi_buffers 4 64k;
fastcgi_busy_buffers_size 128k;
fastcgi_temp_file_write_size 128k;

#gzip模块设置
gzip on; #开启gzip压缩输出
gzip_min_length 1k; #最小压缩文件大小
gzip_buffers 4 16k; #压缩缓冲区
gzip_http_version 1.0; #压缩版本（默认1.1，前端如果是squid2.5请使用1.0）
gzip_comp_level 2; #压缩等级
gzip_types text/plain application/x-javascript text/css application/xml; #压缩类型
gzip_vary on; #增加响应头'Vary: Accept-Encoding'
limit_zone crawler $binary_remote_addr 10m; #开启限制IP连接数的时候需要使用
```



**server块**

配置虚拟主机的相关参数，一个http中可以有多个server。主要包括：

- 配置网络监听

- 配置https服务

- 基于名称的虚拟主机配置

- 基于IP的虚拟主机配置

  ```bash
  #虚拟主机的常见配置
  server {
      listen       80; #配置监听端口
      server_name  localhost; #配置服务名
      charset utf-8; #配置字符集
      access_log  logs/host.access.log  main; #配置本虚拟主机的访问日志
      
      location / {
          root html; #root是配置服务器的默认网站根目录位置，默认为Nginx安装主目录下的html目录
          index index.html index.htm; #配置首页文件的名称
      }
      
      error_page 404             /404.html; #配置404错误页面
      error_page 500 502 503 504 /50x.html; #配置50x错误页面
  }
  
  #配置https服务，安全的网络传输协议，加密传输，端口443
  server {
      listen       443 ssl;
      server_name  localhost;
  
      ssl_certificate      cert.pem;
      ssl_certificate_key  cert.key;
  
      ssl_session_cache    shared:SSL:1m;
      ssl_session_timeout  5m;
  
      ssl_ciphers  HIGH:!aNULL:!MD5;
      ssl_prefer_server_ciphers  on;
  
      location / {
          root   html;
          index  index.html index.htm;
      }
  }
  ```

  

  **location块**

  配置请求的路由，以及各种页面的处理情况。主要包括：

  - 请求根目录配置更改

  - 网站默认首页配置

  - location的URI

    

### nginx.conf 

Nginx.conf文件是Nginx总配置文件，在我们搭建服务器时经常调整的文件。

```bash
user  nginx;  #运行用户，默认即是nginx，可以不进行设置
#Nginx进程，一般设置为和CPU核数一样
worker_processes  1;
#错误日志存放目录
error_log  /var/log/nginx/error.log warn;
#进程pid存放位置
pid        /var/run/nginx.pid;
events {
    worker_connections  1024; # 单个后台进程的最大并发数
}
http {
    include       /etc/nginx/mime.types;   #文件扩展名与类型映射表
    default_type  application/octet-stream;  #默认文件类型
    #设置日志模式
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                    '$status $body_bytes_sent "$http_referer" '
                    '"$http_user_agent" "$http_x_forwarded_for"';
    access_log  /var/log/nginx/access.log  main;   #nginx访问日志存放位置
    sendfile        on;   #开启高效传输模式
    #tcp_nopush     on;    #减少网络报文段的数量
    keepalive_timeout  65;  #保持连接的时间，也叫超时时间
    #gzip  on;  #开启gzip压缩
    include /etc/nginx/conf.d/*.conf; #包含的子配置项位置和文件
```

### conf.d目录

- 接最后一行 进入conf.d目录，然后使用vim default.conf进行查看

  ```bash
  server {
      listen       80;   #配置监听端口
      server_name  localhost;  //配置域名
      #charset koi8-r;
      #access_log  /var/log/nginx/host.access.log  main;
      location / {
          root   /usr/share/nginx/html;     #服务默认启动目录
          index  index.html index.htm;    #默认访问文件
      }
      #error_page  404              /404.html;   # 配置404页面
      # redirect server error pages to the static page /50x.html
      #
      error_page   500 502 503 504  /50x.html;   #错误状态码的显示页面，配置后需要重启
      location = /50x.html {
          root   /usr/share/nginx/html;
      }
      # proxy the PHP scripts to Apache listening on 127.0.0.1:80
      #
      #location ~ \.php$ {
      #    proxy_pass   http://127.0.0.1;
      #}
      # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
      #
      #location ~ \.php$ {
      #    root           html;
      #    fastcgi_pass   127.0.0.1:9000;
      #    fastcgi_index  index.php;
      #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
      #    include        fastcgi_params;
      #}
      # deny access to .htaccess files, if Apache's document root
      # concurs with nginx's one
      #
      #location ~ /\.ht {
      #    deny  all;
      #}
  }
  ```

  ### 请求转发：

  1.根据url前缀，转发到不同的服务器，
  
  ```bash
  server {
      listen       8800;
      #listen       101.43.181.85;
      server_name  101.43.181.85;
  
      location ~ /vod/ {
          proxy_pass http://127.0.0.1:8081;
          root   html;
          index  index.html index.htm;
      }
      location ~ /edu/ {
          proxy_pass http://127.0.0.1:8080;
          root   html;
          index  index.html index.htm;
      }
  
  }
  ```
  
  ```bash
  server {
      listen       80;
      server_name  localhost;
  
      #access_log  /var/log/nginx/host.access.log  main;
  
      location / {
          root /opt/dist;
          index  index.html index.htm;
      }
  
      location /api/v1/ {
          proxy_pass http://127.0.0.1:9801/v1/;
      }
      location /api/v2/ {
          proxy_pass   http://127.0.0.1:9801/v2/;
      }
      location /api/audit/v1/ {
          proxy_pass   http://127.0.0.1:9601/v1/;
      }
      location /api/ {
          proxy_pass http://127.0.0.1:9801/;
      }
  
      location /files/ {
          root /opt/qskm/static;
      }
  
      #error_page  404              /404.html;
  
      # redirect server error pages to the static page /50x.html
      #
      error_page   500 502 503 504  /50x.html;
      location = /50x.html {
          root   /usr/share/nginx/html;
      }
  
      # proxy the PHP scripts to Apache listening on 127.0.0.1:80
      #
      #location ~ \.php$ {
      #    proxy_pass   http://127.0.0.1;
      #}
  
      # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
      #
      #location ~ \.php$ {
      #    root           html;
      #    fastcgi_pass   127.0.0.1:9000;
      #    fastcgi_index  index.php;
      #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
      #    include        fastcgi_params;
      #}
  
      # deny access to .htaccess files, if Apache's document root
      # concurs with nginx's one
      #
      #location ~ /\.ht {
      #    deny  all;
      #}
  }
  ```
  
  

####  location指令匹配规则：

该指令用于匹配URL。

语法：

Location [ =  | ~ | ~* | ^~] uri {

}



通配符：

- =：用于不含正则表达式的uri前，要求请求字符串与uri严格匹配，如果匹配成功，就停止继续向下搜索并立即处理该请求。
- ~：用于表示uri包含正则表达式，并且区分大小写。
- ~*：用于表示uri包含正则表达式，并且不区分大小写。
- ^~：用于不含正则表达式的uri前，要求Nginx服务器找到标识uri和请求字符串匹配度最高的location后，立即使用此location处理请求，而不再使用location块中的正则uri和请求字符串做匹配。

注意：如果uri包含正则表达式，则必须要有~或者~*标识。

 

= :URI必须与指定模式精确匹配。这里的模式为简单字符串，不可使用正则表达式。

```text
# uri必须与/abcd精确匹配（优先级最高），如:
# http://www.website.com/abcd 可用
location = /abcd {
    ...
}
```

(无)：URI必须以指定模式开始，不可使用正则表达式。

```text
#uri必须与指定模式开始，以下uri可用:
# http://www.website.com/abcd
# http://www.website.com/abcd/
# http://www.website.com/abcde
location /abcd {
    ...
}
```

～：URI与指定的正则表达式模式匹配，区分大小写。

```text
#正则表达式^/abcd$,指定模式必须以/开始，后根abc，以d字符结尾。如：
http://website.com/abcd 可用
http://website.com/abcd?param1&param2 可用
location ~ ^/abcd$ {
    ...
}
```

### **修饰符优先级**

1. =
2. 无
3. ^~
4. ~或～*
5. (无)

###  

## Rewrite模块

nginx rewrite模块使用Perl兼容正则表达式库，在安装nginx时需要预先安装perl和perl-dev。

### 正则表达式

**元字符**

- ^ ：以^后跟的字符开始
- $ ：以$前的字符结尾
- . ：匹配任何字符
- [ ] ：组，匹配指定字符集合内的任何字符。
- [^ ] ：否定组，匹配不包括在指定字符集内的字符。
- | ：匹配符号“|”之前或之后的实体。
- ( ) ：分组，组成一组匹配的实体，常和“|”配合使用。
- \ ：对特殊字符进行转义。

**量词**
使用量词可以扩展实体匹配次数。

- *：0或多次
- +：1或多次
- ?：0或1次
- {x}：x次
- {x,}：至少x次
- {x,y}：x到y次

**捕获**
正则表达式最后一个功能是能够捕获自表达式，放在“( )”之间的任何文本，在被捕获后都能用于后面的处理。 例如：

```
server {
    server_name website.com;
    location ~* ^/(downloads|files)/(.*)$ {
        add_header capture1 $1;
        add_header capture2 $2;
    }
    
}
```

### 关于内部请求

nginx有两种类型的内部请求：

1. 内部重定向： nginx在内部重定向客户端的请求。URI被改变，请求可能匹配到其他的location区段。常见的内部重定向指令有error_page、rewrite、try_files。
2. 子请求： 另一种触发内部请求而产生内容的就是子请求，如使用Addition模块或SSI模块的指令。

注意：error_log 日志级别设置为debug时能够捕获到内部请求日志，详情可参阅error_log指令。



### rewrite模块指令

**条件结构**
Rewrite模块引入了if条件结构。
语法：if (condition) { ... }
作用域：server, location

条件运算符说明：:

- 一个变量名 : 如果指定的变量或数据不等于空字符串或0，则条件为真。
- =或!= : 相等或不等比较
- ~ , ~* : 模式匹配，除此之外还有,!,!*
- -f,!-f : 测试文件是否存在
- -d,!-d : 测试一个目录是否存在
- -e,!-e : 测试文件、目录或符号链接是否存在
- -x,!-x : 测试文件是否存在且可执行。

例如：

```text
if ($http_user_agent ~ MSIE) {
        rewrite ^(.*)$ /msie/$1 break;
    }
    
    if ($http_cookie ~* "id=([^;]+)(?:;|$)") {
        set $id $1;
    }
    
    if ($request_method = POST) {
        return 405;
    }
    
    if ($slow) {
        limit_rate 10k;
    }
    
    if ($invalid_referer) {
        return 403;
    }
```

注意：if和location都能实现类似的效果，它们的不同之处在于能够在这两个区段使用的指令不同。换句话说就是有些指令可以用在if区段，有的则不可以，但几乎所有指令都可以在location中使用。比较常见的是在if区段中使用rewrite指令。

**rewrite**
语法：rewrite regex replacement [flag];
如果指定的正则表达式与URI匹配，则URI将按照replacement指定的字符串进行更改。该rewrite指令将按其在配置文件中出现的顺序执行。可以使用flag终止指令的进一步处理。如果replacement以“http://”、“https://”或“$scheme”开头，则处理将停止并将重定向返回给客户端。 flag标志说明：

- last：停止处理后面的rewrite指令集，并使用新的URI重新发起一个内部请求。
- break:停止处理后面的rewrite指令集，相当于执行了break指令。如果重定向后的文件不存在返回404。
- redirect:临时重定向，返回302 Moved temporarily HTTP 重定向响应，Location Header显示跳转后的URL地址。
- permanent：永久重定向，返回301响应。

```text
server {
    ...
    rewrite ^(/download/.*)/media/(.*)\..*$ $1/mp3/$2.mp3 last;
    rewrite ^(/download/.*)/audio/(.*)\..*$ $1/mp3/$2.ra  last;
    return  403;
    ...
}
# 如上配置，如果放到location区块中，则last标志应该替换称break，否则nginx将执行10次循环，并返回500错误。
location /download/ {
    rewrite ^(/download/.*)/media/(.*)\..*$ $1/mp3/$2.mp3 break;
    rewrite ^(/download/.*)/audio/(.*)\..*$ $1/mp3/$2.ra  break;
    return  403;
}
```

注意：

- 如果URI重写后发生循环，则重复该循环，但不超过10次。nginx将返回“500 Internal Server Error”。
- last和break都不再匹配后面的rewrite规则，但last应用完当前规则后，会重新发起一个内部请求查找一个location.

**break**
停止处理后面的rewrite指令集，阻止进一步改写请求URI。

```text
if (-e $request_filename) {
    break;
}
//因为break指令，后面的指令集不会执行
if ($uri ~ ^/search/(.*)$){
    rewrite ^/search.php?q=$1;
}
```

**return**
中断请求处理过程，并返回一个指定的HTTP状态码。
语法：return code [text] | [codURL] ; return URL;

```text
if ($uri ~ ^/admin/) {
    return 403;
    #因为已经完成了请求，此条指令不会执行。
    rewrite ^ http://website.com; 
}
```

**set**
初始化或重定义一个变量。注意有些变量不能修改，如$uri。
语法：set $variable value;

**rewrite_log**
如果设置为on，nginx将在“notice”错误级别对rewrite引擎处理的每个操作记录日志。
默认值：off



### 正向代理

用正向代理来进行上网等功能, 正向代理：如果把局域网外的 Internet 想象成一个巨大的资源库，则局域网中的客户端要访 问 Internet，则需要通过代理服务器来访问，这种代理服务就称为正向代理。

通过代理服务器来访问真实服务器的过程 就叫正向代理。**正向代理需要在客户端进行配置代理服务器**

例如,我们用的vpn ,国内直接访问不了google , 通过vpn的服务器代理我们访问对应的站点,就可以访问



### 反向代理

反向代理 是代理服务器端, 当客户端发送请求时,请求先到代理服务器,代理服务器再把**请求选择**发送到目标服务器,代理服务器获取目标服务器的数据后,在返回给客户端.这个过程代理服务器代理的是真实服务器端,客户端对代理无感知,不知道请求的真实服务器的IP地址.

在不用反向代理的时候, 我们请求的是真实服务器,可以知道真实服务器的ip; 反向代理之后,请求的是代理服务器,代理服务器访问真实服务器,从而隐藏了真实服务器,而暴露的是代理服务器,**代理服务器需要配置真实服务器**



###  负载均衡

nginx向多个服务做负载均衡的转发

通过增加服务器的数量,将请求分发到各个服务器上,将原先的请求集中到单个服务器上的情况改为 将请求分发到多个服务器上,将负载分发到不同的服务器,也就是我们所说的负载均衡

Nginx服务器的负载均衡策略可以划分为两大类：即内置策略和扩展策略。

内置策略主要包含轮询、加权轮询和IP hash三种；

扩展策略主要通过第三方模块实现，种类比较丰富，常见的有url hash、fair等。

在传统的客户端和服务端的交互中, 客户端直接请求服务端。但是随着业务量越来越大,单一架构下服务器需要处理的也越来越多. 服务器的并发能力是有限的。



#### Demo1-轮询算法:

**每个请求按时间顺序逐一分配到不同的后端服务器，如果后端服务器 down 掉，能自动剔除**

监听了8801 端口, 配置文件中 upstream myserver 配置在 http全局块中,在location 中配置proxy_pass,并指向 upstream myserver.

```bash
upstream myserver{
server 101.43.181.85:8080;
server 101.43.181.85:8081;
}

    server {
        listen       8801;
        server_name  101.43.181.85;
        location / {
            proxy_pass http://myserver;
            root   html;
            index  index.html index.htm;
        }
    }

```



#### Demo2-加权轮询(权重)算法（weight)

在组中,url后面多了 weight, 配置了权重.当为服务器指定权重参数时， 权重将作为负载均衡决策的一部分

```bash
upstream myserver{
server 101.43.181.85:8080 weight=5;
server 101.43.181.85:8081 weight=10;
}

    server {
        listen       8801;
        server_name  101.43.181.85;
        location / {
            proxy_pass http://myserver;
            root   html;
            index  index.html index.htm;
        }
    }

```



Demo3-iphash

- ip_hash是根据用户请求过来的ip，然后映射成hash值，然后分配到一个特定的服务器里面；
- 使用ip_hash这种负载均衡以后，可以保证用户的每一次会话都只会发送到同一台特定的服务，它的session不会跨到其他的服务里面去的；



1.首先通过将ip地址映射成一个hash值，然后将hash值对服务的数量3取模，得到服务的索引0、1、2；

2.比如：5%3=2，则把这个请求发送到Tomcat3服务器，以此类推；

3.这样一来，只要用户的IP不发生改变，当前用户的会话就能够一直保持；

4.nginx的ip_hash算法是取ip地址的前三段数字进行hash映射，如果只有最后一段不一样，也会发送到同一个服务里面

```bash
    upstream myserver{
        ip_hash;
        server 101.43.181.85:8080 weight=5;
        server 101.43.181.85:8081 weight=10;
    }

    server {
        listen       8801;
        server_name  101.43.181.85;
        location / {
            proxy_pass http://myserver;
            root   html;
            index  index.html index.htm;
        }
    }
```

##### 注意：

一旦使用了ip_hash，当我们需要移除一台服务器的时候，不能直接删除这个配置项，而是需要在这台服务器配置后面加上关键字down，表示不可用

```bash
upstream myserver { 
	ip_hash; 
	server 192.168.11.73:8080; 
	server 192.168.11.74:8080 down; 
	server 192.168.11.75:8080; 
}
```



项目中demo:

```bash
upstream backend {
    server 127.0.0.1:9801;
    server 10.10.10.171:9801 backup;
}

upstream auditbackend {
    server 127.0.0.1:9601;
    server 10.10.10.171:9601 backup;
}

server {
    listen       80;
    server_name  localhost;

    #access_log  /var/log/nginx/host.access.log  main;

    location / {
        root /opt/dist;
        index  index.html index.htm;
    }

    location /api/v1/ {
        proxy_pass http://backend/v1/;
    }
    location /api/v2/ {
        proxy_pass   http://backend/v2/;
    }
    location /api/audit/v1/ {
        proxy_pass   http://auditbackend/v1/;
    }
    location /api/ {
        proxy_pass http://backend/;
    }

    location /files/ {
        root /opt/qskm/static;
    }

    #error_page  404              /404.html;

    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }

    # proxy the PHP scripts to Apache listening on 127.0.0.1:80
    #
    #location ~ \.php$ {
    #    proxy_pass   http://127.0.0.1;
    #}

    # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
    #
    #location ~ \.php$ {
    #    root           html;
    #    fastcgi_pass   127.0.0.1:9000;
    #    fastcgi_index  index.php;
    #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
    #    include        fastcgi_params;
    #}

    # deny access to .htaccess files, if Apache's document root
    # concurs with nginx's one
    #
    #location ~ /\.ht {
    #    deny  all;
    #}
}
```



## 动静分离

动静分离是指在web服务器架构中，将静态页面与动态页面或者静态内容接口和动态内容接口,分开在不同系统(服务器)访问的架构设计方法，进而提升整个服务访问性能和可维护性。

![img](https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fwww.pianshen.com%2Fimages%2F880%2Fef6f9210df1167b1877812174b2e4610.png&refer=http%3A%2F%2Fwww.pianshen.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1659152975&t=1606f8661cb505b64d59a68cacfc7f43)



Nginx 动静分离简单来说就是把动态跟静态请求分开，不能理解成只是单纯的把动态页面和静态页面物理分离。严格意义上说应该是动态请求跟静态请求分开，可以理解成使用 Nginx 处理静态页面，Tomcat 处理动态页面。动静分离从目前实现角度来讲大致分为两种，



另外一种方法就是动态跟静态文件混合在一起发布，通过 nginx 来分开。通过 location 指定不同的后缀名实现不同的请求转发。

通过 expires 参数设置，可以使浏览器缓存过期时间，减少与服务器之前的请求和流量。

具体 Expires 定义：是给一个资源设定一个过期时间，也就是说无需去服务端验证，直接通过浏览器自身确认是否过期即可，所以不会产生额外的流量。此种方法非常适合不经常变动的资源。（如果经常更新的文件，不建议使用 Expires 来缓存），我这里设置 3d，表示在这 3 天之内访问这个 URL，发送一个请求，比对服务器该文件最后更新时间没有变化，则不会从服务器抓取，返回状态码304，如果有修改，则直接从服务器重新下载，返回状态码 200。

我准备了两个静态资源

/data/test/www/a.html

/data/test/image/001.jpg

```bash
    server {
        listen       8802;
        server_name  101.43.181.85;
        location /www/ {
            root   /data/test/;
            index  index.html index.htm;
        }
        location /image/ {
        root /data/test/;
        autoindex on;
        }
    }
```

上述配置location 部分, 当我们路径为www时, 会匹配到 对应的location,到/data/test/www 路径下,寻找对应资源;

当我们路径为image时, 匹配到对应的location,到/data/test/image路径下,寻找对应资源



## nginx 高可用

nginx作为负载均衡器，所有请求都到了nginx，可见nginx处于非常重点的位置，如果nginx[服务器](https://www.yisu.com/)宕机后端web服务将无法提供服务，影响严重。

为了屏蔽负载均衡服务器的宕机，需要建立一个备份机。主服务器和备份机上都运行高可用(High Availability)监控程序，通过传送诸如“I am alive”这样的信息来监控对方的运行状况。当备份机不能在一定的时间内收到这样的信息时，它就接管主服务器的服务IP并继续提供负载均衡服务;当备份管理器又从主管理器收到“I am alive”这样的信息时，它就释放服务IP地址，这样的主服务器就开始再次提供负载均衡服务。

### **keepalived+nginx实现主备**

keepalived 是集群管理中保证集群高可用的一个服务软件，用来防止单点故障。

Keepalived的作用是检测web服务器的状态，如果有一台web服务器死机，或工作出现故障，Keepalived将检测到，并将有故障的web服务器从系统中剔除，当web服务器工作正常后Keepalived自动将web服务器加入到服务器群中，这些工作全部自动完成，不需要人工干涉，需要人工做的只是修复故障的web服务器。

#### keepalived工作原理

keepalived是以VRRP协议为实现基础的，VRRP全称Virtual Router Redundancy Protocol，即虚拟路由冗余协议。

虚拟路由冗余协议，可以认为是实现路由器高可用的协议，即将N台提供相同功能的路由器组成一个路由器组，这个组里面有一个master和多个backup，master上面有一个对外提供服务的vip(VIP = Virtual IP Address，虚拟IP地址，该路由器所在局域网内其他机器的默认路由为该vip)，master会发组播，当backup收不到VRRP包时就认为master宕掉了，这时就需要根据VRRP的优先级来选举一个backup当master。这样的话就可以保证路由器的高可用了。

keepalived主要有三个模块，分别是core、check和VRRP。core模块为keepalived的核心，负责主进程的启动、维护以及全局配置文件的加载和解析。check负责健康检查，包括常见的各种检查方式。VRRP模块是来实现VRRP协议的。



##### **安装KeepAlived可使用yum直接安装：**

```none
yum -y install keepalived
```

查看KeepAlived配置文件：

```none
cat /etc/keepalived/keepalived.conf
```

接下来就是要对该配置文件进行修改，该配置文件在/etc/keepalived/路径下。在默认的keepalive.conf里面还有 virtual_server，real_server 这样的配置，它是为lvs准备的。

参考：https://www.cnblogs.com/starn/p/starn_nginx.html



## nginx日志



### Log模块

该模块指令作用域：http,server,location
**access_log**:设置日志写入的路径，格式。
语法：access_log path [format [buffer=size] | off;

- path：日志记录的文件路径，路径中可以使用变量。
- format：用log_format指令声明一个模版名称。
- off：关闭日志记录

**log_format**:定义一个模版，用于表述日志中一个条目包含的内容。 语法：log_format template_name format_string; 默认的模版为combined。定义为：

```text
log_format combined '$remote_addr - $remote_user [$time_local] '
                    '"$request" $status $body_bytes_sent '
                    '"$http_referer" "$http_user_agent"';
```





## nginx原理

**Nginx的线程模型？**



Nginx默认采用多进程工作方式，Nginx启动后，会运行一个master进程和多个worker进程。其中master充当整个进程组与用户的交互接口，同时对进程进行监护，管理worker进程来实现重启服务、平滑升级、更换日志文件、配置文件实时生效等功能。worker用来处理基本的网络事件，worker之间是平等的，他们共同竞争来处理来自客户端的请求。



Nginx的进程模型如图所示：



![图片](https://mmbiz.qpic.cn/mmbiz_png/fEsWkVrSk56QLvAACZRLNEWPf4BcfKoicjKBy2C5aABhrtQP0Ixg4P2IgSAHuR8zEyekNdGWO2ia3XXbp4bnMdgA/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)



**8.2、worker的工作模式？**



worker对于连接是采用争抢的模式，谁先抢到就先交给谁处理，如果想要重新更新配置，由于已经抢到任务的worker不会参与争抢，那些空闲的worker就会去争抢连接，拿到连接后会自动更新配置信息，当那些有任务的worker完成任务后，会自动更新配置，这样就实现了无缝热部署。由于每个worker是独立的进程，如果有其中的一个worker出现问题，并不会影响其它worker继续进行争抢，在实现请求的过程，不会造成服务中断，建议worker数和服务器的CPU数相等是最为适宜的。



![图片](https://mmbiz.qpic.cn/mmbiz_png/fEsWkVrSk56QLvAACZRLNEWPf4BcfKoicBZQ1smtiaXq6hRKBfYQZbFpiaa5miaicCEkFQdz8ORkcLTuzprUoicCliapg/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)



**8.3、如何计算worker连接数？**



如果只访问Nginx的静态资源，一个发送请求会占用了woker的2个连接数

而如果是作为反向代理服务器，一个发送请求会占用了woker的4个连接数



**8.4、如何计算最大的并发数？
**



如果只访问nginx的静态资源，最大并发数量应该是：worker_connections * worker_processes / 2



而如果是作为反向代理服务器，最大并发数量应该是：worker_connections * worker_processes / 4



原文链接：https://blog.csdn.net/qq_38490457/article/details/108300342