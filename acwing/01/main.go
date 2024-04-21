package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	fs := strings.Fields(readString)
	n, _ := strconv.Atoi(fs[0])
	m, _ := strconv.Atoi(fs[1])

	res := [110][110]int{}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	x := 0
	y := 0
	d := 0
	k := 1
	for k <= m*n {
		res[x][y] = k
		a := x + dx[d]
		b := y + dy[d]
		if a < 0 || a >= n || b < 0 || b >= m || res[a][b] > 0 {
			d = (d + 1) % 4
			a = x + dx[d]
			b = y + dy[d]
		}
		x = a
		y = b
		k++
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println()
	}
}
