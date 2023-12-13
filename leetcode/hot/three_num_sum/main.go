package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 0, 0, 0}
	// -4 -1 -1 0 1 2
	res := threeSum(arr)
	fmt.Println(res)
}

// 边界条件一定要考虑清楚
func threeSum(nums []int) [][]int {
	// a + b + c = 0 + 去重 ---> 双指针方法
	if nums == nil || len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}

	var result [][]int
	// 做题的时候总是会有各种各样的边界判断
	// 1. 数组越界的判断
	// 2. 队列 | 栈是否为空的判断
	// 3. peek 还是 pop，这是一个问题

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < len(nums)-1 && nums[left] == nums[left+1] {
					left++
				}
				for right > 0 && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return result
}

// 思路是对的，但是细节上有问题，也就是去重上有问题的。
// 没有办法去重
//func threeSum(nums []int) [][]int {
//	var result [][]int
//	// 分类讨论一下
//	sort.Ints(nums)
//
//	L, R := 0, len(nums)-1
//	mid := 0
//
//	for R-L > 1 {
//		if abs(nums[L]) < abs(nums[R]) {
//			mid = L + 1
//
//			for mid >= 0 && nums[mid] <= 0 {
//				if nums[L]+nums[mid]+nums[R] == 0 {
//					result = append(result, []int{nums[L], nums[mid], nums[R]})
//				}
//				mid++
//			}
//			tmp := R
//			R--
//			for R-L > 1 && nums[tmp] == nums[R] {
//				R--
//			}
//		} else {
//			mid = R - 1
//
//			for mid >= 0 && nums[mid] >= 0 {
//				if nums[L]+nums[mid]+nums[R] == 0 {
//					result = append(result, []int{nums[L], nums[mid], nums[R]})
//				}
//				mid--
//			}
//			tmp := L
//			L++
//			for R-L > 1 && nums[tmp] == nums[L] {
//				L++
//			}
//		}
//	}
//
//	return result
//}
//
//func abs(val int) int {
//	if val < 0 {
//		return -val
//	}
//	return val
//}
