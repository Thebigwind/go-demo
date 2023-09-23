#!/usr/bin/bash
stty erase ^H
userip=`env |grep 'SSH_CLIENT'|awk '{print $1}'`
a=`ip a |grep 'scope global'|awk '{print $2}'`
ip=${a%/*}
time=`date |awk '{print $1 $2 $3 $4 $5 }'`
sdfree=`df -h|grep 'root'|awk '{print $4}'`
free=`free -m|grep 'Mem'|awk '{print $4}'`M
cat <<eof
+--------------+-------------------------------
| 当 前 用 户  |$USER
|--------------+-------------------------------
| 当 前 时 间  |$time
|--------------+-------------------------------
| 当前服务器ip |$ip
|--------------+-------------------------------
| ssh远程机ip  |${userip#*=}
|--------------|-------------------------------
| 磁 盘 空 闲  |$sdfree
|--------------|-------------------------------
| 内 存 空 闲  |$free
|----------------------------------------------
eof
	sleep 4
clear

 while 1>0
	do
cat <<eof

	+--------------------------------------+
	| 作者:+_+ 海涛   	               |
	|                                      |
        |          优 化 脚 本 v 2.0 	       |
	|                                      |
	|--------------------------------------|
	|          1.检查网络状态              |
	|          2.检查yum能否正常使用       |
	|          3.关闭防火墙和selinux       |
	|          4.配置阿里源和epel源        |
        |          5.安装基本软件              |
	|          6.重启服务器                |
        |          7.退出脚本                  |
	+--------------------------------------+

eof
echo  "请输入数字选项！" && read var
case "$var" in
"1")
	ping -w1 -c1 www.baidu.com &>/dev/null
        if [ $? -eq 0 ];then
	echo "网络正常，可以上网"
	else
	echo "无网络，请手动检查"
	fi
	wait
	sleep 2
	clear
	;;

"2")
	yum clean all &>/dev/null
	yum makecache
	if [ $? -ne 0 ];then
	echo "yum启动失败.请手动检查"
	sleep 2
	exit 9
else
	echo "yum状态正常.可以使用"
	fi
	sleep 2
	clear
	;;
"3")
	systemctl  stop firewalld
	systemctl  disable firewalld
	setenforce 0
	sed -i s#SELINUX=enforcing#SELINUX=disabled# /etc/selinux/config
if [ $? -eq 0 ];then
	echo '防火墙已关闭'
	sleep 2 && clear
else
	echo '防火墙关闭失败.可能selnux已经关闭.请手动检查'
	fi
	sleep 2
	clear
	;;

"4")
	yum -y install wget
	yum -y install elinks
	mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
	echo "备份本地源成功"
wait
	wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
	echo "下载阿里源成功"
wait
	wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo
if test "$?" = "0";then
        echo "配置阿里epel源成功"
	sleep 2
	clear
        else
        echo "配置失败_请手动检查"
        sleep 2
        exit 6
fi
	;;
"5")
	yum -y install net-tools psmisc tree bash-completion
	sleep 2
	clear
	;;
"6")
	read -p " 正常重启,如不重启请按 ctrl+c 强制退出脚本！"
	sleep 3
	shutdown -r now
	;;
"7")
	exit
	;;
*)
	echo "请输入数字选项,输入有误！"
	sleep 2
	clear
esac

done
