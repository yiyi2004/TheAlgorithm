package main

func main() {

}

// 7 min 做过
func canPartition(nums []int) bool {
	sum := 0
	maxVal := 0
	for _, num := range nums {
		sum += num
		if num > maxVal {
			maxVal = num
		}
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	if maxVal > sum {
		return false
	}
	dp := make([]int, sum+1)
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			dp[j] = maxV(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[sum] == sum
}

func maxV(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}
