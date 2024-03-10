package main

import "fmt"

func main() {
	nums := []int{2, 3, -2, 4, -5}
	res := maxProduct(nums)
	fmt.Println(res)
}

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(nums[i], nums[i]*mx, nums[i]*mn)
		minF = min(nums[i], nums[i]*mx, nums[i]*mn)
		ans = max(maxF, ans)
	}

	return ans
}

func max(arr ...int) int {
	mx := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > mx {
			mx = arr[i]
		}
	}
	return mx
}

func min(arr ...int) int {
	mn := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < mn {
			mn = arr[i]
		}
	}
	return mn
}
