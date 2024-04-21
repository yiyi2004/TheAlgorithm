package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	if len(s) <= len(t) {
		return ""
	}

	l, r := 0, 0
	curSet := [256]int{}
	tSet := [256]int{}
	for _, v := range t {
		tSet[v]++
	}
	minL := len(s)
	minStr := ""
	curSet[s[r]]++
	for r < len(s) {
		for !equal(curSet, tSet) {
			r++
			if r >= len(s) {
				break
			}
			curSet[s[r]]++
		}
		if r-l+1 < minL {
			minL = r - l + 1
			minStr = s[l : r+1]
		}
		curSet[s[l]]--
		l++
		for !inSet(s[l], tSet) {
			curSet[s[l]]--
			l++
		}
		if equal(curSet, tSet) {
			if r-l+1 < minL {
				minL = r - l + 1
				minStr = s[l : r+1]
			}
		}
	}

	return minStr
}

func equal(curSet, tSet [256]int) bool {
	for i, v := range tSet {
		if v > curSet[i] {
			return false
		}
	}
	return true
}

func inSet(c byte, tSet [256]int) bool {
	return tSet[c] > 0
}
