package main

import "fmt"

func main() {
	nums := []int{7, 8, 9, 11, 12}
	res := firstMissingPositive(nums)

	fmt.Println(res)

}
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	for i := 0; i < len(nums); i++ {
		for i+1 != nums[i] && nums[i] > 0 && nums[i] <= len(nums) {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

//func firstMissingPositive(nums []int) int {
//	m := make(map[int]struct{})
//	maxVal := 1
//	for i := 0; i < len(nums); i++ {
//		m[nums[i]] = struct{}{}
//		if nums[i] > maxVal {
//			maxVal = nums[i]
//		}
//	}
//
//	for i := 1; i <= maxVal; i++ {
//		if _, ok := m[i]; !ok {
//			return i
//		}
//	}
//
//	return maxVal + 1
//}
