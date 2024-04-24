package main

func main() {

}

func equalSubstring(s string, t string, maxCost int) int {
	res := 0
	sum := 0
	for l, r := 0, 0; r < len(s); {
		sum += abs(s[r], t[r])
		r++
		if sum <= maxCost {
			if r-l > res {
				res = r - l
			}
		} else {
			sum -= abs(s[l], t[l])
			l++
		}
	}
	return res
}

func abs(a, b byte) int {
	if a < b {
		return int(b - a)
	}
	return int(a - b)
}
