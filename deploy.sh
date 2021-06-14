#!/bin/bash
# arg=`ifconfig eth0 | grep 'inet' | awk '{print $2}' | awk -F ':' '{print $2}'`
if [  -n $WX_SECURITYKEY ];then
go run main.go $WX_SECURITYKEY
fi
