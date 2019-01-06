package main

import (
	"fmt"
	"time"
)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, k := range nums {
		need := target - k
		if _, ok := seen[need]; !ok {
			seen[k] = i
		} else {
			return []int{seen[need], i}
		}
	}
	return nil
}
func twoSum2(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, k := range nums {
		if v, ok := seen[target-k]; ok {
			return []int{v, i}
		}
		seen[k] = i
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	for i, k := range nums {
		for m, n := range nums {
			if m == i {
				continue
			}
			if k+n == target {
				return []int{i, m}
			}
		}
	}
	return nil
}

func main() {
	{
		fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			twoSum([]int{2, 7, 11, 15}, 9)
		}
		fmt.Println(time.Since(start))
	}
	{
		fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			twoSum2([]int{2, 7, 11, 15}, 9)
		}
		fmt.Println(time.Since(start))
	}
}
