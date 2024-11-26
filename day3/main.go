package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}
	lexer := newLexer(string(data))
	fmt.Printf("lexer: %v\n", lexer)
	tokens := tokenize(&lexer)
	for _, token := range tokens {
		printToken(&token)
	}
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

func tokenize(lexer *lexer) []token {
	var tokens []token

	for lexer.pos < len(lexer.runes) {
		skipSpace(lexer)
		r := peek(lexer)
		if r == 0 {
			break
		}

		current := token{
			pos: lexer.pos,
		}

		offset := 0
		if isNum(r) {
			current.tokenType = Number

			var num []rune
			next := peek(lexer)
			for isNum(next) {
				printCurrentRune(lexer)
				num = append(num, next)

				offset++
				next = peekAt(lexer, offset)
			}
			current.len = len(num)
			current.chars = string(num)
			fmt.Printf("current N: %v\n", current.chars)

			tokens = append(tokens, current)

		} else if isSymbol(r) {
			printCurrentRune(lexer)
			current.tokenType = Symbol

			offset++
			sym := []rune(string(r))
			current.len = len(sym)
			current.chars = string(sym)
			fmt.Printf("current S: %v\n", current.chars)

			tokens = append(tokens, current)

		}
		lexer.pos += offset
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
