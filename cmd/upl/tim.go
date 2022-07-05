package upl

import "time"

func bacfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02")
}

func monfmt(tim time.Time) time.Time {
	y, m, _ := tim.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, tim.UTC().Location())
}

func scrfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02 15:04:05")
}
