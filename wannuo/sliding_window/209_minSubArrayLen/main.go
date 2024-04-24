package main

import (
	"fmt"
)

func main() {
	res := minSubArrayLen(4, []int{1, 4, 4})
	fmt.Println(res)
}

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	minLen := 10000010
	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if minLen == 10000010 {
		return 0
	}
	return minLen
}
