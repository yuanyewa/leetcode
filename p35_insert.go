package main

import (
	"fmt"
	"time"
)

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}

func searchInsert2(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return right + 1
}

func main() {
	nums := []int{1, 3, 5, 6}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			searchInsert(nums, 4)
		}
		fmt.Println(time.Since(start), 1)
	}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			searchInsert2(nums, 4)
		}
		fmt.Println(time.Since(start), 1)
	}
}
