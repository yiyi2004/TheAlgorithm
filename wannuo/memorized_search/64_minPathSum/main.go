package main

func main() {

}

// 设置最大值的技巧
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dp[i][j] = -1
		}
	}
	return dfs(0, 0, grid, dp)
}

func dfs(x, y int, grid [][]int, dp [][]int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return 10000000
	}
	if x == len(grid)-1 && y == len(grid[0])-1 {
		return grid[x][y]
	}
	if dp[x][y] != -1 {
		return dp[x][y]
	}
	dp[x][y] = minv(dfs(x+1, y, grid, dp)+grid[x][y], dfs(x, y+1, grid, dp)+grid[x][y])
	return dp[x][y]
}

func minv(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
