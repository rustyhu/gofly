package someframes

import (
	"github.com/gin-gonic/gin"
)

func ginDemo() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// fmt.Println("listen and serve on localhost:8080")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
