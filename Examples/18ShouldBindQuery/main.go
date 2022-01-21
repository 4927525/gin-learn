package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// 只绑定 url 查询字符串
//ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据。参阅详细信息.

type Person struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()
	r.Any("/testing", startPage)
	r.Run()
}

func startPage(c *gin.Context) {
	var person Person

	if err := c.ShouldBindQuery(&person); err == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	} else {
		log.Println(err.Error())
	}
	c.String(200, "success")

}
