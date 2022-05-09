package helper

import (
	"encoding/json"

	"github.com/fadhlimulyana20/golang-echo-server/model"
	"github.com/labstack/echo/v4"
)

func Json(c echo.Context, httpCode int, err error, data interface{}, message string) error {
	// Get Active User
	// auth := new(Auth)
	// u, _ := auth.GetUser(c)

	if err != nil {
		// If response has data
		if data != nil {
			MakeLogEntry(c, data, "").Error(message)
		}

		MakeLogEntry(c, nil, "").Error(message)

		if message != "" {
			return c.JSON(httpCode, &model.Json{
				Status:  "failed",
				Message: message,
			})
		}

		return c.JSON(httpCode, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	dataLog, err := json.Marshal(data)
	if err != nil {
		MakeLogEntry(c, nil, "").Error(err.Error())
		return c.JSON(httpCode, &model.Json{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	MakeLogEntry(c, string(dataLog), "").Info(message)
	return c.JSON(httpCode, &model.Json{
		Status:  "success",
		Data:    data,
		Message: message,
	})

	// If User is not logged in
	// if u.Id.IsZero() {
	// 	if err != nil {
	// 		// If response has data
	// 		if data != nil {
	// 			MakeLogEntry(c, data, "").Error(message)
	// 		}

	// 		MakeLogEntry(c, nil, "").Error(message)

	// 		if message != "" {
	// 			return c.JSON(httpCode, &model.Json{
	// 				Status:  "failed",
	// 				Message: message,
	// 			})
	// 		}

	// 		return c.JSON(httpCode, &model.Json{
	// 			Status:  "failed",
	// 			Message: err.Error(),
	// 		})
	// 	}

	// 	dataLog, err := json.Marshal(data)
	// 	if err != nil {
	// 		MakeLogEntry(c, nil, "").Error(err.Error())
	// 		return c.JSON(httpCode, &model.Json{
	// 			Status:  "failed",
	// 			Message: err.Error(),
	// 		})
	// 	}

	// 	MakeLogEntry(c, string(dataLog), "").Info(message)
	// 	return c.JSON(httpCode, &model.Json{
	// 		Status:  "success",
	// 		Data:    data,
	// 		Message: message,
	// 	})
	// } else {
	// 	// Marshal user data
	// 	user := map[string]interface{}{
	// 		"id":   u.Id,
	// 		"name": u.Name,
	// 	}

	// 	dataUser, _ := json.Marshal(user)

	// 	if err != nil {
	// 		if data != nil {
	// 			MakeLogEntry(c, data, string(dataUser)).Error(message)
	// 		}

	// 		MakeLogEntry(c, nil, string(dataUser)).Error(message)

	// 		if message != "" {
	// 			return c.JSON(httpCode, &model.Json{
	// 				Status:  "failed",
	// 				Message: message,
	// 			})
	// 		}

	// 		return c.JSON(httpCode, &model.Json{
	// 			Status:  "failed",
	// 			Message: err.Error(),
	// 		})
	// 	}

	// 	dataLog, err := json.Marshal(data)
	// 	if err != nil {
	// 		MakeLogEntry(c, nil, string(dataUser)).Error(err.Error())
	// 		return c.JSON(httpCode, &model.Json{
	// 			Status:  "failed",
	// 			Message: err.Error(),
	// 		})
	// 	}

	// 	MakeLogEntry(c, string(dataLog), string(dataUser)).Info(message)
	// 	return c.JSON(httpCode, &model.Json{
	// 		Status:  "success",
	// 		Data:    data,
	// 		Message: message,
	// 	})
	// }

}
