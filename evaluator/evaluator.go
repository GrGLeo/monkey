package evaluator

import (
	"github.com/GrGLeo/monkey/ast"
	"github.com/GrGLeo/monkey/object"
)

var (
  NULL = &object.Null{}
  TRUE = &object.Boolean{Value: true}
  FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
  switch node := node.(type) {
  // Statements
  case *ast.Program:
    return evalStatements(node.Statements)
  case *ast.ExpressionStatement:
    return Eval(node.Expression)
  // Expression
  case *ast.IntegerLiteral:
    return &object.Integer{Value: node.Value}
  case *ast.Boolean:
    return nativeBooToBooleanObject(node.Value)
  case *ast.PrefixExpression:
    right := Eval(node.Right)
    return evalPrefixExpression(node.Operator, right)
  }

  return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
  var result object.Object

  for _, statement := range stmts {
    result = Eval(statement)
  }
  return result
}

func nativeBooToBooleanObject(input bool) *object.Boolean {
  if input {
    return TRUE
  }
  return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
  switch operator {
  case "!":
    return evalBangOperator(right)
  default:
    return NULL
  }
}

func evalBangOperator(right object.Object) object.Object {
  switch right {
  case TRUE:
    return FALSE
  case FALSE:
    return TRUE
  case NULL:
    return TRUE
  default:
    return FALSE
  }
}
