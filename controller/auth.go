package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/fadhlimulyana20/golang-echo-server/database"
	"github.com/fadhlimulyana20/golang-echo-server/helper"
	"github.com/fadhlimulyana20/golang-echo-server/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type loginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	u.Id = primitive.NewObjectID()

	db := database.DbManager()
	_, err := db.Collection("users").InsertOne(ctx, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &model.Json{
		Status: "success",
		Data:   u,
	})
}

func Login(c echo.Context) error {
	l := new(loginDTO)
	u := new(model.User)

	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}
	db := database.DbManager()
	err := db.Collection("users").FindOne(ctx, bson.M{"email": l.Email}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	if l.Password != u.Password {
		return c.JSON(http.StatusUnauthorized, &model.Json{
			Status:  "failed",
			Message: "Password Salah",
		})
	}

	j := new(helper.JWT)
	token, err := j.CreateToken(u.Id.Hex(), 48, "token")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &model.Json{
		Status: "success",
		Data: map[string]interface{}{
			"token": token,
			"user":  u,
		},
	})
}

func Me(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	u := new(model.User)

	if !strings.Contains(authHeader, "Bearer") {
		data := &model.Json{
			Status:  "failed",
			Message: "Anda belum Login",
		}

		return c.JSON(http.StatusUnauthorized, data)
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	j := new(helper.JWT)
	userId, err := j.Parse(tokenString, "token")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	db := database.DbManager()
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	err = db.Collection("users").FindOne(ctx, bson.M{"_id": objectId}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNotFound, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &model.Json{
		Status: "success",
		Data:   u,
	})
}
