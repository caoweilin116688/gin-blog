#体积大
#FROM golang:latest
#WORKDIR $GOPATH/src/gin-blog
#COPY . $GOPATH/src/gin-blog
##处理容器拉取库超时
#RUN go env -w GOPROXY="https://goproxy.cn"
#RUN go env -w GO111MODULE="on"
#RUN go build .
##非scratch 使用 ENTRYPOINT
##ENTRYPOINT ["./gin-blog"]


#体积小
FROM scratch
WORKDIR $GOPATH/src/gin-blog
COPY . $GOPATH/src/gin-blog
EXPOSE 8000
#scratch 版本使用一下 CMD
CMD ["./gin-blog"]

