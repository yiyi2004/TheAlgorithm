package main

func main() {

}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	v1 := dfs(0, nums[1:], dp)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	v2 := dfs(0, nums[:len(nums)-1], dp)
	return maxv(v1, v2)
}

func dfs(index int, nums []int, dp []int) int {
	if index >= len(nums) {
		return 0
	}
	if dp[index] != -1 {
		return dp[index]
	}
	dp[index] = maxv(dfs(index+1, nums, dp), dfs(index+2, nums, dp)+nums[index])
	return dp[index]
}

func maxv(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}
