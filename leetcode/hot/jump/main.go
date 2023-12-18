package main

func main() {

}

// https://leetcode.cn/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-100-liked

// 反正是碰对了，不知道什么原因
//func jump(nums []int) int {
//	count := 0
//	index := len(nums) - 2
//	targetIndex := len(nums) - 1
//	for targetIndex > 0 {
//		for i := targetIndex - 1; i >= 0; i-- {
//			if nums[i] >= targetIndex-i {
//				index = i
//			}
//		}
//		count++
//		targetIndex = index
//	}
//	return count
//}

func jump(nums []int) int {
	count := 0
	targetIndex := len(nums) - 1
	index := len(nums) - 2

	for targetIndex > 0 {
		for i := targetIndex; i >= 0; i-- {
			if nums[i]+i >= targetIndex {
				index = i
			}
		}
		count++
		targetIndex = index
	}

	return count
}
