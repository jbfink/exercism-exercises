// the twofer package contains the ShareWith function.
package twofer

// The ShareWith function takes one argument, the string variable name, and runs an if-else branch on it, returning "One for Alice,
// one for me." if name == "Alice", "One for Bob, one for me." if name == "Bob", and "One for you, one for me." if name is anything else.

func ShareWith(name string) string {
	if name == "Alice" {
		return "One for Alice, one for me."
	} else if name == "Bob" {
		return "One for Bob, one for me."
	} else {
		return "One for you, one for me."

	}
}
