package utils

import (
	"fmt"
	"time"
)

const (
	TimeFormatForFile = "%"
)

func GetFormattedTimeInfo() string {
	t := time.Now()
	return fmt.Sprintf("%02d-%02d-%02d", t.Hour(), t.Minute(), t.Second())
}
