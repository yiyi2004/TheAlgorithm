package main

func main() {

}

func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	res := 0
	// 空一行和一列，然后进行统一处理
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}

func Max(arr ...int) int {
	res := -1
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}
