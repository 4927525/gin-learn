package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBind(&form); err != nil {
			fmt.Println(err.Error())
			return
		}

		if form.User == "user" && form.Password == "password"{
			c.JSON(http.StatusOK, gin.H{"status" : "ok"})
		}else{
			c.JSON(http.StatusOK, gin.H{"status" : "not ok"})
		}
	})

	r.Run()
}
