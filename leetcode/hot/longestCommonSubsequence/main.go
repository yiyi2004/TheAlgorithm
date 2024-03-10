package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("bsbininm",
		"jmjkbkjkv"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if text2 == "" || text1 == "" {
		return 0
	}
	dp := make([][]int, len(text1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2))
	}
	if text2[0] == text1[0] {
		dp[0][0] = 1
	}
	for i := 1; i < len(dp); i++ {
		if text2[0] == text1[i] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < len(dp[0]); i++ {
		if text1[0] == text2[i] {
			dp[0][i] = 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func max(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}
