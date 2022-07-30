package ind

import "time"

type Item struct {
	Buc float32   `json:"buc"`
	Cou int       `json:"cou"`
	Fil string    `json:"fil"`
	Tim time.Time `json:"tim"`
}
