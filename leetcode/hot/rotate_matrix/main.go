package main

func main() {

}

func rotate(matrix [][]int) {
	// 两次对称
	// i, j = j, len(matrix[0])-i
	n := len(matrix[0])
	if matrix == nil {
		return
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] =
				matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}

// 其实不难，但是我不想写了 QAQ，有时间再写吧
//func rotate(matrix [][]int) {
//	if matrix == nil {
//		return
//	}
//	indexi, indexj := 0, 0
//	value := matrix[0][0]
//	bottom := len(matrix)
//	right := len(matrix[0])
//	num := 0
//	numChar := (right + 1) * (bottom + 1)
//
//}
