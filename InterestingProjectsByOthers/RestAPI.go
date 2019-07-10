package InterestingProjectsByOthers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func RestAPI() {
	fmt.Println()
	r := gin.Default()
	r.GET("/", Homepage)
	_ = r.Run()
}
