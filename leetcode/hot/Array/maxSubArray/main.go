package main

import (
	"fmt"
)

func main() {
	nums := []int{-1}
	maxVal := maxSubArray(nums)
	fmt.Println(maxVal)
}

// 这道题竟然是通过动态规划去解题的，没有什么问题啊。
func maxSubArray(nums []int) int {
	return 0
}

// [-1] 的情况是没有办法解决的
//func maxSubArray(nums []int) int {
//	preArray, preMax, flag := getPreArray(nums)
//	if flag {
//		return preMax
//	}
//	nextMaxVal := make([]int, len(nums))
//	maxIndex := 0
//	lastMaxIndex := 0
//	maxVal := math.MinInt
//	realMax := math.MinInt
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > realMax {
//			realMax = nums[i]
//		}
//	}
//
//	for {
//		for i := lastMaxIndex; i < len(preArray); i++ {
//			if preArray[i] > maxVal {
//				maxVal = preArray[i]
//				maxIndex = i
//			}
//		}
//		for i := lastMaxIndex; i <= maxIndex; i++ {
//			nextMaxVal[i] = maxVal
//		}
//		if maxIndex == len(nextMaxVal)-1 {
//			break
//		}
//		lastMaxIndex = maxIndex + 1
//		maxVal = preArray[lastMaxIndex]
//		maxIndex = lastMaxIndex
//	}
//
//	maxIndex = 0
//	maxVal = math.MinInt
//	for i := 0; i < len(nextMaxVal); i++ {
//		if nextMaxVal[i]-preArray[i] > maxVal {
//			maxVal = nextMaxVal[i] - preArray[i]
//		}
//	}
//	if maxVal < realMax {
//		maxVal = realMax
//	}
//	return maxVal
//}
//
//func getPreArray(nums []int) ([]int, int, bool) {
//	pre := make([]int, len(nums))
//	flag := true
//	pre[0] = nums[0]
//	preMax := pre[0]
//	if pre[0] < 0 {
//		flag = false
//	}
//	for i := 1; i < len(nums); i++ {
//		pre[i] = pre[i-1] + nums[i]
//		if pre[i] > preMax {
//			preMax = pre[i]
//		}
//		if pre[i] < 0 {
//			flag = false
//		}
//	}
//	return pre, preMax, flag
//}
