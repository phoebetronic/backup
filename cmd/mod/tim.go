package mod

import "time"

func datfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02")
}

func clofmt(tim time.Time) string {
	return tim.UTC().Format("15:04:05")
}

func scrfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02 15:04:05")
}
