#!/bin/bash
arg=`ifconfig eth0 | grep 'inet' | awk '{print $2}' | awk -F ':' '{print $2}'`
go run main.go $arg $MSG_SCKEY
