package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//在中间件中使用 Goroutine
//当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本。
func main() {
	r := gin.Default()

	r.GET("long_async", func(c *gin.Context) {
		// 创建 goroutine中使用的副本
		cCp := c.Copy()
		go func() {
			// 用 time.sleep模拟长任务
			time.Sleep(5 * time.Second)

			// 请注意您使用的是复制的上下文cCp
			log.Println("done! in path", cCp.Request.URL.Path)
		}()
	})

	r.GET("long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)

		// 请注意您使用的是复制的上下文cCp
		log.Println("done! in path", c.Request.URL.Path)
	})

	r.Run()
}
