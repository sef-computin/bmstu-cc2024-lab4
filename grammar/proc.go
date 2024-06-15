package grammar


func GetParserFromString(input string) (p *Parser){
  p = NewParser(NewLexer(input))
  p.ParseProgram()

  return
}

func PrintPolish(p *Parser) string{
  return p.ParseTree.getPolishString()
}
