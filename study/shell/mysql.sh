#!/bin/bash
set -e

SHELL_SCRIPT_PATH=$(cd $(dirname $0); pwd)
echo "SHELL_SCRIPT_PATH:"$SHELL_SCRIPT_PATH
#读取的变量
Mysql_Env_File=${SHELL_SCRIPT_PATH}/mysql.env
MYSQL_PASS=$(cat $Mysql_Env_File|grep MysqlPass|awk -F '=' '{print $2}')
MYSQL_PORT=$(cat $Mysql_Env_File|grep MysqlPort|awk -F '=' '{print $2}')

MYSQL_DIR=$(cat $Mysql_Env_File|grep MYSQL_DIR|awk -F '=' '{print $2}')
BASE_DIR=$(cat $Mysql_Env_File|grep BASE_DIR|awk -F '=' '{print $2}')
DATA_DIR=$(cat $Mysql_Env_File|grep DATA_DIR|awk -F '=' '{print $2}')
LOG_DIR=$(cat $Mysql_Env_File|grep LOG_DIR|awk -F '=' '{print $2}')

mkdir -p $MYSQL_DIR
mkdir -p $DATA_DIR
mkdir -p $LOG_DIR
result(){
  if [ $? -eq 0 ]; then
    echo "$1成功。"
  else
    echo "$1失败"
    exit 1
  fi
}

#检查防火墙
open_firewall_port(){
  echo "检查防火墙是否开启"
  systemctl status firewalld
  if [ $? = 0 ];then
      echo 'firewall open '$1
      port_status=$(firewall-cmd --query-port=$1/tcp)
      if [ $port_status = 'no' ]; then
          firewall-cmd --zone=public --add-port=$1/tcp --permanent
          firewall-cmd --reload
      fi
  fi
}

Write_Mysql_Service(){
  cat >  /usr/lib/systemd/system/mysql.service << EOF
[Unit]
Description=MySQL Server
After=network.target
[Service]
Type=forking
User=root
Group=root
ExecStart=/bin/sh -c 'env GOTRACEBACK=crash ${BASE_DIR}/bin/mysqld --basedir=${BASE_DIR} --datadir=${DATA_DIR} --plugin-dir=${BASE_DIR}/lib/plugin --user=root --log-error=${LOG_DIR}/mysqld.log --pid-file=${DATA_DIR}/localhost.localdomain.pid --socket=${DATA_DIR}/mysql.sock --port=3306 >> ${LOG_DIR}/mysql_start.log 2>&1 &'
LimitNOFILE=50000
Restart=always
RestartSec=11
[Install]
WantedBy=multi-user.target

EOF

systemctl enable mysql.service
systemctl daemon-reload
systemctl start mysql.service
}

