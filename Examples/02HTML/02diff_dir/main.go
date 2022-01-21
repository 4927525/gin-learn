package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//使用不同目录下名称相同的模板
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title" : "posts",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title" : "users",
		})
	})

	r.Run()
}
