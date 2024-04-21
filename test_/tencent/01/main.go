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
	links := make([][]ele, n+1)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadString('\n')
		fields2 := strings.Fields(s)
		u, _ := strconv.Atoi(fields2[0])
		v, _ := strconv.Atoi(fields2[1])
		color := fields2[2]
		links[u] = append(links[u], ele{
			color: color,
			node:  v,
		})
		links[v] = append(links[v], ele{
			color: color,
			node:  u,
		})
	}
	ans := 0
	for i := 1; i <= n; i++ {
		flag := true
		for _, v := range links[i] {
			if v.color == "W" {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}
	fmt.Println(ans)
}

type ele struct {
	color string
	node  int
}
