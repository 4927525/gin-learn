package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//单文件
func main() {
	r := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 单
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		c.SaveUploadedFile(file, "abcd.xlsx")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Run()
}
