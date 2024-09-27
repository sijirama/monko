package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// NOTE: if it is a keyword return the keyword type
// if not return the IDENT (identifier/variable type)
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

const (
	ILLEGAL = "ILLEGAL" //token we do not know about
	EOF     = "EOF"

	//Identifies + Literals
	IDENT = "IDENT" //identifiers add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)