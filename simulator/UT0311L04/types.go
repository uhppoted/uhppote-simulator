package UT0311L04

import (
	"encoding/json"
	"strings"
	"time"
)

type ReleaseDate time.Time

func DefaultReleaseDate() *ReleaseDate {
	date := ReleaseDate(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local))

	return &date
}

func (d ReleaseDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

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
