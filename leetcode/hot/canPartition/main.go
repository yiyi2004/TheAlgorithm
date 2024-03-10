package main

func main() {

}

//func canPartition(nums []int) bool {
//	if len(nums) == 1 {
//		return false
//	}
//
//	sum := 0
//	for i := 0; i < len(nums); i++ {
//		sum += nums[i]
//	}
//	if sum%2 == 1 {
//		return false
//	}
//	sum = sum / 2
//
//	dp := make([][]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		dp[i] = make([]int, sum+1)
//	}
//
//	for i := 0; i < len(dp); i++ {
//		dp[i][0] = 0
//	}
//	for i := 0; i <= sum; i++ {
//		if i >= nums[0] {
//			dp[0][i] = nums[0]
//		}
//	}
//
//	for i := 1; i < len(nums); i++ {
//		for j := 1; j <= sum; j++ {
//			if j >= nums[i] {
//				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
//			} else {
//				dp[i][j] = dp[i-1][j]
//			}
//		}
//	}
//
//	return dp[len(nums)-1][sum] == sum
//}
//
//func max(v1, v2 int) int {
//	if v1 < v2 {
//		return v2
//	}
//	return v1
//}

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	// 优化部分，减枝
	target := sum / 2
	if max > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		v := nums[i]
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}
