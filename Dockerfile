#体积大
#FROM golang:latest
#体积小
FROM scratch
WORKDIR $GOPATH/src/gin-blog
COPY . $GOPATH/src/gin-blog
#处理容器拉取库超时
RUN go env -w GOPROXY="https://goproxy.cn"
RUN go env -w GO111MODULE="on"
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]
