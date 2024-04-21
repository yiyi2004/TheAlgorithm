package main

func main() {

}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxIndex := nums[0]
	for i := 0; i < len(nums); i++ {
		if maxIndex < i {
			return false
		}
		if i+nums[i] > maxIndex {
			maxIndex = i + nums[i]
		}
	}
	return true
}

//func dfs(index int, nums []int, dp map[int]bool) bool {
//	if index >= len(nums)-1 {
//		return true
//	}
//	if _, ok := dp[index]; ok {
//		return dp[index]
//	}
//	for i := 1; i <= nums[index]; i++ {
//		if dfs(index+i, nums, dp) {
//			dp[index] = true
//			return true
//		}
//	}
//	dp[index] = false
//	return false
//}
