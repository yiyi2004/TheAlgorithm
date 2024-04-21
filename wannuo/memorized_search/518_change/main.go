package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(change(500, []int{3, 5, 7, 8, 9, 10, 11}))
}

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = -1
		}
	}
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	return dfs(0, 0, coins, amount, dp)
}

func dfs(index int, sum int, coins []int, amount int, dp [][]int) int {
	if sum >= amount {
		if sum == amount {
			return 1
		}
		return 0
	}
	if dp[index][sum] != -1 {
		return dp[index][sum]
	}
	res := 0
	for i := index; i < len(coins); i++ {
		if coins[i] <= amount-sum {
			res += dfs(i, sum+coins[i], coins, amount, dp)
		}
	}
	dp[index][sum] = res
	return res
}
