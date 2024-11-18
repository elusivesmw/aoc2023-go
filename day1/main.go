package main

import (
	"fmt"
	"os"
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
	var first int
	var last int
	first, last = getNums(line)

	return first*10 + last
}

func isNum(r rune) bool {
	return (r >= '0' && r <= '9')
}

var nums = [20]string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

type foundNum struct {
	index int
	value int
	text  string
}

func newFoundNum() foundNum {
	f := foundNum{index: -1, value: 0}
	return f
}

func getNums(line string) (int, int) {
	first := newFoundNum()
	last := newFoundNum()

	for i, num := range nums {
		fi := strings.Index(line, num)
		li := strings.LastIndex(line, num)

		if fi != -1 && (fi < first.index || first.index == -1) {
			first.index = fi
			first.value = i % 10
			//first.text = num
			//fmt.Println("first", first)
		}
		if li != -1 && li > last.index {
			last.index = li
			last.value = i % 10
			//last.text = num
			//fmt.Println("last", last)
		}
	}

	return first.value, last.value
}
