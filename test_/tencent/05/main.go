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
	fields := strings.Fields(readString)
	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])
	target := "tencent"
	graph := make([][]byte, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]byte, m)
		str, _ := reader.ReadString('\n')
		for j := 0; j < m; j++ {
			graph[i][j] = str[j]
		}
	}
	ans := 0
	var dfs func(x, y, index int)
	dfs = func(x, y, index int) {
		if index == len(target)-1 {
			ans++
			return
		}
		if inArea(x-1, y, n, m) && graph[x-1][y] == target[index+1] {
			dfs(x-1, y, index+1)
		}
		if inArea(x+1, y, n, m) && graph[x+1][y] == target[index+1] {
			dfs(x+1, y, index+1)
		}
		if inArea(x, y-1, n, m) && graph[x][y-1] == target[index+1] {
			dfs(x, y-1, index+1)
		}
		if inArea(x, y+1, n, m) && graph[x][y+1] == target[index+1] {
			dfs(x, y+1, index+1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] != 't' {
				continue
			}
			dfs(i, j, 0)
		}
	}

	fmt.Println(ans)
}

func inArea(x, y int, n, m int) bool {
	if x >= 0 && x < n && y >= 0 && y < m {
		return true
	}
	return false
}
