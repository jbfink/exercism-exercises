package hamming

import "errors"
import "strings"

// Distance function calculates the Hamming distance of two strings.
// If the two strings are unequal in length, returns error.
// If the two strings are identical, returns Hamming distance of zero.
// If the two strings are of identical length, but differ in composition,
// then a loop happens. When the two strings have different letters at the
// same position, a counter is incremented. At the end of the loop,
// return the contents of that counter.
func Distance(a, b string) (int, error) {
	a1 := strings.Split(a, "")
	b1 := strings.Split(b, "")

	if len(a1) != len(b1) {
		return -1, errors.New("strings are not equal")
	} else if a == b {
		return 0, nil
	}

	distance := 0

	for i := 0; i < len(a1); i++ {
		{
			if a1[i] == b1[i] {
				continue
			}
			if a1[i] != b1[i] {
				distance++
			}
		}

	}
	return distance, nil
}
