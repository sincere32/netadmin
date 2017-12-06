#!/bin/sh
go=$(which go)
local_dir=$(cd `dirname $0`; pwd)
if [ ! -f $GOPATH/bin/bee ];then
   $go get github.com/beego/bee
fi

cd $local_dir && $GOPATH/bin/bee run -gendoc=true -downdoc=true