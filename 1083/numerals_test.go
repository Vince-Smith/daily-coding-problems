package numerals

import "testing"

func TestParseNumeral(t *testing.T) {
	var scenarios = []struct {
		in string
		out int64
	}{
		{"XIV", 14},
		{"M", 1000},
		{"D", 500},
		{"C", 100},
		{"L", 50},
		{"X", 10},
		{"V", 5},
		{"I", 1},
		{"MDC", 1600},
		{"MCD", 1400},
		{"MD", 1500},
	}

	for _, s := range scenarios {
		result := ParseNumeral(s.in)
		if result != s.out {
			t.Errorf("Expected: %d Got: %d", s.out, result)
		}
	}
}
