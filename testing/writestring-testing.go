package main

import (
	"fmt"
	"strings"
)

//i := 3

func Convert(i int) string {
	var rain strings.Builder
	//rain.WriteString("")
	if i%3 == 0 {
		rain.WriteString("Pling ")
	}
	if i%5 == 0 {
		rain.WriteString("Plang ")
	}
	if i%7 == 0 {
		rain.WriteString("Plong")
	}
	if len(rain.String()) == 0 {
		return "poop"
	}
	return rain.String()
}
