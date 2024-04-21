package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const Mod = 1000000007

func main() {
	//sum := 2 + 4 + 12 + 40000000000
	//fmt.Println(sum % Mod)

	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	fields := strings.Fields(readString)
	n, _ := strconv.Atoi(fields[0])
	q, _ := strconv.Atoi(fields[1])

	input := make([]int, n)
	inputStr, _ := reader.ReadString('\n')
	fieldsArr := strings.Fields(inputStr)

	nums := make([]int, n)
	for i := 0; i < len(fieldsArr); i++ {
		v, _ := strconv.Atoi(fieldsArr[i])
		input[i] = v
		nums[i] = q
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Scan(&x)
		nums[x-1]--
	}

	var sum int
	for i := 0; i < len(input); i++ {
		val := (input[i] * int(math.Pow(2.0, float64(nums[i])))) % Mod
		sum += val
		sum = sum % Mod
	}
	fmt.Println(sum)
}
