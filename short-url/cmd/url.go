package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"log"
	"net/http"
	_ "short-url/docs"
	adminService "short-url/internal/controller/admin"
	"short-url/internal/controller/url"
	"short-url/internal/middleware/event"
	"short-url/pkg/pg"
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
	StartMQServer()
	e := echo.New()

	// swagger docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// infra
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Routes
	e.GET("/:token", url.GetLongURL, event.RecordEvent)
	e.POST("/shorten", url.SetSortURL)

	admin := e.Group("/admin")
	admin.GET("/statistic/:password", adminService.GetAdminStatistics)
	admin.POST("/register", adminService.CreateAdmin)

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

func StartMQServer() {
	// start mq server
	go func() {
		for {
			dataChan, err := event.MQ.Consume()
			if err != nil {
				println("mq consume error", err.Error())
			}
			select {
			case data := <-dataChan:
				var e pg.Event
				err = json.Unmarshal(data.Body, &e)
				if err != nil {
					println("record event error", err.Error())
				}
				err = adminService.UserService.CreateEvent(e)
				if err != nil {
					println("record event error", err.Error())
				}
			}
		}
	}()
}
