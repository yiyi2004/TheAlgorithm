package main

func main() {

}

// 杨辉三角
func generate(numRows int) (ans [][]int) {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	ans = [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmp[j] = 1
			} else {
				tmp[j] = ans[i-1][j] + ans[i-1][j-1]
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
