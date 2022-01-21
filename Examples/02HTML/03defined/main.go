package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func formatAsDate(time time.Time) string {
	year, month, day := time.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

//自定义模板渲染器
func main() {
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLFiles("./testdata/template/raw.tmpl")

	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run()
}
