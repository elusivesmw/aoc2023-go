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
	fmt.Printf("input:\n%s\n", input)

	// get all tokens
	tokens := tokenize(&input)
	total := 0

	// print all tokens
	//fmt.Printf("tokens: %v\n", tokens)
	for _, token := range *tokens {
		printToken(&token)
	}
	fmt.Println()

	// add em up
	for _, token := range *tokens {
		printToken(&token)
		num := validNum(tokens, &token)
		if num > 0 {
			fmt.Printf("**%v is valid**\n\n", &token)
		}
		total += num
		//line := token.pos / LineLen
		//fmt.Printf("%v = line %d\n", token, line)
	}
	fmt.Printf("total: %d\n", total)
}

const LineLen = 140 // just hardcode it
func validNum(tokens *map[int]token, token *token) int {
	if hasAdjacentSymbol(tokens, token) {
		n, err := strconv.Atoi(token.chars)
		if err != nil {
			panic("out of the range")
		}
		return n
	}
	return 0
}

func isSymbolToken(token *token) bool {
	return token.tokenType == Symbol
}

func validate(tokens *map[int]token, surrounding int) bool {
	value, ok := (*tokens)[surrounding]
	if !ok {
		fmt.Printf("\t\tno token at %d\n", surrounding)
		return false
	}

	if !isSymbolToken(&value) {
		fmt.Printf("\t\ttoken is not symbol: %v\n", value.tokenType)
		return false
	}

	return true
}

func hasAdjacentSymbol(tokens *map[int]token, center *token) bool {

	//fmt.Printf("all tokens: %v\n", tokens)

	//fmt.Printf("at index %d, center = %v\n", center.pos, string(center.chars))
	// only look at numbers
	if center.tokenType != Number {
		fmt.Println("\tcenter not a number")
		return false
	}

	// get row above
	for i := center.pos - LineLen - 1; i < center.pos-LineLen+len(center.chars)+1; i++ {
		fmt.Printf("\ti (surrounding based) = %d vs center = %d\n", i, center.pos)
		if !validate(tokens, i) {
			continue
		}
		fmt.Printf("\t\tmatch at i = %d\n", i)
		return true
	}

	// get left of
	lo := center.pos - 1
	fmt.Printf("\tlo (surrounding based) = %d vs center = %d\n", lo, center.pos)
	if validate(tokens, lo) {
		fmt.Printf("\t\tmatch at lo = %d\n", lo)
		return true
	}

	// get right of
	ro := center.pos + len(center.chars)
	fmt.Printf("\tro (surrounding based) = %d vs i = %d\n", ro, center.pos)
	if validate(tokens, ro) {
		fmt.Printf("\t\tmatch at ro = %d\n", ro)
		return true
	}

	// get row below
	for i := center.pos + LineLen - 1; i < center.pos+LineLen+len(center.chars)+1; i++ {
		fmt.Printf("\ti (surrounding based) = %d vs i = %d\n", i, center.pos)
		if !validate(tokens, i) {
			continue
		}
		fmt.Printf("\t\tmatch at i = %d\n", i)
		return true
	}
	return false
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

func tokenize(input *string) *map[int]token {
	tokens := make(map[int]token)

	lines := strings.Split(*input, "\n")
	for i, line := range lines {
		current := token{}

		j := 0
		for j < len(line) {
			r := peek(line, j)
			absPos := i*LineLen + j
			//fmt.Printf("pos: %d\n", absPos)

			offset := 0
			if isNum(r) {
				current.tokenType = Number
				current.pos = absPos

				var num []rune
				next := peek(line, j)
				for isNum(next) {
					//printCurrentRune(next)
					num = append(num, next)
					//fmt.Printf("num in progress: %s\n", string(num))

					offset++
					next = peekAt(line, j, offset)
				}
				current.chars = string(num)

				tokens[absPos] = current
				current = token{}
			} else if isSymbol(r) {
				//printCurrentRune(r)
				current.tokenType = Symbol
				current.pos = absPos

				offset++
				sym := []rune(string(r))
				current.chars = string(sym)

				tokens[absPos] = current
				current = token{}
			} else {
				// '.'
				offset++
			}
			//fmt.Printf("offset here is %d\n", offset)
			j += offset
		}
		i++

	}

	return &tokens
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

func printToken(token *token) {
	fmt.Printf("%s\tis a %s at %d (len of %d)\n", token.chars, token.tokenType, token.pos, len(token.chars))
}
