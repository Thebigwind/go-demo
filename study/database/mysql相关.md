## 1.事务隔离级别

  序列化

 可重复读

读已提交      幻读

读未提交      脏读



### MVCC版本控制

MVCC工作原理是使用数据在某个时间点的快照来实现的。这意味着，无论事务运行多长时间，都可以看到数据的一致视图，也意味着不同的事务可以在同一时间看到同一张表中的不同数据。

![image-20230411221455623](./mvcc.png)

InnoDB通过为每个事务在启动时分配一个事务ID来实现MVCC。该ID在事务首次计取任何数据时分配。在该事务中修改记录时，将向 Undo 日志写入一条说明如何恢复更改的 Undo 记录，并且事务的回滚指针指向该 Undo 日志记录。这就是事务如何在需要时执行回滚的方法。
当不同的会话读取聚簇主键索引记录时，InnoDB会将该记录的事务ID与该会话的读取视图进行比较。如果当前状态下的记录不应可见(更改它的事务尚未提交)，那么Undo日志记录将被跟踪并应用，直到会话达到一个符合可见条件的事务ID。这个过程可以一直循环到完全删除这一行的Undo记录，然后向读取视图发出这一行不存在的信号。
事务中的记录可以通过在记录的“info flags”中设置“deleted”位来删除。这在Undo日志中也被作为“删除标记”进行跟踪。
值得注意的是，所有Undo日志写入也都会写入Redo日志，因为Undo日志写入是服务器崩溃恢复过程的一部分，并且是事务性的。这些Redo日志和Undo日志的大小也是高并发事务工作机制中的重要影响因素。
在记录中保留这些额外信息带来的结果是，大多数读取查询都不再需要获取锁。它们只是尽可能快地读取数据，确保仅查询符合条件的行即可。

缺点是存储引擎必须在每一行中存储更多的数据，在检查行时需要做更多的工作，并处理一些额外的内部操作。
MVCC 仅适用于REPEATABLE READ和READ COMMITTED隔离级别。

READ UNCOMMITTED与MVCC不兼容，是因为查询不会读取适合其事务版本的行版本，而是不管怎样都读最新版本。

SERIALIZABLE与MVCC也不兼容，是因为读取会锁定它们返回的每一行。

## 2.索引

###   索引类型

聚簇索引：

非聚簇索引：





###  索引使用场景优化

https://zhuanlan.zhihu.com/p/375828064





### B+树原理





### 修改字段类型

格式：alter table 表名 modify column 字段名 类型;

**实例：**

将users表的registerTime字段改为datetime类型;

```sql
alter` `table` `users ``modify` `column` `registerTime datetime;
```

### 修改字段长度

格式：alter table 表名 modify column 字段名 类型(长度);

**实例：**

将users表的username字段改为varchar类型，长度为30个字节;

```sql
alter` `table` `users ``modify` `column` `username ``varchar``(30);
```



## 3.count

1、COUNT有几种用法？

2、COUNT(字段名)和COUNT(*)的查询结果有什么不同？

3、COUNT(1)和COUNT(*)之间有什么不同？

4、COUNT(1)和COUNT(*)之间的效率哪个更高？

5、为什么《阿里巴巴Java开发手册》建议使用COUNT(*)

6、MySQL的MyISAM引擎对COUNT(*)做了哪些优化？

7、MySQL的InnoDB引擎对COUNT(*)做了哪些优化？

8、上面提到的MySQL对COUNT(*)做的优化，有一个关键的前提是什么？

9、SELECT COUNT(*) 的时候，加不加where条件有差别吗？

10、COUNT(*)、COUNT(1)和COUNT(字段名)的执行过程是怎样的？



COUNT(expr) ，返回SELECT语句检索的行中expr的值不为NULL的数量。结果是一个BIGINT值。

`COUNT(*)` 的统计结果中，会包含值为NULL的行数。

除了`COUNT(id)`和`COUNT(*)`以外，还可以使用`COUNT(常量)`（如`COUNT(1)`）来统计行数，那么这三条SQL语句有什么区别呢？到底哪种效率更高呢？

