package main

import (
	"fmt"

	c "github.com/bysergr/mp3_platform/auth/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", c.Login)
	r.POST("/validate", c.Validate)
	r.POST("/register", c.Register)

	fmt.Println("Server Start")
	r.Run("0.0.0.0:5000")
}
