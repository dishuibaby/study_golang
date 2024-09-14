package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		name := c.DefaultQuery("name", "一张小饼")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		name := c.DefaultQuery("name", "一张小饼")
		age := c.PostForm("age")
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "ok",
			"data": gin.H{
				"name": name,
				"age":  age,
			},
		})
	})

	r.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "del ok",
			"data": gin.H{
				"id": id,
			},
		})

	})

	r.Run(":8080")
}
