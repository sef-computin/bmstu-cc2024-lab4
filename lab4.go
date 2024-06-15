package main

import (
	"fmt"

	"github.com/sef-computin/bmstu-cc2024-lab4/grammar"
)


func main(){
	input := "begin a = true & ~false ! b end"
  // input := "begin a = true & ~false end"

  p := grammar.GetParserFromString(input)

  fmt.Println(grammar.PrintPolish(p))
}
