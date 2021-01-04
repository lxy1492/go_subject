package WEB

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

func Index()  {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Lxy WEB")
	})
	router.Run()
}
