systemctl start firewalld
firewall-cmd --zone=public --add-port=9902/tcp --permanent
firewall-cmd --zone=public --add-port=9903/tcp --permanent
firewall-cmd --zone=public --add-port=9904/tcp --permanent
firewall-cmd --zone=public --add-port=19903/tcp --permanent
firewall-cmd --zone=public --add-port=29903/tcp --permanent
firewall-cmd --zone=public --add-port=19904/tcp --permanent
firewall-cmd --zone=public --add-port=29904/tcp --permanent
#重启防火墙
firewall-cmd --reload


#检查修改是否生效
firewall-cmd --zone= public --query-port=10002/tcp
#开启端口
firewall-cmd --zone=public --add-port=10002/tcp --permanent
# 命令含义：
  –zone #作用域
  –add-port=80/tcp #添加端口，格式为：端口/通讯协议
  –permanent #永久生效，没有此参数重启后失效( —permanent放在前面与后面都行)

#查看已开放端口
firewall-cmd --zone=public --list-ports

#移除指定端口
firewall-cmd --permanent --remove-port=123/tcp