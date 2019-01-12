package main

import (
	"fmt"
	"time"
)

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
func removeElement2(nums []int, val int) int {
	i := 0
	j := len(nums)
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
		} else {
			i++
		}
	}
	return j
}
func main() {
	nums := []int{3, 2, 2, 2, 2, 2, 2, 4, 5, 6, 7}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			removeElement(nums, 3)
		}
		fmt.Println(time.Since(start))
	}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			removeElement2(nums, 3)
		}
		fmt.Println(time.Since(start))
	}
}
