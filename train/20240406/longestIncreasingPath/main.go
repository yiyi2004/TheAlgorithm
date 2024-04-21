package main

import "fmt"

func main() {
	matrix := [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}
	res := longestIncreasingPath(matrix)
	fmt.Println(res)
}

// 135 / 139 个通过的测试用例 超出时间限制
// 12 min
func longestIncreasingPath(matrix [][]int) int {
	if matrix == nil {
		return 0
	}
	n := len(matrix)
	m := len(matrix[0])
	res := 0
	var dfs func(x, y int, count int)
	dfs = func(x, y int, count int) {
		if count > res {
			res = count
		}
		if inArea(x-1, y, n, m) && matrix[x-1][y] != -1 && matrix[x-1][y] > matrix[x][y] {
			tmp := matrix[x][y]
			matrix[x][y] = -1
			dfs(x-1, y, count+1)
			matrix[x][y] = tmp
		}
		if inArea(x+1, y, n, m) && matrix[x+1][y] != -1 && matrix[x+1][y] > matrix[x][y] {
			tmp := matrix[x][y]
			matrix[x][y] = -1
			dfs(x+1, y, count+1)
			matrix[x][y] = tmp
		}
		if inArea(x, y-1, n, m) && matrix[x][y-1] != -1 && matrix[x][y-1] > matrix[x][y] {
			tmp := matrix[x][y]
			matrix[x][y] = -1
			dfs(x, y-1, count+1)
			matrix[x][y] = tmp
		}
		if inArea(x, y+1, n, m) && matrix[x][y+1] != -1 && matrix[x][y+1] > matrix[x][y] {
			tmp := matrix[x][y]
			matrix[x][y] = -1
			dfs(x, y+1, count+1)
			matrix[x][y] = tmp
		}
	}
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			dfs(x, y, 1)
		}
	}
	return res
}

func inArea(x, y int, n, m int) bool {
	if x >= 0 && x < n && y >= 0 && y < m {
		return true
	}
	return false
}
