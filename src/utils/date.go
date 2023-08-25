package utils

import (
	"fmt"
	"time"
)

func DateLogger() string {

	date := time.Now()
	dateFormat := fmt.Sprintf("%d/%d/%d", date.Month(), date.Day(), date.Year())
	finalFormat := fmt.Sprintf("[%s - %d:%d:%d] -", dateFormat, date.Hour(), date.Minute(), date.Second())
	return finalFormat
}
