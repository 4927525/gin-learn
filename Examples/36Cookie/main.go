package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "not set"
			c.SetCookie("gin_cookie", "test", 3600, "/",
				"localhost", false, true)
		}

		fmt.Printf("cookie value %s", cookie)
	})

	r.Run()

}
