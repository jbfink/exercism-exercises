package isogram

import "strings"

// adrplan
// Take string. Split string. ????
// Maybe a map? https://stackoverflow.com/questions/20111170/how-to-use-golang-to-check-repeat-element-in-array
// https://blog.golang.org/go-maps-in-action
// map might be overthinking. What I need is a function to test an array/slice for unique elements and return a bool. Probably.

// IsIsogram takes a string input and determines if the string is an isogram.
func IsIsogram(s string) bool {
	// https://www.dotnetperls.com/duplicates-go ??? maybe
	letters := strings.Split(s, "")
	encountered := map[string]bool{}

	for v := range letters {
		if encountered[letters[v]] == true {
			bool := false
		}
		{
			bool := true

		}

	}
	return bool
}
