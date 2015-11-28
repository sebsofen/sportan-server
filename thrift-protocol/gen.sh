#!/bin/bash
thrift -r --gen go  -out /home/sebastian/go/src/sportan Service.thrift
thrift -r --gen java:generated_annotations=suppress  -out /home/sebastian/sportan/androidapp/app/src/main/java/ Service.thrift
thrift -r --gen py  -out /home/sebastian/sportan/tools/ Service.thrift
