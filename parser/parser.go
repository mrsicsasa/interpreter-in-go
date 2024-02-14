package parser

import(
	"fmt"
	"strconv"
	"github.com/mrsicsasa/interpreter-in-go/token"
	"github.com/mrsicsasa/interpreter-in-go/ast"
	"github.com/mrsicsasa/interpreter-in-go/lexer"
)
const(
	_ int=iota
	LOWEST
	EQUALS //==
	LESSGREATER //> or <
	SUM //+
	PRODUCT //*
	PREFIX // -X or !X
	CALL //myFunction(x)
)
type Parser struct{
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
	errors []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{
		l: l,
		errors:[]string{},
	}
	p.prefixParseFns=make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT,p.parseIdentifier)
	p.registerPrefix(token.INT,p.parseIntegerLiteral)
    // Read two tokens, so curToken and peekToken are both set
    p.NextToken()
    p.NextToken()
    return p
}
func (p *Parser)Errors() []string{
	return p.errors
}
func (p *Parser) NextToken(){
	p.curToken=p.peekToken
	p.peekToken=p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
	program:=&ast.Program{}
	program.Statements=[]ast.Statement{}
	for p.curToken.Type!=token.EOF{
		stmt:=p.parseStatement()
		if stmt!=nil{
			program.Statements=append(program.Statements,stmt)
		}
		p.NextToken()
	}
	return program
}
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
func (p *Parser)parseStatement() ast.Statement{
	switch p.curToken.Type{
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.ParseExpressionStatement()
	}
}
func (p *Parser)parseLetStatement() *ast.LetStatement{
	stmt:=&ast.LetStatement{Token:p.curToken}
	if !p.expectPeek(token.IDENT){
		return nil
	}
	stmt.Name=&ast.Identifier{Token:p.curToken,Value:p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN){
		return nil
	}
	//TODO: We are skipping the expressions until we encounter a semicolone
	for !p.curTokenIs(token.SEMICOLON){
		p.NextToken()
	}
	return stmt
}
func (p *Parser) parseReturnStatement()*ast.ReturnStatement{
	stmt:=&ast.ReturnStatement{Token:p.curToken}
	p.NextToken()
	//TODO: WE are skipping the expressions until we encounter a semicolone
	for !p.curTokenIs(token.SEMICOLON){
		p.NextToken()
	}
	return stmt
}

func (p *Parser) ParseExpressionStatement()*ast.ExpressionStatement{
	stmt:=&ast.ExpressionStatement{Token:p.curToken}
	stmt.Expression=p.ParseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON){
		p.NextToken()
	}
	return stmt
}

func (p *Parser)curTokenIs(t token.TokenType)bool{
	return p.curToken.Type==t
}
func (p *Parser) peekTokenIs(t token.TokenType)bool{
	return p.peekToken.Type==t
}
func (p *Parser) expectPeek(t token.TokenType)bool{
	if p.peekTokenIs(t){
		p.NextToken()
		return true
	}else{
		p.peekError(t)
		return false
	}
}
func (p *Parser) peekError(t token.TokenType){
	msg:=fmt.Sprintf("expected next token to be %s, got %s instead",
t,p.peekToken.Type)
p.errors=append(p.errors,msg)
}

type(
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn){
	p.prefixParseFns[tokenType]=fn
}
func (p *Parser) registrerInfix(tokenType token.TokenType, fn infixParseFn){
	p.infixParseFns[tokenType]=fn
}
func (p *Parser) ParseExpression(precedence int) ast.Expression{
	prefix:=p.prefixParseFns[p.curToken.Type]
	if prefix==nil{
		return nil
	}
	leftExp:=prefix()
	return leftExp
}
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}