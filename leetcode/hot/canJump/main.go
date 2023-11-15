package main

func main() {

}

// 如果
func canJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if k < i {
			return false
		}
		k = max(k, i+nums[i])
	}
	// 不断更新从当前节点出发，能够达到的最远距离。
	return true
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

// 可以的，策略是错误的	。
//func canJump(nums []int) bool {
//	if nums[0] == 0 {
//		return false
//	}
//	countZero := 0
//	index := 0
//	countMaxJump := nums[0]
//	for index != len(nums)-1 {
//		for index < len(nums) && nums[index] != 0 {
//			index++
//			if index < len(nums) {
//				countMaxJump += nums[index]
//			} else {
//				return true
//			}
//		}
//		for index < len(nums) && nums[index] == 0 {
//			countZero++
//			index++
//		}
//		if countMaxJump < countZero {
//			return false
//		}
//		countZero = 0
//		if index < len(nums) {
//			countMaxJump = nums[index]
//		} else {
//			return true
//		}
//	}
//	return true
//}
