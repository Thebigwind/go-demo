#!/bin/bash
#参考 https://cloud.tencent.com/developer/article/1886339
#yum install -y cmake gcc gcc-c++ ncurses ncurses-devel perl
result(){
  if [ $? -eq 0 ]; then
    echo "$1成功。"
  else
    echo "$1失败"
    exit 1
  fi
}

#检查是否已安装mongo
if [ -d "/usr/sbin/mysqld" ]; then
  echo "/usr/sbin/mysqld 已存在"
  exit 1;
fi

#rpm方式安装
#下载 MySQL yum包
wget http://repo.mysql.com/mysql57-community-release-el7-10.noarch.rpm
result "下载mysql yum包"

#安装MySQL源
rpm -Uvh mysql57-community-release-el7-10.noarch.rpm
result "安装MySQL源"
#安装MySQL服务端
yum install -y mysql-community-server
result "安装MySQL服务端"

#启动MySQL
systemctl start mysqld.service
result "启动MySQL"

#检查是否启动成功
systemctl status mysqld.service
result "检查是否启动成功"

#获取临时密码，MySQL5.7为root用户随机生成了一个密码
# 2023-03-28T07:48:04.378809Z 1 [Note] A temporary password is generated for root@localhost: DS58pu?xR,e?
$pwd =  grep 'temporary password' /var/log/mysqld.log | awk -F "root@localhost:" '{print $2}'
#通过临时密码登录MySQ
/usr/bin/mysql  -S /var/lib/mysql/mysql.sock -uroot -p$pwd -e "set global validate_password_policy=0;" &> /dev/null
/usr/bin/mysql  -S /var/lib/mysql/mysql.sock -uroot -p$pwd -e "set global validate_password_length=1;" &> /dev/null
/usr/bin/mysql  -S /var/lib/mysql/mysql.sock -uroot -p$pwd -e "ALTER USER 'root'@'localhost' IDENTIFIED BY '123456';" &> /dev/null
/usr/bin/mysql  -S /var/lib/mysql/mysql.sock -uroot -p$pwd -e "GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '123456' WITH GRANT OPTION;" &> /dev/null
/usr/bin/mysql  -S /var/lib/mysql/mysql.sock -uroot -p$pwd -e "FLUSH PRIVILEGES;" &> /dev/null
result "修改mysql密码"

#开启开机自启动
systemctl enable mysqld
systemctl daemon-reload
result "设置mysql开机自启动"

exit 0