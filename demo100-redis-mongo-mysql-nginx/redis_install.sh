#!/bin/bash
#https://cloud.tencent.com/developer/article/2184759

result(){
  if [ $? -eq 0 ]; then
    echo "$1成功。"
  else
    echo "$1失败"
    exit 1
  fi
}

basepath=/usr/local/redis-6.2.6
#如果文件夹存在，则退出，不执行安装脚本
if [ -d $basepath ]; then
  echo "$basepath 已存在"
  exit 1;
fi

#下载
cd /usr/local && wget http://download.redis.io/releases/redis-6.2.6.tar.gz
result "下载"

#解压包
echo "开始解压redis源码包..."
tar -xzf redis-6.2.6.tar.gz
result "解压"

#编译
echo "开始编译压redis源码..."
cd redis-6.2.6 && make
result "编译"

#安装
echo "开始安装redis。。。"
cd src && make install PREFIX=$basepath/install/
result "安装"

#建立数据目录和日志目录
echo "建立数据目录和日志目录"
mkdir -p /datas/redis/data
mkdir -p /datas/redis/logs

#判断日志文件是否已存在
if [ ! -f "/datas/redis/logs/redis.log" ];then
  touch /datas/redis/logs/redis.log
  result "建立数据目录和日志目录/datas/redis/logs/redis.log"
else
  echo "日志文件已存在，不需要重新创建"
fi

#修改配置：
echo "修改redis.conf配置"
#1、#bind 127.0.0.1 -::1（开头加#）
sed -i '75s/bind 127.0.0.1 -::1/#bind 127.0.0.1 -::1/' $basepath/redis.conf
#2、protected-mode no（修改为 no）
sed -i '94s/protected-mode yes/protected-mode no/' $basepath/redis.conf
#3、daemonize yes （修改为 yes ）
sed -i '257s/daemonize no/daemonize yes/' $basepath/redis.conf
#4、logfile 指定日志文件路径，若log目录不存在则需手动创建log目录
sed -i '302c logfile "/datas/redis/logs/redis.log"' $basepath/redis.conf
#5、dir 指定缓存目录路径，若data目录不存在则需手动创建data目录
sed -i '454c dir /datas/redis/data' $basepath/redis.conf


#复制 redis.conf 到 /usr/local/redis-6.2.6/install/bin/
cp $basepath/redis.conf $basepath/install/bin/
result "拷贝配置文件到$basepath/install/bin/"

#启动 redis 服务
cd $basepath/install/bin/ && ./redis-server redis.conf
result "启动redis"

#注册 redis 为服务，并设置开机自启动
echo "注册 redis 为服务，并设置开机自启动"
mkdir -p  /etc/redis/ && cp $basepath/install/bin/redis.conf /etc/redis/
result "拷贝配置文件"

echo "修改 redis 启动脚本"
#sed -i '14c REDISPORT=6379' $basepath/utils/redis_init_script
sed -i '15c EXEC=/usr/local/redis-6.2.6/install/bin/redis-server' $basepath/utils/redis_init_script
sed -i '16c CLIEXEC=/usr/local/redis-6.2.6/install/bin/redis-cli' $basepath/utils/redis_init_script
sed -i '19c CONF="/etc/redis/redis.conf"' $basepath/utils/redis_init_script
sed -i '19a chkconfig: 2345 80 90' $basepath/utils/redis_init_script


echo "将自启动脚本复制到系统启动目录下，并改名为redis"
cp $basepath/utils/redis_init_script /etc/init.d/redis
result "将自启动脚本复制到系统启动目录下，并改名为redis"

#增加执行权限
chmod a+x /etc/init.d/redis

#将redis注册成服务
echo "将redis注册成服务"
chkconfig --add redis
result "注册成服务"

exit 0