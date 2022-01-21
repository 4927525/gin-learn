package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用 LoadHTMLGlob() 或者 LoadHTMLFiles()
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "main website",
		})
	})

	r.Run()
}
