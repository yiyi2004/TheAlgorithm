package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{2, 3, 6, 7}, 7))
}

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(candidates)
	var backtrace func(path []int, sum int, index int)
	backtrace = func(path []int, sum int, index int) {
		if sum >= target {
			if sum == target {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for i := index; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				return
			}
			path = append(path, candidates[i])
			backtrace(path, sum+candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	backtrace([]int{}, 0, 0)
	return res
}
