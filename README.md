# Netadmin
## What's This?
- Netadmin is built for network engineer to manage the network deivices(for now just juniper and cisco) with [beego](https://beego.me/) and [FastAdmin](http://fastadmin.net)
- Netadmin is just platform which provides management for deivces, scripts, users,roles, gitlab
- Using gitlab as the store of scripts and Netadmin gets the scripts from it by gitlab user token, so generate it first
- Script or command is sent to [netadmin-driver](https://github.com/pippozq/netadmin-driver)

## Build  By

Package | Page
---|---
Beego| https://github.com/astaxie/beego
FastAdmin-html | https://github.com/karsonzhang/fastadmin-html

## Warning
Please Reading and Following the protocol of FastAdmin-html.
It can not be used for commercial purposes.Read [Details](https://github.com/karsonzhang/fastadmin-html)

## Deploy Gitlab

```
docker run -d --hostname gitlab --publish 8443:443 --publish 8080:80 --publish 2222:22 --name gitlab --restart always --privileged=true --env TZ=Asia/Shanghai -v /opt/gitlab/config:/etc/gitlab -v /opt/gitlab/logs:/var/log/gitlab -v /opt/gitlab/data:/var/opt/gitlab gitlab/gitlab-ce:latest

```


## Build
At first you need to copy conf/app.conf.example to conf/app.conf and
update the following configurations

```
# Netadmin Database
user = "netadmin"
password = "nfsetso12fdds9s"
db_name = "netadmin"
host = "postgres_host" // your host
port = "5432"
max_idle_conn = 20
max_open_conn = 50

# debug
debug = false

# netadmin driver
netadmin_driver_url = "http://host:port/netdriver" // netadmin driver url


# session
sessionproviderconfig = "redis_host:redis_port" // redis host port
sessionhashkey = "abcdefghijklmnopqrstuvwxyz" // use your hashkey
```

You can build it esaily using Dockerfile
```
docker build -t <your docker repository url>/netadmin:<your version> .

```
Run it
```
docker run -d --name netadmin -p 8081:8081 -p 8088:8088 -e ENV_MODE=dev  <your docker repository url>/netadmin:<your version>

```

ENV_MODE is the mode in  configuration. check [app.conf.example](https://github.com/pippozq/netadmin/blob/master/conf/app.conf.example)

---

then view http://ip:8081

---
Default User
```
name:admin
password:password

```


## License
source code is licensed under the Apache Licence, Version 2.0 (http://www.apache.org/licenses/LICENSE-2.0.html).