package main

func main() {

}

// 递归方式实现
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	nums[0], nums[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

// 动态规划解决问题，我不知道你怎么想的
