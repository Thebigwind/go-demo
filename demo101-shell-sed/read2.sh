#!/bin/bash

echo -n "输入一些文本 > "
read text
echo "你的输入：$text"


#read可以接受用户输入的多个值。 如果用户的输入项少于read命令给出的变量数目，那么额外的变量值为空。如果用户的输入项多于定义的变量，那么多余的输入项会包含到最后一个变量中。
echo Please, enter your firstname and lastname
read FN LN
echo "Hi! $LN, $FN !"

