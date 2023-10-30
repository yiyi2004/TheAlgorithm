package main

// 原地交换的写法
func findRepeatNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		} else {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}
