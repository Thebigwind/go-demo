#!/bin/bash
echo "检查zdlz用户是否存在"
id zdlz
if [ $? = 0 ];then
  echo "zdlz用户已存在"
else
  echo "建zdlz用户"
  useradd zdlz
fi