package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")

		// 请求前
		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的是status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()

	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// 打印："12345"
		log.Println(example)
	})

	r.Run()
}
