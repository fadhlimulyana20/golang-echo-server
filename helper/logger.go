package helper

import (
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func MakeLogEntry(c echo.Context, data interface{}, user interface{}) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at":   time.Now().Format("2006-01-02 15:04:05"),
			"data": data,
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
		"data":   data,
		"by":     user,
	})
}

// func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		MakeLogEntry(c).Info("incoming request")
// 		return next(c)
// 	}
// }

// func ErrorHandler(err error, c echo.Context) {
// 	report, ok := err.(*echo.HTTPError)
// 	if ok {
// 		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
// 	} else {
// 		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	MakeLogEntry(c).Error(report.Message)
// 	c.HTML(report.Code, report.Message.(string))
// }
