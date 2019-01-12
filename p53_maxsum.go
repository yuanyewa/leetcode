package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m, c := nums[0], nums[0]
	for _, v := range nums[1:] {
		if c > 0 {
			c += v
		} else {
			c = v
		}
		if m < c {
			m = c
		}
	}
	return m
}

func main() {
	fmt.Println(maxSubArray([]int{-2, -1}))

}
