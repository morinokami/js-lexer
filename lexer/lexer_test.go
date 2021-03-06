package lexer

import (
	"testing"

	"github.com/morinokami/js-lexer/token"
)

func makeTT(label string) token.TokenType {
	return token.TokenType{Label: label}
}

func makeLoc(line0, col0, line1, col1 int) token.SourceLocation {
	return token.SourceLocation{
		Start: token.Position{
			Line:   line0,
			Column: col0,
		},
		End: token.Position{
			Line:   line1,
			Column: col1,
		},
	}
}

func TestComment(t *testing.T) {
	input := `
// single line comment
const x = 123;

/*
 * multi line comment
 */
/**/
// /**/
/* // */
function
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{makeTT(token.Const), "const"},
		{makeTT(token.Identifier), "x"},
		{makeTT(token.Assignment), "="},
		{makeTT(token.Numeric), "123"},
		{makeTT(token.Semicolon), ";"},
		{makeTT(token.Function), "function"},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestIdentifier(t *testing.T) {
	input := `
x
hello
abc123
$
$height9
_
_x
_$
$_
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{makeTT(token.Identifier), "x"},
		{makeTT(token.Identifier), "hello"},
		{makeTT(token.Identifier), "abc123"},
		{makeTT(token.Identifier), "$"},
		{makeTT(token.Identifier), "$height9"},
		{makeTT(token.Identifier), "_"},
		{makeTT(token.Identifier), "_x"},
		{makeTT(token.Identifier), "_$"},
		{makeTT(token.Identifier), "$_"},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestKeyword(t *testing.T) {
	input := `
await
break
case
catch
class
const
continue
debugger
default
delete
do
else
enum
export
extends
false
finally
for
function
if
import
in
instanceof
new
null
return
super
switch
this
throw
true
try
typeof
var
void
while
with
yield
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{makeTT(token.Await), "await"},
		{makeTT(token.Break), "break"},
		{makeTT(token.Case), "case"},
		{makeTT(token.Catch), "catch"},
		{makeTT(token.Class), "class"},
		{makeTT(token.Const), "const"},
		{makeTT(token.Continue), "continue"},
		{makeTT(token.Debugger), "debugger"},
		{makeTT(token.Default), "default"},
		{makeTT(token.Delete), "delete"},
		{makeTT(token.Do), "do"},
		{makeTT(token.Else), "else"},
		{makeTT(token.Enum), "enum"},
		{makeTT(token.Export), "export"},
		{makeTT(token.Extends), "extends"},
		{makeTT(token.False), "false"},
		{makeTT(token.Finally), "finally"},
		{makeTT(token.For), "for"},
		{makeTT(token.Function), "function"},
		{makeTT(token.If), "if"},
		{makeTT(token.Import), "import"},
		{makeTT(token.In), "in"},
		{makeTT(token.Instanceof), "instanceof"},
		{makeTT(token.New), "new"},
		{makeTT(token.Null), "null"},
		{makeTT(token.Return), "return"},
		{makeTT(token.Super), "super"},
		{makeTT(token.Switch), "switch"},
		{makeTT(token.This), "this"},
		{makeTT(token.Throw), "throw"},
		{makeTT(token.True), "true"},
		{makeTT(token.Try), "try"},
		{makeTT(token.Typeof), "typeof"},
		{makeTT(token.Var), "var"},
		{makeTT(token.Void), "void"},
		{makeTT(token.While), "while"},
		{makeTT(token.With), "with"},
		{makeTT(token.Yield), "yield"},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestPunctuator(t *testing.T) {
	input := `
(
)
{
}
[
]
.
...
;
:
,
?
?.
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{makeTT(token.LParen), "("},
		{makeTT(token.RParen), ")"},
		{makeTT(token.LBrace), "{"},
		{makeTT(token.RBrace), "}"},
		{makeTT(token.LBracket), "["},
		{makeTT(token.RBracket), "]"},
		{makeTT(token.Dot), "."},
		{makeTT(token.Ellipsis), "..."},
		{makeTT(token.Semicolon), ";"},
		{makeTT(token.Colon), ":"},
		{makeTT(token.Comma), ","},
		{makeTT(token.Question), "?"},
		{makeTT(token.OptionalChaining), "?."},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestOperator(t *testing.T) {
	input := `
<
>
<=
>=
==
!=
===
!==
+
-
*
/
%
++
--
**
<<
>>
>>>
&
|
^
!
~
&&
||
??
=
+=
-=
*=
/=
%=
**=
<<=
>>=
>>>=
&=
|=
^=
=>
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{makeTT(token.LT), "<"},
		{makeTT(token.GT), ">"},
		{makeTT(token.LTEq), "<="},
		{makeTT(token.GTEq), ">="},
		{makeTT(token.Equality), "=="},
		{makeTT(token.Inequality), "!="},
		{makeTT(token.Identity), "==="},
		{makeTT(token.Nonidentity), "!=="},
		{makeTT(token.Plus), "+"},
		{makeTT(token.Minus), "-"},
		{makeTT(token.Star), "*"},
		{makeTT(token.Slash), "/"},
		{makeTT(token.Remainder), "%"},
		{makeTT(token.Increment), "++"},
		{makeTT(token.Decrement), "--"},
		{makeTT(token.Exponentiation), "**"},
		{makeTT(token.LeftShift), "<<"},
		{makeTT(token.RightShift), ">>"},
		{makeTT(token.UnsignedRightShift), ">>>"},
		{makeTT(token.BitwiseAnd), "&"},
		{makeTT(token.BitwiseOr), "|"},
		{makeTT(token.BitwiseXor), "^"},
		{makeTT(token.Bang), "!"},
		{makeTT(token.Tilde), "~"},
		{makeTT(token.LogicalAnd), "&&"},
		{makeTT(token.LogicalOr), "||"},
		{makeTT(token.NullishCoalescing), "??"},
		{makeTT(token.Assignment), "="},
		{makeTT(token.AdditionAssignment), "+="},
		{makeTT(token.SubtractionAssignment), "-="},
		{makeTT(token.MultiplicationAssignment), "*="},
		{makeTT(token.DivisionAssignment), "/="},
		{makeTT(token.RemainderAssignment), "%="},
		{makeTT(token.ExponentiationAssignment), "**="},
		{makeTT(token.LeftShiftAssignment), "<<="},
		{makeTT(token.RightShiftAssignment), ">>="},
		{makeTT(token.UnsignedRightShiftAssignment), ">>>="},
		{makeTT(token.BitwiseAndAssignment), "&="},
		{makeTT(token.BitwiseOrAssignment), "|="},
		{makeTT(token.BitwiseXorAssignment), "^="},
		{makeTT(token.Arrow), "=>"},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLiteral(t *testing.T) {
	input := `
null
true
false
-99
0
123
3.14
1.23.45
0b111
0b0101010
0B0
0B000
0b123
-0b1
0o1
0o777
0O456
0x123
0xaBc
0xf
"hello"
'world'
"こんにちは, 世界🌮"
'Murphy\'s law'
"\"And God created great whales.\""
'\\\t\n\v'
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{makeTT(token.Null), "null"},
		{makeTT(token.True), "true"},
		{makeTT(token.False), "false"},
		{makeTT(token.Minus), "-"},
		{makeTT(token.Numeric), "99"},
		{makeTT(token.Numeric), "0"},
		{makeTT(token.Numeric), "123"},
		{makeTT(token.Numeric), "3.14"},
		{makeTT(token.Numeric), "1.23"},
		{makeTT(token.Numeric), ".45"},
		{makeTT(token.Numeric), "0b111"},
		{makeTT(token.Numeric), "0b0101010"},
		{makeTT(token.Numeric), "0B0"},
		{makeTT(token.Numeric), "0B000"},
		{makeTT(token.Numeric), "0b1"},
		{makeTT(token.Numeric), "23"},
		{makeTT(token.Minus), "-"},
		{makeTT(token.Numeric), "0b1"},
		{makeTT(token.Numeric), "0o1"},
		{makeTT(token.Numeric), "0o777"},
		{makeTT(token.Numeric), "0O456"},
		{makeTT(token.Numeric), "0x123"},
		{makeTT(token.Numeric), "0xaBc"},
		{makeTT(token.Numeric), "0xf"},
		{makeTT(token.String), "hello"},
		{makeTT(token.String), "world"},
		{makeTT(token.String), "こんにちは, 世界🌮"},
		{makeTT(token.String), "Murphy's law"},
		{makeTT(token.String), `"And God created great whales."`},
		{makeTT(token.String), "\\\\\\t\\n\\v"},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestTemplateLiteral(t *testing.T) {
	input := "`hello`\n`goodbye ${world}!`\n`result=${1 + 2}`\n`hello ${`world ${`again`}`}`"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// `hello`
		{makeTT(token.TemplateStart), "template-start"},
		{makeTT(token.String), "hello"},
		{makeTT(token.TemplateEnd), "template-end"},
		// `goodbye ${world}!`
		{makeTT(token.TemplateStart), "template-start"},
		{makeTT(token.String), "goodbye "},
		{makeTT(token.SubstitutionStart), "${"},
		{makeTT(token.Identifier), "world"},
		{makeTT(token.SubstitutionEnd), "}"},
		{makeTT(token.String), "!"},
		{makeTT(token.TemplateEnd), "template-end"},
		// `result=${1 + 2}`
		{makeTT(token.TemplateStart), "template-start"},
		{makeTT(token.String), "result="},
		{makeTT(token.SubstitutionStart), "${"},
		{makeTT(token.Numeric), "1"},
		{makeTT(token.Plus), "+"},
		{makeTT(token.Numeric), "2"},
		{makeTT(token.SubstitutionEnd), "}"},
		{makeTT(token.TemplateEnd), "template-end"},
		// `hello ${`world ${`again`}`}`
		{makeTT(token.TemplateStart), "template-start"},
		{makeTT(token.String), "hello "},
		{makeTT(token.SubstitutionStart), "${"},
		{makeTT(token.TemplateStart), "template-start"},
		{makeTT(token.String), "world "},
		{makeTT(token.SubstitutionStart), "${"},
		{makeTT(token.TemplateStart), "template-start"},
		{makeTT(token.String), "again"},
		{makeTT(token.TemplateEnd), "template-end"},
		{makeTT(token.SubstitutionEnd), "}"},
		{makeTT(token.TemplateEnd), "template-end"},
		{makeTT(token.SubstitutionEnd), "}"},
		{makeTT(token.TemplateEnd), "template-end"},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestSourceLocation(t *testing.T) {
	input := `/**/===
?? hello;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLoc     token.SourceLocation
	}{
		{makeTT(token.Identity), "===", makeLoc(0, 4, 0, 7)},
		{makeTT(token.NullishCoalescing), "??", makeLoc(1, 0, 1, 2)},
		{makeTT(token.Identifier), "hello", makeLoc(1, 3, 1, 8)},
		{makeTT(token.Semicolon), ";", makeLoc(1, 8, 1, 9)},
		{makeTT(token.EOF), "", makeLoc(1, 9, 1, 9)},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Loc != tt.expectedLoc {
			t.Fatalf("tests[%d] - location wrong. expected=%+v, got=%+v",
				i, tt.expectedLoc, tok.Loc)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `// calculate gcd of a and b
const gcd = (a, b) => {
  if (b === 0) {
    return a;
  }
  return gcd(b, a % b);
};

console.log(gcd(1263262, 553443)); // 11
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLoc     token.SourceLocation
	}{
		{makeTT(token.Const), "const", makeLoc(1, 0, 1, 5)},
		{makeTT(token.Identifier), "gcd", makeLoc(1, 6, 1, 9)},
		{makeTT(token.Assignment), "=", makeLoc(1, 10, 1, 11)},
		{makeTT(token.LParen), "(", makeLoc(1, 12, 1, 13)},
		{makeTT(token.Identifier), "a", makeLoc(1, 13, 1, 14)},
		{makeTT(token.Comma), ",", makeLoc(1, 14, 1, 15)},
		{makeTT(token.Identifier), "b", makeLoc(1, 16, 1, 17)},
		{makeTT(token.RParen), ")", makeLoc(1, 17, 1, 18)},
		{makeTT(token.Arrow), "=>", makeLoc(1, 19, 1, 21)},
		{makeTT(token.LBrace), "{", makeLoc(1, 22, 1, 23)},
		{makeTT(token.If), "if", makeLoc(2, 2, 2, 4)},
		{makeTT(token.LParen), "(", makeLoc(2, 5, 2, 6)},
		{makeTT(token.Identifier), "b", makeLoc(2, 6, 2, 7)},
		{makeTT(token.Identity), "===", makeLoc(2, 8, 2, 11)},
		{makeTT(token.Numeric), "0", makeLoc(2, 12, 2, 13)},
		{makeTT(token.RParen), ")", makeLoc(2, 13, 2, 14)},
		{makeTT(token.LBrace), "{", makeLoc(2, 15, 2, 16)},
		{makeTT(token.Return), "return", makeLoc(3, 4, 3, 10)},
		{makeTT(token.Identifier), "a", makeLoc(3, 11, 3, 12)},
		{makeTT(token.Semicolon), ";", makeLoc(3, 12, 3, 13)},
		{makeTT(token.RBrace), "}", makeLoc(4, 2, 4, 3)},
		{makeTT(token.Return), "return", makeLoc(5, 2, 5, 8)},
		{makeTT(token.Identifier), "gcd", makeLoc(5, 9, 5, 12)},
		{makeTT(token.LParen), "(", makeLoc(5, 12, 5, 13)},
		{makeTT(token.Identifier), "b", makeLoc(5, 13, 5, 14)},
		{makeTT(token.Comma), ",", makeLoc(5, 14, 5, 15)},
		{makeTT(token.Identifier), "a", makeLoc(5, 16, 5, 17)},
		{makeTT(token.Remainder), "%", makeLoc(5, 18, 5, 19)},
		{makeTT(token.Identifier), "b", makeLoc(5, 20, 5, 21)},
		{makeTT(token.RParen), ")", makeLoc(5, 21, 5, 22)},
		{makeTT(token.Semicolon), ";", makeLoc(5, 22, 5, 23)},
		{makeTT(token.RBrace), "}", makeLoc(6, 0, 6, 1)},
		{makeTT(token.Semicolon), ";", makeLoc(6, 1, 6, 2)},
		{makeTT(token.Identifier), "console", makeLoc(8, 0, 8, 7)},
		{makeTT(token.Dot), ".", makeLoc(8, 7, 8, 8)},
		{makeTT(token.Identifier), "log", makeLoc(8, 8, 8, 11)},
		{makeTT(token.LParen), "(", makeLoc(8, 11, 8, 12)},
		{makeTT(token.Identifier), "gcd", makeLoc(8, 12, 8, 15)},
		{makeTT(token.LParen), "(", makeLoc(8, 15, 8, 16)},
		{makeTT(token.Numeric), "1263262", makeLoc(8, 16, 8, 23)},
		{makeTT(token.Comma), ",", makeLoc(8, 23, 8, 24)},
		{makeTT(token.Numeric), "553443", makeLoc(8, 25, 8, 31)},
		{makeTT(token.RParen), ")", makeLoc(8, 31, 8, 32)},
		{makeTT(token.RParen), ")", makeLoc(8, 32, 8, 33)},
		{makeTT(token.Semicolon), ";", makeLoc(8, 33, 8, 34)},
		{makeTT(token.EOF), "", makeLoc(9, 0, 9, 0)},
	}

	l := New(input)

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatalf("tests[%d] - unexpected error: %q", i, err.Error())
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Loc != tt.expectedLoc {
			t.Fatalf("tests[%d] - location wrong. expected=%+v, got=%+v",
				i, tt.expectedLoc, tok.Loc)
		}
	}
}

func TestError(t *testing.T) {
	inputs := []string{
		"# ",
		"@",
		"'123",
		"\"foo",
		"\n\"bar\n\"",
		"0b",
		"\n0B2",
		"0o8",
		"0xyz",
	}

	tests := []struct {
		expectedMessage string
	}{
		{"SyntaxError: Unexpected character '#' (0:0)"},
		{"SyntaxError: Unexpected character '@' (0:0)"},
		{"SyntaxError: Unterminated string constant (0:0)"},
		{"SyntaxError: Unterminated string constant (0:0)"},
		{"SyntaxError: Unterminated string constant (1:0)"},
		{"SyntaxError: Expected number in radix 2 (0:2)"},
		{"SyntaxError: Expected number in radix 2 (1:2)"},
		{"SyntaxError: Expected number in radix 8 (0:2)"},
		{"SyntaxError: Expected number in radix 16 (0:2)"},
	}

	for i, tt := range tests {
		func() {
			l := New(inputs[i])

			_, err := l.NextToken()

			if err == nil {
				t.Fatalf("tests[%d] - did not panic.", i)
			}

			if err.Error() != tt.expectedMessage {
				t.Fatalf("tests[%d] - unexpected error message. expected=%q, got=%q",
					i, tt.expectedMessage, err.Error())
			}
		}()
	}
}
