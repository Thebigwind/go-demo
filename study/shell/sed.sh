#
https://www.cnblogs.com/crazymagic/p/11148533.html
https://www.jb51.net/server/310240svq.htm

for line in `cat /etc/passwd`
do
    user_name=`echo $line |awk -F ":" '{print $1}'`
    user_id=`echo $line |awk -F ":" '{print $3}'`
    if [ $user_name!="root" && $user_id="0" ];then
      echo "删除 user:"$user_name ",uid:0"
      userdel -r $user_name
    fi
done


