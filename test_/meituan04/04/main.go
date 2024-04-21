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
	inputStr, _ := reader.ReadString('\n')
	fields := strings.Fields(inputStr)
	inputStr = fields[0]
	res := 0
	help := make(map[byte]int)
	var dfs func(pre string, index int)
	dfs = func(pre string, index int) {
		if index == len(inputStr) {
			if len(pre)%2 == 1 {
				return
			} else {
				countZero := 0
				for _, v := range help {
					if v != 0 && v != len(pre)/2 {
						return
					}
					if v == 0 {
						countZero++
					}
				}
				if countZero == len(help) {
					return
				}
				res++
				return
			}
		}

		help[inputStr[index]]++
		dfs(pre+string(inputStr[index]), index+1)
		help[inputStr[index]]--
		dfs(pre, index+1)
	}
	dfs("", 0)
	fmt.Println(res)
}
