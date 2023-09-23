#!/bin/bash

  #获取临时密码
  #pwd=$(grep 'temporary password' $datadir/mysqld.log | awk -F "root@localhost:" '{print $2}')
  #echo "临时密码是:"$pwd
  echo "请输入屏幕上显示的临时密码:"
  TEMP_PASSWORD=$(bash -c 'read  -p  "temp Password:" pass; echo $pass')

mysql -uroot  -p${TEMP_PASSWORD} -P ${MYSQL_PORT}  --connect-expired-password  <<EOF
ALTER USER 'root'@'localhost' IDENTIFIED BY '${MYSQL_PASS}'; FLUSH PRIVILEGES; use mysql;update user set host='%' where user='root'; FLUSH PRIVILEGES; grant all privileges on *.* to 'root'@'%';FLUSH PRIVILEGES;
EOF