package UT0311L04

import (
	"fmt"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

func TestCheckTimeProfileWithValidProfile(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(0, 0)
	end := types.NewHHmm(23, 59)

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  true,
			time.Sunday:    true,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := true
	ok := checkTimeProfile(profile, 0)

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for a valid profile - expected: %v", ok, expected)
	}
}

func TestCheckTimeProfileInTimeSegment(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	now := time.Now()
	start := types.HHmmFromTime(now.Add(-10 * time.Minute))
	end := types.HHmmFromTime(now.Add(10 * time.Minute))

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  true,
			time.Sunday:    true,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := true
	ok := checkTimeProfile(profile, 0)

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an in-bounds time segment - expected: %v", ok, expected)
	}
}

// FIXME: won't work before 00:10
func TestCheckTimeProfileBeforeTimeSegment(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	now := time.Now()
	start := types.HHmmFromTime(now.Add(10 * time.Minute))
	end := types.HHmmFromTime(now.Add(60 * time.Minute))

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  true,
			time.Sunday:    true,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := false
	ok := checkTimeProfile(profile, 0)

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an out-of-bounds time segment - expected: %v", ok, expected)
	}
}

// FIXME: won't work after 23:50
func TestCheckTimeProfileAfterTimeSegment(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	now := time.Now()
	start := types.HHmmFromTime(now.Add(-60 * time.Minute))
	end := types.HHmmFromTime(now.Add(-10 * time.Minute))

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  true,
			time.Sunday:    true,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := false
	ok := checkTimeProfile(profile, 0)

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an out-of-bounds time segment - expected: %v", ok, expected)
	}
}

func TestCheckTimeProfileWithValidProfileAndTimeOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(0, 0)
	end := types.NewHHmm(23, 59)

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  true,
			time.Sunday:    true,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := true
	ok := checkTimeProfile(profile, entities.Offset(-8*time.Hour))

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for a valid profile + offset - expected: %v", ok, expected)
	}
}

func TestCheckTimeProfileInTimeSegmentWithOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(11, 30)
	end := types.NewHHmm(12, 30)

	now := time.Now()
	datetime := fmt.Sprintf("%v 11:57:32", now.Format("2006-01-02"))

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  true,
			time.Sunday:    true,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := true
	ok := checkTimeProfile(profile, offset(datetime))

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an in-bounds time segment + offset - expected: %v", ok, expected)
	}
}

func TestCheckTimeProfileBeforeTimeSegmentWithOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(11, 30)
	end := types.NewHHmm(12, 30)

	now := time.Now()
	datetime := fmt.Sprintf("%v 11:05:32", now.Format("2006-01-02"))

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  true,
			time.Sunday:    true,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := false
	ok := checkTimeProfile(profile, offset(datetime))

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an out-of-bounds time segment + offset - expected: %v", ok, expected)
	}
}

func TestCheckTimeProfileAfterTimeSegmentWithOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(11, 30)
	end := types.NewHHmm(12, 30)

	now := time.Now()
	datetime := fmt.Sprintf("%v 13:13:13", now.Format("2006-01-02"))

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  true,
			time.Sunday:    true,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := false
	ok := checkTimeProfile(profile, offset(datetime))

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an out-of-bounds time segment + offset - expected: %v", ok, expected)
	}
}

func offset(datetime string) entities.Offset {
	utc, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
	now := time.Now()
	delta := utc.Sub(now)

	return entities.Offset(delta)
}
