package hamming

//import "errors"
import "strings"

//import "fmt"

func Distance(a, b string) (int, error) {
	a1 := strings.Split(a, "")
	b1 := strings.Split(b, "")

	if len(a1) != len(b1) {
		return -1, nil
	}
	return 0, nil
}
