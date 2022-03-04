package main

import (
	"github.com/gin-gonic/gin"
	"grayRelease/src/router"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	r := gin.Default()
	router.CustomRoute(r)
	r.Run(":8000")
}
