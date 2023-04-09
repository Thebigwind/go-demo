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