package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	ans := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ans)
}

//func findMedianSortedArrays(nums1 []int, nums2 []int) (ans float64) {
//	index1, index2, index := 0, 0, 0
//	preVal, val := 0, 0
//	target := (len(nums1) + len(nums2)) / 2
//
//	for index != target+1 && index1 < len(nums1) && index2 < len(nums2) {
//		if nums1[index1] < nums2[index2] {
//			preVal = val
//			val = nums1[index1]
//			index1++
//		} else {
//			preVal = val
//			val = nums2[index2]
//			index2++
//		}
//
//		index++
//	}
//
//	for index != target+1 && index1 < len(nums1) {
//		preVal = val
//		val = nums1[index1]
//		index1++
//		index++
//	}
//
//	for index != target+1 && index2 < len(nums2) {
//		preVal = val
//		val = nums2[index2]
//		index2++
//		index++
//	}
//
//	if is(len(nums1) + len(nums2)) {
//		return float64((preVal + val)) / 2.0
//	}
//
//	return float64(val)
//}
//
//func is(val int) bool {
//	return val%2 == 0
//}

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	totalLength := len(nums1) + len(nums2)
//	if totalLength%2 == 1 {
//		return findK(nums1, nums2, totalLength/2)
//	}
//	return (findK(nums1, nums2, totalLength/2) + findK(nums1, nums2, totalLength-1-totalLength/2)) / 2.0
//}
//
//func findK(nums1, nums2 []int, k int) float64 {
//	// 找到 nums1 k/2 位置的数字
//	// 找到 nums2 k/2 位置的数字
//
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}
	midIndex1, midIndex2 := totalLength/2-1, totalLength/2
	return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