```
COUNT(expr)`用于做行数统计，统计的是expr不为NULL的行数，那么`COUNT(列名)`、 `COUNT(常量)` 和 `COUNT(*)`这三种语法中，expr分别是`列名`、 `常量` 和 `*
列名、 常量 和 *这三个条件中，常量 是一个固定值，肯定不为NULL。*可以理解为查询整行，所以肯定也不为NULL，那么就只有列名的查询结果有可能是NULL了。

所以， COUNT(常量) 和 COUNT(*)表示的是直接查询符合条件的数据库表的行数。而COUNT(列名)表示的是查询符合条件的列的值不为NULL的行数。
COUNT(*)相比COUNT(常量) 和 COUNT(列名)来讲，COUNT(*)是SQL92定义的标准统计行数的语法，因为他是标准语法，所以MySQL数据库对他进行过很多优化。

```





## 4.高可用





## MyISAM和InnoDB区别

**MyISAM不支持事务，MyISAM中的锁是表级锁；而InnoDB支持事务，并且支持行级锁。**

因为MyISAM的锁是表级锁，所以同一张表上面的操作需要串行进行，所以，**MyISAM做了一个简单的优化，那就是它可以把表的总行数单独记录下来，如果从一张表中使用COUNT(\*)进行查询的时候，可以直接返回这个记录下来的数值就可以了，当然，前提是不能有where条件。**

MyISAM之所以可以把表中的总行数记录下来供COUNT(*)查询使用，那是因为MyISAM数据库是表级锁，不会有并发的数据库行数修改，所以查询得到的行数是准确的。



从MySQL 8.0.13开始，针对InnoDB的`SELECT COUNT(*) FROM tbl_name`语句，确实在扫表的过程中做了一些优化。前提是查询语句中不包含WHERE或GROUP BY等条件。

COUNT(*)的目的只是为了统计总行数，所以，他根本不关心自己查到的具体值，所以，他如果能够在扫表的过程中，选择一个成本较低的索引进行的话，那就可以大大节省时间。

我们知道，InnoDB中索引分为聚簇索引（主键索引）和非聚簇索引（非主键索引），聚簇索引的叶子节点中保存的是整行记录，而非聚簇索引的叶子节点中保存的是该行记录的主键的值。相比之下，非聚簇索引要比聚簇索引小很多，所以**MySQL会优先选择最小的非聚簇索引来扫表。所以，当我们建表的时候，除了主键索引以外，创建一个非主键索引还是有必要的。**



官方：

> InnoDB handles SELECT COUNT(*) and SELECT COUNT(1) operations in the same way. There is no performance difference.
>
> 画重点：`same way` , `no performance difference`。**所以，对于COUNT(1)和COUNT(\*)，MySQL的优化是完全一样的，根本不存在谁比谁快！**

建议使用`COUNT(*)`！因为这个是SQL92定义的标准统计行数的语法。

##### COUNT(字段) 他的查询就比较简单粗暴了，就是进行全表扫描，然后判断指定字段的值是不是为NULL，不为NULL则累加。相比`COUNT(*)`，`COUNT(字段)`多了一个步骤就是判断所查询的字段是否为NULL，所以他的性能要比`COUNT(*)`慢。



## 开窗函数



## 自定义函数

https://www.jb51.net/article/70677.htm



## 存储过程

https://www.jb51.net/article/70677.htm

存储过程和函数是在数据库中定义一些SQL语句的集合，然后直接调用这些存储过程和函数来执行已经定义好的SQL语句。存储过程和函数可以避免开发人员重复的编写相同的SQL语句。而且，存储过程和函数是在MySQL服务器中存储和执行的，可以减少客户端和服务器端的数据传输。





## 游标



## 备份

https://www.cnblogs.com/wuqiuyin/p/15412286.html

### 备份什么

一般情况下, 我们需要备份的数据分为以下几种

- 二进制日志, InnoDB事务日志
- 代码(存储过程、存储函数、触发器、事件调度器)
- 服务器配置文件

### 备份的类型

按照备份时数据库的运行状态，可以分为三种，分别是：冷备、温备、热备。、

- 冷备：停库、停服务来备份，即当数据库进行备份时, 数据库不能进行读写操作, 即数据库要下线。
- 温备：不停库、不停服务来备份，会(锁表)阻止用户的写入，即当数据库进行备份时, 数据库的读操作可以执行, 但是不能执行写操作 。
- 热备：不停库、不停服务来备份，也不会(锁表)阻止用户的写入 即当数据库进行备份时, 数据库的读写操作均不是受影响 。



### 逻辑备份与物理备份

按照备份的内容分，可以分为两种，分别是逻辑备份与物理备份

