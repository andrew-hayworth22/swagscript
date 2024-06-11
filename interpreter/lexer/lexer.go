package lexer

import (
	"swagscript/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.char {
	case '=':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = newToken(token.ASSIGN, lexer.char)
		}
	case '+':
		tok = newToken(token.PLUS, lexer.char)
	case ',':
		tok = newToken(token.COMMA, lexer.char)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.char)
	case '(':
		tok = newToken(token.LPAREN, lexer.char)
	case ')':
		tok = newToken(token.RPAREN, lexer.char)
	case '{':
		tok = newToken(token.LBRACE, lexer.char)
	case '}':
		tok = newToken(token.RBRACE, lexer.char)
	case '-':
		tok = newToken(token.MINUS, lexer.char)
	case '!':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(token.BANG, lexer.char)
		}
	case '*':
		tok = newToken(token.ASTERISK, lexer.char)
	case '/':
		tok = newToken(token.SLASH, lexer.char)
	case '<':
		tok = newToken(token.LT, lexer.char)
	case '>':
		tok = newToken(token.GT, lexer.char)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(lexer.char) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookUpKeyword(tok.Literal)
			return tok
		} else if isDigit(lexer.char) {
			tok.Literal = lexer.readInteger()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readIdentifier() string {
	start_position := lexer.position

	for isLetter(lexer.char) {
		lexer.readChar()
	}

	return lexer.input[start_position:lexer.position]
}

func (lexer *Lexer) readInteger() string {
	start_position := lexer.position

	for isDigit(lexer.char) {
		lexer.readChar()
	}

	return lexer.input[start_position:lexer.position]
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\n' || lexer.char == '\t' || lexer.char == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}