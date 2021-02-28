package lexer

import (
	"testing"

	"github.com/morinokami/js-lexer/token"
)

func makeTT(label string) token.TokenType {
	return token.TokenType{Label: label}
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
