#!/bin/bash
#read命令除了读取键盘输入，可以用来读取文件。
filename='/etc/hosts'

while read myline
do
  echo "$myline"
done < $filename

#通过read命令，读取一个文件的内容。done命令后面的定向符<，将文件内容导向read命令，每次读取一行，存入变量myline，直到文件读取完毕。