- 1、物理备份：直接将底层物理文件备份
- 2、逻辑备份：通过特定的工具从数据库中导出sql语句或者数据，可能会丢失数据精度

### 备份方式之全量、差异、增量

按照每次备份的数据量，可以分为全量备份、差异备份以及增量备份。

- 全量备份/完全备份（Full Backup）：备份整个数据集( 即整个数据库 )
- 部分备份：备份部分数据集(例如: 只备份一个表的变化)
- - 差异备份
    - 增量备份

```
# 1、差异备份（Differential Backup）
每次备份时，都是基于第一次完全备份的内容，只备份有差异的数据(新增的、修改的、删除的)，例如

第一次备份：完全备份
第二次备份：以当前时间节点的数据为基础，备份与第一次备份内容的差异
第三次备份：以当前时间节点的数据为基础，备份与第一次备份内容的差异
第四次备份：以当前时间节点的数据为基础，备份与第一次备份内容的差异
第五次备份：以当前时间节点的数据为基础，备份与第一次备份内容的差异
。。。

# 2、增量备份（Incremental Backup ）
每次备份时，都是基于上一次备份的内容（注意是上一次，而不是第一次），只备份有差异的数据(新增的、修改的、删除的)，所以增量备份的结果是一条链，例如

第一次备份：完全备份
第二次备份：以当前时间节点的数据为基础，备份与第一次备份内容的差异
第三次备份：以当前时间节点的数据为基础，备份与第二次备份内容的差异
第四次备份：以当前时间节点的数据为基础，备份与第三次备份内容的差异
第五次备份：以当前时间节点的数据为基础，备份与第四次备份内容的差异
。。。
```

```
# 1、全量备份的数据恢复
只需找出指定时间点的那一个备份文件即可，即只需要找到一个文件即可

# 2、差异备份的数据恢复
需要先恢复第一次备份的结果，然后再恢复最近一次差异备份的结果，即需要找到两个文件

# 3、增量备份的数据恢复
需要先恢复第一次备份的结果，然后再依次恢复每次增量备份，直到恢复到当前位置，即需要找到一条备份链

综上，对比三种备份方案
1、占用空间：全量 > 差异 > 增量
2、恢复数据过程的复杂程度：增量 > 差异 > 全量

```

### 备份的工具

| **备份工具**       | **备份速度** | **恢复速度** | **便捷性**               | **适用存储引擎**                   | **支持的备份类型**                                           | **功能** | **应用场景**       |
| :----------------- | :----------- | :----------- | :----------------------- | :--------------------------------- | :----------------------------------------------------------- | :------- | :----------------- |
| cp、tar等（物理）  | 快           | 快           | 一般                     | 所有                               | 冷备、全量、差异、增量                                       | 很弱     | 少量数据备份       |
| lvm2快照（物理）   | 快           | 快           | 一般                     | 所有                               | 支持几乎热备（即差不多是热备，哈哈），是借助文件系统管理工具进行的备份 | 一般     | 中小型数据量的备份 |
| xtrabackup（物理） | 较快         | 较快         | 是一款非常强大的热备工具 | 由percona提供，只支持InnoDB/XtraDB | 热备、全量、差异、增量                                       | 强大     | 较大规模的备份     |
| mysqldump（逻辑）  | 慢           | 慢           | 一般                     | 所有                               | 支持温备、完全备份、部分备份、对于**InnoDB**存储引擎支持热备 | 一般     | 中小型数据量的备份 |

如果考虑到增量备份，还需要结合binlog日志（binlog只属于增量恢复），需要用到工具mysqlbinlog，相当于逻辑备份的一种。



## 使用cp进行备份

```sql
mysql> FLUSH TABLES WITH READ LOCK; 
Query OK, 0 rows affected (0.00 sec)

[root@localhost ~]# mkdir /data
[root@localhost ~]# cp -a /usr/local/mysql-5.7.34/data/* /data
```

## 使用mysqldump

MySQL数据库自带的一个很好用的备份命令。是逻辑备份，导出 的是SQL语句。也就是把数据MySQL库中以逻辑的SQL语句的形式直接输出或生成备份的文件的过程。

### 语法

```sql
mysqldump  -h 服务器  -u用户名  -p密码  选项与参数 > 备份文件.sql
```

### 参数

