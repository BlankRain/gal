package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	//
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NEQ       = "!="
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	//array
	LBRACKET = "["
	RBRACKET = "]"
	// key word
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	// gal feature
	PACKAGE  = "PACKAGE"
	NODETYPE = "NODETYPE"
	EDGETYPE = "EDGETYPE"
	SHOW     = ":"
	AS       = "AS"
	AT       = "@"
	GQL      = "GQL"
	FROM     = "FROM"
	MAKE     = "MAKE"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"function": FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"package":  PACKAGE,
	"NodeType": NODETYPE,
	"EdgeType": EDGETYPE,
	"@":        AT,
	"as":       AS,
	"make":     MAKE,
	"from":     FROM,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
