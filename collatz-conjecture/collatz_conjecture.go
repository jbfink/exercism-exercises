package collatzconjecture

import "fmt"

// from readme:
// Take any positive integer n. If n is even, divide n by 2 to get n / 2. If n is
// odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely.
// The conjecture states that no matter which number you start with, you will
// always reach 1 eventually.

// so:
// take n. Initial test for positive and test for integer. Test for evenness by modulo, like, divide by 2 == modulo 0? then even. Otherwise odd.
// Start loop with modulo test:
// Is n 1? then exit.
// path 1: even, divide by 2, return to 1 test
// path 2: odd, (n*3)+1. Return to 1 test.
func CollatzConjecture(a) int {
	fmt.Println("programming!")
}
