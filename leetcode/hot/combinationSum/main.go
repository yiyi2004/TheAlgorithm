package main

import (
	"fmt"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Println(res)
}

//
//var combinations [][]int
//
//func combinationSum(candidates []int, target int) [][]int {
//	if len(candidates) == 0 {
//		return [][]int{}
//	}
//	combinations = [][]int{}
//	process(candidates, target, 0, []int{})
//	return combinations
//}
//
//func process(candidates []int, target int, index int, list []int) {
//	if target < 0 {
//		return
//	}
//	if target == 0 {
//		combinations = append(combinations, append([]int{}, list...))
//		return
//	}
//
//	for i := index; i < len(candidates); i++ {
//		// 下一个的起点仍然是 i
//		process(candidates, target-candidates[i], i, append(list, candidates[i]))
//	}
//}

func combinationSum(candidates []int, target int) (res [][]int) {
	if len(candidates) == 0 {
		return nil
	}
	tmp := make([]int, 0)
	var dfs func(index int, target int)
	dfs = func(index int, target int) {
		if target < 0 || index >= len(candidates) {
			return
		}
		if index < len(candidates) && target == 0 {
			res = append(res, append([]int{}, tmp...))
			return
		}
		tmp = append(tmp, candidates[index])
		dfs(index, target-candidates[index])
		tmp = tmp[:len(tmp)-1]
		dfs(index+1, target)
	}
	dfs(0, target)
	return res
}
