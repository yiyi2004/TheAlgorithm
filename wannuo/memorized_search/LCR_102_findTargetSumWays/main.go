package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2010)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = -1
		}
	}
	return dfs(nums, 0, 0, target, dp)
}

func dfs(nums []int, index int, sum int, target int, dp [][]int) int {
	if index >= len(nums) {
		if index == len(nums) && sum == target {
			return 1
		}
		return 0
	}
	if dp[index][sum+1000] != -1 {
		return dp[index][sum+1000]
	}
	dp[index][sum+1000] = dfs(nums, index+1, sum+nums[index], target, dp) + dfs(nums, index+1, sum-nums[index], target, dp)
	return dp[index][sum+1000]
}
