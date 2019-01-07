package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}
	y := 0
	r := 0
	for x >= 10 {
		x, r = x/10, x%10
		y = 10*y + r
	}
	y = 10*y + x
	if y > math.MaxInt32 {
		return 0
	}
	if isNeg {
		return -y
	}
	return y
}

func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(math.MaxInt32)
}
