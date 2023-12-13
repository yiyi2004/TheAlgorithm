package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	res := searchRange(nums, target)
	fmt.Println(res)
}

//	func searchRange(nums []int, target int) []int {
//		if nums == nil {
//			return nil
//		}
//		if len(nums) == 0 {
//			return []int{-1, -1}
//		}
//		return binarySearch(nums, 0, len(nums)-1, target)
//	}
//
//	func binarySearch(nums []int, left, right int, target int) []int {
//		if right < left {
//			return []int{-1, -1}
//		}
//		mid := (left + right) >> 1
//		if target == nums[mid] {
//			leftIndex := mid
//			rightIndex := mid
//			for leftIndex > 0 {
//				if nums[leftIndex-1] == target {
//					leftIndex--
//				} else {
//					break
//				}
//			}
//			for rightIndex < len(nums)-1 {
//				if nums[rightIndex+1] == target {
//					rightIndex++
//				} else {
//					break
//				}
//			}
//			return []int{leftIndex, rightIndex}
//		}
//
//		if target < nums[mid] {
//			return binarySearch(nums, left, mid-1, target)
//		}
//		return binarySearch(nums, mid+1, right, target)
//	}
//
// 只不过一个用递归实现，一个用非递归是西安
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			lIndex, rIndex := mid, mid
			for lIndex >= 0 && nums[lIndex] == target {
				lIndex--
			}
			for rIndex < len(nums) && nums[rIndex] == target {
				rIndex++
			}
			return []int{lIndex + 1, rIndex - 1}
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}
