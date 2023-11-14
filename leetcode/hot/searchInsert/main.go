package main

func main() {

}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	mid := (left + right) >> 1
	if target == nums[mid] {
		return mid
	}
	if mid+1 < len(nums) && target > nums[mid] && target < nums[mid+1] {
		return mid + 1
	}
	if target < nums[mid] {
		return binarySearch(nums, left, mid-1, target)
	}
	return binarySearch(nums, mid+1, right, target)
}

// 回溯的方法我已经基本学会了，并没有很高的难度
