package token

type TokenType struct {
	Label string
}

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"

	// Keywords
	Await      = "await"
	Break      = "break"
	Case       = "case"
	Catch      = "catch"
	Class      = "class"
	Const      = "const"
	Continue   = "continue"
	Debugger   = "debugger"
	Default    = "default"
	Delete     = "delete"
	Do         = "do"
	Else       = "else"
	Enum       = "enum"
	Export     = "export"
	Extends    = "extends"
	False      = "false"
	Finally    = "finally"
	For        = "for"
	Function   = "function"
	If         = "if"
	Import     = "Import"
	In         = "in"
	Instanceof = "instanceof"
	New        = "new"
	Null       = "null"
	Return     = "return"
	Super      = "super"
	Switch     = "switch"
	This       = "this"
	Throw      = "throw"
	True       = "true"
	Try        = "try"
	Typeof     = "typeof"
	Var        = "var"
	Void       = "void"
	While      = "while"
	With       = "with"
	Yield      = "yield"

	// Punctuators
	LParen           = "("
	RParen           = ")"
	LBrace           = "{"
	RBrace           = "}"
	LBracket         = "["
	RBracket         = "]"
	Dot              = "."
	Ellipsis         = "..."
	Semicolon        = ";"
	Colon            = ":"
	Comma            = ","
	Question         = "?"
	OptionalChaining = "?."

	// Operators
	LT                           = "<"
	GT                           = ">"
	LTEq                         = "<="
	GTEq                         = ">="
	Equality                     = "=="
	Inequality                   = "!="
	Identity                     = "==="
	Nonidentity                  = "!=="
	Plus                         = "+"
	Minus                        = "-"
	Star                         = "*"
	Slash                        = "/"
	Remainder                    = "%"
	Increment                    = "++"
	Decrement                    = "--"
	Exponentiation               = "**"
	LeftShift                    = "<<"
	RightShift                   = ">>"
	UnsignedRightShift           = ">>>"
	BitwiseAnd                   = "&"
	BitwiseOr                    = "|"
	BitwiseXor                   = "^"
	Bang                         = "!"
	Tilde                        = "~"
	LogicalAnd                   = "&&"
	LogicalOr                    = "||"
	NullishCoalescing            = "??"
	Assignment                   = "="
	AdditionAssignment           = "+="
	SubtractionAssignment        = "-="
	MultiplicationAssignment     = "*="
	DivisionAssignment           = "/="
	RemainderAssignment          = "%="
	ExponentiationAssignment     = "**="
	LeftShiftAssignment          = "<<="
	RightShiftAssignment         = ">>="
	UnsignedRightShiftAssignment = ">>>="
	BitwiseAndAssignment         = "&="
	BitwiseOrAssignment          = "|="
	BitwiseXorAssignment         = "^="
	Arrow                        = "=>"

	// Literals
	Numeric  = "numeric"
	String   = "string"
	Regex    = "regex"
	Template = "template"
)

var keywords = map[string]TokenType{
	"await":      {Await},
	"break":      {Break},
	"case":       {Case},
	"catch":      {Catch},
	"class":      {Class},
	"const":      {Const},
	"continue":   {Continue},
	"debugger":   {Debugger},
	"default":    {Default},
	"delete":     {Delete},
	"do":         {Do},
	"else":       {Else},
	"enum":       {Enum},
	"export":     {Export},
	"extends":    {Extends},
	"false":      {False},
	"finally":    {Finally},
	"for":        {For},
	"function":   {Function},
	"if":         {If},
	"import":     {Import},
	"in":         {In},
	"instanceof": {Instanceof},
	"new":        {New},
	"null":       {Null},
	"return":     {Return},
	"super":      {Super},
	"switch":     {Switch},
	"this":       {This},
	"throw":      {Throw},
	"true":       {True},
	"try":        {Try},
	"typeof":     {Typeof},
	"var":        {Var},
	"void":       {Void},
	"while":      {While},
	"with":       {With},
	"yield":      {Yield},
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TokenType{IDENT}
}
