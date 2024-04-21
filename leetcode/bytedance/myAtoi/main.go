package main

import "math"

func main() {

}

func myAtoi(s string) int {
	if len(s) <= 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] == '+' {
			index := i + 1
			for j := i + 1; j < len(s); j++ {
				if s[j] != '0' {
					index = j
					break
				}
			}
			return count(true, s[index:])
		} else if s[i] == '-' {
			index := i + 1

			for j := i + 1; j < len(s); j++ {
				if s[j] != '0' {
					index = j
					break
				}
			}
			return count(false, s[index:])
		} else if s[i] >= '1' && s[i] <= '9' {
			index := i
			for j := i; j < len(s); j++ {
				if s[j] != '0' {
					index = j
					break
				}
			}
			return count(true, s[index:])
		} else {
			break
		}
	}
	return 0
}

func count(isplus bool, s string) int {
	var res int64

	for i := 0; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' {
			if res*10+int64(s[i]-'0') <= math.MaxInt32 {
				res = res*10 + int64(s[i]-'0')
			} else {
				if isplus {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
		} else {
			if isplus {
				return int(res)
			} else {
				return -int(res)
			}
		}
	}

	if isplus {
		return int(res)
	} else {
		return -int(res)
	}
}
