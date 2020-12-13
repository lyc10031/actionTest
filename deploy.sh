#!/bin/bash
# arg=`ifconfig eth0 | grep 'inet' | awk '{print $2}' | awk -F ':' '{print $2}'`
if [  -n $MSG_SCKEY ] ;then
go run main.go $MSG_SCKEY
fi
