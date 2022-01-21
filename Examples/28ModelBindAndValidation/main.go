package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `json:"user" form:"user" xml:"user" binding:"required"`
	Password string `json:"password" form:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "test" || json.Password != "test" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	})

	r.Run()
}
