#!/bin/bash
#参数：https://www.cnblogs.com/tangkaishou/p/10018244.html
result(){
  if [ $? -eq 0 ]; then
    echo "$1成功。"
  else
    echo "$1失败"
    exit 1
  fi
}

basepath=/usr/local/mongodb
#检查是否已安装mongo
if [ -d $basepath ]; then
  echo "$basepath 已存在"
  exit 1;
fi

#下载包
wget  https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-rhel70-5.0.14.tgz
result "下载包"

#解压包
tar -zxvf mongodb-linux-x86_64-rhel70-5.0.14.tgz
result "解压包"


#移动包到/usr/local/mongodb
mv mongodb-linux-x86_64-rhel70-5.0.14 $basepath
result "移动包到$basepath"

#建数据目录和日志目录
mkdir -p /datas/mongo/db
mkdir -p /datas/mongo/logs

#配置文件
echo "开始配置文件$basepath/bin/mongodb.conf"
cat>$basepath/bin/mongodb.conf<<EOF
#数据文件存放目录
dbpath = /datas/mongo/db
#日志文件存放目录
logpath = /datas/mongo/logs/mongodb.log
port = 27017  #端口
fork = true  #以守护程序的方式启用，即在后台运行
auth=true #
bind_ip=0.0.0.0
EOF
result "写配置文件mongodb.conf"

#环境变量
export PATH=$PATH:$basepath/bin
source /etc/profile
result "环境变量生效"

#服务启动
mongod -f $basepath/bin/mongodb.conf
result "服务启动"

#设置自启动
cat>/etc/init.d/mongodb<<EOF
#!/bin/sh
#
#chkconfig: 2345 80 90
#description: mongodb

if test -f /sys/kernel/mm/transparent_hugepage/enabled; then
   echo never > /sys/kernel/mm/transparent_hugepage/enabled
fi
if test -f /sys/kernel/mm/transparent_hugepage/defrag; then
   echo never > /sys/kernel/mm/transparent_hugepage/defrag
fi

start() {
/usr/local/mongodb/bin/mongod -f /usr/local/mongodb/bin/mongodb.conf
}

stop() {
/usr/local/mongodb/bin/mongod -f /usr/local/mongodb/bin/mongodb.conf --shutdown
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  restart)
    stop
    start
    ;;
  *)
 echo $"Usage: $0 {start|stop|restart}"
 exit 1
esac
EOF

result "设置自启动/etc/init.d/mongodb"

#可执行权限
chmod -R a+x /etc/init.d/mongodb
result "修改执行权限"

# 通过chkconfig 添加为自启动服务
chkconfig --add mongodb
result "添加为自启动服务"

chkconfig mongodb on
result "mongo自启动服务开启"

exit 0