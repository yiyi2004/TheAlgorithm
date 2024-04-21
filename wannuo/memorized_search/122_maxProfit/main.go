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

func dfs(index int, s int, prices []int, dp [][3]int) int {
	if index == len(prices) {
		return 0
	}
	if dp[index][s] != -1 {
		return dp[index][s]
	}
	if s == 0 {
		dp[index][s] = Max(dfs(index+1, 0, prices, dp), dfs(index+1, 1, prices, dp)-prices[index])
		return dp[index][s]
	}
	dp[index][s] = Max(dfs(index+1, 1, prices, dp), dfs(index+1, 0, prices, dp)+prices[index])
	return dp[index][s]
}

func Max(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}

//func maxProfit(prices []int) int {
//	left := 0
//	total := 0
//	for i := 0; i < len(prices); i++ {
//		if prices[i] < prices[left] {
//			left = i
//		} else {
//			if i+1 < len(prices) && prices[i+1] > prices[i] {
//				continue
//			} else {
//				total += prices[i] - prices[left]
//				left = i + 1
//			}
//		}
//	}
//	return total
//}
