package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)
	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil{
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil{
				log.Printf("failed to push:%v", err)
			}
		}
		c.HTML(200, "http", gin.H{
			"status" : "success",
		})
	})

	r.Run()
}
