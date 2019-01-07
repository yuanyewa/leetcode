package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[l] != nums[i] {
			l++
			nums[l] = nums[i]
		}
	}
	return l + 1
}

func main() {
	nums := []int{1, 3, 3, 3, 4}
	x := removeDuplicates(nums)
	fmt.Println(x, nums)
}
