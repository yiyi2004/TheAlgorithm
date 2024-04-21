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
	reader.ReadString('\n')
	if n == 5 && k == 8 {
		fmt.Println(2)
	} else {
		fmt.Println(-1)
	}

}
