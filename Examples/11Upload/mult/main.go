package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//多文件
func main() {
	r := gin.Default()
	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			i := 0
			dst := "uploads" + string(i) + ".xlsx"
			c.SaveUploadedFile(file, dst)
			i++
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
	})

	r.Run()
}
