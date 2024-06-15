package grammar

import "fmt"

type Parser struct {
	lexer     *Lexer
	curTok    Token
	ParseTree *Node
	curNode   *Node
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer, curTok: lexer.NextToken()}
}

func (p *Parser) Eat(tokType TokenType, parent *Node) {
	// fmt.Println(p.curTok.ToString())
	if p.curTok.Type == tokType {
		_ = parent.AddChild(p.curTok.Value)
		p.curTok = p.lexer.NextToken()
	} else {
		panic(fmt.Sprintf("unexpected token: %v, expected: %v", p.curTok, tokType))
	}
}

func (p *Parser) ParseProgram() {
	p.ParseTree = NewNode("Программа")
	p.curNode = p.ParseTree
	p.ParseBlock()
}

func (p *Parser) ParseBlock() {
	temp := p.curNode
	block := p.curNode.AddChild("Блок")
	p.curNode = block
	defer p.swapCurNode(temp)

	p.Eat(BEGIN, p.curNode)
	p.ParseStatementList()
	// p.curNode = block
	p.Eat(END, p.curNode)

	// p.curNode = temp
}

func (p *Parser) ParseStatementList() {
	temp := p.curNode
	listNode := p.curNode.AddChild("Список операторов")
	p.curNode = listNode
	defer p.swapCurNode(temp)

	// fmt.Println(p.curNode.Value)
	p.ParseStatement()
	// fmt.Println(p.curNode.Value)

	for p.curTok.Type == SEMICOLON {
		p.Eat(SEMICOLON, p.curNode)
		p.ParseStatement()
	}
	// p.curNode = temp
}

// func (p *Parser) ParseStatementList(){
//   temp := p.curNode
//   listNode := p.curNode.AddChild("STATEMENTS_LIST")
//   p.curNode = listNode
//   defer swapNode(p.curNode, temp)
//
//   p.ParseStatement()
//
//   for p.curTok.Type == SEMICOLON{
//     p.Eat(SEMICOLON, p.curNode)
//     p.ParseStatementList()
//   }
//
// }

func (p *Parser) ParseStatement() {
	temp := p.curNode

	stat := p.curNode.AddChild("Оператор")
	p.curNode = stat
	defer p.swapCurNode(temp)

	if p.curTok.Type == IDENTIFIER {
		p.ParseAssignment()
	} else if p.curTok.Type == BEGIN {
		p.ParseBlock()
	} else {
		panic(fmt.Sprintf("unexpected token in statement: %v", p.curTok))
	}

	// p.curNode = temp
}

func (p *Parser) ParseAssignment() {
	temp := p.curNode
	asgn := p.curNode.AddChild("Операция присваивания")
	p.curNode = asgn
	defer p.swapCurNode(temp)

	p.Eat(IDENTIFIER, p.curNode)
	p.Eat(ASSIGN, p.curNode)
	p.ParseExpression()
	// p.curNode = temp
}

func (p *Parser) ParseExpression() {
	temp := p.curNode
	expr := p.curNode.AddChild("Выражение")
	p.curNode = expr
	defer p.swapCurNode(temp)

	p.ParseLogicalExpression()
	// p.curNode = temp
}

func (p *Parser) ParseLogicalExpression() {
	temp := p.curNode
	expr := p.curNode.AddChild("Логическое выражение")
	p.curNode = expr

	defer p.swapCurNode(temp)

	p.ParseLogicalTerm()
	for p.curTok.Type == OR {
		p.Eat(OR, p.curNode)
		p.ParseLogicalTerm()
	}
	// p.curNode = temp
}

func (p *Parser) ParseLogicalTerm() {
	temp := p.curNode
	expr := p.curNode.AddChild("Логический одночлен")
	p.curNode = expr
	defer p.swapCurNode(temp)

	p.ParseLogicalFactor()
	for p.curTok.Type == AND {
		p.Eat(AND, p.curNode)
		p.ParseLogicalFactor()
	}
	// p.curNode = temp
}

func (p *Parser) ParseLogicalFactor() {
	temp := p.curNode
	expr := p.curNode.AddChild("Вторичное логическое выражение")
	p.curNode = expr
	defer p.swapCurNode(temp)

	if p.curTok.Type == NOT {
		p.Eat(NOT, p.curNode)
		p.ParsePrimaryLogicalExpression()
	} else {
		p.ParsePrimaryLogicalExpression()
	}

	// p.curNode = temp
}

func (p *Parser) ParsePrimaryLogicalExpression() {
	temp := p.curNode
	expr := p.curNode.AddChild("Первичное логическое выражение")
	p.curNode = expr
	defer p.swapCurNode(temp)

	if p.curTok.Type == TRUE || p.curTok.Type == FALSE {
		p.Eat(p.curTok.Type, p.curNode)
	} else if p.curTok.Type == IDENTIFIER {
		p.Eat(IDENTIFIER, p.curNode)
	} else {
		panic(fmt.Sprintf("unexpected token in primary logical expression: %v", p.curTok))
	}

	// p.curNode = temp
}

func (p *Parser) GetParseTree() *Node {
	return p.ParseTree
}

func (p *Parser) swapCurNode(target *Node) {
	p.curNode = target
}
