package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

func lengthOfLongestSubstring(s string) int {
	// write your code here.
	if s == "" {
		return 0
	}
	help := make(map[byte]struct{})
	l := 0
	maxLen := 1
	for i := 0; i < len(s); i++ {
		if _, ok := help[s[i]]; !ok {
			help[s[i]] = struct{}{}
			if i-l+1 > maxLen {
				maxLen = i - l + 1
			}
		} else {
			for {
				_, ok = help[s[i]]
				if ok {
					delete(help, s[l])
					l++
				} else {
					help[s[i]] = struct{}{}
					break
				}
			}
		}
	}
	return maxLen
}
