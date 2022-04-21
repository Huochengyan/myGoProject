#!/usr/bin/env bash

dbhost=127.0.0.1
dbport=27017
dbname=publicChain
dbuser=test
dbpwd=test

collections=("nodeType" "conf" "chain" "exchangeType" "files" "nickName" "setting" "stacType" "user" "person")

DATE=`date +%Y_%m_%d_%H_%M` #获取当前系统时间
mongodump -h ${dbhost}  --port ${dbport} -u  ${dbuser}  -p ${dbpwd} -d ${dbname}    -o ./${DATE}


# 57 14 * * * root /home/thk/go/src/myProject/backmongdb.sh
#每分钟执行      * * * * *
#每五分钟执行    */5 * * * *
#每小时执行      0 * * * *
#每天执行       0 0 * * *
#每周执行       0 0 * * 0
#每月执行       0 0 1 * *
#每年执行       0 0 1 1 *

