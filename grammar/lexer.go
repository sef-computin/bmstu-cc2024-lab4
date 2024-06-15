package grammar

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
    BEGIN TokenType = iota
    END
    IDENTIFIER
    ASSIGN
    TRUE
    FALSE
    NOT
    AND
    OR
    SEMICOLON
    EOF    
    INVALID
)

type Token struct {
    Type  TokenType
    Value string
}

func (t Token) ToString() string{
  return fmt.Sprintf("{%d, %s}", t.Type, t.Value)
}

type Lexer struct {
    input string
    pos   int
}

func NewLexer(input string) *Lexer {
    return &Lexer{input: input}
}

func (l *Lexer) NextToken() Token {
    for l.pos < len(l.input) {
        ch := l.input[l.pos]
        switch {
        case unicode.IsSpace(rune(ch)):
            l.pos++
        case strings.HasPrefix(l.input[l.pos:], "begin"):
            l.pos += 5
            return Token{Type: BEGIN, Value: "begin"}
        case strings.HasPrefix(l.input[l.pos:], "end"):
            l.pos += 3
            return Token{Type: END, Value: "end"}
        case strings.HasPrefix(l.input[l.pos:], "true"):
            l.pos += 4
            return Token{Type: TRUE, Value: "true"}
        case strings.HasPrefix(l.input[l.pos:], "false"):
            l.pos += 5
            return Token{Type: FALSE, Value: "false"}
        case ch == '=':
            l.pos++
            return Token{Type: ASSIGN, Value: "="}
        case ch == '!':
            l.pos++
            return Token{Type: OR, Value: "!"}
        case ch == '&':
            l.pos++
            return Token{Type: AND, Value: "&"}
        case ch == '~':
            l.pos++
            return Token{Type: NOT, Value: "~"}
        case ch == ';':
            l.pos++
            return Token{Type: SEMICOLON, Value: ";"}
        case unicode.IsLetter(rune(ch)):
            start := l.pos
            for l.pos < len(l.input) && unicode.IsLetter(rune(l.input[l.pos])) {
                l.pos++
            }
            return Token{Type: IDENTIFIER, Value: l.input[start:l.pos]}
        default:
            l.pos++
            return Token{Type: INVALID, Value: string(ch)}
        }
    }
    return Token{Type: EOF, Value: ""}
}

func matchTokenType(value string) (ttype TokenType){
  switch value {
  case "begin":
    ttype = BEGIN
  case "end":
    ttype = END
  case "=":
    ttype = ASSIGN
  case "true":
    ttype = TRUE
  case "false":
    ttype = FALSE
  case "~":
    ttype = NOT
  case "&":
    ttype = AND
  case "!":
    ttype = OR
  case ";":
    ttype = SEMICOLON
  default:
    ttype = IDENTIFIER
  }
  return
}
