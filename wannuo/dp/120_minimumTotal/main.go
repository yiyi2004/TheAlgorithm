package main

func main() {

}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = minVal(triangle[i][j]+triangle[i+1][j], triangle[i][j]+triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func minVal(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
