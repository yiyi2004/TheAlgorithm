package main

import (
	"fmt"
	"math"
)

func main() {
	res := numSquares2(12)
	fmt.Println(res)
}

//给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
//完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
//例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//https://leetcode.cn/problems/perfect-squares/?envType=study-plan-v2&envId=top-100-liked
//可能是我的递归写的不是很好

// 应该用动态规划解决问题，但是我还没想好要怎么用动态规划去解决问题
func numSquares(n int) int {
	helpMap := make(map[int]int)
	for i := 0; i <= n; i++ {
		if i*i <= n {
			helpMap[i*i] = i
		} else {
			break
		}
	}
	minCount := math.MaxInt
	var dfs func(helpMap map[int]int, count, n int)
	dfs = func(helpMap map[int]int, count, n int) {
		if n == 0 {
			if count < minCount {
				minCount = count
			}
			return
		}
		for i := n; i > 0; i-- {
			if exists(helpMap, i) {
				dfs(helpMap, count+1, n-i)
			}
		}
	}
	dfs(helpMap, 0, n)
	return minCount
}

func exists(helpMap map[int]int, v int) bool {
	_, ok := helpMap[v]
	return ok
}

func numSquares2(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		minVal := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minVal = min(minVal, dp[i-j*j]+1)
		}
		dp[i] = minVal
	}
	return dp[n]
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

// 动态规划解法
