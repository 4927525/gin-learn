package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//01AsciiJSON
//使用 01AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。
func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		//{"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run()
}
