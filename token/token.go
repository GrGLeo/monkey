package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifiers + literals
  IDENT = "IDENT" // add, foobar, x, y
  INT = "INT"
  STRING = "STRING"

  // Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"


  // Conditional
  LT = "<"
  GT = ">"
  EQ = "=="
  NOT_EQ = "!="

  // Delimeters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keyword
  FUNCTION = "FUNCTION"
  LET = "LET"
  RETURN = "RETURN"
  IF = "IF"
  ELSE = "ELSE"
  TRUE = "TRUE"
  FALSE = "FALSE"
)

var keywords = map[string]TokenType {
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

func LookUpIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
