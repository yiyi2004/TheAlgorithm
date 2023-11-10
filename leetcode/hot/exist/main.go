package main

import "fmt"

func main() {
	board := [][]byte{{'a', 'a', 'a'}, {'A', 'A', 'A'}, {'a', 'a', 'a'}}
	res := exist(board, "aAaaaAaaA")
	fmt.Println(res)
}

type pair struct {
	x, y int
}

func exist(board [][]byte, word string) bool {
	if board == nil {
		return false
	}
	onPath := make(map[pair]bool)
	var flag = false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != word[0] {
				continue
			} else {
				flag = process(board, i, j, word, 0, onPath)
				if flag == true {
					return true
				}
			}
		}

	}
	return flag
}

func process(board [][]byte, row, col int, word string, index int, onPath map[pair]bool) bool {
	if index == len(word) {
		return true
	}
	if !inArea(board, row, col) {
		return false
	}
	if board[row][col] != word[index] {
		return false
	}
	p := pair{row, col}
	onPath[p] = true
	res := false

	if !onPath[pair{row - 1, col}] {
		res = res || process(board, row-1, col, word, index+1, onPath)
	}
	if !onPath[pair{row + 1, col}] {
		res = res || process(board, row+1, col, word, index+1, onPath)
	}
	if !onPath[pair{row, col - 1}] {
		res = res || process(board, row, col-1, word, index+1, onPath)
	}
	if !onPath[pair{row, col + 1}] {
		res = res || process(board, row, col+1, word, index+1, onPath)
	}
	if index+1 == len(word) {
		return true
	}
	onPath[p] = false
	return res
}

func inArea(board [][]byte, row, col int) bool {
	if row >= 0 && row < len(board) && col >= 0 && col < len(board[0]) {
		return true
	}
	return false
}
