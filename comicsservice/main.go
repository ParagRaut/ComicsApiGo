package main

import (
	"./service"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/getComic", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"comicURL": service.GetComicURL(),
		})
	})
	r.Run()
}
