package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	index := search(nums, target)
	fmt.Println(index)
}

// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) >> 1
		if target == nums[mid] {
			return mid
		}
		if nums[l] > nums[mid] {
			// 左边无序
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if nums[l] <= nums[mid] {
			// 左边有序
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

//func search(nums []int, target int) int {
//	if nums == nil {
//		return -1
//	}
//	if len(nums) == 0 {
//		return -1
//	}
//	//if len(nums) == 1 && nums[0] == target {
//	//	return 0
//	//} else if len(nums) == 1 && nums[0] != target {
//	//	return -1
//	//}
//	var index int
//	if !sort.IsSorted(sort.IntSlice(nums)) {
//		index = findIndex(nums, 0, len(nums)-1)
//	} else {
//		index = 0
//	}
//	if target < nums[index] {
//		return -1
//	}
//	if index-1 > 0 && target > nums[index-1] {
//		return -1
//	}
//	if target >= nums[index] && target <= nums[len(nums)-1] {
//		return binarySearch(nums, index, len(nums)-1, target)
//	}
//	return binarySearch(nums, 0, index-1, target)
//}
//
//func findIndex(nums []int, left, right int) int {
//	if right < left {
//		return 0
//	}
//	mid := (left + right) >> 1
//	if mid-1 >= 0 && nums[mid] < nums[mid-1] {
//		return mid
//	}
//	if nums[left] < nums[mid] {
//		return findIndex(nums, mid, right)
//	}
//	return findIndex(nums, left, mid)
//}
//
//// return index
//func binarySearch(nums []int, left, right, target int) int {
//	if right < left {
//		return -1
//	}
//	mid := (left + right) >> 1
//	if nums[mid] == target {
//		return mid
//	}
//	if target < nums[mid] {
//		return binarySearch(nums, left, mid-1, target)
//	}
//	return binarySearch(nums, mid+1, right, target)
//}

// 分为乱序和有序两个部分，然后分别处理
//func search(nums []int, target int) int {
//	if len(nums) == 0 {
//		return -1
//	}
//	l, r := 0, len(nums)-1
//	for l <= r {
//		mid := (l + r) >> 1
//		if nums[mid] == target {
//			return mid
//		}
//
//		if nums[mid] >= nums[l] { //
//			// 左边有序的，右边是无序的
//			if target >= nums[l] && target < nums[mid] {
//				r = mid - 1
//			} else {
//				l = mid + 1
//			}
//		} else if nums[mid] < nums[l] {
//			// 右边是有序的
//			if target > nums[mid] && target <= nums[r] {
//				l = mid + 1
//			} else {
//				r = mid - 1
//			}
//		}
//	}
//	return -1
//}
