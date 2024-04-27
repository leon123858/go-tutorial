package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

/**
 * Documentation: https://github.com/gin-gonic/gin/blob/v1.9.1/docs/doc.md
 */
func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		panic("Error loading .env file")
	}
	if os.Getenv("GO_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/error", func(c *gin.Context) {
		panic("error")
	})

	api := r.Group("/api")

	api.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	hostUrl := fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT"))
	err := r.Run(hostUrl)
	if err != nil {
		panic(err)
	}
}
