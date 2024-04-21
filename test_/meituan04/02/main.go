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
	reader.ReadString('\n')
	inputStr, _ := reader.ReadString('\n')
	fields := strings.Fields(inputStr)
	input := make([]int, len(fields))
	maxV := -1
	for i := 0; i < len(input); i++ {
		val, _ := strconv.Atoi(fields[i])
		if val > maxV {
			maxV = val
		}
		input[i] = val
	}
	for i := 0; i < len(input); i++ {
		if input[i]*2 > maxV {
			fmt.Printf("%d ", input[i]*2)
		} else {
			fmt.Printf("%d ", maxV)
		}
	}
}
