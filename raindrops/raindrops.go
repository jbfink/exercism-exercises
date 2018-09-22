package raindrops

// from README.md
/*  # Raindrops

Convert a number to a string, the contents of which depend on the number's factors.

- If the number has 3 as a factor, output 'Pling'.
- If the number has 5 as a factor, output 'Plang'.
- If the number has 7 as a factor, output 'Plong'.
- If the number does not have 3, 5, or 7 as a factor,
  just pass the number's digits straight through.

## Examples

- 28's factors are 1, 2, 4, **7**, 14, 28.
 - In raindrop-speak, this would be a simple "Plong".
- 30's factors are 1, 2, **3**, **5**, 6, 10, 15, 30.
 - In raindrop-speak, this would be a "PlingPlang".
- 34 has four factors: 1, 2, 17, and 34.
 - In raindrop-speak, this would be "34".
*/

// adrplan:
// First, test that the input is an integer. Possibly a positive integer.
// Next, figure out how to get a factors of an integer. (look this up, it's probably in a math package))
// Convert factors to array
// No test will contain more than one pling or plang or plong, but could contain a combo of each or nothing at all.
// First test if array does NOT contain 3, 5, 7. If so, *return* input.
// in a loop:
// go through each array element. If array element contains 3, 5, 7, append to new string the appropriate raindrop
// at the end, return the raindrop string.

// note: if you HAVE to write a fuckin' thing to get factors jesus christ something like this:
// Take integer i
// Generate array of all numbers between 1 to i
// loop:
// Divide i by array[n]
//	Is it modulo 0? Stick n into new array, factors[].
//	Is it not modulo 0? Continue.
//	Return factors[].
// OR:::
// You don't need to get the full range of factors -- you only need to test if there is 3, 5, 7. So instead:
// loop: Does 3 modulo 0? then insert "Pling" into raindrop string.
//	 Does 5 modulo 0? then insert "Plang" into raindrop string.
// 	 Does 7 modulo 0? then insert "Plong" into raindrop string.
//	 Are there no raindrops? Return i.

import "fmt"

// Raindrops function takes a single input, an integer. If the integer's factors contain a 3, it inserts "Pling" into a string.
// If an integer's factors contain a 5, it inserts "Plang" into a string. If the integer's factors contain a 7, it
// inserts "Plong" into a string. If the integer contains none of those, it returns the integer itself.
func Raindrops(i int) {
	// building strings might be hard. see: https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
	rain := ""
	if i%3 == 0 {
		fmt.Println("Pling")
	}
	if i%5 == 0 {
		fmt.Println("Plang")
	}
	if i%7 == 0 {
		fmt.Println("Plong")
	}

}
