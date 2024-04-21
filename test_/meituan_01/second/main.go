package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int64
	var n, k int
	var countZero int64
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fields := strings.Fields(text)
	n, _ = strconv.Atoi(fields[0])
	k, _ = strconv.Atoi(fields[1])

	arrStr, _ := reader.ReadString('\n')
	splitArr := strings.Fields(arrStr)
	if len(splitArr) != n {
		return
	}

	for _, v := range splitArr {
		val, _ := strconv.Atoi(v)
		if val == 0 {
			countZero++
		}
		sum += int64(val)
	}

	for i := 0; i < k; i++ {
		text, _ = reader.ReadString('\n')
		split := strings.Fields(text)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		fmt.Println(sum+int64(l)*countZero, sum+int64(r)*countZero)
	}
}
