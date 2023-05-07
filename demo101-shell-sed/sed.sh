#在455后增加一行，内容为：dir /datas/redis/data
sed  -i '455a "dir /datas/redis/data"' redis.conf
#删除455行
sed  -i '455d' redis.conf
#把455行的内容替换为：dir /datas/redis/data
sed -i '454c dir /datas/redis/data' redis.conf

#把protected-mode yes 替换为protected-mode no
sed -i 's/protected-mode yes/protected-mode no/' $basepath/redis.conf