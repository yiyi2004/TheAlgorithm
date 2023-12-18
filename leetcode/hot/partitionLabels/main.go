package main

func main() {

}

//func partitionLabels(s string) (ans []int) {
//	if s == "" {
//		return nil
//	}
//
//	locationMap := make(map[uint8]int)
//
//	for i := len(s) - 1; i >= 0; i-- {
//		if _, ok := locationMap[s[i]]; ok {
//			continue
//		}
//		locationMap[s[i]] = i
//	}
//
//	tmpMaxIndex := -1
//	pre := -1
//	index := 0
//	for index < len(s) {
//		if v, ok := locationMap[s[index]]; ok {
//			if v > tmpMaxIndex {
//				tmpMaxIndex = v
//			}
//		}
//		if index == tmpMaxIndex {
//			ans = append(ans, tmpMaxIndex-pre)
//			pre = tmpMaxIndex
//		}
//		index++
//	}
//	return ans
//}

// 贪心策略，记录每个字母出现的最终位置
// 从头开始遍历字符串，i == end 的时候得到一个子字符串
func partitionLabels(s string) (ans []int) {
	endLoc := [26]int{}
	for i, c := range s {
		endLoc[c-'a'] = i
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if endLoc[s[i]-'a'] > end {
			end = endLoc[s[i]-'a']
		}
		if end == i {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
