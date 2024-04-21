package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := restoreIpAddresses("101023")
	fmt.Println(res)
}

// 17 min
func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}
	res := make([]string, 0)
	var dfs func(prePath string, index, k int)
	dfs = func(prePath string, index, k int) {
		if k > 4 {
			return
		}
		if len(s)+3 == len(prePath) && k == 4 {
			res = append(res, prePath)
			return
		}
		for i := index; i <= len(s) && i <= index+3; i++ {
			if isValid(s[index:i]) {
				if len(prePath) == 0 {
					dfs(s[index:i], i, k+1)
				} else {
					dfs(prePath+"."+s[index:i], i, k+1)
				}
			}
		}
	}
	dfs("", 0, 0)
	return res
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	sV, _ := strconv.Atoi(s)
	if sV >= 0 && sV <= 255 {
		return true
	}
	return false
}
