package main

import "fmt"

func main() {
	fmt.Println(minDistance("pneumonoultramicroscopicsilicovolcanoconiosis",
		"ultramicroscopically"))
}

func minDistance(word1 string, word2 string) int {
	if word2 == "" {
		return len(word1)
	}
	if word1 == "" {
		return len(word2)
	}

	dp := make([][]int, len(word1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2))
	}
	if word1[0] != word2[0] {
		dp[0][0] = 1
	}
	for i := 1; i < len(dp); i++ {
		if word1[i] == word2[0] {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = dp[i-1][0] + 1
		}
	}

	for i := 1; i < len(dp[0]); i++ {
		if word2[i] == word1[0] {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = dp[0][i-1] + 1
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minv(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func minv(arr ...int) int {
	minV := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < minV {
			minV = arr[i]
		}
	}
	return minV
}
