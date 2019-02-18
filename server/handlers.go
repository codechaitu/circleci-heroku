package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	fmt.Println("user taxi")
	c.String(200, "User, you taxi is confirmed.")
}
