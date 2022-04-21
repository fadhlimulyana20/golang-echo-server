package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/fadhlimulyana20/golang-echo-server/helper"
	"github.com/fadhlimulyana20/golang-echo-server/model"
	"github.com/labstack/echo/v4"
)

var ctx = context.Background()

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")

		// Check Bearer
		if !strings.Contains(authorizationHeader, "Bearer") {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Anda belum Login.",
			})
		}

		// Extract token
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		j := new(helper.JWT)
		userId, err := j.Parse(tokenString, "token")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, &model.Json{
				Status:  "error",
				Message: err.Error(),
			})
		}

		fmt.Print(userId)

		// Get User by ID
		// u := new(model.User)
		// if err := u.Get(userId); err != nil {
		// 	return c.JSON(http.StatusInternalServerError, &model.Json{
		// 		Status:  "fail",
		// 		Message: err.Error(),
		// 	})
		// }

		return next(c)
	}
}
