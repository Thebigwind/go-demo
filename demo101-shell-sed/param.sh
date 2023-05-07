#!/bin/bash
# script.sh

echo "全部参数：" $@
echo "命令行参数数量：" $#
echo '$0 = ' $0
echo '$1 = ' $1
echo '$2 = ' $2
echo '$3 = ' $3

#执行结果如下。
#
#$ ./script.sh a b c
#全部参数：a b c
#命令行参数数量：3
#$0 =  script.sh
#$1 =  a
#$2 =  b
#$3 =  c

#用户可以输入任意数量的参数，利用for循环，可以读取每一个参数。
for i in "$@"; do
  echo $i
done

#如果多个参数放在双引号里面，视为一个参数。
#$ ./script.sh "a b" Bash 会认为"a b"是一个参数，$1会返回a b。注意，返回时不包括双引号。
