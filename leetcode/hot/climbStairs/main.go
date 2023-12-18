package main

func main() {

}

// 1. dp[i][j] = val
// 2. 递归关系式
// 3. 初始条件
// 4. 确认遍历顺序
// 5. print
// 遍历顺序在其他的题目中也非常重要

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
