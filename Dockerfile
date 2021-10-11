FROM golang:1.16
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

ENV TZ=Asia/Shanghai

# 创建代码目录
WORKDIR  /src
# copy代码放入代码目录
COPY . .
# 将代码编译为二进制可执行文件
RUN go build -o main .
# 创建运行运行环境
WORKDIR /bin
#将二进制文件从 /src 移动到/bin
RUN cp /src/main .
# 将项目所需配置和static资源copy至该目录
RUN cp -r /src/config .
RUN cp -r /src/public .
RUN cp -r /src/views .

# 暴露端口对外服务
EXPOSE 8080

# 启动容器运行命令
CMD ["/bin/main"]



