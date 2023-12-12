package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	res := exist(board, "ABCB")
	fmt.Println(res)
}

//	type pair struct {
//		x, y int
//	}
//
//	func exist(board [][]byte, word string) bool {
//		if board == nil {
//			return false
//		}
//		onPath := make(map[pair]bool)
//		var flag = false
//		for i := 0; i < len(board); i++ {
//			for j := 0; j < len(board[0]); j++ {
//				if board[i][j] != word[0] {
//					continue
//				} else {
//					flag = process(board, i, j, word, 0, onPath)
//					if flag == true {
//						return true
//					}
//				}
//			}
//
//		}
//		return flag
//	}
//
//	func process(board [][]byte, row, col int, word string, index int, onPath map[pair]bool) bool {
//		if index == len(word) {
//			return true
//		}
//		if !inArea(board, row, col) {
//			return false
//		}
//		if board[row][col] != word[index] {
//			return false
//		}
//		p := pair{row, col}
//		onPath[p] = true
//		res := false
//
//		if !onPath[pair{row - 1, col}] {
//			res = res || process(board, row-1, col, word, index+1, onPath)
//		}
//		if !onPath[pair{row + 1, col}] {
//			res = res || process(board, row+1, col, word, index+1, onPath)
//		}
//		if !onPath[pair{row, col - 1}] {
//			res = res || process(board, row, col-1, word, index+1, onPath)
//		}
//		if !onPath[pair{row, col + 1}] {
//			res = res || process(board, row, col+1, word, index+1, onPath)
//		}
//		if index+1 == len(word) {
//			return true
//		}
//		onPath[p] = false
//		return res
//	}
// 第二次写的
//func inArea(board [][]byte, row, col int) bool {
//	if row >= 0 && row < len(board) && col >= 0 && col < len(board[0]) {
//		return true
//	}
//	return false
//}
//
//type pair struct {
//	x, y int
//}
//
//func exist(board [][]byte, word string) bool {
//	if board == nil {
//		return false
//	}
//	var res = false
//	onPath := make(map[pair]struct{})
//	var dfs func(index int, i, j int) bool
//	dfs = func(index int, i, j int) bool {
//		if index == len(word) {
//			return true
//		}
//		if inArea(board, i, j) && word[index] == board[i][j] {
//			if exi(onPath, pair{i, j}) {
//				return false
//			}
//			onPath[pair{i, j}] = struct{}{}
//			tmp := dfs(index+1, i+1, j) || dfs(index+1, i-1, j) || dfs(index+1, i, j+1) || dfs(index+1, i, j-1)
//			delete(onPath, pair{i, j})
//			return tmp
//		}
//		return false
//	}
//
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[0]); j++ {
//			res = res || dfs(0, i, j)
//		}
//	}
//	return res
//}
//
//func exi(help map[pair]struct{}, v pair) bool {
//	_, ok := help[v]
//	return ok
//}

// 79. Word Search
// 79. 单词搜索
// 思路：DFS + 回溯
// time O(mnp) space O(p)
// 记录 + 标记 + 恢复的过程
func exist(board [][]byte, word string) bool {
	found := false
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		// 超出索引范围
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		// 走过，不能再走
		if board[i][j] == '*' {
			return
		}
		// 元素不相等
		if board[i][j] != word[k] {
			return
		}
		// 元素相等 && 长度相等，标记找到
		if k == len(word)-1 {
			found = true
			return
		}
		// 标记走过
		tmp := board[i][j]
		board[i][j] = '*'

		// 继续往后走
		dfs(i-1, j, k+1)
		dfs(i+1, j, k+1)
		dfs(i, j-1, k+1)
		dfs(i, j+1, k+1)

		// 从 [i,j] 为中心，向四个方向发散走，进行验证
		// 走完之后需回溯状态，以便于下一个点的验证
		board[i][j] = tmp
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0 // index of the word
			dfs(i, j, k)
		}
	}
	return found
}
