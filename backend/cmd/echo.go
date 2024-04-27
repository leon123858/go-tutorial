package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var (
	users = []string{"Joe", "Veer", "Zion"}
)

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		panic("Error loading .env file")
	}
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}:${uri} => ${status} , from ${remote_ip}, ${latency_human} (${bytes_in}/${bytes_out})\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	e.GET("/error", func(c echo.Context) error {
		panic("error")
	})

	api := e.Group("/api")

	api.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello world",
		})
	})

	api.GET("/users", getUsers)

	hostUrl := fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT"))
	err := e.Start(hostUrl)
	if err != nil {
		return
	}
}
