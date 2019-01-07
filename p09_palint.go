package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	d := []int{}
	l := 0
	for x > 0 {
		l++
		d = append(d, x%10)
		x /= 10
	}
	fmt.Println(d)
	for i := 0; i < l/2; i++ {
		if d[i] != d[l-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome4(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10
}

func isPalindrome2(x int) bool {

	if x < 0 {
		return false
	}
	for x >= 10 {
		b := int(math.Pow10(int(math.Floor(math.Log10(float64(x))))))
		r, l := x/b, x%10
		fmt.Println(r, l, b)
		if r == l {
			x = (x - r*b) / 10
			fmt.Println(x)
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome4(1000021))
	fmt.Println(isPalindrome4(123321))
}
