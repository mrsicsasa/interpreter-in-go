package parser

import(
	"github.com/mrsicsasa/interpreter-in-go/token"
	"github.com/mrsicsasa/interpreter-in-go/ast"
	"github.com/mrsicsasa/interpreter-in-go/lexer"
)

type Parser struct{
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}
    // Read two tokens, so curToken and peekToken are both set
    p.NextToken()
    p.NextToken()
    return p
}
func (p *Parser) NextToken(){
	p.curToken=p.peekToken
	p.peekToken=p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
	return nil
}