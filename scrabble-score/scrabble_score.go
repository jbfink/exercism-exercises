package scrabble

import "strings"
import "errors"

// Lookup takes a single character and returns its Scrabble score. It downcases all input, and throws an error if it is passed more than
// one letter.
func Lookup(s string) (int, error) {
	if len(s) != 1 {

		return -1, errors.New("must be one letter only")

	}

	s = strings.ToLower(s)

	switch s {
	case
		"a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
		return 1, nil
	case
		"d", "g":
		return 2, nil
	case
		"b", "c", "m", "p":
		return 3, nil
	case
		"f", "h", "v", "w", "y":
		return 4, nil
	case
		"k":
		return 5, nil
	case
		"j", "x":
		return 8, nil
	case
		"q", "z":
		return 10, nil
	}
	return 0, nil
}

// Strip removes the error message from a given input, unless error is not nil, in which case it throws a panic.
func Strip(i int, err error) int {
	if err != nil {
		panic(err)
	}
	return i
}

// Score returns the Scrabble score for a word.
func Score(word string) int {
	letter := strings.Split(word, "")
	score := 0
	for i := 0; i < len(letter); i++ {
		score += Strip(Lookup(letter[i]))
	}

	return score
}
