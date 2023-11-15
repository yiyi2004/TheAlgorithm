package main

import (
	"fmt"
	"math"
)

func main() {
	res := numSquares(13)
	fmt.Println(res)
}

//给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
//完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
//例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

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
	var dfs func(n int, helpMap map[int]int, preCount int)
	dfs = func(n int, helpMap map[int]int, preCount int) {
		if n == 0 {
			if preCount < minCount {
				minCount = preCount
			}
			return
		}
		for i := n; i > 0; i-- {
			if exists(helpMap, i) {
				dfs(n-i, helpMap, preCount+1)
			}
		}
	}
	dfs(n, helpMap, 0)
	return minCount
}

func exists(helpMap map[int]int, v int) bool {
	_, ok := helpMap[v]
	return ok
}

// 动态规划解法
