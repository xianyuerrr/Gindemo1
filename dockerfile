from golang

run mkdir /go/src/demo
COPY . /go/src/demo

run go get github.com/spf13/cast
run go get github.com/gin-gonic/gin
run go get github.com/jinzhu/gorm

run cd /go/src/demo
    && export GO111MODULE="on"
    && export GOPROXY="https://mirrors.aliyun.com/goproxy/"
    && go build -o greyRelease demo1