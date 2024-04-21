package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	fields := strings.Fields(readString)
	n, _ := strconv.Atoi(fields[0])
	k, _ := strconv.Atoi(fields[1])
	s, _ := reader.ReadString('\n')
	arr := strings.Fields(s)
	input := make([]int, len(arr))
	sum := 0
	for i := 0; i < n; i++ {
		v, _ := strconv.Atoi(arr[i])
		sum += v
		input[i] = v
	}
	if n == k || k == 1 {
		fmt.Println(sum)
		return
	}
	maxV := -1
	var dfs func(preSum int, curSum int, index int, curK int)
	dfs = func(preSum int, curSum int, index int, curK int) {
		if index == len(input) {
			if curK == k {
				if preSum >= maxV {
					maxV = preSum
				}
			}
			return
		}
		dfs(preSum, curSum^input[index], index+1, curK)
		preSum += curSum
		dfs(preSum, 0, index+1, curK+1)
	}
	dfs(input[0], 0, 1, 0)
	fmt.Println(maxV)
}
