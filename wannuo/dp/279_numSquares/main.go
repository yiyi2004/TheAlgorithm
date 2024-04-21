package main

func main() {

}

func numSquares(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = i
	}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < len(dp); i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = minVal(dp[i], 1+dp[i-j*j])
		}
	}
	return dp[n]
}

func minVal(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
