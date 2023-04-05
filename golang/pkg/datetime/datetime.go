package datetime

import "time"

type DateTime struct {
	Unix int64
	Date time.Time
}

// NewDateTime creates a new instance of DateTime.
func NewDateTime(date time.Time) DateTime {
	return DateTime{
		Date: date,
		Unix: date.Unix(),
	}
}

// NewDateTimeFromUnix creates a new instance of DateTime from unix.
func NewDateTimeFromUnix(unix int64) DateTime {
	date := time.Unix(unix, 0)

	return DateTime{
		Date: date,
		Unix: date.Unix(),
	}
}

func Now() DateTime {
	date := time.Now().UTC()

	return NewDateTime(date)
}
