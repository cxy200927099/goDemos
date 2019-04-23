#!/bin/bash
srcDir=/home/cxy/redis/code/demos
ip=192.168.126.223

if [ "$1" = "" ]; then
    echo "usage: ./build.sh <file_name>"
    echo "      file_name not include .go"
    exit 1
fi

binFile=$1

curTime=`date "+%Y-%m-%d %H:%M:%S"`
echo "cur time:$curTime"

bin_dir=./bin
echo "mkdir -p $bin_dir"
mkdir -p $bin_dir

echo "rm $bin_dir/$binFile"
rm $bin_dir/$binFile

echo "Building $binFile"
GOOS=linux GOARCH=amd64 go build -o $bin_dir/$binFile -i ./${1}.go
echo "Build Complete"

echo "ls -l $bin_dir/$binFile"
ls -l $bin_dir/$binFile

echo "deploy to server: $ip"
sshpass -p 8tFmAIj256Yw scp -r $bin_dir/$binFile root@$ip:$srcDir
echo "Deploy $binFile complete"

#sshpass -p 8tFmAIj256Yw ssh -o StrictHostKeyChecking=no root@$ip  'nohup '$srcDir'/fp_writer &'