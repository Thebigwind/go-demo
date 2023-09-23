#!/bin/bash
#mysql目录
MYSQL_DIR="/xxx/mysql"
LOGFILE=${MYSQL_DIR}"/masterha/email.log"
:>"$LOGFILE"
exec 1>"$LOGFILE"
exec 2>&1
SMTP_server='smtp.mxhichina.com'
username='xxx'
password='ooo'
from_email_address='xxxx'
to_email_address='xxxxxxxx'

message_subject_utf8="MHA集群主库故障转移提醒"

HTML_PATH=html_path
echo "<h2 style="color:red">">$HTML_PATH
echo "MHA集群主节点发生故障，进行节点故障转移，请及时解决查看！！！">>$HTML_PATH
 echo "</h2>">>$HTML_PATH
 echo "<p>以下为MHA集群的相关信息:</p>">>$HTML_PATH
 echo "<table border="1" cellspacing="0"  width="700"><tr><th>节点</th><th>角色</th>     <th>作用</th></tr><tr><td>10.6.110.170</td><td>MHA manager</td> <td>MHA监控节点</td></tr><tr><td>10.8.40.77</td><td>master/master.bak</td>      <td>主库或者主备</td></tr><tr><td>10.8.40.68</td><td>master/master.bak</td>     <td>主库或者主备</td></tr><tr><td>10.6.119.241</td><td>slave</td>       <td>从库</td></tr><tr><td>10.8.40.79</td><td>VIP</td>   <td>虚拟ip</td></tr></table>">>$HTML_PATH
 echo "<br>">>$HTML_PATH

message_body_utf8=$(cat $HTML_PATH)

#message_body_utf8="mysql的MHA集群主节点发生故障，进行节点故障转移，请及时解决查看！！！"
# 转换邮件标题为GB2312，解决邮件标题含有中文，收到邮件显示乱码的问题。
message_subject_gb2312=`iconv -t GB2312 -f UTF-8 << EOF
$message_subject_utf8
EOF`
[ $? -eq 0 ] && message_subject="$message_subject_gb2312" || message_subject="$message_subject_utf8"
# 转换邮件内容为GB2312，解决收到邮件内容乱码
message_body_gb2312=`iconv -t GB2312 -f UTF-8 << EOF
$message_body_utf8
EOF`
[ $? -eq 0 ] && message_body="$message_body_gb2312" || message_body="$message_body_utf8"
# 发送邮件
sendEmail='/usr/bin/sendEmail'
set -x
$sendEmail -s "$SMTP_server" -xu "$username" -xp "$password" -f "$from_email_address" -t "$to_email_address" -u "$message_subject" -m "$message_body" -o message-content-type=html -o message-charset=gb2312
