package controllers

import "github.com/gin-gonic/gin"

func ExampleView(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"name": "Fulan",
	})
}

func ExampleApi(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  true,
		"message": "",
		"data":    nil,
		"code":    "",
	})
}
