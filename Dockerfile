FROM golang

RUN mkdir /go/src/demo
COPY . /go/src/demo

RUN go get github.com/spf13/cast
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm

RUN cd /go/src/demo
    && export GO111MODULE="on"
    && export GOPROXY="https://mirrors.aliyun.com/goproxy/"
    && go build -o greyRelease demo1