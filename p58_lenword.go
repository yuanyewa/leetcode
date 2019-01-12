package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	w := strings.Fields(s)
	if len(w) == 0 {
		return 0
	}
	return len(w[len(w)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
