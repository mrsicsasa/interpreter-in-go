package token

type TokenType string

const(
	ILLEGAL="ILLEGAL"
	EOF="EOF"

	//Identifiers + literals
	IDENT="IDENT" //add,footbar,x,y
	INT="INT" 

	//Operators
	ASSIGN="="
	PLUS="+"
	MINUS="-"
	BANG="!"
	ASTERISK="*"
	SPLASH="/"

	LT="<"
	GT=">"
	
	EQ="=="
	NOT_EQ="!="

	//Delimiters
	COMMA=","
	SEMICOLON=";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION="FUNCTION"
	LET="LET"
	TRUE="TRUE"
	FALSE="FALSE"
	IF="IF"
	ELSE="ELSE"
	RETURN="RETURN"
)

type Token struct{
	Type TokenType
	Literal string
}