package win

import "time"

func scrfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02 15:04:05")
}
