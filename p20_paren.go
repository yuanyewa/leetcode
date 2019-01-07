package main

import "fmt"

func isValid(s string) bool {
	m := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	buf := []rune{}
	for _, r := range s {
		if _, ok := m[r]; ok {
			buf = append(buf, r)
		} else if len(buf) == 0 || r != m[buf[len(buf)-1]] {
			return false
		} else {
			buf = buf[:len(buf)-1]
		}
	}
	return len(buf) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("((])"))
}
