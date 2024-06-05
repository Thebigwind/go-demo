#!/bin/bash
set -e

SHELL_SCRIPT_PATH=$(cd $(dirname $0); pwd)
echo "SHELL_SCRIPT_PATH:"$SHELL_SCRIPT_PATH

rpm -qa|grep wget
if [ $? -ne 0 ];then
    echo "not exit"
fi
echo "already exit"