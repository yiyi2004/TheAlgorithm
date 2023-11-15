package main

func main() {

}

func partitionLabels(s string) (ans []int) {
	if s == "" {
		return nil
	}

	locationMap := make(map[uint8]int)

	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := locationMap[s[i]]; ok {
			continue
		}
		locationMap[s[i]] = i
	}

	tmpMaxIndex := -1
	pre := -1
	index := 0
	for index < len(s) {
		if v, ok := locationMap[s[index]]; ok {
			if v > tmpMaxIndex {
				tmpMaxIndex = v
			}
		}
		if index == tmpMaxIndex {
			ans = append(ans, tmpMaxIndex-pre)
			pre = tmpMaxIndex
		}
		index++
	}
	return ans
}
