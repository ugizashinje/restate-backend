package utils

import (
	"time"
	"warrant-api/pkg/config"

	"gopkg.in/guregu/null.v4"
)

func ParseNullTime(timeRequested string, valid bool) null.Time {
	var result null.Time
	if valid {
		t, err := time.Parse(config.Format.TimeFormat, timeRequested)
		if err == nil {
			result = null.TimeFrom(t)
		}
	}
	return result
}

func ParseNullDate(dateRequested string, valid bool) null.Time {
	var result null.Time
	if valid {
		t, err := time.Parse(config.Format.DateFormat, dateRequested)
		if err == nil {
			result = null.TimeFrom(t)
		}
	}
	return result
}
