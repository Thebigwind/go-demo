#!/bin/bash
#source命令用于执行一个脚本，通常用于重新加载一个配置文件。
#source命令最大的特点是在当前 Shell 执行脚本，不像直接执行脚本时，会新建一个子 Shell。所以，source命令执行脚本时，不需要export变量。

## 当前 Shell 新建一个变量 foo
#$ foo=1
#
## 打印输出 1
#$ source test.sh
#1
#
## 打印输出空字符串
#$ bash test.sh

#source命令的另一个用途，是在脚本内部加载外部库。
source ./lib.sh

function_from_lib

#source有一个简写形式，可以使用一个点（.）来表示。
. .bashrc

