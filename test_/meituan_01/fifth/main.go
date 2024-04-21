package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fa map[int]int

func find(x int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func union(x, y int) {
	fa[find(x)] = find(y)
}

// 暂时放弃叭
func main() {
	fa = make(map[int]int)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, q int
	fmt.Scan(&n, &m, &q)

	edges := make(map[int]map[int]bool)
	ops := make([][3]int, q)
	delEdges := make(map[int]map[int]bool)

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		vals := strings.Fields(line)
		u, _ := strconv.Atoi(vals[0])
		v, _ := strconv.Atoi(vals[1])
		if edges[u] == nil {
			edges[u] = make(map[int]bool)
		}
		edges[u][v] = true
	}

	for i := 0; i < q; i++ {
		line, _ := reader.ReadString('\n')
		vals := strings.Fields(line)
		op, _ := strconv.Atoi(vals[0])
		u, _ := strconv.Atoi(vals[1])
		v, _ := strconv.Atoi(vals[2])
		ops[i] = [3]int{op, u, v}
		if op == 1 {
			if delEdges[u] == nil {
				delEdges[u] = make(map[int]bool)
			}
			delEdges[u][v] = true
			if delEdges[v] == nil {
				delEdges[v] = make(map[int]bool)
			}
			delEdges[v][u] = true
		}
	}

	// 逆向构图
	for u, connected := range edges {
		for v := range connected {
			if delEdges[u] != nil && delEdges[u][v] {
				continue // 要删除的边
			}
			if fa[u] == 0 {
				fa[u] = u
			}
			if fa[v] == 0 {
				fa[v] = v
			}
			union(u, v)
		}
	}

	ans := []string{}
	// 逆向遍历
	for i := q - 1; i >= 0; i-- {
		op, u, v := ops[i][0], ops[i][1], ops[i][2]
		if op == 2 {
			if find(u) == find(v) {
				ans = append(ans, "Yes")
			} else {
				ans = append(ans, "No")
			}
		} else {
			if fa[u] == 0 {
				fa[u] = u
			}
			if fa[v] == 0 {
				fa[v] = v
			}
			union(u, v)
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprintln(writer, ans[i])
	}
}
