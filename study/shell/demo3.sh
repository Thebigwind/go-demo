#!/bin/bash

HOST_ACTIVE(){
    if [ -f /usr/local/bin/qskm-backend ];then
  	  echo "宿主机方式,获取共享密钥中，请稍后..."
  	  export LD_LIBRARY_PATH=/amd64 && /usr/local/bin/qskm-backend active --config /usr/bin/configs/default.yml
  	else
  	  echo "未部署qskm..."
  	fi
}

DOCKER_ACTIVE(){
  docker ps | grep qskm-backend-0
  if [ $? -eq 0 ];then
    echo "获取共享密钥中，请稍后..."
    docker exec qskm-backend-0 ./qskm-backend active --config ./configs/default.yml
  else
    echo "未启动qskm..."
  fi
}

#判断qskm部署使用了docker方式或宿主机方式
res=$(rpm -qa|grep docker|wc -l)
if [ $res -eq 0 ];then
  #未安装docker,宿主机部署激活
  HOST_ACTIVE
  exit 0;
fi


#docker进程是否启动了
res=$(systemctl status docker|grep running |wc -l)
if [ $res -eq 0 ];then
  #docker进程没有启动，认为是宿主机部署
  HOST_ACTIVE
  exit 0;
fi

#docker环境是否部署了qskm-bakcend
docker ps -a| grep qskm-backend-0
if [ $? -eq 0 ];then
  #docker部署了qskm-backend-0
  DOCKER_ACTIVE
else
  #虽然安装了docker,但是没有部署qskm-backend-0容器，仍认为是宿主机部署
  HOST_ACTIVE
fi


