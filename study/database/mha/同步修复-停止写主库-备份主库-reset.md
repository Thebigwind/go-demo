参考： https://blog.csdn.net/weixin_42771643/article/details/117765906



1.在从数据库服务器上面，使用mysql -h服务器IP -P数据库端口 -u用户 -p密码;登入数据库输入show slave [status](https://so.csdn.net/so/search?q=status&spm=1001.2101.3001.7020) \G,能看到现在从数据库已经不同步了。。



2.停止 slave服务器的主从同步
使用mysql -h服务器IP -P数据库端口 -u用户 -p密码；登入数据库，输入**stop slave**；停止slave上同步服务，千万要做，否则恢复备份的时候服务器的磁盘空间会炸！



3.对master服务器的数据库加锁
这步是为了避免备份的时候对数据库进行操作，所以需要对master数据库进行枷锁；
还是登入数据库的前提下输入下面的命令哦：**flush tables with read lock;**



4.备份master服务器上的数据
这步不需要登入数据库，直接在服务器上输入下面命令：
mysqldump -u数据库用户名 -p数据库密码 -h服务器IP -P数据库端口 --quick --events --routines -B 需要备份的数据库名称 >./mysql_master_bizframe.sql
如：mysqldump -uroot -pMysql@123 -h127.0.0.1 -P3306 --quick --events --routines -B bizframe >./mysql_master_bizframe.sql 

PS：建议数据库数据较多的话，进行单个数据库逐一备份（满满心酸泪。。。）



5.将master的数据库备份文件拷贝到slave服务器上
scp -P 端口 /root/备份的文件名 root@服务器IP:/root/



6.删除slave服务器上旧数据
删除前最好备份下数据

DROP DATABASE 数据库名称;



7.slave数据库导入备份数据
导入前需要进入数据库执行下以下命令，防止恢复备份的时候报错，唯一的坏处就是会丢数据表。

set global log_bin_trust_function_creators =1; 

mysql -u数据库用户 -p数据库密码 -h服务器IP -P数据库端口 <已备份的sql文件
举个栗子：`mysql -uroot -pMysql@123 -h127.0.0.1 -P3306 <215mysql_master_bak.sql`



8.重置master和slave服务
在主数据库服务器上面执行reset master；，并对master服务器数据库解锁unlock tables;
在从数据库服务器上面执行reset slave；或者reset slave all；

9.开启slave服务
在从数据库上面执行start slave；
