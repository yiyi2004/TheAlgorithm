package main

import (
	"fmt"
)

func main() {
	res := minWindow("ADOBECODEBANC", "ABC")
	fmt.Println(res)
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tmp := [128]int{}
	target := [128]int{}
	for i := 0; i < len(t); i++ {
		target[t[i]-'A']++
	}
	st := -1
	minLen := 10000010
	for left, right := 0, 0; right < len(s); right++ {
		tmp[s[right]-'A']++
		for left <= right && contain(tmp, target) {
			if right-left+1 < minLen {
				st = left
				minLen = right - left + 1
			}
			tmp[s[left]-'A']--
			left++
		}
	}
	if st == -1 {
		return ""
	}
	return s[st : st+minLen]
}

func contain(tmp, target [128]int) bool {
	for i := 0; i < len(tmp); i++ {
		if tmp[i] < target[i] {
			return false
		}
	}
	return true
}

func inTarget(c byte, target [128]int) bool {
	if target[c-'A'] != 0 {
		return true
	}
	return false
}
