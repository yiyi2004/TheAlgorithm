package main

import "fmt"

func main() {
	res := integerReplacement(200000000)
	fmt.Println(res)
}

// 改造成记忆化搜索
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	dp := make(map[int]int)

	if n%2 == 0 {
		return o(n, dp)
	}
	return ino(n, dp)
}

func o(n int, dp map[int]int) int {
	n = n / 2
	if n == 1 {
		return 1
	}
	if _, ok := dp[n]; ok {
		return dp[n]
	}
	res := 0
	if n%2 == 0 {
		res = 1 + o(n, dp)
	} else {
		res = 1 + ino(n, dp)
	}
	dp[n] = res
	return dp[n]
}

func ino(n int, dp map[int]int) int {
	if _, ok := dp[n]; ok {
		return dp[n]
	}
	dp[n] = minv(o(n-1, dp), o(n+1, dp)) + 1
	return dp[n]
}

func minv(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
