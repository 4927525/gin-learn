package main

import "github.com/gin-gonic/gin"

//查询字符串参数
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		first := c.DefaultQuery("first", "guest")
		last := c.Query("last")

		c.String(200, "%s%s", first, last)
	})

	r.ru
}
