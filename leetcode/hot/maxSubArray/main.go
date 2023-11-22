package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := 0

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}

	return maxSum
}

// 递归实现
// 核心思想是以 i 位置结尾的f(i)表示以i位置结尾的最大子数组的和的最大值。
//func maxSubArray(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	return process(nums)
//}
//
//func process(nums []int, i int) int {
//	if i == 0 {
//		return nums[0]
//	}
//	return max(process(nums, i-1), nums[i])
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// 动态规划 + 重复利用数组
//func maxSubArray2(nums []int) int {
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i]+nums[i-1] > nums[i] {
//			nums[i] += nums[i-1]
//		}
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	return max
//}

//func maxSubArray2(nums []int) int {
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i]+nums[i-1] > nums[i] {
//			nums[i] = nums[i] + nums[i-1]
//		}
//		if max < nums[i] {
//			max = nums[i]
//		}
//	}
//	return max
//}

// 分治
