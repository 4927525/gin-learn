package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//从 reader 读取数据
func main() {
	r := gin.Default()
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://www.baidu.com/img/flexible/logo/pc/result@2.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="baidu.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
}