#安装数据库
MYSQL_INSTALL(){
  rm -f /usr/bin/mysql
  rm -f /usr/bin/mysqlbinlog
  #是否安装wget
  res=$(rpm -qa|grep wget|wc -l)
  if [ $res -eq 0 ];then
    #yum install wget -y
    echo "use local package install ..."
  fi

  #检查是否已安装mysql
  if [ -d "${BASE_DIR}" ]; then
    echo "${BASE_DIR} already exist,please remove first: rm -rf ${BASE_DIR}; rm -rf ${DATA_DIR}; rm /etc/my.cnf; rm /etc/rc.d/init.d/mysqld"
    exit 1;
  fi

  #下载包
  package=../../pkg/mysql-8.0.24-linux-glibc2.12-x86_64.tar.xz
  if [ -e $package ] ; then
    echo "Found $package"
  else
    echo "开始拷贝..."
    #wget https://downloads.mysql.com/archives/get/p/23/file/$package
    result "拷贝包"
  fi

  #解压
  echo "开始解压..."
  tar -xvf $package -C ${MYSQL_DIR}
  result "解压包"
  #重命名
  mv ${MYSQL_DIR}/mysql-8.0.24-linux-glibc2.12-x86_64 ${MYSQL_DIR}/mysql-8.0.24


  #创建数据目录
  echo "开始创建数据目录"
  mkdir -p "${DATA_DIR}"
  mkdir -p "${LOG_DIR}"
  echo "" > ${LOG_DIR}/mysqld.log

  echo "开始设置my_default.cnf..."
  #在 $datadir/support-files目录下创建my_default.cnf
  touch ${BASE_DIR}/support-files/my_default.cnf
  cat>${BASE_DIR}/support-files/my_default.cnf<<EOF
[mysqld]
port=${MYSQL_PORT}
basedir=${BASE_DIR}
datadir=${DATA_DIR}
socket=${DATA_DIR}/mysql.sock
lower_case_table_names=1
character-set-server=utf8
collation-server=utf8_general_ci
#performance_schema_max_table_instances=400
#table_definition_cache=400
#table_open_cache=256
#datadir=/var/lib/mysql
#socket=/var/lib/mysql/mysql.sock
# Disabling symbolic-links is recommended to prevent assorted security risks
#symbolic-links=0
# Settings user and group are ignored when systemd is used.
# If you need to run mysqld under a different user or group,
# customize your systemd unit file for mariadb according to the
# instructions in http://fedoraproject.org/wiki/Systemd
max_allowed_packet=1024M
max_connections=1000
# Recommended in standard MySQL setup
sql_mode=NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
wait_timeout=2147483
interactive_timeout=2147483
connect_timeout=20
thread_cache_size=256
lower_case_table_names=1
innodb_strict_mode=0
# 创建新表时将使用的默认存储引擎
default-storage-engine=INNODB
default_authentication_plugin=mysql_native_password
innodb_file_per_table=1
log_bin_trust_function_creators=1
binlog_expire_logs_seconds=604800
innodb_flush_log_at_trx_commit=1
sync_binlog=1














[mysqld_safe]
log-error=${DATA_DIR}/mysqld.log
pid-file=${DATA_DIR}/mysqld.pid
default-character-set=utf8
[mysql]
default-character-set=utf8
[mysql.server]
default-character-set=utf8
[client]
#default-character-set=utf8
socket=${DATA_DIR}/mysql.sock
#
# include all files from the config directory
#
#!includedir /etc/my.cnf.d

EOF

  result "配置my_default.cnf"

  #拷贝配置文件
  \cp ${BASE_DIR}/support-files/my_default.cnf /etc/my.cnf
  result "拷贝my_default.cnf -> /etc/my.cnf"
  echo "设置只允许本机连接"
  sed -i "6c bind-address=127.0.0.1" /etc/my.cnf
  result "设置只允许本机连接"
  #初始化Mysql,并获取登录密码
  ${BASE_DIR}/bin/mysqld --initialize --user=zdlz --basedir=${BASE_DIR} --datadir=${DATA_DIR}
  result "初始化mysql"
  #echo "临时密码是:"$pwd

  #获取临时密码
  #pwd=$(grep 'temporary password' $datadir/mysqld.log | awk -F "root@localhost:" '{print $2}')
  #echo "临时密码是:"$pwd
  echo "请输入屏幕上显示的临时密码:"
  TEMP_PASSWORD=$(bash -c 'read  -p  "temp Password:" pass; echo $pass')
  #TEMP_PASSWORD=$(eval echo $TEMP_PASSWORD)
  length=${#TEMP_PASSWORD}
  echo "临时密码正确长度12，当前临时密码长度:"$length
  TEMP_PASSWORD=${TEMP_PASSWORD:0:12}
  echo "取临时密码的前12位"

  echo "配置mysql service"
  Write_Mysql_Service

  echo "mysql客户端加入/usr/bin"
  ln -s ${BASE_DIR}/bin/mysql /usr/bin
  ln -s ${BASE_DIR}/bin/mysqlbinlog /usr/bin
  echo "ln -s ${BASE_DIR}/bin/mysql /usr/bin"
  echo "ln -s ${BASE_DIR}/bin/mysqlbinlog /usr/bin"

  #建数据库用户和修改数据库连接
  #通过临时密码登录MySQ
 # EXEC_SQL="ALTER USER 'root'@'localhost' IDENTIFIED BY '${MYSQL_PASS}'; FLUSH PRIVILEGES; use mysql;update user set host='%' where user='root'; FLUSH PRIVILEGES; grant all privileges on *.* to 'root'@'%';FLUSH PRIVILEGES;"
echo "连接mysql修改密码和配置权限"
sleep 10
mysql -uroot  -p${TEMP_PASSWORD} -P ${MYSQL_PORT}  --connect-expired-password  <<EOF
ALTER USER 'root'@'localhost' IDENTIFIED BY '${MYSQL_PASS}'; FLUSH PRIVILEGES; use mysql;update user set host='%' where user='root'; FLUSH PRIVILEGES; grant all privileges on *.* to 'root'@'%';FLUSH PRIVILEGES;
EOF
  result "修改mysql密码"

  echo "安装mysql已完成"
  echo "mysql_master_slave_config">schedule.log
}

##################################################################################################
#open_firewall_port $MYSQL_PORT
# mysql安装
MYSQL_INSTALL

if [ $? -eq 0 ]; then
  echo "********************************** 当前节点配置mysql成功 **********************************"
fi

