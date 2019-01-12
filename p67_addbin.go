package main

import "fmt"

func addBinary(a string, b string) string {
	c := ""
	m, n := len(a), len(b)
	var x, y int
	carry := 0
	for i := 0; ; i++ {
		if i >= m {
			x = 0
		} else {
			if a[m-i-1] == '1' {
				x = 1
			} else {
				x = 0
			}
		}
		if i >= n {
			y = 0
		} else {
			if b[n-i-1] == '1' {
				y = 1
			} else {
				y = 0
			}
		}
		if i >= m && i >= n && carry == 0 {
			break
		}
		z := x + y + carry
		if z >= 2 {
			z -= 2
			carry = 1
		} else {
			carry = 0
		}
		if z == 0 {
			c = "0" + c
		} else {
			c = "1" + c
		}
	}
	return c
}

func main() {
	fmt.Println(addBinary("1011", "1")) // expect 1100
}
