package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	arrStr, _ := reader.ReadString('\n')
	fields := strings.Fields(arrStr)
	if len(fields) <= 2 {
		res := ""
		for i := 0; i < len(fields); i++ {
			res += fields[i]
		}
		fmt.Println(res)
		return
	}

	p, mid, q := 0, 1, 2
	for q < len(fields) {
		if fields[p] == fields[mid] && fields[mid] == fields[q] {
			if q+1 < len(fields) {
				fields = append(fields[:p], fields[q+1:]...)
			} else {
				fields = fields[:p]
			}
			if p-2 >= 0 {
				p = p - 2
				mid = p + 1
				q = mid + 1
			} else {
				p, mid, q = 0, 1, 2
			}
		} else {
			p++
			mid++
			q++
		}
	}

	res := ""
	for i := 0; i < len(fields); i++ {
		res += fields[i]
		if i < len(fields)-1 {
			res += " "
		}
	}
	if res == "" {
		fmt.Println(0)
		return
	}
	fmt.Println(res)
	return
}
