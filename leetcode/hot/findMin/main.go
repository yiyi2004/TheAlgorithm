package main

func main() {

}

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot - 1
		} else {
			// 这里为什么
			low = pivot + 1
		}
	}
	return nums[low]
}

//作者：力扣官方题解
//链接：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 错误答案，你根本就没有搞懂啊
//func findMin(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	if len(nums) == 2 {
//		if nums[0] < nums[1] {
//			return nums[0]
//		} else {
//			return nums[1]
//		}
//	}
//	l := 0
//	r := len(nums) - 1
//	index := 0
//	for l <= r {
//		mid := (l + r) >> 1
//
//		if nums[l] > nums[mid] {
//			index = mid
//			r = mid
//		} else if nums[l] <= nums[mid] {
//			l = mid
//		}
//	}
//	return nums[index]
//}
