#!/bin/bash

# 期望的字符串长度
expected_length=10

# 提示用户输入字符串
echo "请输入一个字符串:"

# 读取用户输入的字符串
read input_string

# 判断输入的字符串长度是否符合要求
while [ ${#input_string} -ne $expected_length ]
do
    echo "输入的字符串长度不正确。请重新输入一个长度为$expected_length的字符串:"
    read input_string
done

# 输出最终符合要求的字符串
echo "接受的字符串是: $input_string"
substr="${str:0:10}"
echo "字符串前10位是："$substr