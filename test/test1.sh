#!/bin/bash

olddb=''
newdb=''
if [ $1 ]; then
   echo "oldDB is $1"
   olddb=$1
else
   echo "oldDb is empty"
   exit 0;
fi

if [ $2 ]; then
   echo "new db is $2"
   newdb=$2
else
   echo "new db is empty"
   exit 0;
fi

echo "move database $olddb to $newdb "

my_connect="mysql -h 10.10.10.162 -P33307 -u root -p123456"
echo "$my_connect"

$my_connect -e "create database if not exists $newdb"

list_table=$( $my_connect -Nse "select table_name from information_schema.tables where table_schema='$olddb'" )

echo "start to move $olddb tables ... "
for table in $list_table
do
   $my_connect -e "rename table $olddb.$table to $newdb.$table"
done

echo "move $olddb tables finished "
#echo "drop database $olddb ..."
#$my_connect -e "drop database if exists $olddb"

echo "move success"