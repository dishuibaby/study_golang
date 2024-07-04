package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"text/template"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLGlob("tpl/**/*")

	r.GET("/user/show", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title":   "这是一行标题",
			"Content": "用户页面",
			"Now":     time.Now(),
		})
	})

	r.GET("/post/show", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post/index.tmpl", gin.H{
			"title":   "这是一行标题",
			"Content": "文章页面",
		})
	})

	r.Run(":8080")
}
