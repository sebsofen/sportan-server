#!/bin/bash
thrift -r --gen go  -out /home/sebastian/go/src/sportan Service.thrift
thrift -r --gen java:generated_annotations=suppress  -out ../../androidapp/app/src/main/java/ Service.thrift
thrift -r --gen py  -out ../../tools/ Service.thrift
