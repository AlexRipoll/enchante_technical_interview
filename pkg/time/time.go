package time

import "time"

func Current() string {
	return time.Now().Format(time.RFC3339)
}

func CurrentFormat(format string) string {
	return time.Now().Format(format)
}

