package main

import (
	"./service"
	"github.com/gin-gonic/gin"
)

func main() {

	data := service.GetComicURL()

	r := gin.Default()
	r.GET("/getComic", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"comicURL": data,
		})
	})
	r.Run()
}
