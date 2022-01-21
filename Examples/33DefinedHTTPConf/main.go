package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//自定义 HTTP 配置
//直接使用 http.ListenAndServe()，如下所示：
func main() {
	r := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
