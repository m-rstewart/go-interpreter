package lexer

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
