package main

func main() {

}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func Max(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}

// 面试的时候更好的解法
// 二分查找
