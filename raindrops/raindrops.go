package raindrops

import "strings"
import "strconv"

// Convert function takes a single input, an integer. If the integer's factors contain a 3, it inserts "Pling" into a string.
// If an integer's factors contain a 5, it inserts "Plang" into a string. If the integer's factors contain a 7, it
// inserts "Plong" into a string. If the integer contains none of those, it returns the integer itself.
func Convert(i int) string {
	var rain strings.Builder
	if i%3 == 0 {
		rain.WriteString("Pling")
	}
	if i%5 == 0 {
		rain.WriteString("Plang")
	}
	if i%7 == 0 {
		rain.WriteString("Plong")
	}
	if rain.String() == "" {
		return strconv.Itoa(i)
	}
	return rain.String()
}
