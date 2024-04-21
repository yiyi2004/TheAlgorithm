package main

import (
	"fmt"
)

func main() {
	res := coinChange([]int{1, 2, 5}, 11)
	fmt.Println(res)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	val := dfs(0, coins, amount, dp)
	if val == 1000010 {
		return -1
	}
	return val
}

func dfs(sum int, coins []int, amount int, dp []int) int {
	if sum >= amount {
		if sum == amount {
			return 0
		}
		return 1000010
	}
	if dp[sum] != -1 {
		return dp[sum]
	}

	minVal := 1000010
	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount-sum {
			minVal = minv(minVal, dfs(sum+coins[i], coins, amount, dp)+1)
		}
	}
	dp[sum] = minVal
	return dp[sum]
}

func minv(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
