#!/bin/bash
#参考：https://www.cnblogs.com/wendy-0901/p/12673705.html
#yum install -y cmake gcc gcc-c++ ncurses ncurses-devel perl

#以下卸载命令手动执行
#检测系统是否自带Mysql
# rpm -qa|grep mysql
#如果有进行强行卸载
# rpm -e --nodeps mysql-libs-5.1.52-1.el6_0.1.x86_64

#检测系统是否自带mariadb
# rpm -qa|grep mariadb
#rpm -e --nodeps xxx

result(){
  if [ $? -eq 0 ]; then
    echo "$1成功。"
  else
    echo "$1失败"
    exit 1
  fi
}



#下载包
 wget https://downloads.mysql.com/archives/get/p/23/file/mysql-5.7.28-linux-glibc2.12-x86_64.tar.gz
result "下载包"
#解压
tar -zxvf mysql-5.7.28-linux-glibc2.12-x86_64.tar.gz
result "解压包"
#重命名
mv mysql-5.7.28-linux-glibc2.12-x86_64 mysql-5.7.28
mv mysql-5.7.28 /usr/local/

#检查mysql组和用户是否存在，如果没有则创建
cat /etc/group|grep mysql

groupadd mysql
useradd -r -g mysql mysql
result "添加mysql用户和组"

#安装数据库
#创建数据目录
mkdir -p  /datas/mysql/data

#将/usr/local/mysql-5.7.28的所有者及所属组改为mysql
chown -R mysql:mysql /usr/local/mysql-5.7.28

#在/usr/local/mysql-5.7.28/support-files目录下创建my_default.cnf
cat>/usr/local/mysql-5.7.28/support-files/my_default.cnf<<EOF
[mysqld]

#设置mysql的安装目录
basedir=/usr/local/mysql-5.7.28/
#设置mysql数据库的数据存放目录
datadir=/datas/mysql/data/
#设置端口
port = 3306

socket = /tmp/mysql.sock
#设置字符集
character-set-server=utf8
#日志存放目录
log-error = $datadir/mysqld.log
pid-file = $datadir/mysqld.pid
#允许时间类型的数据为零(去掉NO_ZERO_IN_DATE,NO_ZERO_DATE)
sql_mode=ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
#ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
EOF
result "配置my_default.cnf"
#拷贝配置文件
cp $basedir/support-files/my_default.cnf /etc/my.cnf
result "拷贝my_default.cnf -> /etc/my.cnf"
#初始化Mysql
$basedir/bin/mysqld --initialize --user=mysql --basedir=$basedir --datadir=$datadir
result "初始化mysql"
#如果报错：./bin/mysqld: error while loading shared libraries: libaio.so.1: cannot open shared object file: No such file or 就安装yum install libaio，如果没有则跳过


#获取临时密码
$pwd =  grep 'temporary password' $datadir/mysqld.log | awk -F "root@localhost:" '{print $2}'
echo "临时密码是:"$pwd
#把启动脚本放到开机初始化目录
cp $basedir/support-files/mysql.server /etc/init.d/mysql
#46:basedir,47:datadir
sed -i '46c basedir=/usr/local/mysql-5.7.28/' /etc/init.d/mysql
sed -i '47c datadir=/datas/mysql/data/'/etc/init.d/mysql

#启动mysql
service mysql start
result "启动mysql"
