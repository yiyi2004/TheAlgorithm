package main

import "fmt"

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
		if target == matrix[i][0] || target == matrix[i][len(matrix[i])-1] {
			return true
		}
		if target > matrix[i][0] && target < matrix[i][len(matrix[i])-1] {
			return binarySearch(matrix[i], 0, len(matrix[i])-1, target)
		}
	}
	return false
}

func binarySearch(arr []int, left, right, target int) bool {
	if right < left {
		return false
	}
	mid := (left + right) >> 1
	if arr[mid] == target {
		return true
	}
	if target < arr[mid] {
		return binarySearch(arr, left, mid-1, target)
	}
	return binarySearch(arr, mid+1, right, target)
}
