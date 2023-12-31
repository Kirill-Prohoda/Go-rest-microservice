package ginapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinApp() {
	r := gin.Default()
	r.Routes()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
