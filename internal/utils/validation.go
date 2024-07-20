package utils

import "time"

func TimeIsBetween(t, min, max time.Time) bool {
	if min.After(max) {
		min, max = max, min
	}
	return (t.Equal(min) || t.After(min)) && (t.Equal(max) || t.Before(max))
}

func DateIsAfter(t1, t2 time.Time) bool {
	date1 := t1.Truncate(24 * time.Hour)
	date2 := t2.Truncate(24 * time.Hour)
	return date1.After(date2)
}
