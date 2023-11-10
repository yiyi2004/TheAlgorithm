package main

func main() {

}

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '1' {
				continue
			} else {
				res++
				process(grid, i, j)
			}
		}
	}
	return res
}

func process(grid [][]byte, row, col int) {
	if !inArea(grid, row, col) {
		return
	}
	if grid[row][col] != '1' {
		return
	}
	grid[row][col] = '2'
	process(grid, row-1, col)
	process(grid, row, col+1)
	process(grid, row+1, col)
	process(grid, row, col-1)
}

func inArea(grid [][]byte, row, col int) bool {
	if row >= 0 && row <= len(grid)-1 && col >= 0 && col <= len(grid[0])-1 {
		return true
	}
	return false
}
