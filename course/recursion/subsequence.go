package recursion

import (
	"fmt"
)

func SubSequence(str string) {
	var res []string
	var path string
	process(0, str, res, path)
}

func process(local int, str string, res []string, path string) {
	if local == len(str) {
		fmt.Println(path)
		res = append(res, path)
		return
	}

	process(local+1, str, res, path)
	process(local+1, str, res, path+string(str[local]))
}
