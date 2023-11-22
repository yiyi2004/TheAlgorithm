package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	res := searchMatrix(matrix, 13)
	fmt.Println(res)
}

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) >> 1
		if arr[mid] == target {
			return true
		}
		if target > arr[mid] {
			l = mid + 1
		}
		if target < arr[mid] {
			r = mid - 1
		}
	}
	return false
}
