package upl

import "time"

func bacfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02")
}

func scrfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02 15:04:05")
}
