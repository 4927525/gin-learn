package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//映射查询字符串或表单参数
func main() {
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.QueryMap("names")

		fmt.Printf("ids:%v, names:%v", ids, names)
	})

	r.Run()
}
