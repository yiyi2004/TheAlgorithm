package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	ans := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ans)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) (ans float64) {
	index1, index2, index := 0, 0, 0
	preVal, val := 0, 0
	target := (len(nums1) + len(nums2)) / 2

	for index != target+1 && index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] < nums2[index2] {
			preVal = val
			val = nums1[index1]
			index1++
		} else {
			preVal = val
			val = nums2[index2]
			index2++
		}

		index++
	}

	for index != target+1 && index1 < len(nums1) {
		preVal = val
		val = nums1[index1]
		index1++
		index++
	}

	for index != target+1 && index2 < len(nums2) {
		preVal = val
		val = nums2[index2]
		index2++
		index++
	}

	if is(len(nums1) + len(nums2)) {
		return float64((preVal + val)) / 2.0
	}

	return float64(val)
}

func is(val int) bool {
	return val%2 == 0
}
