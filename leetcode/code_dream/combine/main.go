package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var dfs func(index int, pre []int)
	dfs = func(index int, pre []int) {
		if len(pre) == k {
			tmp := make([]int, len(pre))
			copy(tmp, pre)
			res = append(res, tmp)
			return
		}

		if index > n {
			return
		}

		for i := index + 1; i <= n; i++ {
			pre = append(pre, i)
			dfs(i, pre)
			pre = pre[:len(pre)-1]
		}
	}
	dfs(0, []int{})
	return res
}
