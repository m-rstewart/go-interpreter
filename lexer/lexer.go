package lexer

import "github.com/m-rstewart/go-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	lex := &Lexer{
		input: input,
	}
	lex.readChar()
	return lex
}

// Give us the next character and advance our position in the input string
// Only supports ASCII for now
func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		// terminate
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

// Gives us the token of the current character under examination.
// Advances the lexer pointers before returning the token.
func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lex.ch {
	case '=':
		tok = newToken(token.ASSIGN, lex.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	lex.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
