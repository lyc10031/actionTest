#!/bin/bash
ifconfig eth0 | grep 'inet' | awk '{print $2}' | head -1
