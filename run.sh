#!/bin/sh
app="netadmin"
go=$(which go)
local_dir=$(cd `dirname $0`; pwd)

cd $local_dir && $go build && ./netadmin
