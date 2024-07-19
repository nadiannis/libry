package utils

import (
	"time"
)

func FormatDate(datetime time.Time) string {
	format := "02 Jan 2006"
	return datetime.Format(format)
}
