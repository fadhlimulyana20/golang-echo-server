package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fadhlimulyana20/golang-echo-server/config"
	"github.com/fadhlimulyana20/golang-echo-server/database"
	"github.com/fadhlimulyana20/golang-echo-server/middleware"
	"github.com/fadhlimulyana20/golang-echo-server/router"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Initialize config
	config.Init()

	// Logger middleware
	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	// Cors Middleware
	e.Use(m.CORSWithConfig(m.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	// Database Connection
	err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	// Router Init
	api := new(router.Api)
	api.Init(e)

	// Start HTTP Server
	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(":5000") }(lock)

	middleware.MakeLogEntry(nil).Info(config.TimeZone)

	time.Sleep(1 * time.Millisecond)
	middleware.MakeLogEntry(nil).Warning("application started without ssl/tls enabled")

	err = <-lock
	if err != nil {
		middleware.MakeLogEntry(nil).Panic("failed to start application")
	}
}
