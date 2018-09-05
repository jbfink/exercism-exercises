// the twofer package contains the ShareWith function.
package twofer

// The ShareWith function takes one argument, the string variable name, and runs an if-else branch on it, returning "One for name, one for me."
// if name is blank, returns "One for you, one for me."

func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."

	}

}
