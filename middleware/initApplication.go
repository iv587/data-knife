package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func initApplication(c *gin.Context) {
	fmt.Println("hello")
	c.Next()
}
