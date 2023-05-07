#!/bin/bash
j=$1
for ((i=1; i<=j; i++))
do
touch file$i && echo file $i is ok
done

#产生10个随机数
for i in {0..9};do echo $RANDOM;done

for i in $(seq 10);do echo $RANDOM;done

#倒数5秒
echo "准备倒数5秒："
for i in $(seq 5 -1 1)
do
     echo -en "$i";sleep 1
done
echo -e "开始"

#批量添加用户
#!/bin/bash
for i in $(cat /root/users.txt)        #--》从列表文件读取文件名
do
    useradd $i
    echo "123456" | passwd --stdin $i #--》通过管道指定密码字串
done

#


