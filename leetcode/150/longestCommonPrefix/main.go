package main

import "fmt"

func main() {
	a := []string{"1", "2"}
	fmt.Println(a[0:1])
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	minL := len(strs[0])
	for _, v := range strs {
		if len(v) < minL {
			minL = len(v)
		}
	}
	for i := 1; i <= minL; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][0:i] != strs[j][0:i] {
				return strs[0][0 : i-1]
			}
		}
	}
	return strs[0][0:minL]
}
