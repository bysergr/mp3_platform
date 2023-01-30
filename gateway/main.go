package main

import (
	c "github.com/bysergr/mp3_platform/gateway/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.POST("/login", c.Login)
	r.POST("/upload", c.Upload)
	r.GET("/download", c.Download)

	r.Run("0.0.0.0:8080")
}
