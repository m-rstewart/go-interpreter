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

	lex.skipWhitespace()

	switch lex.ch {
	case '=':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			literal := string(ch) + string(lex.ch)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.ASSIGN, lex.ch)
		}
	case '/':
		tok = newToken(token.SLASH, lex.ch)
	case '*':
		tok = newToken(token.ASTERISK, lex.ch)
	case '<':
		tok = newToken(token.LT, lex.ch)
	case '>':
		tok = newToken(token.GT, lex.ch)
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '-':
		tok = newToken(token.MINUS, lex.ch)
	case '!':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			literal := string(ch) + string(lex.ch)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, lex.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.ch) {
			// could either be a keyword (fn, let) or a identifier (foo, bar) at this point
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lex.ch) {
			tok.Type = token.INT
			tok.Literal = lex.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.ch)
		}
	}
	lex.readChar()
	return tok
}

func (lex *Lexer) readIdentifier() string {
	position := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func (lex *Lexer) readNumber() string {
	position := lex.position
	for isDigit(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

// if next char is whitespace, advance the pointers until a valid char is being read
func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

func (lex *Lexer) peekChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	} else {
		return lex.input[lex.readPosition]
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
