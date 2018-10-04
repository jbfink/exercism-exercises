package scrabble

/* from README.md
## Letter Values

You'll need these:

```text
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
```

## Examples

"cabbage" should be scored as worth 14 points:

- 3 points for C
- 1 point for A, twice
- 3 points for B, twice
- 2 points for G
- 1 point for E

And to total:

- `3 + 2*1 + 2*3 + 2 + 1`
- = `3 + 2 + 6 + 3`
- = `5 + 9`
- = 14

*/

// adrplan:
// first, an key-value store array thing for each letter, a=1, d=2 etc. Maybe in its own function, like func Lookup (s string) int
// Then, split a word into letter chunks. Apply Lookup to each letter, storing the value in an int, adding the next letter value to the int.
// Return the int.

func Score(s string) int {
}
