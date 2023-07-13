package weeks

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestISO8601(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
	}
	LoadTestData(t, "iso8601_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := ISO8601.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := ISO8601.WeekEndExclusive(tt.Year, tt.Week)
			wantEnd := start.AddDate(0, 0, 7)
			if !end.Equal(wantEnd) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, end, wantEnd)
			}
		})
	}
}

func TestMYSQLMode0(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	LoadTestData(t, "mysqlmode0_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := MySQLMode0.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := MySQLMode0.WeekEndExclusive(tt.Year, tt.Week)
			if !end.Equal(tt.End) {
				t.Errorf("WeekEndExclusive(%d, %d) = %v, want %v", tt.Year, tt.Week, end, tt.End)
			}
		})
	}
}

func TestMYSQLMode1(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	LoadTestData(t, "mysqlmode1_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := MySQLMode1.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := MySQLMode1.WeekEndExclusive(tt.Year, tt.Week)
			if !end.Equal(tt.End) {
				t.Errorf("WeekEndExclusive(%d, %d) = %v, want %v", tt.Year, tt.Week, end, tt.End)
			}
		})
	}
}

func TestMYSQLMode2(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	LoadTestData(t, "mysqlmode2_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := MySQLMode2.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := MySQLMode2.WeekEndExclusive(tt.Year, tt.Week)
			if !end.Equal(tt.End) {
				t.Errorf("WeekEndExclusive(%d, %d) = %v, want %v", tt.Year, tt.Week, end, tt.End)
			}
		})
	}
}

func TestMYSQLMode3(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	LoadTestData(t, "mysqlmode3_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := MySQLMode3.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := MySQLMode3.WeekEndExclusive(tt.Year, tt.Week)
			if !end.Equal(tt.End) {
				t.Errorf("WeekEndExclusive(%d, %d) = %v, want %v", tt.Year, tt.Week, end, tt.End)
			}
		})
	}
}

func TestMYSQLMode4(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	LoadTestData(t, "mysqlmode4_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := MySQLMode4.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := MySQLMode4.WeekEndExclusive(tt.Year, tt.Week)
			if !end.Equal(tt.End) {
				t.Errorf("WeekEndExclusive(%d, %d) = %v, want %v", tt.Year, tt.Week, end, tt.End)
			}
		})
	}
}

func TestMYSQLMode5(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	LoadTestData(t, "mysqlmode5_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := MySQLMode5.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := MySQLMode5.WeekEndExclusive(tt.Year, tt.Week)
			if !end.Equal(tt.End) {
				t.Errorf("WeekEndExclusive(%d, %d) = %v, want %v", tt.Year, tt.Week, end, tt.End)
			}
		})
	}
}

func TestMYSQLMode6(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	LoadTestData(t, "mysqlmode6_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := MySQLMode6.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := MySQLMode6.WeekEndExclusive(tt.Year, tt.Week)
			if !end.Equal(tt.End) {
				t.Errorf("WeekEndExclusive(%d, %d) = %v, want %v", tt.Year, tt.Week, end, tt.End)
			}
		})
	}
}

func TestMYSQLMode7(t *testing.T) {
	// Load the test data.
	var tests []struct {
		Year  int       `json:"year"`
		Week  int       `json:"week"`
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}
	LoadTestData(t, "mysqlmode7_weeks.json", &tests)

	// Run the tests
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%vW%v", tt.Year, tt.Week), func(t *testing.T) {
			start := MySQLMode7.WeekStart(tt.Year, tt.Week)
			if !start.Equal(tt.Start) {
				t.Errorf("WeekStart(%d, %d) = %v, want %v", tt.Year, tt.Week, start, tt.Start)
			}

			end := MySQLMode7.WeekEndExclusive(tt.Year, tt.Week)
			if !end.Equal(tt.End) {
				t.Errorf("WeekEndExclusive(%d, %d) = %v, want %v", tt.Year, tt.Week, end, tt.End)
			}
		})
	}
}

// LoadTestData loads a JSON file from the testdata directory into dest. The value of dest
// must be compatible with json.Unmarshal.
func LoadTestData(t *testing.T, file string, dest any) {
	t.Helper()

	f, err := os.Open(path.Join("testdata", file))
	if err != nil {
		t.Fatalf("failed to open testdata: %v", err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(dest); err != nil {
		t.Fatalf("failed to decode testdata: %v", err)
	}
}
