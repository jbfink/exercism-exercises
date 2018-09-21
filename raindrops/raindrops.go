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
// in a possible loop:
// First test if array does NOT contain 3, 5, 7. If so, *return* input.
// Then go through each array element. If array element contains 3, 5, 7, append to new string the appropriate raindrop
// at the end, return the raindrop string.

import "fmt"

func HelloRaindrops() {
	fmt.Println("Hello Raindrops")
}
