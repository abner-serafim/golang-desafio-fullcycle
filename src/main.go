package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)


var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>{{ .developer }} - {{ .message }}</title>
</head>
<body>
	<b>{{ .message }}</b>
<br /><br />
  <b>{{ .developer }}</b>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		a := c.Query("a");

		if s, err := strconv.ParseFloat(a, 64); err == nil {
			b, _ := strconv.ParseFloat(c.Query("b"), 64)
			c.JSON(200, gin.H{
				"resultado": s + b,
			})
		} else {
			c.HTML(200, "https", gin.H{
				"message": "Code.education Rocks!",
				"developer": "Abner Serafim Arede",
			})
		}
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

