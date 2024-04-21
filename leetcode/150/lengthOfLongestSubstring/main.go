package main

import "fmt"

func main() {
	s := "1234"
	fmt.Println(s[:2])
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	help := make(map[byte]struct{}, 0)
	p, q := 0, 1
	help[s[p]] = struct{}{}
	maxL := 1
	for p < len(s) && q < len(s) {
		if _, ok := help[s[q]]; ok {
			for p <= q {
				if _, exist := help[s[q]]; exist {
					delete(help, s[p])
					p++
				} else {
					break
				}
			}
		} else {
			help[s[q]] = struct{}{}
			if q-p+1 > maxL {
				maxL = q - p + 1
			}
			q++
		}
	}
	return maxL
}
