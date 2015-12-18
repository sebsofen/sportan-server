#!/bin/bash
cd /home/blog/go/src/sportan
./shutdown.sh
git pull -f -u origin master
go build main.go
./startup.sh
