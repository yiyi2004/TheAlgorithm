package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	res := make([]int, len(nums))
	length := len(nums)
	for i := 0; i < len(nums); i++ {
		res[i] = nums[mod(i-k, length)]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}
}

func mod(a, b int) int {
	if a < 0 {
		for a < 0 {
			a = a + b
		}
		return a
	}
	return a % b
}
