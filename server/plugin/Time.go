package plugin

import (
	"time"
)

func GetDateTime(format string) string {
	dateTime := time.Now()
	return dateTime.Format(format)
}
