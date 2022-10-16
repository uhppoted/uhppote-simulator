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
	now := time.Now().UTC()
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
	now := time.Now().UTC()
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
	now := time.Now().UTC()
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

func TestCheckTimeProfileInTimeSegmentWithTimeOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(11, 30)
	end := types.NewHHmm(12, 30)

	now := time.Now().UTC()
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

func TestCheckTimeProfileBeforeTimeSegmentWithTimeOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(11, 30)
	end := types.NewHHmm(12, 30)

	now := time.Now().UTC()
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

func TestCheckTimeProfileAfterTimeSegmentWithTimeOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(11, 30)
	end := types.NewHHmm(12, 30)

	now := time.Now().UTC()
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

func TestCheckTimeProfileWeekdayInDateTimeSegmentWithDateTimeOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(16, 45)
	end := types.NewHHmm(17, 30)

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   false,
			time.Wednesday: false,
			time.Thursday:  false,
			time.Friday:    false,
			time.Saturday:  false,
			time.Sunday:    false,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := true
	ok := checkTimeProfile(profile, offset("2022-10-03 16:55:00"))

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an in-bounds time segment + offset - expected: %v", ok, expected)
	}
}

func TestCheckTimeProfileWeekdayNotInDateTimeSegmentWithDateTimeOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(16, 30)
	end := types.NewHHmm(17, 30)

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    false,
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
	ok := checkTimeProfile(profile, offset("2022-10-03 16:35:00"))

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an in-bounds time segment + offset - expected: %v", ok, expected)
	}
}

func TestCheckTimeProfileWeekdayBeforeTimeSegmentWithDateTimeOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(16, 45)
	end := types.NewHHmm(17, 30)

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   false,
			time.Wednesday: false,
			time.Thursday:  false,
			time.Friday:    false,
			time.Saturday:  false,
			time.Sunday:    false,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := false
	ok := checkTimeProfile(profile, offset("2022-10-03 16:35:00"))

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an out-of-bounds time segment + offset - expected: %v", ok, expected)
	}
}

func TestCheckTimeProfileWeekdayAfterTimeSegmentWithDateTimeOffset(t *testing.T) {
	from := types.ToDate(2000, time.January, 1)
	to := types.ToDate(2099, time.December, 31)
	start := types.NewHHmm(16, 45)
	end := types.NewHHmm(17, 30)

	profile := types.TimeProfile{
		ID:              37,
		LinkedProfileID: 0,
		From:            &from,
		To:              &to,
		Weekdays: types.Weekdays{
			time.Monday:    true,
			time.Tuesday:   false,
			time.Wednesday: false,
			time.Thursday:  false,
			time.Friday:    false,
			time.Saturday:  false,
			time.Sunday:    false,
		},
		Segments: types.Segments{
			1: types.Segment{
				Start: start,
				End:   end,
			},
		},
	}

	expected := false
	ok := checkTimeProfile(profile, offset("2022-10-03 17:35:00"))

	if ok != expected {
		t.Errorf("checkTimeProfile returned %v for an out-of-bounds time segment + offset - expected: %v", ok, expected)
	}
}

// |----|-------------------------------|-----------------------------------------------------------------------------|--------|
// |    | controller                    | time profile                                                                | access |
// |----|-------------------------------|-----------------------------------------------------------------------------|--------|
// |    | date       | time  | weekeday | start date | end date   | M | T | W | T | F | S | S | start time | end time |        |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 1  | 2022-10-03 | 16:35 | Mon      | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | Y      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 2  | 2022-10-03 | 16:35 | Mon      | 2022-10-05 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | N      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 3  | 2022-10-03 | 16:35 | Mon      | 2000-01-01 | 2022-09-30 | Y | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | N      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 4  | 2022-10-03 | 16:35 | Mon      | 2000-01-01 | 2099-12-31 | N | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | N      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 5  | 2022-10-03 | 16:35 | Mon      | 2000-01-01 | 2099-12-31 | Y | N | N | N | N | N | N | 00:00      | 23:59    | Y      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 6  | 2022-10-03 | 16:35 | Mon      | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 16:45      | 17:30    | N      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 7  | 2022-10-03 | 16:55 | Mon      | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 16:45      | 17:30    | Y      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 8  | 2022-10-03 | 17:35 | Mon      | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 16:45      | 17:30    | N      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 9  | 2022-10-03 | 16:55 | Mon      | 2000-01-01 | 2099-12-31 | Y | N | N | N | N | N | N | 16:45      | 17:30    | Y      |
// |----|------------|-------|----------|------------|------------|---|---|---|---|---|---|---|------------|----------|--------|
// | 10 | 2022-10-04 | 16:55 | Tue      | 2000-01-01 | 2099-12-31 | Y | N | N | N | N | N | N | 16:45      | 17:30    | N      |
// |----|-------------------------------|-----------------------------------------------------------------------------|--------|

func TestCheckTimeProfile(t *testing.T) {
	tests := []struct {
		datetime  string
		startDate string
		endDate   string
		monday    bool
		tuesday   bool
		wednesday bool
		thursday  bool
		friday    bool
		saturday  bool
		sunday    bool
		startTime string
		endTime   string
		allowed   bool
	}{
		{"2022-10-03 16:35:00", "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "00:00", "23:59", true},
		{"2022-10-03 16:35:00", "2022-10-05", "2099-12-31", true, true, true, true, true, true, true, "00:00", "23:59", false},
		{"2022-10-03 16:35:00", "2000-01-01", "2022-09-30", true, true, true, true, true, true, true, "00:00", "23:59", false},
		{"2022-10-03 16:35:00", "2000-01-01", "2099-12-31", false, true, true, true, true, true, true, "00:00", "23:59", false},
		{"2022-10-03 16:35:00", "2000-01-01", "2099-12-31", true, false, false, false, false, false, false, "00:00", "23:59", true},
		{"2022-10-03 16:35:00", "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "16:45", "17:30", false},
		{"2022-10-03 16:55:00", "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "16:45", "17:30", true},
		{"2022-10-03 17:35:00", "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "16:45", "17:30", false},
		{"2022-10-03 16:55:00", "2000-01-01", "2099-12-31", true, false, false, false, false, false, false, "16:45", "17:30", true},
		{"2022-10-04 16:55:00", "2000-01-01", "2099-12-31", true, false, false, false, false, false, false, "16:45", "17:30", false},
	}

	for ix, test := range tests {
		from, _ := types.DateFromString(test.startDate)
		to, _ := types.DateFromString(test.endDate)
		start, _ := types.HHmmFromString(test.startTime)
		end, _ := types.HHmmFromString(test.endTime)

		profile := types.TimeProfile{
			ID:              37,
			LinkedProfileID: 0,
			From:            &from,
			To:              &to,
			Weekdays: types.Weekdays{
				time.Monday:    test.monday,
				time.Tuesday:   test.tuesday,
				time.Wednesday: test.wednesday,
				time.Thursday:  test.thursday,
				time.Friday:    test.friday,
				time.Saturday:  test.saturday,
				time.Sunday:    test.sunday,
			},
			Segments: types.Segments{
				1: types.Segment{
					Start: *start,
					End:   *end,
				},
			},
		}

		expected := test.allowed
		ok := checkTimeProfile(profile, offset(test.datetime))

		if ok != expected {
			t.Errorf("checkTimeProfile:%d returned incorrect access - expected: %v, got:%v", ix+1, expected, ok)
		}
	}
}

func offset(datetime string) entities.Offset {
	utc, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.UTC)
	now := time.Now().UTC()
	delta := utc.Sub(now)

	return entities.Offset(delta)
}
