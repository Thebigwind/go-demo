**MySQL程序之mysql参数详解**

mysql 是一个命令行客户程序，用于交互式或以批处理模式执行SQL语句

**用法:**

```
mysql [OPTIONS] [database]
```

**参数：**

```bash
1、-? --help                                      # 查看mysql的帮助并退出
2、-I                                             # 与-？同义
3、--auto-rehash                                  # 启动命令补全功能，默认是开启的。可以通过 --skip-auto-rehash 关闭 注意：功能需要使用readline库编译的MySQL客户端。
4、-A, --no-auto-rehash                           # 关闭命令补全功能，会使mysql启动更快
5、--auto-vertical-output                         # 如果结果比终端宽度宽，则自动切换到垂直输出模式。否则用表格模式。适用于；或 \G 结尾的语
6、-B, --batch                                    # 不使用历史文件，使用制表符作为列分隔符打印结果，每行位于一行上。不使用表格 （mysql -B -e 'select * from oldboy.student'）
7、--bind-address                                 # 在存在多个网卡的计算机上，通过该选项选择连接MySQL的IP
8、--binary-as-hex                                # 使用十六进制符号（0xvalue）显示二进制数据。
9、--binary-mode                                  # 默认情况下，ASCII'\0'是不允许的，而'\r\n'是翻译成“\n”。这个开关关闭这两个功能。在处理mysqlbinlog的时候该选项会有帮助
10、--character-sets-dir=name                     # 安装字符集的目录
11、 --column-names                               # 在结果中写入列名，默认是开启的
12、--column-type-info                            # 显示结果集元数据，展示字段类型信息
13、-c, --comments                                # 保存注释，发送注释到服务器。默认值是跳过注释 --skip-comments 。
14、-C, --compress                                # 如果客户端和服务器都支持压缩，则压缩在客户端和服务器之间发送的所有信息
15、--connect-expired-password                    # 通知服务器客户端用沙箱模式处理过期密码
16、-D, --database=name                           # 指定要使用的数据库
17、--debug[=debug_options], -# [debug_options]   # 编写调试日志，这个选项只有在MySQL使用WITH_DEBUG构建时才可用
18、--debug-check                                 # 在程序退出时打印一些调试信息
19、-T, --debug-info                              # 当程序退出时，打印调试信息、内存和CPU使用统计信息
20、--default-auth=name                           # 默认使用的客户端身份验证插件
21、--default-character-set=name                  # 设置客户端和连接的默认字符集
22、--defaults-extra-file=file_name               # 在全局选项文件之后读取此选项文件
23、--defaults-file=file_name                     # 只读取给定文件中的默认选项
24、--defaults-group-suffix=str                   # 不仅要读取通常的选项组，还要读取具有名称str后缀的组例如，mysql通常读取[客户机]和[mysql]组。如defaults-group =_other, mysql也读取[client_other]和[mysql_other]
25、--delimiter=str                               # 设置语句分隔符。默认是分号字符(;)
26、--disable-named-commands                      # 禁用命名命令。只使用\*表单，或者只在以分号(;)结尾的行开头使用命名命令。mysql默认启用这个选项
27、 --enable-cleartext-plugin                    # 启用/禁用明文身份验证插件           
28、-e, --execute=name                            # 执行命令并退出
29、-E, --vertical                                # 垂直打印行的输出
30、-f, --force                                   # 即使sql出现错误，仍然继续执行。
31、 --histignore=name                            # 用冒号分隔的模式列表来保存语句
32、-h, --host=name                               # 指定连接的IP
33、-H, --html                                    # 生成HTML输出
34、-i, --ignore-spaces                           # 忽略函数名后面的空格
35、--init-command=name                           # 连接到MySQL服务器时执行的SQL命令。重新连接时会自动重新执行
36、--line-numbers                                # 给错误写行号（默认开启，可以通过 --skip-line-numbers 来关闭）
37、--local-infile=[{1|0}]                        # 开启/关闭 加载本地数据文件
38、--login-path=#                                # 从登陆文件中读取此路径
39、-G, --named-commands                          # 启用命名命令。命名命令意味着这个程序内部命令；参见mysql>help。启用后，命名命令可以在查询的任何行中使用，否则只从第一行开始
40、-b, --no-beep                                 # 错误发生时不要发生蜂鸣
41、--no-defaults                                 # 不要从任何选项文件中读取默认选项，除了登录文件
42、-o, --one-database                            # 忽略语句，除非默认数据库是命令行中指定的数据库时发生的语句
43、-X, --xml                                     # 生成XML输出
44、-L, --skip-line-numbers                       # 错误不打印行号
45、-n, --unbuffered                              # 在每次执行语句后刷新缓冲区，默认开启
46、-N, --skip-column-names                       # 不在结果中写入列名
47、--sigint-ignore                               # 忽略 Ctrl-c
48、--pager[=name]                      
49、-p, --password[=name]                         # 连接服务器时指定登陆密码
50、-P, --port=#                                  # 用于连接的端口号，默认3306
51、--prompt=name                                 # 将mysql prompt设置为此值
52、--protocol=name                               # 用于连接的协议（TCP、套接字、管道、内存）
53、-q, --quick                                   # 不缓存结果，逐行打印。这可能会慢下来。如果输出被挂起，则关闭服务器。不使用历史文件。
54、-r, --raw                                     # 在不转换的情况下写入字段，与 --batch同时使用
55、--reconnect                                   # 当连接断开进行重连，用--disable-reconnect 关闭该功能。默认开启，可以通过--skip-reconnect使之失效。
56、-s, --silent                                  # 以制表符作为分隔符打印结果
57、-S, --socket=name                             # 指定建立连接时的套接字文件
58、--ssl-mode=name                               # SSL连接模式
59、--ssl-ca=name                                 # CA证书
60、--ssl-capath=name                             # CA路径
61、--ssl-cert=name                               # PEM格式的X509证书
62、--ssl-cipher=name                             # 要使用的SSL密码
63、--ssl-key=name                                # PEM格式的X509密钥
64、--ssl-crl=name                                # 证书吊销列表
65、--ssl-crlpath=name                            # 证书吊销列表路径
66、--tls-version=name                            # 要使用的TLS版本，允许值为：tlsv1、tlsv1.1
67、-t, --table                                   # 以表格的形式输出
68、--tee=name                                    # 将输出拷贝添加到给定的文件中，禁时用--disable-tee
69、-u, --user=name                               # 登陆用户
70、-U, --safe-updates                            # 使用更新和删除必须加where条件
71、-v, --verbose                                 # 多写。（-v-v-v给出表格输出格式）
72、-V, --version                                 # 输出版本信息
73、-w, --wait                                    # 如果连接断开，等待并重新连接
74、--connect-timeout=#                           # 连接时等待的时间，如果超过报超时错误
75、--max-allowed-packet=#                        # 发送到服务器或从服务器接收的最大数据包长度
76、--net-buffer-length=#                         # TCP/IP和套接字通信的缓冲区大小。
77、--select-limit=#                              # 使用--safe-updates时SELECT语句的自动限制
78、--max-join-size=#                             # 使用--safe-updates时联接中的行的自动限
79、--secure-auth                                 # 拒绝用(pre-4.1.1)的方式连接到数据库
80、--server-arg=name                             # 将此作为参数发送到嵌入式服务器
81、--show-warnings                               # 在每条语句后显示警告
82、-j, --syslog                                  # 将筛选的交互命令记录到syslog。除了默认模式外，命令筛选还取决于通过histignore选项提供的模式
83、--plugin-dir=name                             # 客户端插件目录
84、--print-defaults                              # 打印程序的参数列表
```