| **参数**                | **解释**                                                     |
| :---------------------- | :----------------------------------------------------------- |
| -A --all-databases      | 导出全部数据库                                               |
| -Y --all-tablespaces    | 导出全部表空间                                               |
| --add-drop-database     | 每个数据库创建之前添加drop数据库语句。                       |
| --add-drop-table        | 每个数据表创建之前添加drop数据表语句。(默认为打开状态，使用--skip-add-drop-table取消选项) |
| --add-locks             | 在每个表导出之前增加LOCK TABLES并且之后UNLOCK TABLE。(默认为打开状态，使用--skip-add-locks取消选项) |
| --comments              | 附加注释信息。默认为打开，可以用--skip-comments取消          |
| --compact               | 导出更少的输出信息(用于调试)。去掉注释和头尾等结构。可以使用选项：--skip-add-drop-table --skip-add-locks --skip-comments --skip-disable-keys |
| -c --complete-insert    | 使用完整的insert语句(包含列名称)。这么做能提高插入效率，但是可能会受到max_allowed_packet参数的影响而导致插入失败。 |
| -C --compress           | 在客户端和服务器之间启用压缩传递所有信息                     |
| -B --databases          | 导出几个数据库。参数后面所有名字参量都被看作数据库名。       |
| --debug                 | 输出debug信息，用于调试。                                    |
| --debug-info            | 输出调试信息并退出                                           |
| --default-character-set | 设置默认字符集，默认值为utf8                                 |
| --delayed-insert        | 采用延时插入方式（INSERT DELAYED）导出数据                   |
| -E --events             | 导出事件。                                                   |
| --master-data           | 在备份文件中写入备份时的binlog文件。值为1时，binlog文件名和位置没有注释，为2时，则在备份文件中将binlog的文件名和位置进行注释。 |
| --flush-logs            | 开始导出之前刷新日志。请注意：假如一次导出多个数据库(使用选项--databases或者--all-databases)，将会逐个数据库刷新日志。除使用--lock-all-tables或者--master-data外。在这种情况下，日志将会被刷新一次，相应的所以表同时被锁定。因此，如果打算同时导出和刷新日志应该使用--lock-all-tables 或者--master-data 和--flush-logs。 |
| --flush-privileges      | 在导出mysql数据库之后，发出一条FLUSH PRIVILEGES 语句。为了正确恢复，该选项应该用于导出mysql数据库和依赖mysql数据库数据的任何时候。 |
| --force                 | 在导出过程中忽略出现的SQL错误。                              |
| -h --host               | 需要导出的主机信息                                           |
| --ignore-table          | 不导出指定表。指定忽略多个表时，需要重复多次，每次一个表。每个表必须同时指定数据库和表名。例如：--ignore-table=database.table1 --ignore-table=database.table2 …… |
| -x --lock-all-tables    | 提交请求锁定所有数据库中的所有表，以保证数据的一致性。这是一个全局读锁，并且自动关闭--single-transaction 和--lock-tables 选项。 |
| -l --lock-tables        | 开始导出前，锁定所有表。用READ LOCAL锁定表以允许MyISAM表并行插入。对于支持事务的表例如InnoDB和BDB，--single-transaction是一个更好的选择，因为它根本不需要锁定表。请注意当导出多个数据库时，--lock-tables分别为每个数据库锁定表。因此，该选项不能保证导出文件中的表在数据库之间的逻辑一致性。不同数据库表的导出状态可以完全不同。 |
| --single-transaction    | 适合innodb事务数据库的备份。保证备份的一致性，原理是设定本次会话的隔离级别为Repeatable read，来保证本次会话（也就是dump）时，不会看到其它会话已经提交了的数据。 |
| -F                      | 刷新binlog，如果binlog打开了，-F参数会在备份时自动刷新binlog进行切换。 |
| -n --no-create-db       | 只导出数据，而不添加CREATE DATABASE 语句。                   |
| -t --no-create-info     | 只导出数据，而不添加CREATE TABLE 语句。                      |
| -d --no-data            | 不导出任何数据，只导出数据库表结构。                         |
| -p --password           | 连接数据库密码                                               |
| -P --port               | 连接数据库端口号                                             |
| -u --user               | 指定连接的用户名。                                           |
| -R                      | 备份存储过程和函数数据（如果开发写了函数和存储过程，就备，没写就不备） |
| --triggers              | 备份触发器数据（现在都是开发写触发器                         |

