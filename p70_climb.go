package main

import "fmt"

func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

var m = make(map[int]int)

func climbStairs(n int) int {
	if v, ok := m[n]; ok {
		return v
	}
	if n <= 2 {
		m[n] = n
		return n
	}
	v := climbStairs(n-1) + climbStairs(n-2)
	m[n] = v
	return v
}

func main() {
	fmt.Println(climbStairs2(5))
}
