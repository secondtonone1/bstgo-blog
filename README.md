# bstgo-blog
前后一体化博客系统

## 初衷
市面上很多博客系统，为什么还要自己造个轮子，我也曾经问过自己，早些年用hexo自动化发布了个人博客，后来发现最好可以做一个索引式的博客供自己查询，也便于交流分享。看了一些博客页面不尽满意，遂决心自学前端写一个全栈博客系统，常年从事服务器开发，学习前端是件苦差事，于2021年初用了四个月时间学习CSS,html,js,jquery,bootstrap，后端采用go开发，初步完成了博客系统，当然后期还会接触一些新的前端知识，vue之类的以后有时间会学习。或许，当生活枯燥的时候，就尝试做一些自己不擅长的事情，突破舒适区后带来的成就感会不断激励自己螺旋上升。

## 准备工作
如果想体验代码和效果可直接跳过以下内容，代码连接的是我的远程数据库和redis
如果需要配置自己的数据库和redis可以参考以下步骤
###  mongo容器
1 先拉取mongo镜像
docker pull mongo:latest
2 创建挂载目录
``` cmd
mkdir -p /data/mongo/data
mkdir -p /data/mongo/log
mkdir -p /data/mongo/conf
touch /data/mongo/log/mongodb.log
chmod 777 /data/mongo/log/mongodb.log
cd /data/mongo/conf
vim mongod.conf
添加如下配置
```

``` cmd
# 数据库文件存储位置
dbpath = /data/db/
# log文件存储位置
logpath = /data/log/mongodb.log
# 使用追加的方式写日志
logappend = true
# 是否以守护进程方式运行
# fork = true
# 端口号
port = 27017
# 是否启用认证
auth = true
# 设置oplog的大小(MB)
oplogSize = 2048 
```

启动docker
```
docker run -itd --name blogmongo  --restart=always --privileged -p 27017:27017 -v /data/mongo/data:/data/db  -v /data/mongo/conf:/data/configdb  -v /data/mongo/log:/data/log/  mongo  -f /data/configdb/mongod.conf --bind_ip_all  
# -v 指定配置文件启动
# --bind_ip_all 允许所以IP访问
# ----restart=always  Docker服务重启容器也启动
# --privileged  拥有真正的root权限 
```

创建用户
```
#创建用户admin，具有读写权限
docker exec -it blogmongo mongo admin
# 创建一个名为 admin，密码为 123456 的用户。
>  db.createUser({ user:'admin',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});
# 尝试使用上面创建的用户信息进行连接。
> db.auth('admin', '123456')
```
   
远程连接mongo
``` 
#1  url
mongodb://admin:123456@81.68.86.146:27017/?authSource=admin
#2  用户名 admin
密码 123456
```

集合表设置index

```
#登录用户admin
docker exec -it blogmongo mongo admin
db.auth('admin','123456')
#切换数据库
use blog
#创建索引
db.artcontents.createIndex({"id":1})
db.articles.createIndex({"cat":1, "subcat":1 })
db.articles.createIndex({"title":1,"subtitle":1 })
db.articles.createIndex({"content":1})
db.articles.createIndex({"createdAt":-1})
db.articles.createIndex({"id":1})
db.articles.createIndex({"index":1})
db.baseinfos.createIndex({"type":1})
db.comments.createIndex({"id":1})
db.comments.createIndex({"parent":1})
db.loginfaileds.createIndex({"createdAt":-1})
db.loginfaileds.createIndex({"email":1})
db.menus.createIndex({"catid":1})
db.menus.createIndex({"index":1})
db.menus.createIndex({"parent":1})
db.sessions.createIndex({"createdAt":-1})
db.sessions.createIndex({"sid":-1})
```

###  redis容器

创建目录
```
mkdir -p /data/redis/conf
mkdir -p /data/redis/data
```

拉取最新镜像
```
docker pull redis:latest
```

redis的配置文件在项目的docker目录中配置好了，修改了如下配置

```
#注释掉这部分，这是限制redis只能本地访问
#bind 127.0.0.1 

#默认yes，开启保护模式，限制为本地访问
protected-mode no 

#默认no，改为yes意为以守护进程方式启动，可后台运行，除非kill进程，改为yes会使配置文件方式启动redis失败
daemonize no

#数据库个数（可选），我修改了这个只是查看是否生效
databases 16

#输入本地redis数据库存放文件夹（可选）
dir  ./ 

#redis持久化（可选）
appendonly yes 
```

docker 启动redis
```
docker run -p 6379:6379 --name blogredis -v /data/redis/conf/redis.conf:/etc/redis/redis.conf -v /data/redis/data:/data -d redis redis-server /etc/redis/redis.conf --appendonly yes
```

配置密码
```
docker exec -it blogredis bash
进入目录 cd /usr/local/bin
运行命令：redis-cli
查看现有的redis密码：config get requirepass
设置redis密码config set requirepass 123456（123456为你要设置的密码）

若出现(error) NOAUTH Authentication required.错误，则使用 auth 密码 来认证密码
```
## 启动项目
### 测试环境
进入项目目录 go run ./main.go 直接启动，自动连接数据库和redis

## 管理后台
url = /admin
用户名"secondtonone1@163.com"
密码 "123456"

## 主页
url = /home
![主页](./demopic/2.png)
## 分类页
url = /category?catid=
![分类页](./demopic/3.png)
## 详情页
url = /article?id=
![详情页](./demopic/1.png)
## 测试版效果

