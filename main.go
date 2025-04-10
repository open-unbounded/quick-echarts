package main

import (
	_ "embed"
	"flag"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed chart.html
	charHtml []byte
	addr     = flag.String("addr", "0.0.0.0:8080", "post")
)

func main() {
	engine := gin.New()
	flag.Parse()
	html := template.Must(template.New("chart.html").Parse(string(charHtml)))
	engine.SetHTMLTemplate(html)
	engine.GET("/chart", func(c *gin.Context) {
		cc := c.Query("c")
		c.HTML(200, "chart.html", gin.H{
			"option": cc,
		})
	})
	err := engine.Run(*addr)
	if err != nil {
		log.Panic(err)
	}
}
