package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([]int, n)
	var dfs func(index int, col [20]bool, diag1 [20]bool, diag2 [20]bool)
	dfs = func(index int, col [20]bool, diag1 [20]bool, diag2 [20]bool) {
		if index == n {
			res = append(res, getQueens(path))
			return
		}
		for i := 0; i < n; i++ {
			if col[i] || diag1[index+i] || diag2[index-i+9] {
				continue
			}
			col[i], diag1[index+i], diag2[index-i+9] = true, true, true
			path[index] = i
			dfs(index+1, col, diag1, diag2)
			col[i], diag1[index+i], diag2[index-i+9] = false, false, false
		}
	}
	dfs(0, [20]bool{}, [20]bool{}, [20]bool{})
	return res
}

func getQueens(path []int) []string {
	res := make([]string, 0)
	for i := 0; i < len(path); i++ {
		tmp := ""
		for j := 0; j < len(path); j++ {
			if j == path[i] {
				tmp += "Q"
			} else {
				tmp += "."
			}
		}
		res = append(res, tmp)
	}
	return res
}
