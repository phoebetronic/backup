package upl

import "time"

const (
	filfmt = "06-01-02.15-04-05"
)

func timfmt(tim time.Time) string {
	return tim.UTC().Truncate(time.Minute).Format(filfmt)
}
