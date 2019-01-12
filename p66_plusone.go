package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 1
	i := len(digits)
	for ; i > 0; i-- {
		c := digits[i-1]
		if c+carry == 10 {
			digits[i-1] = 0
			carry = 1
		} else {
			digits[i-1] = c + carry
			carry = 0
			break
		}
	}
	if i == 0 && carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 0, 9}))

}
