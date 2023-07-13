// Package weeks provides utilities for working with weeks according to various week
// definitions.
package weeks

import (
	"fmt"
	"time"
)

// Definition holds the properties that describe a definition of a week.
type Definition struct {
	// FirstDayOfWeek controls which day is considered the first day of the week.
	FirstDayOfWeek time.Weekday

	// Numbering controls how weeks are numbered.
	Numbering Numbering

	// FirstFullWeekStrategy is the strategy for finding the first full week of a year.
	// All other weeks are calculated relative to this week.
	FirstFullWeekStrategy FirstFullWeekStrategy
}

// WeekStart returns the first moment of the week for the given year and week number.
func (d Definition) WeekStart(year, week int) time.Time {
	firstOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	firstWeekOfYear, hasPartialWeek := d.FirstFullWeekStrategy.FirstFullWeekForYear(year)

	switch d.Numbering {

	case StartAtZero:
		return firstWeekOfYear.AddDate(0, 0, week*7)

	case StartAtOne:
		return firstWeekOfYear.AddDate(0, 0, (week-1)*7)

	case PartialWeekZero:
		switch {
		// Week zero always starts on the first day of the year.
		case week == 0:
			return firstOfYear

		// If this is week 1 and there is no week 0, so it always starts on the
		// first day of the year.
		case week == 1 && !hasPartialWeek:
			return firstOfYear

		default:
			return firstWeekOfYear.AddDate(0, 0, (week-1)*7)
		}
	default:
		panic(fmt.Errorf("weeks: unkown numbering schema %v", d.Numbering))
	}
}

// WeekEndExclusive returns the first moment after the week for the given year and week
// number.
func (d Definition) WeekEndExclusive(year, week int) time.Time {
	switch {
	case d.Numbering == PartialWeekZero && week == 0:
		firstWeekOfYear, _ := d.FirstFullWeekStrategy.FirstFullWeekForYear(year)
		return firstWeekOfYear
	case d.Numbering == PartialWeekZero:
		firstWeekOfYear, _ := d.FirstFullWeekStrategy.FirstFullWeekForYear(year)
		end := firstWeekOfYear.AddDate(0, 0, week*7)

		if end.Year() != year {
			return time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC)
		}

		return firstWeekOfYear.AddDate(0, 0, week*7)
	}

	return d.WeekStart(year, week).AddDate(0, 0, 7)
}

// WeekEndInclusive returns the Last moment of the week for the given year and week.
func (d Definition) WeekEndInclusive(year, week int, precision time.Duration) time.Time {
	return d.WeekEndExclusive(year, week).Add(-1 * precision)
}

type Numbering int

const (
	// StartAtZero is a numbering scheme where the first week of the year is week zero.
	// StartAtZero does not handle partial weeks.
	StartAtZero Numbering = 0

	// StartAtOne is a numbering scheme where the first week of the year is week one.
	// StartAtOne does not handle partial weeks.
	StartAtOne Numbering = 1

	// PartialWeekZero is a numbering scheme where the first week of the year is week one.
	// If a year starts with a partial week, that week is numbered zero.
	PartialWeekZero Numbering = 2
)

// FirstFullWeekStrategy is a strategy for finding the first full week of a year.
type FirstFullWeekStrategy interface {

	// FirstFullWeekForYear should return the first moment of the first full week of a
	// year. When a year starts with a partial week hasPartialWeek should be returned as
	// true.
	FirstFullWeekForYear(year int) (start time.Time, hasPartialWeek bool)
}

type FirstWeekFunc func(year int) (start time.Time, hasPartialWeek bool)

func (f FirstWeekFunc) FirstFullWeekForYear(year int) (start time.Time, hasPartialWeek bool) {
	return f(year)
}

// WeekContainingFirstOccurrenceOfWeekday returns a FirstFullWeekStrategy where the first week
// of the year is considered to be the first week that contains the given weekday.
func WeekContainingFirstOccurrenceOfWeekday(firstDayOfWeek, weekContaining time.Weekday) FirstFullWeekStrategy {

	correction := weekdayPrecedingWeekdayOffset(firstDayOfWeek, weekContaining)

	return FirstWeekFunc(func(year int) (time.Time, bool) {
		// Find the number of days between the first day of the year and the first
		// occurrence of `weekContaining`.
		offset := firstOccurrenceOfWeekdayOffset(year, weekContaining)

		hasPartialWeek := offset > 0

		// Now correct the offset to find `firstDayOfWeek` preceding the `weekContaining`.
		offset = offset + correction

		start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, offset)

		return start, hasPartialWeek
	})
}

func FirstWeekWithNDays(firstDayOfWeek time.Weekday, n int) FirstFullWeekStrategy {
	return FirstWeekFunc(func(year int) (time.Time, bool) {

		// Find how many days until the start of the first 7-day week.
		daysUntilWeekStart := firstOccurrenceOfWeekdayOffset(year, firstDayOfWeek)

		// If the first day of the year is the first day of the week, then we're done.
		if daysUntilWeekStart == 0 {
			return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC), false
		}

		hasPartialWeek := true
		offset := daysUntilWeekStart

		if daysUntilWeekStart >= n {
			hasPartialWeek = false
			offset = offset - 7
		}

		start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, offset)

		return start, hasPartialWeek
	})
}

func firstOccurrenceOfWeekdayOffset(year int, weekday time.Weekday) int {
	//  firstOfYear  | Sunday Monday Tuesday Wednesday Thursday Friday Saturday
	//  Sunday       | 0      1      2       3         4        5      6
	//  Monday       | 6      0      1       2         3        4      5
	//  Tuesday      | 5      6      0       1         2        3      4
	//  Wednesday    | 4      5      6       0         1        2      3
	//  Thursday     | 3      4      5       6         0        1      2
	//  Friday       | 2      3      4       5         6        0      1
	//  Saturday     | 1      2      3       4         5        6      0

	firstOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	offset := int(weekday) - int(firstOfYear.Weekday())
	if offset < 0 {
		offset = 7 + offset
	}

	return offset
}

func weekdayPrecedingWeekdayOffset(preceding, start time.Weekday) int {
	// Calculate the correction between `weekContaining` and the preceding `firstDayOfWeek`.
	//
	//                 | preceding
	//  start          | Sunday Monday Tuesday Wednesday Thursday Friday Saturday
	//  Sunday         | 0      -6     -5      -4        -3       -2    -1
	//  Monday         | ?      0      ?       ?         ?        ?      ?
	//  Tuesday        | ?      ?      0       ?         ?        ?      ?
	//  Wednesday      | ?      ?      ?       0         ?        ?      ?
	//  Thursday       | -4     -3     -2      -1        0        -6     -5
	//  Friday         | ?      ?      ?       ?         ?        0      ?
	//  Saturday       | ?      ?      ?       ?         ?        ?      0

	// preceding       start
	// 0 (sunday)    - 4 (Thursday) = -4
	// 1 (monday)    - 4 (Thursday) = -3
	// 2 (tuesday)   - 4 (Thursday) = -2
	// 3 (wednesday) - 4 (Thursday) = -1
	// 4 (thursday)  - 4 (Thursday) = 0
	// 5 (friday)    - 4 (Thursday) = 1
	// 6 (saturday)  - 4 (Thursday) = 2

	offset := int(preceding) - int(start)
	if offset > 0 {
		offset -= 7
	}

	return offset
}
