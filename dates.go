package common

import "time"

// CurrentYear is a helper to get current year.
// Use tz parameter to provide timezone shift
func CurrentYear(tz int) int {
	dt := time.Now().UTC().Add(time.Duration(tz) * time.Hour)
	return dt.Year()
}

// CurrentMonth is a helper to get current month.
// Use tz parameter to provide timezone shift
func CurrentMonth(tz int) int {
	dt := time.Now().UTC().Add(time.Duration(tz) * time.Hour)
	return int(dt.Month())
}

// CurrentDay is a helper to get current day.
// Use tz parameter to provide timezone shift
func CurrentDay(tz int) int {
	dt := time.Now().UTC().Add(time.Duration(tz) * time.Hour)
	return dt.Day()
}

// CurrentDay is a helper to get current date.
// Use tz parameter to provide timezone shift and format to provide custom format ("2006-01-02" by default)
func CurrentDateStr(tz int, format string) string {
	if format == "" {
		format = "2006-01-02"
	}
	dt := time.Now().UTC().Add(time.Duration(tz) * time.Hour)
	return dt.Format(format)
}
