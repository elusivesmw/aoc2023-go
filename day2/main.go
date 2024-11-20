package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	print(input)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parseLine(line)
	}
}

type game struct {
	game  int
	grabs []grab
}

type grab struct {
	red   int
	green int
	blue  int
}

func parseLine(input string) {

	//var game game
	gg := strings.Split(input, ":")
	if len(gg) < 2 {
		return
	}
	_ = parseGame(strings.TrimSpace(gg[0]))
	_ = parseGrabs(strings.TrimSpace(gg[1]))

}

func parseGame(input string) int {
	println(input)
	return 0
}

func parseGrabs(input string) []grab {
	grabs := make([]grab, 0)
	gs := strings.Split(input, ";")
	for _, g := range gs {
		grab := parseGrab(strings.TrimSpace(g))
		grabs = append(grabs, grab)
	}

	return grabs
}

func parseGrab(input string) grab {
	grab := grab{}

	gs := strings.Split(input, ",")
	for _, g := range gs {
		cube := strings.Split(strings.TrimSpace(g), " ")
		if len(cube) < 2 {
			println("cube < 2")
			continue
		}
		count, err := strconv.Atoi(cube[0])
		if err != nil {
			println(err)
			continue
		}
		color := cube[1]

		switch color {
		case "red":
			grab.red = count
			break
		case "green":
			grab.green = count
			break
		case "blue":
			grab.blue = count
			break
		}
	}
	fmt.Printf("grab: %v\n", grab)
	return grab
}
