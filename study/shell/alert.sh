#!/bin/bash
#命令提示符通常是美元符号$，对于根用户则是井号#。这个符号是环境变量PS1决定的，执行下面的命令，可以看到当前命令提示符的定义。
echo $PS1

#Bash 允许用户自定义命令提示符，只要改写这个变量即可。改写后的PS1，可以放在用户的 Bash 配置文件.bashrc里面，
# 以后新建 Bash 对话时，新的提示符就会生效。要在当前窗口看到修改后的提示符，可以执行下面的命令。
source ~/.bashrc

#命令提示符的定义，可以包含特殊的转义字符，表示特定内容。
#\a：响铃，计算机发出一记声音。
#\d：以星期、月、日格式表示当前日期，例如“Mon May 26”。
#\h：本机的主机名。
#\H：完整的主机名。
#\j：运行在当前 Shell 会话的工作数。
#\l：当前终端设备名。
#\n：一个换行符。
#\r：一个回车符。
#\s：Shell 的名称。
#\t：24小时制的hours:minutes:seconds格式表示当前时间。
#\T：12小时制的当前时间。
#\@：12小时制的AM/PM格式表示当前时间。
#\A：24小时制的hours:minutes表示当前时间。
#\u：当前用户名。
#\v：Shell 的版本号。
#\V：Shell 的版本号和发布号。
#\w：当前的工作路径。
#\W：当前目录名。
#\!：当前命令在命令历史中的编号。
#\#：当前 shell 会话中的命令数。
#\$：普通用户显示为$字符，根用户显示为#字符。
#\[：非打印字符序列的开始标志。
#\]：非打印字符序列的结束标志。
#举例来说，[\u@\h \W]\$这个提示符定义，显示出来就是[user@host ~]$（具体的显示内容取决于你的系统）。
#
#[user@host ~]$ echo $PS1
#[\u@\h \W]\$
#改写PS1变量，就可以改变这个命令提示符。
#
#$ PS1="\A \h \$ "
#17:33 host $
#注意，$后面最好跟一个空格，这样的话，用户的输入与提示符就不会连在一起。






#颜色
 #默认情况下，命令提示符是显示终端预定义的颜色。Bash 允许自定义提示符颜色。
 #
 #使用下面的代码，可以设定其后文本的颜色。
 #
 #\033[0;30m：黑色
 #\033[1;30m：深灰色
 #\033[0;31m：红色
 #\033[1;31m：浅红色
 #\033[0;32m：绿色
 #\033[1;32m：浅绿色
 #\033[0;33m：棕色
 #\033[1;33m：黄色
 #\033[0;34m：蓝色
 #\033[1;34m：浅蓝色
 #\033[0;35m：粉红
 #\033[1;35m：浅粉色
 #\033[0;36m：青色
 #\033[1;36m：浅青色
 #\033[0;37m：浅灰色
 #\033[1;37m：白色
 #举例来说，如果要将提示符设为红色，可以将PS1设成下面的代码。
 #
 #PS1='\[\033[0;31m\]<\u@\h \W>\$'
 #但是，上面这样设置以后，用户在提示符后面输入的文本也是红色的。为了解决这个问题， 可以在结尾添加另一个特殊代码\[\033[00m\]，表示将其后的文本恢复到默认颜色。
 #
 #PS1='\[\033[0;31m\]<\u@\h \W>\$\[\033[00m\]'
 #除了设置前景颜色，Bash 还允许设置背景颜色。
 #
 #\033[0;40m：蓝色
 #\033[1;44m：黑色
 #\033[0;41m：红色
 #\033[1;45m：粉红
 #\033[0;42m：绿色
 #\033[1;46m：青色
 #\033[0;43m：棕色
 #\033[1;47m：浅灰色
 #下面是一个带有红色背景的提示符。
 #
 #PS1='\[\033[0;41m\]<\u@\h \W>\$\[\033[0m\] '








# 环境变量 PS2，PS3，PS4
# 除了PS1，Bash 还提供了提示符相关的另外三个环境变量。
#
# 环境变量PS2是命令行折行输入时系统的提示符，默认为>。
#
# $ echo "hello
# > world"
# 上面命令中，输入hello以后按下回车键，系统会提示继续输入。这时，第二行显示的提示符就是PS2定义的>。
#
# 环境变量PS3是使用select命令时，系统输入菜单的提示符。
#
# 环境变量PS4默认为+。它是使用 Bash 的-x参数执行脚本时，每一行命令在执行前都会先打印出来，并且在行首出现的那个提示符。
#
# 比如下面是脚本test.sh。
#
# #!/bin/bash
#
# echo "hello world"
# 使用-x参数执行这个脚本。
#
# $ bash -x test.sh
# + echo 'hello world'
# hello world
# 上面例子中，输出的第一行前面有一个+，这就是变量PS4定义的。



