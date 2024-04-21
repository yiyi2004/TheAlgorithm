package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

// 记忆化搜索
//func combinationSum4(nums []int, target int) int {
//	dp := make([]int, target+1)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = -1
//	}
//	return dfs(0, nums, target, dp)
//}
//
//func dfs(sum int, nums []int, target int, dp []int) int {
//	if sum >= target {
//		if sum == target {
//			return 1
//		}
//		return 0
//	}
//	if dp[sum] != -1 {
//		return dp[sum]
//	}
//	res := 0
//	for i := 0; i < len(nums); i++ {
//		res += dfs(sum+nums[i], nums, target, dp)
//	}
//	dp[sum] = res
//	return res
//}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	// dp[i] = sum(dp[i-nums[j]])
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
