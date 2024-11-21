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

	input := string(data)
	print(input)

	var games []game
	var total int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		game := parseLine(line)
		if isValid(&game) {
			total += game.num
		}
		games = append(games, game)
	}

	fmt.Printf("total: %d\n", total)
}

type game struct {
	num   int
	grabs []grab
}

type grab struct {
	red   int
	green int
	blue  int
}

var maxGrab = grab{red: 12, green: 13, blue: 14}

func isValid(game *game) bool {
	for _, grab := range game.grabs {
		if grab.red > maxGrab.red || grab.green > maxGrab.green || grab.blue > maxGrab.blue {
			return false
		}
	}
	return true
}

func parseLine(input string) game {
	var game game
	gg := strings.Split(input, ":")
	if len(gg) < 2 {
		println("invalid game input")
		return game
	}
	game.num = parseGame(strings.TrimSpace(gg[0]))
	game.grabs = parseGrabs(strings.TrimSpace(gg[1]))
	fmt.Printf("game: %v\n", game)

	return game
}

func parseGame(input string) int {
	gs := strings.Split(input, " ")
	if len(gs) < 2 {
		println("game < 2")
		return 0
	}
	gameNum, err := strconv.Atoi(gs[1])
	if err != nil {
		println(err)
		return 0
	}
	return gameNum
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
	return grab
}
