package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}
	if numRows == 1 {
		return res
	}
	for i := 1; i < len(res); i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = row
	}
	return res
}
