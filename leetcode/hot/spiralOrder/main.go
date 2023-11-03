package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}

func spiralOrder(matrix [][]int) []int {
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	indexi, indexj := 0, 0
	num := 0
	res := make([]int, 0)
	numCharacter := (bottom + 1) * (right + 1)
	for num < numCharacter {
		for indexj <= right {
			res = append(res, matrix[indexi][indexj])
			num++
			if indexj != right {
				indexj++
			} else {
				break
			}
		}
		if num == numCharacter {
			break
		}
		indexi++
		for indexi <= bottom {
			res = append(res, matrix[indexi][indexj])
			num++
			if indexi != bottom {
				indexi++
			} else {
				break
			}
		}
		if num == numCharacter {
			break
		}
		indexj--
		for indexj >= left {
			res = append(res, matrix[indexi][indexj])
			num++
			if indexj != left {
				indexj--
			} else {
				break
			}
		}
		if num == numCharacter {
			break
		}
		indexi--
		for indexi > top {
			res = append(res, matrix[indexi][indexj])
			num++
			if indexi != top+1 {
				indexi--
			} else {
				break
			}
		}
		if num == numCharacter {
			break
		}
		indexj++
		top++
		bottom--
		left++
		right--
	}

	return res
}
