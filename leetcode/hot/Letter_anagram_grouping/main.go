package main

import (
	"sort"
)

func main() {

}

func groupAnagrams(strs []string) [][]string {
	tmp := make(map[string][]string)

	for _, value := range strs {
		byteStr := []byte(value)

		sort.Slice(byteStr, func(i, j int) bool {
			return byteStr[i] < byteStr[j]
		})

		tmp[string(byteStr)] = append(tmp[string(byteStr)], value)
	}

	res := make([][]string, len(tmp))
	index := 0
	for _, val := range tmp {
		res[index] = append(res[index], val...)
		index++
	}

	return res
}
