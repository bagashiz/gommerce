package helper

import "time"

const (
	// ddmmyyyyLayout is the date format used for the application.
	ddmmyyyyLayout = "02/01/2006"
)

// ParseTime parses a date string into a time.Time object.
func ParseTime(dateStr string) (time.Time, error) {
	t, err := time.Parse(ddmmyyyyLayout, dateStr)
	if err != nil {
		return t, err
	}
	return t, nil
}
