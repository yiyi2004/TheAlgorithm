package main

import "fmt"

func main() {
	res := longestPalindrome("aaaa")
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(dp); i++ {
		dp[i][i] = true
	}
	for i := 0; i < len(dp)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	for j := 2; j < len(dp); j++ {
		startI := 0
		startJ := j
		for startI < len(dp)-2 && startJ < len(dp) {
			if s[startI] == s[startJ] && dp[startI+1][startJ-1] {
				dp[startI][startJ] = true
			}
			startI++
			startJ++
		}
	}
	res := string(s[0])
	maxLen := 1
	for i := 0; i < len(dp); i++ {
		for j := i; j < len(dp); j++ {
			if dp[i][j] && j-i+1 > maxLen {
				res = s[i : j+1]
				maxLen = j - i + 1
			}
		}
	}
	return res
}
