package main

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		m := l + (r-l)/2
		if m*m == x {
			return m
		} else if m*m > x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r
}

func main() {
	fmt.Println(mysqrt(10))
}
