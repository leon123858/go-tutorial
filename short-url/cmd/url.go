package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	e.GET("/:token", func(c echo.Context) error {
		// redirect to the origin URL
		token := c.Param("token")
		return c.String(http.StatusOK, token)
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
	s := http.Server{
		Addr:        "0.0.0.0:8080",
		Handler:     e,
		ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
