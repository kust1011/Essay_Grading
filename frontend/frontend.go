package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("../templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "interface.html", nil)
	})

	r.Run(":8081")
}
