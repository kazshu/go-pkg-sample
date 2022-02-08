package days

import "time"

func AddDays(value int) time.Time {
	t := time.Now()
	t.AddDate(0, 0, value)
	return t
}
