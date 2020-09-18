package main

import (
	"fmt"

	"github.com/manoj-gupta/toolkit/algo"
)

func main() {
	text := "bacbabababcaabababca"
	pattern := "abababca"

	result := algo.KmpMatcher(text, pattern)
	fmt.Println(result)
}
