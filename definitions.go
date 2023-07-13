package weeks

import (
	"time"
)

// ISO8601 is a week definition that follows the ISO 8601 "week date" standard.See
// https://en.wikipedia.org/wiki/ISO_week_date for more information.
var ISO8601 = Definition{
	FirstDayOfWeek:        time.Monday,
	Numbering:             StartAtOne,
	FirstFullWeekStrategy: WeekContainingFirstOccurrenceOfWeekday(time.Monday, time.Thursday),
}

// MySQL has 8 different week definitions.
//
//	Mode  First day of week  Range  Week 1 is the first week
//	0     Sunday             0-53   with a Sunday in this year
//	1     Monday             0-53   with 4 or more days this year
//	2     Sunday             1-53   with a Sunday in this year
//	3     Monday             1-53   with 4 or more days this year
//	4     Sunday             0-53   with 4 or more days this year
//	5     Monday             0-53   with a Monday in this year
//	6     Sunday             1-53   with 4 or more days this year
//	7     Monday             1-53   with a Monday in this year
//
// See https://dev.mysql.com/doc/refman/8.0/en/date-and-time-functions.html#function_week
var (

	// MySQLMode0 is a week definition that conform to MySQL mode 0.
	//
	//	Mode  First day of week  Range  Week 1 is the first week
	//	0     Sunday             0-53   with a Sunday in this year
	MySQLMode0 = Definition{
		FirstDayOfWeek:        time.Sunday,
		Numbering:             PartialWeekZero,
		FirstFullWeekStrategy: WeekContainingFirstOccurrenceOfWeekday(time.Sunday, time.Sunday),
	}

	// MySQLMode1 is a week definition that conform to MySQL mode 1.
	//
	//	Mode  First day of week  Range  Week 1 is the first week
	//	1     Monday             0-53   with 4 or more days this year
	MySQLMode1 = Definition{
		FirstDayOfWeek:        time.Monday,
		Numbering:             PartialWeekZero,
		FirstFullWeekStrategy: FirstWeekWithNDays(time.Monday, 4),
	}

	// MySQLMode2 is a week definition that conform to MySQL mode 2.
	//
	//	Mode  First day of week  Range  Week 1 is the first week
	//	2     Sunday             1-53   with a Sunday in this year
	MySQLMode2 = Definition{
		FirstDayOfWeek:        time.Sunday,
		Numbering:             StartAtOne,
		FirstFullWeekStrategy: WeekContainingFirstOccurrenceOfWeekday(time.Sunday, time.Sunday),
	}

	// MySQLMode3 is a week definition that conform to MySQL mode 3.
	//
	//	Mode  First day of week  Range  Week 1 is the first week
	//	3     Monday             1-53   with 4 or more days this year
	MySQLMode3 = Definition{
		FirstDayOfWeek:        time.Monday,
		Numbering:             StartAtOne,
		FirstFullWeekStrategy: FirstWeekWithNDays(time.Monday, 4),
	}

	// MySQLMode4 is a week definition that conform to MySQL mode 4.
	//
	//	Mode  First day of week  Range  Week 1 is the first week
	//	4     Sunday             0-53   with 4 or more days this year
	MySQLMode4 = Definition{
		FirstDayOfWeek:        time.Sunday,
		Numbering:             PartialWeekZero,
		FirstFullWeekStrategy: FirstWeekWithNDays(time.Sunday, 4),
	}

	// MySQLMode5 is a week definition that conform to MySQL mode 5.
	//
	//	Mode  First day of week  Range  Week 1 is the first week
	//	5     Monday             0-53   with a Monday in this year
	MySQLMode5 = Definition{
		FirstDayOfWeek:        time.Monday,
		Numbering:             PartialWeekZero,
		FirstFullWeekStrategy: WeekContainingFirstOccurrenceOfWeekday(time.Monday, time.Monday),
	}

	// MySQLMode6 is a week definition that conform to MySQL mode 6.
	//
	//	Mode  First day of week  Range  Week 1 is the first week
	//	6     Sunday             1-53   with 4 or more days this year
	MySQLMode6 = Definition{
		FirstDayOfWeek:        time.Sunday,
		Numbering:             StartAtOne,
		FirstFullWeekStrategy: FirstWeekWithNDays(time.Sunday, 4),
	}

	// MySQLMode7 is a week definition that conform to MySQL mode 7.
	//
	//	Mode  First day of week  Range  Week 1 is the first week
	//	7     Monday             1-53   with a Monday in this year
	MySQLMode7 = Definition{
		FirstDayOfWeek:        time.Monday,
		Numbering:             StartAtOne,
		FirstFullWeekStrategy: WeekContainingFirstOccurrenceOfWeekday(time.Monday, time.Monday),
	}
)
