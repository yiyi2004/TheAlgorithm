package main

import "fmt"

func main() {
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	return dfs(n, dp)
}

// 拆分，要么拆成 2 个数字，要么拆成 1 + f
// dfs 代表拆分逻辑
func dfs(n int, dp []int) int {
	if n == 2 {
		return 1
	}
	if dp[n] != -1 {
		return dp[n]
	}
	res := 0
	for i := 1; i <= n-2; i++ {
		res = maxv(res, i*(n-i), i*dfs(n-i, dp))
	}
	dp[n] = res
	return dp[n]
}

func maxv(arr ...int) int {
	res := -1
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}
