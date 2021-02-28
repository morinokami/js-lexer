package lexer

import (
	"testing"

	"github.com/morinokami/js-lexer/token"
)

func makeTT(label string) token.TokenType {
	return token.TokenType{Label: label}
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
		tok := l.NextToken()

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
		tok := l.NextToken()

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
		tok := l.NextToken()

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
		tok := l.NextToken()

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

func TestOperators(t *testing.T) {
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
		tok := l.NextToken()

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
0
123
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
		{makeTT(token.Numeric), "0"},
		{makeTT(token.Numeric), "123"},
		{makeTT(token.String), "hello"},
		{makeTT(token.String), "world"},
		{makeTT(token.String), "こんにちは, 世界🌮"},
		{makeTT(token.String), "Murphy's law"},
		{makeTT(token.String), `"And God created great whales."`},
		{makeTT(token.String), "\\\\\\t\\n\\v"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

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

func TestNextToken(t *testing.T) {
	input := `
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
	}{
		{makeTT(token.Const), "const"},
		{makeTT(token.Identifier), "gcd"},
		{makeTT(token.Assignment), "="},
		{makeTT(token.LParen), "("},
		{makeTT(token.Identifier), "a"},
		{makeTT(token.Comma), ","},
		{makeTT(token.Identifier), "b"},
		{makeTT(token.RParen), ")"},
		{makeTT(token.Arrow), "=>"},
		{makeTT(token.LBrace), "{"},
		{makeTT(token.If), "if"},
		{makeTT(token.LParen), "("},
		{makeTT(token.Identifier), "b"},
		{makeTT(token.Identity), "==="},
		{makeTT(token.Numeric), "0"},
		{makeTT(token.RParen), ")"},
		{makeTT(token.LBrace), "{"},
		{makeTT(token.Return), "return"},
		{makeTT(token.Identifier), "a"},
		{makeTT(token.Semicolon), ";"},
		{makeTT(token.RBrace), "}"},
		{makeTT(token.Return), "return"},
		{makeTT(token.Identifier), "gcd"},
		{makeTT(token.LParen), "("},
		{makeTT(token.Identifier), "b"},
		{makeTT(token.Comma), ","},
		{makeTT(token.Identifier), "a"},
		{makeTT(token.Remainder), "%"},
		{makeTT(token.Identifier), "b"},
		{makeTT(token.RParen), ")"},
		{makeTT(token.Semicolon), ";"},
		{makeTT(token.RBrace), "}"},
		{makeTT(token.Semicolon), ";"},
		{makeTT(token.Identifier), "console"},
		{makeTT(token.Dot), "."},
		{makeTT(token.Identifier), "log"},
		{makeTT(token.LParen), "("},
		{makeTT(token.Identifier), "gcd"},
		{makeTT(token.LParen), "("},
		{makeTT(token.Numeric), "1263262"},
		{makeTT(token.Comma), ","},
		{makeTT(token.Numeric), "553443"},
		{makeTT(token.RParen), ")"},
		{makeTT(token.RParen), ")"},
		{makeTT(token.Semicolon), ";"},
		{makeTT(token.EOF), ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

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
