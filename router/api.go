package router

import (
	"github.com/fadhlimulyana20/golang-echo-server/controller"
	"github.com/gofiber/fiber/v2"
)

type Api struct{}

func (a *Api) Init(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Get("/me", controller.Me)
}
