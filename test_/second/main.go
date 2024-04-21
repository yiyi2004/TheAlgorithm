package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	tmpStr := ""
	fmt.Scan(&tmpStr)
	index := 1
	begin := 2
	end := 2
	inputStr := ""

	tmpByte := tmpStr[0]
	for index < len(tmpStr) {
		if tmpStr[index] == '(' {
			begin = index
			for i := index + 1; i < len(tmpStr); i++ {
				if tmpStr[i] == ')' {
					end = i
					break
				}
			}
		}
		index = end + 1
		num, _ := strconv.Atoi(tmpStr[begin+1 : end])
		inputStr += strings.Repeat(string(tmpByte), num)
		if index < len(tmpStr) {
			tmpByte = tmpStr[index]
		}
		index++
	}
	preArr := make([]int, len(inputStr))
	backArr := make([]int, len(inputStr))
	helpPreMap := make(map[byte]struct{})
	helpBackMap := make(map[byte]struct{})
	preArr[0] = 1
	helpPreMap[inputStr[0]] = struct{}{}
	backArr[len(backArr)-1] = 1
	helpBackMap[inputStr[len(inputStr)-1]] = struct{}{}

	for i := 1; i < len(preArr); i++ {
		if _, ok := helpPreMap[inputStr[i]]; ok {
			preArr[i] = preArr[i-1]
		} else {
			preArr[i] = preArr[i-1] + 1
			helpPreMap[inputStr[i]] = struct{}{}
		}
	}
	for i := len(backArr) - 2; i >= 0; i-- {
		if _, ok := helpBackMap[inputStr[i]]; ok {
			backArr[i] = backArr[i+1]
		} else {
			backArr[i] = backArr[i+1] + 1
			helpBackMap[inputStr[i]] = struct{}{}
		}
	}
	ans := 0
	if len(inputStr)*preArr[len(preArr)-1] >= k {
		ans = 1
	} else {
		fmt.Println(-1)
		return
	}
	for i := 0; i < len(inputStr); i++ {
		if i+1 < len(inputStr) && (i-0+1)*preArr[i] >= k && (len(inputStr)-i-1)*backArr[i+1] >= k {
			ans++
		}
	}
	fmt.Println(ans)
}
