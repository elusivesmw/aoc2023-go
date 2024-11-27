package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}
	lexer := newLexer(string(data))
	fmt.Printf("lexer: %v\n", lexer)

	// get all tokens
	tokens := tokenize(&lexer)
	total := 0

	//fmt.Printf("tokens: %v\n", tokens)

	// add em up
	for _, token := range *tokens {
		printToken(&token)
		if token.tokenType == Symbol {
			continue
		}

		//line := token.pos / LineLen
		//fmt.Printf("%v = line %d\n", token, line)
	}
	fmt.Printf("total: %d\n", total)
}

const LineLen = 10 // just hardcode it
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

func hasAdjacentSymbol(tokens *map[int]token, star *token) bool {

	//fmt.Printf("all tokens: %v\n", tokens)

	for i, t := range *tokens {
		fmt.Printf("at index %d, t = %v\n", i, t)
		// dimiss other numbers
		if t.tokenType == Number {

			fmt.Println("  number")
			continue
		}

		// get row above
		for j := star.pos - LineLen - 1; j < star.pos-LineLen+star.len; j++ {
			fmt.Printf("  j (star based) = %d vs i = %d\n", j, i)
			if j == i {
				fmt.Println("    match")
				//if value, ok := tokens[i]

				return true
			}
		}

		// get left of
		ab := star.pos - 1
		fmt.Printf("  ab (star based) = %d vs i = %d\n", ab, i)
		if ab == i {
			fmt.Println("    match")
			return true
		}

		// get right of
		aa := star.pos + star.len
		if aa == i {
			fmt.Printf("  aa (star based) = %d vs i = %d\n", aa, i)
			fmt.Println("    match")
			return true
		}

		// get row below
		for j := star.pos + LineLen + 1; j < star.pos+LineLen+star.len; j++ {
			fmt.Printf("  j (star based) = %d vs i = %d\n", j, i)
			if j == i {
				fmt.Println("    match")
				return true
			}
		}
	}
	return false
}

func validIndex(i int, length int) bool {
	return i >= 0 && i < length
}

type lexer struct {
	runes []rune
	pos   int
}

func newLexer(input string) lexer {
	return lexer{
		runes: []rune(input),
		pos:   0,
	}
}

func peek(lexer *lexer) rune {
	return peekAt(lexer, 0)
}

func peekAt(lexer *lexer, offset int) rune {
	pos := lexer.pos + offset
	if pos >= len(lexer.runes) {
		return 0
	}
	return lexer.runes[pos]
}

func skipSpace(lexer *lexer) {
	p := peek(lexer)
	for p == '.' || p == '\n' {
		lexer.pos++
		p = peek(lexer)
	}
}

func printCurrentRune(lexer *lexer) {
	fmt.Printf("lexer.runes[%d] = %d = %s\n", lexer.pos, lexer.runes[lexer.pos], string(lexer.runes[lexer.pos]))
}

func tokenize(lexer *lexer) *map[int]token {
	tokens := make(map[int]token)

	for lexer.pos < len(lexer.runes) {
		skipSpace(lexer)
		r := peek(lexer)
		if r == 0 {
			break
		}

		key := lexer.pos
		current := token{
			pos: lexer.pos,
		}

		offset := 0
		if isNum(r) {
			current.tokenType = Number

			var num []rune
			next := peek(lexer)
			for isNum(next) {
				//printCurrentRune(lexer)
				num = append(num, next)

				lexer.pos++
				next = peekAt(lexer, offset)
			}
			current.len = len(num)
			current.chars = string(num)

			tokens[key] = current

		} else if isSymbol(r) {
			//printCurrentRune(lexer)
			current.tokenType = Symbol

			offset++
			sym := []rune(string(r))
			current.len = len(sym)
			current.chars = string(sym)

			tokens[key] = current

		}
		lexer.pos += offset
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
	len       int
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r rune) bool {
	return r != '.' && !isNum(r)
}

func printToken(token *token) {
	fmt.Printf("%s\tis a %s at %d (len of %d)\n", token.chars, token.tokenType, token.pos, token.len)
}