```bash
# 在命令行执行命令，进行全量备份
[root@localhost mysql]# mysqldump -uroot -p123456 -A -R --triggers --master-data=2 --single-transaction | gzip > /tmp/full.sql.gz
mysqldump: [Warning] Using a password on the command line interface can be insecure.
[root@localhost mysql]# ll /tmp/full.sql.gz 
-rw-r--r--. 1 root root 191963 Oct 14 00:01 /tmp/full.sql.gz

# 在命令行执行命令，刷新binlog，便于日后查找
[root@localhost mysql]# mysql -uroot -p123456 -e "flush logs"
mysql: [Warning] Using a password on the command line interface can be insecure.

# 登录数据库，再插入一些数据，模拟增量，这些数据写入了新的binlog
mysql> insert t1 values(4),(5),(6);
Query OK, 3 rows affected (0.01 sec)
Records: 3  Duplicates: 0  Warnings: 0


# 案例2：要求每天凌晨3点半的时候，做数据库备份

1、编写脚本
[root@localhost ~]# cat mysqldump.sh 
#!/bin/bash

USERNAME=root
PASSWORD=123456
DATABASE=linux14

/usr/local/mysql/bin/mysqldump -u${USERNAME} -p${PASSWORD} -R --triggers -B ${DATABASE} --master-data=2 --single-transaction  | gzip > /tmp/MySQL_`date +"%F".sql.gz`


2、加入定时任务
30 03 * * *  /root/mysqldump.sh

```

### 模拟数据丢失



```bash
# 模拟数据丢失
mysql> drop database db1;

# 恢复数据
# 1、mysql数据导入时，临时关闭binlog，不要将恢复数据的写操作也记入
mysql> set sql_log_bin=0;

# 2、先恢复全量
mysql> source /tmp/MySQL_2021-10-15.sql

如果是压缩包呢，那就这么做
mysql> system zcat /tmp/MySQL_2021-10-15.sql.gz | mysql -uroot -p123456

# 3、模拟恢复数据
mysql> drop database db01;
Query OK, 1 row affected (0.01 sec)

mysql> set sql_log_bin=0;
Query OK, 0 rows affected (0.00 sec)

mysql> system zcat /tmp/MySQL_2021-10-15.sql.gz | mysql -uroot -p123456
mysql: [Warning] Using a password on the command line interface can be insecure.
mysql> use db01;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> select * from t1;
+------+
| id   |
+------+
|    1 |
|    2 |
|    3 |
+------+
3 rows in set (0.00 sec)
```

### 数据的快速导入与导出

在公司中，如果运营或者产品手里有几千万甚至几亿条数据，要求你将其导入数据中，请问如何做？

如果你依据运营或产品交给你的数据文件直接使用insert语句，一行一行地批量插入，那至少需要1-2天时间才能插入完毕，显然是不可行的。

此时我们可以用LOAD DATA INFILE语句。LOAD DATA INFILE语句可以从一个文本文件中，将数据以很高的速度读入一个表中。MySQL官方文档也说明了，该方法比一次性插入一条数据性能快20倍。

此外，mysql也支持快速导出语句SELECT INTO OUTFILE，使用MySQL的SELECT INTO OUTFILE 、LOAD DATA INFILE快速导出导入数据，12G的数据导出用时3分钟左右，导入用时4分钟左右（执行时间根据机器的配置会有所不同，不具有参考价值）。

- 快速导出

```bash
语法：
SELECT... INTO OUTFILE 导出文本文件

要想导出成功，需要设置安全目录才行
vim /etc/my.cnf
[mysqld]
secure-file-priv=/tmp

示例：
SELECT * FROM db1.t1
    INTO OUTFILE '/tmp/db1_t1.txt'
    FIELDS TERMINATED BY ','      -- 定义字段分隔符
    OPTIONALLY ENCLOSED BY '"'    -- 定义字符串使用什么符号括起来
    LINES TERMINATED BY '\n';     -- 定义换行符
```

- 快速导入

```bash
语法
LOAD DATA INFILE 导入的文本文件路径

示例
mysql> DELETE FROM student1;
mysql> create table new_t1(表结构与文件中数据保持一致);
mysql> LOAD DATA INFILE '/tmp/db1_t1.txt'
            INTO TABLE new_db.new_t1
            FIELDS TERMINATED BY ','
            OPTIONALLY ENCLOSED BY '"'
            LINES TERMINATED BY '\n';
```
