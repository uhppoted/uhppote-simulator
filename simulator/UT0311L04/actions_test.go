package UT0311L04

import (
	"fmt"
	"testing"
	"time"

	"github.com/uhppoted/uhppote-core/types"
	"github.com/uhppoted/uhppote-simulator/entities"
)

// |----|-------------------------------|-----------------------------------------------------------------------------|--------|
// |    | controller                    | time profile                                                                | access |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// |    | date       | time    | weekeday | start date | end date   | M | T | W | T | F | S | S | start time | end time |        |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 1  | 2022-10-03 | 16:35   | Mon      | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | Y      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 2  | 2022-10-03 | 16:35   | Mon      | 2022-10-05 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 3  | 2022-10-03 | 16:35   | Mon      | 2000-01-01 | 2022-09-30 | Y | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 4  | 2022-10-03 | 16:35   | Mon      | 2000-01-01 | 2099-12-31 | N | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 5  | 2022-10-03 | 16:35   | Mon      | 2000-01-01 | 2099-12-31 | Y | N | N | N | N | N | N | 00:00      | 23:59    | Y      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 6  | 2022-10-03 | 16:35   | Mon      | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 16:45      | 17:30    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 7  | 2022-10-03 | 16:55   | Mon      | 2000-01-01 | 2099-12-31 | Y | N | N | N | N | N | N | 16:45      | 17:30    | Y      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 8  | 2022-10-03 | 17:35   | Mon      | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 16:45      | 17:30    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 9  | 2022-10-03 | 16:55   | Mon      | 2000-01-01 | 2099-12-31 | Y | N | N | N | N | N | N | 16:45      | 17:30    | Y      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 10 | 2022-10-04 | 16:55   | Tue      | 2000-01-01 | 2099-12-31 | Y | N | N | N | N | N | N | 16:45      | 17:30    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 11 | today      | 16:35   | -        | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 16:45      | 17:30    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 12 | today      | 16:55   | -        | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 16:45      | 17:30    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 13 | today      | 17:35   | -        | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 16:45      | 17:30    | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 14 | today      | now     | -        | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | 00:00      | 23:59    | Y      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 15 | today      | now     | -        | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | now-10m    | now+10m  | Y      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 16 | today      | now-15m | -        | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | now-10m    | now+10m  | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|
// | 17 | today      | now+15m | -        | 2000-01-01 | 2099-12-31 | Y | Y | Y | Y | Y | Y | Y | now-10m    | now+10m  | N      |
// |----|---------------------------------|-----------------------------------------------------------------------------|--------|

func TestCheckTimeProfile(t *testing.T) {
	utc := time.Now().UTC()
	today1635 := fmt.Sprintf("%v 16:35:00", utc.Format("2006-01-02"))
	today1655 := fmt.Sprintf("%v 16:55:00", utc.Format("2006-01-02"))
	today1735 := fmt.Sprintf("%v 17:35:00", utc.Format("2006-01-02"))
	now := utc.Format("2006-01-02 15:04:05")
	nowN15 := utc.Add(-15 * time.Minute).Format("15:04")
	nowN10 := utc.Add(-10 * time.Minute).Format("15:04")
	nowP10 := utc.Add(+10 * time.Minute).Format("15:04")
	nowP15 := utc.Add(+15 * time.Minute).Format("15:04")

	tests := []struct {
		ID        int
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
		{1, "2022-10-03 16:35:00", "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "00:00", "23:59", true},
		{2, "2022-10-03 16:35:00", "2022-10-05", "2099-12-31", true, true, true, true, true, true, true, "00:00", "23:59", false},
		{3, "2022-10-03 16:35:00", "2000-01-01", "2022-09-30", true, true, true, true, true, true, true, "00:00", "23:59", false},
		{4, "2022-10-03 16:35:00", "2000-01-01", "2099-12-31", false, true, true, true, true, true, true, "00:00", "23:59", false},
		{5, "2022-10-03 16:35:00", "2000-01-01", "2099-12-31", true, false, false, false, false, false, false, "00:00", "23:59", true},
		{6, "2022-10-03 16:35:00", "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "16:45", "17:30", false},
		{7, "2022-10-03 16:55:00", "2000-01-01", "2099-12-31", true, false, false, false, false, false, false, "16:45", "17:30", true},
		{8, "2022-10-03 17:35:00", "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "16:45", "17:30", false},
		{9, "2022-10-03 16:55:00", "2000-01-01", "2099-12-31", true, false, false, false, false, false, false, "16:45", "17:30", true},
		{10, "2022-10-04 16:55:00", "2000-01-01", "2099-12-31", true, false, false, false, false, false, false, "16:45", "17:30", false},
		{11, today1635, "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "16:45", "17:30", false},
		{12, today1655, "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "16:45", "17:30", true},
		{13, today1735, "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "16:45", "17:30", false},
		{14, now, "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, "00:00", "23:59", true},
		{15, now, "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, nowN10, nowP10, true},
		{16, nowN15, "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, nowN10, nowP10, false},
		{17, nowP15, "2000-01-01", "2099-12-31", true, true, true, true, true, true, true, nowN10, nowP10, false},
	}

	for _, test := range tests {
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
			t.Errorf("checkTimeProfile:%v returned incorrect access - expected: %v, got:%v", test.ID, expected, ok)
		}
	}
}

func offset(datetime string) entities.Offset {
	utc, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.UTC)
	now := time.Now().UTC()
	delta := utc.Sub(now)

	return entities.Offset(delta)
}
