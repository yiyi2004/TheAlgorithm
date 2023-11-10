package main

func main() {

}

func orangesRotting(grid [][]int) int {
	num1 := check(grid)
	if num1 == 0 {
		return 0
	}

	n := 0
	for num1 > 0 {
		n++
		change(grid)
		biu(grid)
		tmpNum1 := check(grid)
		if tmpNum1 == num1 {
			return -1
		}
		num1 = tmpNum1
	}
	return n
}

func check(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func change(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				if i-1 >= 0 && grid[i-1][j] == 1 {
					grid[i-1][j] = 3
				}
				if i+1 < len(grid) && grid[i+1][j] == 1 {
					grid[i+1][j] = 3
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					grid[i][j-1] = 3
				}
				if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
					grid[i][j+1] = 3
				}
			}
		}
	}
}

func biu(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 3 {
				grid[i][j] = 2
			}
		}
	}
}
