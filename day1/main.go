package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	var total int
	for _, line := range lines {
		fmt.Println(line)
		num := parseLine(line)
		fmt.Println(num)
		total += num
	}

	fmt.Println(total)
}

func parseLine(line string) int {
	var first rune
	var last rune
	for i, fr := range line {
		li := len(line) - 1 - i
		lr := rune(line[li])

		if first == 0 && isNum(fr) {
			first = fr
		}
		if last == 0 && isNum(lr) {
			last = lr
		}
	}
	str := fmt.Sprintf("%s%s", string(first), string(last))

	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func isNum(r rune) bool {
	return (r >= '0' && r <= '9')
}
