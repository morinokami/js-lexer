package lexer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/morinokami/js-lexer/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// EOF
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	l.column += 1
}

func (l *Lexer) peekChar(n int) byte {
	if l.readPosition+n >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition+n]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (string, error) {
	if l.ch == '0' {
		next := l.peekChar(0)
		if next == 'B' || next == 'b' {
			return l.readBinaryNumber()
		} else if next == 'X' || next == 'x' {
			return l.readHexadecimalNumber()
		}
		//else {
		//	return l.readOctalNumber()
		//}
	}

	position := l.position
	readDecimalPoint := false
	for isDigit(l.ch) || (l.ch == '.' && !readDecimalPoint) {
		if l.ch == '.' {
			readDecimalPoint = true
		}
		l.readChar()
	}
	return l.input[position:l.position], nil
}

func (l *Lexer) readBinaryNumber() (string, error) {
	position := l.position
	l.readChar() // skip '0'
	l.readChar() // skip 'b' or 'B'

	if l.ch != '0' && l.ch != '1' {
		return "", errors.New(fmt.Sprintf("SyntaxError: Expected number in radix 2 (%d:%d)", l.line, l.column-1))
	}

	for l.ch == '0' || l.ch == '1' {
		l.readChar()
	}
	return l.input[position:l.position], nil
}

func (l *Lexer) readOctalNumber() (string, error) {
	// TODO
	return "", nil
}

func (l *Lexer) readHexadecimalNumber() (string, error) {
	position := l.position
	l.readChar()
	l.readChar()

	if !isHexChar(l.ch) {
		return "", errors.New(fmt.Sprintf("SyntaxError: Expected number in radix 16 (%d:%d)", l.line, l.column-1))
	}

	for isHexChar(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position], nil
}

func (l *Lexer) readString(quote byte) (string, error) {
	literalStart := l.column - 1
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == quote {
			break
		} else if l.ch == '\\' && l.peekChar(0) == quote {
			l.readChar()
		} else if l.ch == 0 || l.ch == '\n' {
			return "", fmt.Errorf("SyntaxError: Unterminated string constant (%d:%d)", l.line, literalStart)
		}
	}
	escapedQuote := fmt.Sprintf("\\%s", string(quote))
	return strings.ReplaceAll(l.input[position:l.position], escapedQuote, string(quote)), nil
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '$'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isHexChar(ch byte) bool {
	return isDigit(ch) || ('a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F')
}

func (l *Lexer) skipWhitespace() {
	for {
		if l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
			if l.ch == '\n' {
				l.line += 1
				l.column = 0
			}
			l.readChar()
			continue
		} else if l.ch == '/' && l.peekChar(0) == '/' {
			l.skipSingleLineComment()
			continue
		} else if l.ch == '/' && l.peekChar(0) == '*' {
			l.skipMultiLineComment()
			continue
		}
		break
	}
}

func (l *Lexer) skipSingleLineComment() {
	for {
		if l.ch == '\n' || l.ch == 0 {
			break
		}
		l.readChar()
	}
}

func (l *Lexer) skipMultiLineComment() {
	l.readChar()
	l.readChar()
	for {
		if l.ch == '*' && l.peekChar(0) == '/' || l.ch == 0 {
			l.readChar()
			l.readChar()
			break
		}
		l.readChar()
	}
}

