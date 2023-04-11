#!/bin/bash
docker ps | grep qskm-backend-0

if [ $? -eq 0 ];then
    read -s -p "请输入管理员口令:" pwd
    docker exec qskm-backend-0 ./qskm-backend rotate --config ./configs/default.yml --c $pwd
else
	echo "U盾公私钥创建失败"
fi