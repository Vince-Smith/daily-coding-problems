package numerals

var alphabet = map[string]int64 {
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

var edge = map[string]int64 {
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func ParseNumeral(s string) int64 {
	sum := int64(0)
	l := len(s)
	for i := 0; i < l; i++ {
		// Check to see if we're looking at an edge case
		if l > i+1 {
			seg := s[i:i+2]
			val, found := edge[seg]
			if found {
				sum += val
				i++
				continue
			}	
		}
		sum += alphabet[string(s[i])]
	}

	return sum
}
