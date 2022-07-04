package upl

import "fmt"

func (r *run) siz(siz int64) string {
	const unit = 1000

	if siz < unit {
		return fmt.Sprintf("%d B", siz)
	}

	div, exp := int64(unit), 0

	for n := siz / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(siz)/float64(div), "kMGTPE"[exp])
}