**参数默认值：**

```bash
auto-rehash                       TRUE
auto-vertical-output              FALSE
bind-address                      (No default value)
binary-as-hex                     FALSE
character-sets-dir                (No default value)
column-type-info                  FALSE
comments                          FALSE
compress                          FALSE
database                          (No default value)
default-character-set             utf8
delimiter                         ;
enable-cleartext-plugin           FALSE
vertical                          FALSE
force                             FALSE
histignore                        (No default value)
named-commands                    FALSE
ignore-spaces                     FALSE
init-command                      (No default value)
local-infile                      FALSE
no-beep                           FALSE
host                              (No default value)
html                              FALSE
xml                               FALSE
line-numbers                      TRUE
unbuffered                        FALSE
column-names                      TRUE
sigint-ignore                     FALSE
port                              3306
prompt                            mysql>
quick                             FALSE
raw                               FALSE
reconnect                         TRUE
socket                            /data/3306/mysql.sock
ssl                               TRUE
ssl-verify-server-cert            FALSE
ssl-ca                            (No default value)
ssl-capath                        (No default value)
ssl-cert                          (No default value)
ssl-cipher                        (No default value)
ssl-key                           (No default value)
ssl-crl                           (No default value)
ssl-crlpath                       (No default value)
tls-version                       (No default value)
table                             FALSE
user                              (No default value)
safe-updates                      FALSE
i-am-a-dummy                      FALSE
connect-timeout                   0
max-allowed-packet                16777216
net-buffer-length                 16384
select-limit                      1000
max-join-size                     1000000
secure-auth                       TRUE
show-warnings                     FALSE
plugin-dir                        (No default value)
default-auth                      (No default value)
binary-mode                       FALSE
connect-expired-password          FALSE
　　
```