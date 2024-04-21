package main

import "fmt"

func main() {
	n := 0
	m := 0
	fmt.Scan(&n, &m)
	input := make([][]int, n)
	for i := 0; i < n; i++ {
		input[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		tmpStr := ""
		fmt.Scan(&tmpStr)
		for j := 0; j < len(tmpStr); j++ {
			if tmpStr[j] == '0' {
				input[i][j] = 0
			} else {
				input[i][j] = 1
			}
		}
	}
	preArr := make([][]int, n)
	for i := 0; i < n; i++ {
		preArr[i] = make([]int, m)
	}
	if input[0][0] == 1 {
		preArr[0][0] = 1
	}
	for i := 1; i < n; i++ {
		if input[i][0] == 1 {
			preArr[i][0] = preArr[i-1][0] + 1
		} else {
			preArr[i][0] = preArr[i-1][0]
		}
	}
	for i := 1; i < m; i++ {
		if input[0][i] == 1 {
			preArr[0][i] = preArr[0][i-1] + 1
		} else {
			preArr[0][i] = preArr[0][i-1]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			preArr[i][j] = preArr[i][j-1] + preArr[i-1][j] - preArr[i-1][j-1] + input[i][j]
		}
	}

	if n <= 1 || m <= 1 {
		fmt.Println(0)
	} else {
		ans := 0
		for k := 2; k <= min(m, n) && k%2 == 0; k++ {
			for i := k - 1; i < n; i++ {
				for j := k - 1; j < m; j++ {
					flagA := false
					flagB := false
					if i-k+1 == 0 {
						flagA = true
					}
					if j-k+1 == 0 {
						flagB = true
					}
					if flagA && flagB {
						if preArr[i][j] == k*k/2 {
							ans++
						}
					} else if flagA && !flagB {
						if preArr[i][j]-preArr[i][j-k] == k*k/2 {
							ans++
						}
					} else if !flagA && flagB {
						if preArr[i][j]-preArr[i-k][j] == k*k/2 {
							ans++
						}
					} else {
						if preArr[i][j]-preArr[i-k][j]-preArr[i][j-k]+preArr[i-k][j-k] == k*k/2 {
							ans++
						}
					}
				}
			}
		}
		fmt.Printf("%d\n", ans)
	}
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
