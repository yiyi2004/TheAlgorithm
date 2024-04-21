package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Union struct {
	Fathers []int
}

func NewUnion(n int) Union {
	u := Union{
		Fathers: make([]int, n+1),
	}
	for i := 1; i <= n; i++ {
		u.Fathers[i] = i
	}
	return u
}

func (u Union) Uni(x, y int) {
	xF := u.Find(x)
	yF := u.Find(y)
	u.Fathers[xF] = yF
}

func (u Union) Find(x int) int {
	if x != u.Fathers[x] {
		y := u.Find(u.Fathers[x])
		u.Fathers[x] = y
	}
	return u.Fathers[x]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	fields := strings.Fields(readString)
	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])
	uni := NewUnion(n)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadString('\n')
		i2 := strings.Fields(s)
		u, _ := strconv.Atoi(i2[0])
		v, _ := strconv.Atoi(i2[1])
		uni.Uni(u, v)
	}
	help := make(map[int]int)
	for i := 1; i <= n; i++ {
		help[uni.Find(i)]++
	}
	if len(help) > 2 {
		fmt.Println(0)
	} else if len(help) == 2 {
		res := 1
		for _, v := range help {
			res *= v
		}
		fmt.Println(res)
	} else {
		fmt.Println(0)
	}
}
