package main

import (
	"fmt"
	"math"
)

func main() {
	res := equalSubstring("krpgjbjjznpzdfy", "nxargkbydxmsgby", 14)
	fmt.Println(res)
}

// 32 / 37 个通过的测试用例
func equalSubstring(s string, t string, maxCost int) int {
	if len(s) == 0 {
		return 0
	}
	l, r := 0, 0
	maxL := 0
	preCost := 0
	for l < len(s) && r < len(s) && l <= r {
		for preCost <= maxCost && r < len(s) {
			preCost += int(math.Abs(float64(t[r]) - float64(s[r])))
			r++
		}

		if r-l-1 > maxL {
			maxL = r - l - 1
		}

		for preCost > maxCost {
			preCost -= int(math.Abs(float64(t[l]) - float64(s[l])))
			l++
		}
	}
	return maxL
}
