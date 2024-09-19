#!/bin/bash

#创建数组
ARRAY=(value1 value2 value3)
echo ${ARRAY[1]}
#value2

#
array[0]=va1
array[1]=val2
echo ${array[1]}

#采用上面方式创建数组时，可以按照默认顺序赋值，也可以在每个值前面指定位置。
days=(Sun Mon Tue Wed Thu Fri Sat)
#如果读取数组成员时，没有读取指定哪一个位置的成员，默认使用0号位置。
echo $days
echo ${days[2]}
days=([0]=Sun [1]=Mon [2]=Tue [3]=Wed [4]=Thu [5]=Fri [6]=Sat)
#如果读取数组成员时，没有读取指定哪一个位置的成员，默认使用0号位置。
echo $days
echo ${days[2]}

echo "-------------"
#读取所有成员
#@和*是数组的特殊索引，表示返回数组的所有成员。
echo ${days[@]}

for i in "${days[@]}"; do
  echo $i
done

echo "-------------"
#@和*放不放在双引号之中，是有差别的
#一般把${activities[@]}放在双引号之中
activities=( swimming "water skiing" canoeing "white-water rafting" surfing )
for act in "${activities[@]}";
do
echo "Activity: $act";
done

#拷贝一个数组的最方便方法，就是写成下面这样。
 #
hobbies=( "${activities[@]}" )
echo ${hobbies[@]}

#这种写法也可以用来为新数组添加成员。
hobbies=( "${activities[@]}" diving )
#新数组hobbies在数组activities的所有成员之后，又添加了一个成员


#数组的长度
echo "----------"
echo ${#array[*]}
echo ${#array[@]}

echo ${#hobbies[*]}
echo ${#hobbies[@]}

echo "------------------"
#提取数组序号
#${!array[@]}或${!array[*]}，可以返回数组的成员序号，即哪些位置是有值的
arr=([5]=a [9]=b [23]=c)
echo ${!arr[@]}
#利用这个语法，也可以通过for循环遍历数组。
arr=(a b c d)
for i in ${!arr[@]};do
  echo ${arr[i]}
done

echo "-------------------"

#提取数组成员
#${array[@]:position:length}的语法可以提取数组成员。
food=( apples bananas cucumbers dates eggs fajitas grapes )
echo ${food[@]:1:1} #从数组1号位置开始的1个成员
echo ${food[@]:1:3} #从数组1号位置开始的3个成员
##如果省略长度参数length，则返回从指定位置开始的所有成员。
echo ${food[@]:4}

echo "--------------"
#追加数组成员
foo=(a b c)
echo ${foo[@]}
foo+=(d e f)
echo ${foo[@]}

echo "---------------"
#删除数组
#删除一个数组成员，使用unset命令。
foo=(a b c d e f)
echo ${foo[@]}
unset foo[2]
echo ${foo[@]}

#将某个成员设为空值，可以从返回值中“隐藏”这个成员
#这里是“隐藏”，而不是删除，因为这个成员仍然存在，只是值变成了空值
echo "--------------------"
#unset ArrayName可以清空整个数组。
unset foo
echo ${foo[@]}
echo "-------------"
#关联数组
#Bash 的新版本支持关联数组。关联数组使用字符串而不是整数作为数组索引。
#declare -A可以声明关联数组。
#关联数组必须用带有-A选项的declare命令声明创建。相比之下，整数索引的数组，可以直接使用变量名创建数组，关联数组就不行。
declare -A colors
colors["red"]="#ff0000"
colors["green"]="#00ff00"
colors["blue"]="#0000ff"

echo ${colors["blue"]}
