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
	strand1 := strings.Split(a, "")
	strand2 := strings.Split(b, "")

	if len(strand1) != len(strand2) {
		return -1, errors.New("strings are not equal in length")
	} else if a == b {
		return 0, nil
	}

	distance := 0

	for i := 0; i < len(strand1); i++ {
		{
			if strand1[i] == strand2[i] {
				continue
			}
			if strand1[i] != strand2[i] {
				distance++
			}
		}

	}
	return distance, nil
}
