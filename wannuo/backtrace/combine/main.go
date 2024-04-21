package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtrace func(path []int, index int)
	backtrace = func(path []int, index int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := index; i <= n; i++ {
			path = append(path, i)
			backtrace(path, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrace([]int{}, 1)
	return res
}
