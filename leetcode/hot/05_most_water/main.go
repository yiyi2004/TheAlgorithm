package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 18, 17, 6}
	res := maxArea(arr)
	fmt.Println(res)
}

func maxArea(height []int) int {
	L := 0
	R := len(height) - 1

	maxValue := min(height[L], height[R]) * (R - L)

	for L != R {
		if height[L] < height[R] {
			L++
		} else {
			R--
		}
		tmpArea := min(height[L], height[R]) * (R - L)
		maxValue = max(maxValue, tmpArea)
	}

	return maxValue
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}
