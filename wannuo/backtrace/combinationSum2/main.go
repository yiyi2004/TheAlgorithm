package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// 同层去重

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	res := make([][]int, 0)
	var dfs func(path []int, index int, sum int)
	dfs = func(path []int, index int, sum int) {
		if sum >= target {
			if sum == target {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		// 排序 + 同层去重的方法
		//help := make(map[int]struct{})
		for i := index; i < len(candidates); i++ {
			// 基于排序的剪枝
			if candidates[i]+sum > target {
				return
			}
			//if _, ok := help[candidates[i]]; ok {
			//	continue
			//} else {
			//	help[candidates[i]] = struct{}{}
			//}
			// i > index
			// index 代表什么
			// 这一层从 index 开始选择，
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(path, i+1, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0, 0)
	return res
}
