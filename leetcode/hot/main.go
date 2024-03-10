package main

import "fmt"

func main() {
	nums := []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97}
	fmt.Println(canPartition(nums))
}

// https://leetcode.cn/problems/partition-equal-subset-sum/?envType=study-plan-v2&envId=top-100-liked
// 这是一个背包问题，我背包问题没有解决啊
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	for i := 0; i < len(nums); i++ {
		if nums[i] > sum {
			return false
		}
		if nums[i] == sum {
			return true
		}
	}

	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j <= sum; j++ {
		}
	}

	return dp[len(nums)][sum]

	//var dfs func(i int, pre int) bool
	//dfs = func(i int, pre int) bool {
	//	if pre == sum {
	//		return true
	//	}
	//	if i == len(nums) {
	//		if pre == sum {
	//			return true
	//		} else {
	//			return false
	//		}
	//	}
	//
	//	return dfs(i+1, pre+nums[i]) || dfs(i+1, pre)
	//}
	//return dfs(0, 0)
}

func canPartition1(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	var dfs func(i int, pre int) bool
	dfs = func(i int, pre int) bool {
		if pre == sum {
			return true
		}
		if i == len(nums) {
			if pre == sum {
				return true
			} else {
				return false
			}
		}

		return dfs(i+1, pre+nums[i]) || dfs(i+1, pre)
	}
	return dfs(0, 0)
}
