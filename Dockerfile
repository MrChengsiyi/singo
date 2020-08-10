FROM golang as build
ADD . /usr/local/go/src/giligili
WORKDIR /usr/local/go/src/giligili
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7
ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=3000

#第一行是为了换源 第二行是更新 第三行是安装一个包 这个包可以使用https 第四行是为了解决一个bug 第五行是创建一个conf文件夹
RUN echo -e "https://mirrors.aliyun.com/alpine/v3.7/main" > /etc/apk/repositories &&\
    apk update && \
    apk add ca-certificates && \
    echo "hosts:files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /usr/local/go/src/giligili/api_server /usr/bin/api_server

#将配置文件转入
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]
#usr/bin/路径下的算是path路径 可以省略