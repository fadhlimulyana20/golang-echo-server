package config

import (
	"fmt"
	"os"
	"time"
)

var TimeZone = "Asia/Jakarta"

func TimeInit() string {
	tz := os.Getenv("TIME_ZONE")

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return err.Error()
	}

	time.Local = loc
	TimeZone = tz

	return fmt.Sprintf("Set Time Zone to %s", tz)
}
