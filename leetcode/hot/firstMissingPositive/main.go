package main

import "fmt"

func main() {
	nums := []int{3, 4, -1, 1}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}

func firstMissingPositive(nums []int) int {
	m := make(map[int]struct{})
	maxVal := 1
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	for i := 1; i <= maxVal; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return maxVal + 1
}
