package main

import (
	"fmt"
)

func main() {
	res := solveNQueens(4)
	//quenes := []string{".Q..", "...Q", "Q...", "..Q."}
	//res := check(quenes, 4)
	//checkSingleFlag := checkSingle(quenes, 3, 2, 4)
	fmt.Println(res)
}

// 需要优化时间复杂度
// 通过是通过了，但是还需要进行时间和空间复杂度的优化
func solveNQueens(n int) (ans [][]string) {
	var dfs func(n, row int, pre []string)
	dfs = func(n, row int, pre []string) {
		if row == n {
			if check(pre, n) {
				ans = append(ans, append([]string{}, pre...))
			}
			return
		}
		for i := 0; i < n; i++ {
			tmpStr := ""
			for j := 0; j < n; j++ {
				if j == i {
					tmpStr += "Q"
				} else {
					tmpStr += "."
				}
			}
			pre = append(pre, tmpStr)
			if check(pre, n) {
				dfs(n, row+1, pre)
			}
			pre = pre[:len(pre)-1]
		}
	}
	dfs(n, 0, []string{})
	return

}

func check(res []string, n int) (flag bool) {
	flag = true
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			if res[i][j] == 'Q' {
				flag = flag && checkSingle(res, i, j, n)
				if flag == false {
					return false
				}
			}
		}
	}
	return flag
}

func checkSingle(res []string, row, col int, n int) (flag bool) {
	for i := 0; i < len(res); i++ {
		if i == row {
			continue
		}
		if res[i][col] == 'Q' {
			return false
		}
	}
	for i := 1; i < n; i++ {
		if row+i < len(res) && col+i < len(res[0]) && res[row+i][col+i] == 'Q' {
			return false
		}
		if row-i >= 0 && col-i >= 0 && res[row-i][col-i] == 'Q' {
			return false
		}
		if row+i < len(res) && col-i >= 0 && res[row+i][col-i] == 'Q' {
			return false
		}
		if row-i >= 0 && col+i < len(res[0]) && res[row-i][col+i] == 'Q' {
			return false
		}
	}
	return true
}

// 斜线是有问题的
// 你这样做肯定是有问题的
