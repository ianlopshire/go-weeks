package weeks

import (
	"fmt"
	"testing"
	"time"
)

func Test_firstOccurrenceOfWeekdayOffset(t *testing.T) {
	for _, tt := range []struct {
		years   []int
		offsets map[time.Weekday]int
	}{
		{
			years: []int{1995, 2006, 2012}, // Sunday
			offsets: map[time.Weekday]int{
				time.Sunday:    0,
				time.Monday:    1,
				time.Tuesday:   2,
				time.Wednesday: 3,
				time.Thursday:  4,
				time.Friday:    5,
				time.Saturday:  6,
			},
		},
		{
			years: []int{1990, 1996, 2001}, // Monday
			offsets: map[time.Weekday]int{
				time.Sunday:    6,
				time.Monday:    0,
				time.Tuesday:   1,
				time.Wednesday: 2,
				time.Thursday:  3,
				time.Friday:    4,
				time.Saturday:  5,
			},
		},
		{
			years: []int{1991, 2002, 2008}, // Tuesday
			offsets: map[time.Weekday]int{
				time.Sunday:    5,
				time.Monday:    6,
				time.Tuesday:   0,
				time.Wednesday: 1,
				time.Thursday:  2,
				time.Friday:    3,
				time.Saturday:  4,
			},
		},
		{
			years: []int{1992, 1997, 2003}, // Wednesday
			offsets: map[time.Weekday]int{
				time.Sunday:    4,
				time.Monday:    5,
				time.Tuesday:   6,
				time.Wednesday: 0,
				time.Thursday:  1,
				time.Friday:    2,
				time.Saturday:  3,
			},
		},
		{
			years: []int{1998, 2004, 2009}, // Thursday
			offsets: map[time.Weekday]int{
				time.Sunday:    3,
				time.Monday:    4,
				time.Tuesday:   5,
				time.Wednesday: 6,
				time.Thursday:  0,
				time.Friday:    1,
				time.Saturday:  2,
			},
		},
		{
			years: []int{1993, 1999, 2010}, // Friday
			offsets: map[time.Weekday]int{
				time.Sunday:    2,
				time.Monday:    3,
				time.Tuesday:   4,
				time.Wednesday: 5,
				time.Thursday:  6,
				time.Friday:    0,
				time.Saturday:  1,
			},
		},
		{
			years: []int{1994, 2000, 2005}, // Saturday
			offsets: map[time.Weekday]int{
				time.Sunday:    1,
				time.Monday:    2,
				time.Tuesday:   3,
				time.Wednesday: 4,
				time.Thursday:  5,
				time.Friday:    6,
				time.Saturday:  0,
			},
		},
	} {
		for _, year := range tt.years {
			// Find the weekday that the year starts with. We only use this to name the test run.
			yearStartsWith := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Weekday()

			t.Run(fmt.Sprintf("%v (%v)", year, yearStartsWith), func(t *testing.T) {
				for weekday, wantOffset := range tt.offsets {
					t.Run(fmt.Sprintf("%v to %v", yearStartsWith, weekday), func(t *testing.T) {
						gotOffset := firstOccurrenceOfWeekdayOffset(year, weekday)
						if gotOffset != wantOffset {
							t.Errorf("firstOccurrenceOfWeekdayOffset(%v, %v) = %v, want %v", year, weekday, gotOffset, wantOffset)
						}
					})
				}
			})
		}
	}
}

func Test_weekdayPrecedingWeekdayOffset(t *testing.T) {
	for _, tt := range []struct {
		start   time.Weekday
		offsets map[time.Weekday]int
	}{
		{
			start: time.Sunday,
			offsets: map[time.Weekday]int{
				time.Sunday:    0,
				time.Monday:    -6,
				time.Tuesday:   -5,
				time.Wednesday: -4,
				time.Thursday:  -3,
				time.Friday:    -2,
				time.Saturday:  -1,
			},
		},
		{
			start: time.Monday,
			offsets: map[time.Weekday]int{
				time.Sunday:    -1,
				time.Monday:    0,
				time.Tuesday:   -6,
				time.Wednesday: -5,
				time.Thursday:  -4,
				time.Friday:    -3,
				time.Saturday:  -2,
			},
		},
		{
			start: time.Thursday,
			offsets: map[time.Weekday]int{
				time.Sunday:    -4,
				time.Monday:    -3,
				time.Tuesday:   -2,
				time.Wednesday: -1,
				time.Thursday:  0,
				time.Friday:    -6,
				time.Saturday:  -5,
			},
		},
	} {

		for weekday, wantOffset := range tt.offsets {
			t.Run(fmt.Sprintf("%v preceding %v", weekday, tt.start), func(t *testing.T) {
				gotOffset := weekdayPrecedingWeekdayOffset(weekday, tt.start)
				if gotOffset != wantOffset {
					t.Errorf("weekdayPrecedingWeekdayOffset(%v, %v) = %v, want %v", weekday, tt.start, gotOffset, wantOffset)
				}
			})
		}

	}
}
