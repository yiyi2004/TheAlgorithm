package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	t, _ := reader.ReadString('\n')
	res1 := convert(s, t, 1)
	res2 := convert(t, s, 2)
	if res1.count < res2.count {
		fmt.Println(res1.count)
		for _, r := range res1.rows {
			fmt.Println(r.i, r.j, string(r.c))
		}
	} else {
		fmt.Println(res2.count)
		for _, r := range res2.rows {
			fmt.Println(r.i, r.j, string(r.c))
		}
	}
}

type result struct {
	count int
	rows  []row
}

type row struct {
	i int
	j int
	c byte
}

func convert(s, t string, rowid int) result {
	index := len(s) - 1
	count := 0
	var res result
	for index >= 0 {
		if s[index] == t[index] {
			index--
			continue
		} else {
			count++
			res.rows = append(res.rows, row{
				i: rowid,
				j: index + 1,
				c: t[index],
			})
			index--
		}
	}
	res.count = count
	return res
}
