package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const LineLen = 140 // just hardcode it

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	fmt.Printf("input:\n%s\n", input)

	// get all tokens
	tokens := tokenize(&input)
	total := 0

	// print all tokens
	printTokens(tokens)

	// add em up
	for i, token := range tokens {
		printToken(token, i)
		ratio := getGearRatio(tokens, token)
		if ratio > 0 {
			fmt.Printf("**%v is valid**\n\n", &token)
		}
		total += ratio
	}
	fmt.Printf("total: %d\n", total)
}

func printTokens(tokens map[int]*token) {
	fmt.Println("tokens:")
	for i, token := range tokens {
		printToken(token, i)
	}
	fmt.Println()
}

func validate(tokens map[int]*token, surrounding int) bool {
	value, ok := tokens[surrounding]
	if !ok {
		//fmt.Printf("\t\tno token at %d\n", surrounding)
		return false
	}

	if value.tokenType != Number {
		//fmt.Printf("\t\ttoken is not number: %v\n", value.tokenType)
		return false
	}

	return true
}

// NOTE: this function is flawed in that when getting relative token,
// it will move up or down a line if center token is at line index 0
// or line index LineLen-1.
// returns: the gear ratio, else 0 when 2 adjacent numbers are not found
func getGearRatio(tokens map[int]*token, center *token) int {

	//fmt.Printf("all tokens: %v\n", tokens)

	//fmt.Printf("at index %d, center = %v\n", center.pos, string(center.chars))
	// only look at symbols
	if center.tokenType != Symbol && center.chars != "*" {
		//fmt.Println("\tcenter not a * symbol")
		return 0
	}

	// hold adjacent numbers
	var adjacents []*token

	// get row above
	for i := center.pos - LineLen - 1; i < center.pos-LineLen+len(center.chars)+1; i++ {
		//fmt.Printf("\ti (surrounding based) = %d vs center = %d\n", i, center.pos)
		if !validate(tokens, i) {
			continue
		}
		//fmt.Printf("\t\tmatch at i = %d\n", i)
		found := tokens[i]
		adjacents = appendWithoutDupes(adjacents, found)
	}

	// get left of
	lo := center.pos - 1
	//fmt.Printf("\tlo (surrounding based) = %d vs center = %d\n", lo, center.pos)
	if validate(tokens, lo) {
		//fmt.Printf("\t\tmatch at lo = %d\n", lo)
		found := tokens[lo]
		adjacents = appendWithoutDupes(adjacents, found)
	}

	// get right of
	ro := center.pos + len(center.chars)
	//fmt.Printf("\tro (surrounding based) = %d vs i = %d\n", ro, center.pos)
	if validate(tokens, ro) {
		//fmt.Printf("\t\tmatch at ro = %d\n", ro)
		found := tokens[ro]
		adjacents = appendWithoutDupes(adjacents, found)
	}

	// get row below
	for i := center.pos + LineLen - 1; i < center.pos+LineLen+len(center.chars)+1; i++ {
		//fmt.Printf("\ti (surrounding based) = %d vs i = %d\n", i, center.pos)
		if !validate(tokens, i) {
			continue
		}
		//fmt.Printf("\t\tmatch at i = %d\n", i)
		found := tokens[i]
		adjacents = appendWithoutDupes(adjacents, found)
	}

	// discard *'s with more or less than 2 adjacents
	adjacentsCount := len(adjacents)
	//fmt.Printf("\tadjacentsCount: %d\n", adjacentsCount)
	if adjacentsCount != 2 {
		return 0
	}

	//fmt.Printf("adjacents: %v\n", adjacents)
	total := 1
	for _, a := range adjacents {
		value, err := strconv.Atoi(a.chars)
		fmt.Printf("\tgear value: %d\n", value)
		if err != nil {
			panic(err)
		}
		total *= value
	}
	fmt.Printf("gear total: %d\n", total)

	return total
}

func appendWithoutDupes(adjacents []*token, token *token) []*token {
	//fmt.Printf("adjacents: %v\n", adjacents)
	//fmt.Printf("token: %v\n", token)
	for _, v := range adjacents {
		if v == token {
			return adjacents
		}
	}
	return append(adjacents, token)
}

func validIndex(i int, length int) bool {
	return i >= 0 && i < length
}

func peek(line string, pos int) rune {
	return peekAt(line, pos, 0)
}

func peekAt(line string, pos int, offset int) rune {
	next := pos + offset
	if next >= len(line) {
		return 0
	}
	return []rune(line)[next]
}

func printCurrentRune(r rune) {
	fmt.Printf("r: %s\n", string(r))
}

func tokenize(input *string) map[int]*token {
	tokens := make(map[int]*token)

	lines := strings.Split(*input, "\n")
	for i, line := range lines {
		current := &token{}

		j := 0
		for j < len(line) {
			r := peek(line, j)
			absPos := i*LineLen + j
			fmt.Printf("pos: %d\n", absPos)

			offset := 0
			if isNum(r) {
				current.tokenType = Number
				current.pos = absPos

				var num []rune
				next := peek(line, j)
				for isNum(next) {
					//printCurrentRune(next)
					num = append(num, next)
					tokens[absPos+offset] = current
					//fmt.Printf("tokens[%d+%d] = %v\n", absPos, offset, current)
					//fmt.Printf("num in progress: %s, at: %d\n", string(num), current.pos)
					//printTokens(tokens)

					offset++
					next = peekAt(line, j, offset)
				}
				current.chars = string(num)
				fmt.Printf("current.chars: %s\n", current.chars)

				current = &token{}
			} else if isSymbol(r) {
				//printCurrentRune(r)
				current.tokenType = Symbol
				current.pos = absPos

				offset++
				sym := []rune(string(r))
				current.chars = string(sym)

				tokens[absPos] = current
				current = &token{}
			} else {
				// '.'
				offset++
			}
			//fmt.Printf("offset here is %d\n", offset)
			j += offset
		}
		i++

	}

	return tokens
}

type TokenType string

const (
	Number TokenType = "number"
	Symbol           = "symbol"
)

type token struct {
	tokenType TokenType
	chars     string
	pos       int
}

func isNum(r rune) bool {
	//fmt.Printf("r = %s is num \n", string(r))
	return r >= '0' && r <= '9'
}

func isSymbol(r rune) bool {
	//fmt.Printf("r = %s is sym \n", string(r))
	return r != '.' && !isNum(r)
}

func printToken(token *token, i int) {
	fmt.Printf("%s\tis a %s at %d(%d) (len of %d)\n", token.chars, token.tokenType, i, token.pos, len(token.chars))
}
