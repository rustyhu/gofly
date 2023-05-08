package someframes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GinDemo() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	confaddr := ":8080"
	fmt.Println("listen and serve on", confaddr)
	// default port of gin
	r.Run() // 0.0.0.0:8080
}
