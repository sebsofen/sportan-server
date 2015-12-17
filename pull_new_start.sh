#!/bin/bash
./shutdown.sh
git pull -u origin master
go build main.go
./startup.sh
