# Go Weeks

```go
import "github.com/ianlopshire/go-weeks"
```

Package weeks provides utilities for working with weeks according to various week definitions.

```go
// Get the interval [start, end) for the first week of 2020 according to ISO8601
start := weeks.ISO8601.WeekStart(2018, 1)
end := weeks.ISO8601.WeekEndExclusive(2018, 1)

fmt.Printf("2018W1 (ISO 8601) = [%v, %v)\n", start, end)
// Output: 2018W1 (ISO 8601) = [2018-01-01 00:00:00 +0000 UTC, 2018-01-08 00:00:00 +0000 UTC)
```
```go
// Get the interval [start, end] for the first week of 2020 according to ISO8601
start = weeks.ISO8601.WeekStart(2018, 1)
end = weeks.ISO8601.WeekEndInclusive(2018, 1, time.Second)

fmt.Printf("2018W1 (ISO 8601) = [%v, %v]\n", start, end)
// Output: 2018W1 (ISO 8601) = [2018-01-01 00:00:00 +0000 UTC, 2018-01-07 23:59:59 +0000 UTC]
```

### Builtin Week Definitions

The following week definitions are provided by this package:
- ISO 8601 (`weeks.ISO8601`)
- MySQL mode 0 (`weeks.MySQLMode0`)
- MySQL mode 1 (`weeks.MySQLMode1`)
- MySQL mode 2 (`weeks.MySQLMode2`)
- MySQL mode 3 (`weeks.MySQLMode3`)
- MySQL mode 4 (`weeks.MySQLMode4`)
- MySQL mode 5 (`weeks.MySQLMode5`)
- MySQL mode 6 (`weeks.MySQLMode6`)
- MySQL mode 7 (`weeks.MySQLMode7`)