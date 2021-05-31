## 镜像

docker images

REPOSITORY          仓库

TAG                 标签

IMAGE ID            id

CREATED             创建事件

SIZE                大小

-a 列出所有镜像

-q 只显示镜像id

docker search

--filter=STARTS=1000

docker pull 镜像名 【tag】

    docker pull mysql
    Using default tag: latest #默认tag
    latest: Pulling from library/mysql
    d121f8d1c412: Pull complete # 分层下载，联合文件系统
    f3cebc0b4691: Pull complete
    1862755a0b37: Pull complete
    489b44f3dbb4: Pull complete
    690874f836db: Pull complete
    baa8be383ffb: Pull complete
    55356608b4ac: Pull complete
    dd35ceccb6eb: Pull complete
    429b35712b19: Pull complete
    162d8291095c: Pull complete
    5e500ef7181b: Pull complete
    af7528e958b6: Pull complete
    Digest:     sha256:e1bfe11693ed2052cb3b4e5fa356c65381129e87e38551c6cd6ec532ebe0e808 #签名
    Status: Downloaded newer image for mysql:latest
    docker.io/library/mysql:latest #真实地址

指定版本

    docker pull mysql:5.7

删除镜像

    docker rmi {id|name}

删除全部镜像

    docker rmi $(docker images -aq)

## 容器

新建容器并启动

    docker run 【可选参数】image

    --name="" 名字

    -d 后台运行

    -it 交互方式运行

    -p 指定端口

        -p ip：主机端口：容器端口
        -p 主机端口：容器端口
        -p 容器端口
        容器端口

    -P 随机端口

列出所有运行的容器

    docker ps

    -a 列出所有在运行的容器，带出历史运行容器

    -n=? 显示最近？个容器

    -q 只显示容器编号

退出容器

    exit：退出并停止

    Ctrl+P+Q：退出不停止

删除容器

    docker rm {容器ID}

启动停止容器

    docker start restart stop


## 常用其他命令

后台启动容器

必须有一个前台进程

查看日志

    docker logs 

    docker logs -f -t --tail 10 id

查看容器中的进程信息

    docker top id

查看容器信息

    docker inspect id

进入当前正在运行的容器

    docker exec -it id {shell}#开启新终端

    docker attach id#进入正在执行的容器

## 从容器内文件拷贝出来

    docker cp 容器id：地址 主机地址


