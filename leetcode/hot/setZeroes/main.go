package main

func main() {

}
func setZeroes(matrix [][]int) {
	if matrix == nil {
		return
	}
	rowZero, colZero := false, false
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			rowZero = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			colZero = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if colZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

/*func setZeroes(matrix [][]int) {
	if matrix == nil {
		return
	}
	rowMap := make(map[int]struct{})
	colMap := make(map[int]struct{})

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = struct{}{}
				colMap[j] = struct{}{}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if exists(rowMap, i) || exists(colMap, j) {
				matrix[i][j] = 0
			}
		}
	}
}

func exists(m map[int]struct{}, val int) bool {
	_, ok := m[val]
	return ok
}
*/
