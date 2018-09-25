package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("Dude")
	b.WriteString(" ")
	b.WriteString("Sweet")
	fmt.Println(b.String())
}
