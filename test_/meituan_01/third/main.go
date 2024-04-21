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
	mt, _ := reader.ReadString('\n')
	countMT := 0
	for _, v := range mt {
		if v == 'M' || v == 'T' {
			countMT++
		}
	}
	if countMT+k >= n {
		fmt.Println(n)
		return
	}
	fmt.Println(countMT + k)
}
