package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}

//func subsets(nums []int) [][]int {
//	if nums == nil {
//		return nil
//	}
//	var res [][]int
//	path := make([]int, len(nums))
//	onPath := make([]bool, len(nums))
//
//	res = append(res, []int{})
//	var dfs func(i int, left int)
//	dfs = func(i int, left int) {
//		if i == left {
//			res = append(res, append([]int{}, path[:left]...))
//		}
//
//		for j, on := range onPath {
//			if !on {
//				path[i] = nums[j]
//				onPath[j] = true
//				dfs(i+1, left)
//				onPath[j] = false
//			}
//		}
//	}
//
//	for i := 1; i <= len(nums); i++ {
//		dfs(0, i)
//	}
//
//	return res
//}

// 你连题目都没有看懂

//func subsets(nums []int) (ans [][]int) {
//	set := []int{}
//	var dfs func(i int)
//	dfs = func(i int) {
//		if i == len(nums) {
//			ans = append(ans, append([]int{}, set...))
//			return
//		}
//		set = append(set, nums[i])
//		dfs(i + 1)
//		set = set[:len(set)-1]
//		dfs(i + 1)
//	}
//	dfs(0)
//	return
//}

//func subsets(nums []int) (ans [][]int) {
//	set := []int{}
//	var dfs func(int)
//	dfs = func(cur int) {
//		if cur == len(nums) {
//			ans = append(ans, append([]int(nil), set...))
//			return
//		}
//		set = append(set, nums[cur])
//		dfs(cur + 1)
//		set = set[:len(set)-1]
//		dfs(cur + 1)
//	}
//	dfs(0)
//	return
//}

// 这种解决方案是错误的，人家是去重的
//func subsets(nums []int) [][]int {
//	if len(nums) == 0 {
//		return nil
//	}
//	res := make([][]int, 0)
//	path := make([]int, len(nums))
//	onPath := make(map[int]bool, len(nums))
//	for i := 0; i < len(nums); i++ {
//		onPath[i] = false
//	}
//
//	res = append(res, []int{})
//	var dfs func(index int, end int)
//	dfs = func(index int, end int) {
//		if index == end {
//			res = append(res, append([]int{}, path[:end]...))
//			return
//		}
//		for i, ok := range onPath {
//			if !ok {
//				fmt.Println(index)
//				path[index] = nums[i]
//				onPath[i] = true
//				dfs(index+1, end)
//				onPath[i] = false
//			}
//		}
//	}
//	for i := 1; i <= len(nums); i++ {
//		dfs(0, i)
//	}
//	return res
//}

// 解题思路是，放或者不放当前元素
func subsets(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return nil
	}
	set := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, set...))
			return
		}
		set = append(set, nums[i])
		dfs(i + 1)
		set = set[:len(set)-1]
		dfs(i + 1)
	}
	dfs(0)
	return res
}
