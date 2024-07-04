package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"ID":   "Bing.X",
			"Name": "小饼",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
