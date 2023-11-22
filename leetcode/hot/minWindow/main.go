package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	res := minWindow(s, t)
	fmt.Println(res)
}

func minWindow(s string, t string) string {
	arrs := [129]int{}
	arrt := [129]int{}
	for i := 0; i < len(t); i++ {
		arrt[t[i]]++
	}
	left, right := 0, 0
	minW := math.MaxInt
	minLeft, minRight := 0, len(s)-1
	for right < len(s) {
		// 如果不满住条件，右移 right 指针
		arrs[s[right]]++
		for right < len(s) && !isBingo(arrs, arrt) {
			right++
			if right < len(s) {
				arrs[s[right]]++
			}
		}
		if right < len(s) && right-left+1 < minW {
			minW = right - left + 1
			minLeft = left
			minRight = right
			minW = right - left + 1
		}
		for left <= right {
			if isBingo(arrs, arrt) {
				arrs[s[left]]--
				left++
			} else {
				left--
				arrs[s[left]]++
				break
			}
		}
		if right-left+1 < minW {
			minW = right - left + 1
			minLeft = left
			minRight = right
		}
		arrs[s[left]]--
		left++
	}
	return s[minLeft : minRight+1]
}

func isBingo(arrs, arrt [129]int) bool {
	for i := 0; i < len(arrs); i++ {
		if arrs[i] < arrt[i] {
			return false
		}
	}
	return true
}

//func minWindow(s string, t string) string {
//	ori, cnt := map[byte]int{}, map[byte]int{}
//	for i := 0; i < len(t); i++ {
//		ori[t[i]]++
//	}
//
//	sLen := len(s)
//	len := math.MaxInt32
//	ansL, ansR := -1, -1
//
//	check := func() bool {
//		for k, v := range ori {
//			if cnt[k] < v {
//				return false
//			}
//		}
//		return true
//	}
//	for l, r := 0, 0; r < sLen; r++ {
//		if r < sLen && ori[s[r]] > 0 {
//			cnt[s[r]]++
//		}
//		for check() && l <= r {
//			if (r - l + 1 < len) {
//				len = r - l + 1
//				ansL, ansR = l, l + len
//			}
//			if _, ok := ori[s[l]]; ok {
//				cnt[s[l]] -= 1
//			}
//			l++
//		}
//	}
//	if ansL == -1 {
//		return ""
//	}
//	return s[ansL:ansR]
//}
//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/minimum-window-substring/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
