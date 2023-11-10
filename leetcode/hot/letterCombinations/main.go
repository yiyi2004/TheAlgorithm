package main

func main() {

}

var combinations []string

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combinations = []string{}
	process(digits, 0, "")
	return combinations
}

func process(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
		return
	}
	chars := phoneMap[string(digits[index])]
	for i := 0; i < len(chars); i++ {
		process(digits, index+1, combination+string(chars[i]))
	}
}
