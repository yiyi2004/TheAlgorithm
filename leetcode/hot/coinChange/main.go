package main

import (
	"fmt"
	"math"
)

func main() {
	//coins := []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}
	coins := []int{1, 2, 5}
	amounts := 11
	res := coinChange(coins, amounts)
	fmt.Println(res)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

//// 把 dp 的问题都背一背，应该是可以的。
//func coinChange(coins []int, amount int) int {
//	help := make(map[int]bool)
//	for i := 0; i < len(coins); i++ {
//		help[coins[i]] = true
//	}
//	dp := make([]int, amount+1)
//	dp[0] = 0
//	for i := 1; i <= amount; i++ {
//		minVal := math.MaxInt
//		for j := i; j > 0; j-- {
//			if help[j] {
//				if dp[i-j] == -1 {
//					continue
//				}
//				minVal = min(minVal, dp[i-j]+1)
//			}
//		}
//		if minVal == math.MaxInt {
//			dp[i] = -1
//		} else {
//			dp[i] = minVal
//		}
//	}
//	return dp[amount]
//}
//
//func min(v1, v2 int) int {
//	if v1 < v2 {
//		return v1
//	}
//	return v2
//}
//
//// 这道题值得你认真学习一下的。我现在脑子比较糊涂。
//func coinChange1(coins []int, amount int) int {
//	dp := make([]int, amount+1)
//	// 初始化dp[0]
//	dp[0] = 0
//	// 初始化为math.MaxInt32
//	for j := 1; j <= amount; j++ {
//		dp[j] = math.MaxInt32
//	}
//
//	// 遍历物品
//	for i := 0; i < len(coins); i++ {
//		// 遍历背包
//		for j := coins[i]; j <= amount; j++ {
//			if dp[j-coins[i]] != math.MaxInt32 {
//				// 推导公式
//				dp[j] = min(dp[j], dp[j-coins[i]]+1)
//				//fmt.Println(dp,j,i)
//			}
//		}
//	}
//	// 没找到能装满背包的, 就返回-1
//	if dp[amount] == math.MaxInt32 {
//		return -1
//	}
//	return dp[amount]
//}

//作者：代码随想录
//链接：https://leetcode.cn/problems/coin-change/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
