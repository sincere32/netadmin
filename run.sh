#!/bin/sh
app="netadmin"
go=$(which go)
local_dir=$(cd `dirname $0`; pwd)

# install lib
$go get github.com/astaxie/beego
$go get github.com/lib/pq
$go get github.com/astaxie/beego/session
$go get github.com/garyburd/redigo/redis

export ENV_MODE=$1

cd $local_dir && $go build && ./netadmin
