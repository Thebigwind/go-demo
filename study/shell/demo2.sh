
#检查文件是否存在
FILE=/etc/resolv.conf
if test -f "$FILE"; then
    echo "$FILE exist"
fi

FILE=/etc/resolv.conf
if [ -f "$FILE" ]; then
    echo "$FILE exist"
fi

FILE=/etc/resolv.conf
if [[ -f "$FILE" ]]; then
    echo "$FILE exist"
fi

FILE=/etc/resolv.conf
if [ -f "$FILE" ]; then
    echo "$FILE exist"
else
    echo "$FILE does not exist"
fi

#Linux系统中运算符-d允许你测试一个文件是否时目录。
#例如检查/etc/docker目录是否存在，你可以使用如下脚本
FILE=/etc/docker
if [ -d "$FILE" ]; then
    echo "$FILE is a directory"
fi
[ -d /etc/docker ] && echo "$FILE is a directory"


#检查文件是否不存在
FILE=/etc/docker
if [ ! -f "$FILE" ]; then
    echo "$FILE does not exist"
fi
[ ! -f /etc/docker ] && echo "$FILE does not exist"

#检查是否存在多个文件
#不使用复杂的嵌套if/else构造，您可以使用-a（或带[[的&&预算符)来测试是否存在多个文件，示例如下：
if [ -f /etc/resolv.conf -a -f /etc/hosts ]; then
    echo "Both files exist."
fi
if [[ -f /etc/resolv.conf && -f /etc/hosts ]]; then
    echo "Both files exist."
fi

#文件test命令运算符
test命令包含以下文件操作运算符，这些运算符允许你测试不同类型的文件：

-b FILE - 如果文件存在并且是块特殊文件，则为True。
-c FILE - 如果文件存在并且是特殊字符文件，则为True。
-d FILE - 如果文件存在并且是目录，则为True。
-e FILE - 如果文件存在并且是文件，则为True，而不考虑类型（节点、目录、套接字等）。
-f FILE - 如果文件存在并且是常规文件（不是目录或设备），则为True。
-G FILE - 如果文件存在并且与运行命令的用户具有相同的组，则为True。
-h FILE - 如果文件存在并且是符号链接，则为True。
-g FILE - 如果文件存在并已设置组id（sgid）标志，则为True。
-k FILE - 如果文件存在并设置了粘滞位标志，则为True。
-L FILE - 如果文件存在并且是符号链接，则为True。
-O FILE - 如果文件存在并且由运行该命令的用户拥有，则为True。
-p FILE - 如果文件存在并且是管道，则为True。
-r FILE - 如果文件存在且可读，则为True。
-S FILE - 如果文件存在并且是套接字，则为True。
-s FILE - 如果文件存在且大小不为零，则为True。
-u FILE - 如果文件存在并且设置了（suid）标志，则为True。
-w FILE - 如果文件存在且可写，则为True。
-x FILE - 如果文件存在且可执行，则为True。


