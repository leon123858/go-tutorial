package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"log"
	"net/http"
	_ "short-url/docs"
	"time"
)

//	@title			Simple Short URL API Server
//	@version		1.0
//	@description	This is a sample in go tutorial for building a short URL service

//	@contact.name	Leon Lin
//	@contact.url	github.com/leon123858
//	@contact.email	a0970785699@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		127.0.0.1:8080
// @BasePath	/
func main() {
	e := echo.New()

	// swagger docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.GET("/:token", func(c echo.Context) error {
		// get the URL
		return c.String(http.StatusOK, "get")
	})
	e.POST("/shorten", func(c echo.Context) error {
		// shorten the URL
		return c.String(http.StatusOK, "shorten")
	})

	admin := e.Group("/admin")
	admin.GET("/:password", func(c echo.Context) error {
		return c.String(http.StatusOK, "analytics")
	})
	admin.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "password")
	})

	fmt.Println("Starting server at http://127.0.0.1:8080")
	fmt.Println("Swagger docs at http://127.0.0.1:8080/swagger/index.html")
	s := http.Server{
		Addr:        "0.0.0.0:8080",
		Handler:     e,
		ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