func (l *Lexer) makeSourceLocation(lineStart, colStart, adjustment int) token.SourceLocation {
	return token.SourceLocation{
		Start: token.Position{
			Line:   lineStart,
			Column: colStart,
		},
		End: token.Position{
			Line:   l.line,
			Column: l.column + adjustment,
		},
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	lineStart := l.line
	colStart := l.column - 1

	switch l.ch {

	// Punctuators
	case '(':
		tok = newToken(token.LParen, l.ch)
	case ')':
		tok = newToken(token.RParen, l.ch)
	case '{':
		tok = newToken(token.LBrace, l.ch)
	case '}':
		tok = newToken(token.RBrace, l.ch)
	case '[':
		tok = newToken(token.LBracket, l.ch)
	case ']':
		tok = newToken(token.RBracket, l.ch)
	case '.':
		if l.peekChar(0) == '.' && l.peekChar(1) == '.' {
			// Spread syntax
			tok = makeMultiCharToken(l, token.Ellipsis, 2)
		} else if isDigit(l.peekChar(0)) {
			tok.Type = token.TokenType{Label: token.Numeric}
			var err error
			tok.Literal, err = l.readNumber()
			if err != nil {
				panic(err.Error())
			}
			tok.Loc = l.makeSourceLocation(lineStart, colStart, -1)
			return tok
		} else {
			tok = newToken(token.Dot, l.ch)
		}
	case ';':
		tok = newToken(token.Semicolon, l.ch)
	case ':':
		tok = newToken(token.Colon, l.ch)
	case ',':
		tok = newToken(token.Comma, l.ch)
	case '?':
		if l.peekChar(0) == '?' {
			// Nullish coalescing
			tok = makeMultiCharToken(l, token.NullishCoalescing, 1)
		} else if l.peekChar(0) == '.' {
			// Optional chaining
			tok = makeMultiCharToken(l, token.OptionalChaining, 1)
		} else {
			tok = newToken(token.Question, l.ch)
		}

	// Operators
	case '<':
		if l.peekChar(0) == '<' && l.peekChar(1) == '=' {
			// Left shift assignment
			tok = makeMultiCharToken(l, token.LeftShiftAssignment, 2)
		} else if l.peekChar(0) == '<' {
			// Left shift
			tok = makeMultiCharToken(l, token.LeftShift, 1)
		} else if l.peekChar(0) == '=' {
			// Less than or equal
			tok = makeMultiCharToken(l, token.LTEq, 1)
		} else {
			// Less than
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar(0) == '>' && l.peekChar(1) == '>' && l.peekChar(2) == '=' {
			// Unsigned right shift assignment
			tok = makeMultiCharToken(l, token.UnsignedRightShiftAssignment, 3)
		} else if l.peekChar(0) == '>' && l.peekChar(1) == '>' {
			// Unsigned right shift
			tok = makeMultiCharToken(l, token.UnsignedRightShift, 2)
		} else if l.peekChar(0) == '>' && l.peekChar(1) == '=' {
			// Right shift assignment
			tok = makeMultiCharToken(l, token.RightShiftAssignment, 2)
		} else if l.peekChar(0) == '>' {
			// Right shift
			tok = makeMultiCharToken(l, token.RightShift, 1)
		} else if l.peekChar(0) == '=' {
			// Greater than or equal
			tok = makeMultiCharToken(l, token.GTEq, 1)
		} else {
			// Greater than
			tok = newToken(token.GT, l.ch)
		}
	case '=':
		if l.peekChar(0) == '=' && l.peekChar(1) == '=' {
			// Identity
			tok = makeMultiCharToken(l, token.Identity, 2)
		} else if l.peekChar(0) == '=' {
			// Equality
			tok = makeMultiCharToken(l, token.Equality, 1)
		} else if l.peekChar(0) == '>' {
			// Arrow
			tok = makeMultiCharToken(l, token.Arrow, 1)
		} else {
			// Assignment
			tok = newToken(token.Assignment, l.ch)
		}
	case '!':
		if l.peekChar(0) == '=' && l.peekChar(1) == '=' {
			// Nonidentity
			tok = makeMultiCharToken(l, token.Nonidentity, 2)
		} else if l.peekChar(0) == '=' {
			// Inequality
			tok = makeMultiCharToken(l, token.Inequality, 1)
		} else {
			// Logical NOT
			tok = newToken(token.Bang, l.ch)
		}
	case '+':
		if l.peekChar(0) == '+' {
			// Increment
			tok = makeMultiCharToken(l, token.Increment, 1)
		} else if l.peekChar(0) == '=' {
			// Addition assignment
			tok = makeMultiCharToken(l, token.AdditionAssignment, 1)
		} else {
			// Addition
			tok = newToken(token.Plus, l.ch)
		}
	case '-':
		if l.peekChar(0) == '-' {
			// Decrement
			tok = makeMultiCharToken(l, token.Decrement, 1)
		} else if l.peekChar(0) == '=' {
			// Subtraction assignment
			tok = makeMultiCharToken(l, token.SubtractionAssignment, 1)
		} else {
			// Subtraction
			tok = newToken(token.Minus, l.ch)
		}
	case '*':
		if l.peekChar(0) == '*' && l.peekChar(1) == '=' {
			// Exponentiation assignment
			tok = makeMultiCharToken(l, token.ExponentiationAssignment, 2)
		} else if l.peekChar(0) == '*' {
			// Exponentiation
			tok = makeMultiCharToken(l, token.Exponentiation, 1)
		} else if l.peekChar(0) == '=' {
			// Multiplication assignment
			tok = makeMultiCharToken(l, token.MultiplicationAssignment, 1)
		} else {
			// Multiplication
			tok = newToken(token.Star, l.ch)
		}
	case '/':
		if l.peekChar(0) == '=' {
			// Division assignment
			tok = makeMultiCharToken(l, token.DivisionAssignment, 1)
		} else {
			// Division
			tok = newToken(token.Slash, l.ch)
		}
	case '%':
		if l.peekChar(0) == '=' {
			// Remainder assignment
			tok = makeMultiCharToken(l, token.RemainderAssignment, 1)
		} else {
			// Remainder
			tok = newToken(token.Remainder, l.ch)
		}
	case '&':
		if l.peekChar(0) == '&' {
			// Logical AND
			tok = makeMultiCharToken(l, token.LogicalAnd, 1)
		} else if l.peekChar(0) == '=' {
			// Bitwise AND assignment
			tok = makeMultiCharToken(l, token.BitwiseAndAssignment, 1)
		} else {
			// Bitwise AND
			tok = newToken(token.BitwiseAnd, l.ch)
		}
	case '|':
		if l.peekChar(0) == '|' {
			// Logical OR
			tok = makeMultiCharToken(l, token.LogicalOr, 1)
		} else if l.peekChar(0) == '=' {
			// Bitwise OR assignment
			tok = makeMultiCharToken(l, token.BitwiseOrAssignment, 1)
		} else {
			// Bitwise OR
			tok = newToken(token.BitwiseOr, l.ch)
		}
	case '^':
		if l.peekChar(0) == '=' {
			tok = makeMultiCharToken(l, token.BitwiseXorAssignment, 1)
		} else {
			// Bitwise XOR
			tok = newToken(token.BitwiseXor, l.ch)
		}
	case '~':
		// Bitwise NOT
		tok = newToken(token.Tilde, l.ch)

	// Literals
	case '"', '\'':
		// String
		tok.Type = token.TokenType{Label: token.String}
		var err error
		tok.Literal, err = l.readString(l.ch)
		if err != nil {
			panic(err.Error())
		}
	// TODO: Other numeric literals, regex, template literal, ...

	// EOF
	case 0:
		tok.Type = token.TokenType{Label: token.EOF}
		tok.Literal = ""
		tok.Loc = l.makeSourceLocation(lineStart, colStart, -1)
		return tok

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Loc = l.makeSourceLocation(lineStart, colStart, -1)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.TokenType{Label: token.Numeric}
			var err error
			tok.Literal, err = l.readNumber()
			if err != nil {
				panic(err.Error())
			}
			tok.Loc = l.makeSourceLocation(lineStart, colStart, -1)
			return tok
		} else {
			panic(fmt.Sprintf("SyntaxError: Unexpected character '%s' (%d:%d)",
				string(l.ch), lineStart, colStart))
		}
	}

	tok.Loc = l.makeSourceLocation(lineStart, colStart, 0)

	l.readChar()

	return tok
}

func newToken(label string, ch byte) token.Token {
	return token.Token{
		Type:    token.TokenType{Label: label},
		Literal: string(ch),
	}
}

func makeMultiCharToken(l *Lexer, label string, n int) token.Token {
	literal := ""
	for i := 0; i < n; i++ {
		literal += string(l.ch)
		l.readChar()
	}
	literal += string(l.ch)

	return token.Token{Type: token.TokenType{Label: label}, Literal: literal}
}
