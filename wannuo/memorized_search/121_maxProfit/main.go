package main

func main() {

}

func maxProfit(prices []int) int {
	dp := make([][3]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = [3]int{-1, -1, -1}
	}
	return dfs(0, 0, prices, dp)
}

// 返回值代表收益
func dfs(index int, s int, prices []int, dp [][3]int) int {
	if index == len(prices) || s == 2 {
		return 0
	}
	if dp[index][s] != -1 {
		return dp[index][s]
	}
	if s == 0 {
		dp[index][s] = Max(dfs(index+1, 0, prices, dp), dfs(index+1, 1, prices, dp)-prices[index])
		return dp[index][s]
	}
	dp[index][s] = Max(dfs(index+1, 1, prices, dp), dfs(index+1, 2, prices, dp)+prices[index])
	return dp[index][s]
}

func Max(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}
