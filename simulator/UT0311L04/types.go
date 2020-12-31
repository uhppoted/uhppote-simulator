package UT0311L04

import (
	"encoding/json"
	"strings"
	"time"
)

// Firmware release date.
type ReleaseDate time.Time

// Returns a fixed firmware release date to 2020-01-01 for default simulator initialisation.
func DefaultReleaseDate() *ReleaseDate {
	date := ReleaseDate(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local))

	return &date
}

// Custom JSON marshaller to format the release date as a string of the form 'YYYY-mm-dd'.
func (d ReleaseDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

// Custom JSON unmarshaller to parse a date string of the form 'YYYY-mm-dd'.
func (d *ReleaseDate) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	if strings.TrimSpace(s) == "" {
		*d = ReleaseDate(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local))
		return nil
	}

	date, err := time.ParseInLocation("2006-01-02", s, time.Local)
	if err != nil {
		return err
	}

	*d = ReleaseDate(date)

	return nil
}
