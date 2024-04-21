package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	p, q := 0, 1
	for q < len(nums) {
		for q < len(nums) && nums[p] == nums[q] {
			q++
		}
		if q >= len(nums) {
			break
		}
		nums[p+1] = nums[q]
		p++
	}

	return p + 1
}
