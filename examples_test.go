package weeks_test

import (
	"fmt"
	"time"

	"github.com/ianlopshire/go-weeks"
)

func Example_builtin_ISO8601() {
	// Get the interval [start, end) for the first week of 2020 according to ISO8601
	start := weeks.ISO8601.WeekStart(2018, 1)
	end := weeks.ISO8601.WeekEndExclusive(2018, 1)
	fmt.Printf("2018W1 (ISO 8601) = [%v, %v)\n", start, end)

	// Get the interval [start, end] for the first week of 2020 according to ISO8601
	start = weeks.ISO8601.WeekStart(2018, 1)
	end = weeks.ISO8601.WeekEndInclusive(2018, 1, time.Second)
	fmt.Printf("2018W1 (ISO 8601) = [%v, %v]\n", start, end)

	// Output:
	// 2018W1 (ISO 8601) = [2018-01-01 00:00:00 +0000 UTC, 2018-01-08 00:00:00 +0000 UTC)
	// 2018W1 (ISO 8601) = [2018-01-01 00:00:00 +0000 UTC, 2018-01-07 23:59:59 +0000 UTC]
}
