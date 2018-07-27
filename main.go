package main

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	// Tmpl 模版
	Tmpl *template.Template
)

// GetBase 默认属性
func GetBase() gin.H {
	return gin.H{}
}

func renderHTMLFront(c *gin.Context, name string, data gin.H) {
	err := Tmpl.ExecuteTemplate(c.Writer, "./views/home.html", data)
	if err != nil {
		panic(err)
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
		tmpl, err := template.ParseFiles("views/index.html")

		if err != nil {
			panic(err)
		}

		tmpl.ExecuteTemplate(c.Writer, "index.html", GetBase())

		c.Header("Content-Type", "text/html; charset=utf-8")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
