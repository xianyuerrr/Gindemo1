package main

import (
	"demo1/src/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	// 使用 pprof，可能需要把 6060 端口打开
	// go tool pprof --seconds 30 默认采集 30s
	// 使用 go tool pprof http://152.136.103.47:6060/debug/pprof/profile 来查看相关信息
	// 使用 go tool pprof -http=:8080 http://152.136.103.47:6060/debug/pprof/profile 来查看相关信息
	// 使用 go tool pprof -http=:8080 "http://152.136.103.47:6060/debug/pprof/heap" 来查看图
	// http://152.136.103.47:6060/debug/pprof
	// 然后应该就能使用 top 和 web 来查看信息了

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	
	r := gin.Default()
	router.CustomRoute(r)
	//model.Mysql()
	if err := r.Run(":8000"); err != nil {
		return
	}
}
