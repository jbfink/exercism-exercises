package hamming

import "errors"
import "strings"

//import "fmt"

func Distance(a, b string) (int, error) {
	a1 := strings.Split(a, "")
	b1 := strings.Split(b, "")

	if len(a1) != len(b1) {
		return -1, errors.New("Strings are not equal.")
	} else if a == b {
		return 0, nil
	}

	// Here, a loop of some sort, nested?
	// Like, loop elements from 0 to len(a1) (by this point a1 and b1 will have identical len)
	// for each element:
	//	if a1[x] == b1[x] then add 0 to distance variable
	//      if a1[x] != b1[x] then add 1 to distance variable
	// after len(a1), return variable, nil
	// (why error at all? buh)
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

		//		return distance, nil

	}
	return distance, nil
}
