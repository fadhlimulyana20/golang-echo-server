package router

import (
	"github.com/fadhlimulyana20/golang-echo-server/controller"
	"github.com/labstack/echo/v4"
)

type Api struct{}

func (a *Api) Init(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.POST("/register", controller.Register)
	auth.POST("/login", controller.Login)
	auth.GET("/me", controller.Me)
}
