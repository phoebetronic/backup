package dow

import "time"

func ordfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02.15-04-05")
}

func trafmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02")
}

func scrfmt(tim time.Time) string {
	return tim.UTC().Format("06-01-02 15:04:05")
}
