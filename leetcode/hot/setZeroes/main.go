package main

func main() {

}

func setZeroes(matrix [][]int) {
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
