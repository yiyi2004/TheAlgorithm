package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 40 12
// a * 0 2
// a * 1 2
// b a 0 3
// b a 1 5
// c a 1 3
// d a 0 1
// d a 1 3
// e b 0 2
// f * 0 8
// f * 1 10
// g f 1 2
// h * 0 4

type node struct {
	val      string
	parent   string
	children []string
	midNum   int
	hardNum  int
	total    int
}

//rowStr, _ := reader.ReadString('\n')
//row := strings.Fields(rowStr)
//if v, ok := rows[row[0]]; !ok {
//no := new(node)
//val, par := row[0], row[1]
//pro, _ := strconv.Atoi(row[2])
//proNum, _ := strconv.Atoi(row[3])
//no.val = val
//if par != "*" {
//no.parent, ok = rows[par]
//if ok {
//parent := rows[par]
//if parent.children == nil {
//parent.children = make([]*node, 0)
//}
//parent.children = append(parent.children, no)
//} else {
//parent := new(node)
//parent.val = val
//parent.children = make([]*node, 0)
//parent.children = append(parent.children, no)
//parent.total = -1
//rows[par] = parent
//no.parent = parent
//}
//}
//if pro == 0 {
//no.hardNum += proNum
//} else {
//no.midNum += proNum
//}
//rows[val] = no
//no.total = -1
//} else {
//val, par := row[0], row[1]
//pro, _ := strconv.Atoi(row[2])
//proNum, _ := strconv.Atoi(row[3])
//if pro == 0 {
//v.hardNum += proNum
//} else {
//v.midNum += proNum
//}
//v.val = val
//v.parent, ok = rows[par]
//if ok {
//parent := rows[par]
//if parent.children == nil {
//parent.children = make([]*node, 0)
//}
//parent.children = append(parent.children, v)
//} else {
//parent := new(node)
//parent.val = val
//parent.children = make([]*node, 0)
//parent.children = append(parent.children, v)
//parent.total = -1
//rows[par] = parent
//v.parent = parent
//}
//}

func main() {
	reader := bufio.NewReader(os.Stdin)
	mn, _ := reader.ReadString('\n')
	fields := strings.Fields(mn)
	throld, _ := strconv.Atoi(fields[0])
	n, _ := strconv.Atoi(fields[1])
	rows := make(map[string]*node)
	for i := 0; i < n; i++ {
		rowStr, _ := reader.ReadString('\n')
		row := strings.Fields(rowStr)
		val, par := row[0], row[1]
		pro, _ := strconv.Atoi(row[2])
		proNum, _ := strconv.Atoi(row[3])

		if src, ok := rows[val]; !ok {
			no := new(node)
			no.val = val
			no.parent = par
			no.total = -1
			if pro == 0 {
				no.hardNum += proNum
			} else {
				no.midNum += proNum
			}
			no.children = make([]string, 0)
			rows[val] = no
		} else {
			if pro == 0 {
				src.hardNum += proNum
			} else {
				src.midNum += proNum
			}
		}
	}
	for _, no := range rows {
		if no.parent != "*" {
			rows[no.parent].children = append(rows[no.parent].children, no.val)
		}
	}
	for _, no := range rows {
		if no.total == -1 {
			count(no, rows)
		}
	}
	res := 0
	for _, no := range rows {
		if no.parent == "*" {
			if no.total > throld {
				res++
			}
		}
	}
	fmt.Println(res)
}

func count(root *node, rows map[string]*node) int {
	if root == nil {
		return 0
	}
	if root.total != -1 {
		return root.total
	}
	res := 5*root.hardNum + 2*root.midNum
	for _, v := range root.children {
		res += count(rows[v], rows)
	}
	root.total = res
	return res
}
