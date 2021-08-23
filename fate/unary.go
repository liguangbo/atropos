package fate

import "fmt"

const (
	opMinus       = "-"
	opExclamation = "!"
)

func unaryExpr(exprValue interface{}, op string) interface{} {
	switch op {
	case opMinus:
		return unaryMinus(exprValue)
	case opExclamation:
		return unaryExclamation(exprValue)
	default:
		panic(fmt.Sprintf("not recognize unary operation %s", op))
	}
}

func unaryMinus(exprValue interface{}) interface{} {
	switch expr := exprValue.(type) {
	case int:
		return -expr
	case int64:
		return -expr
	case float64:
		return -expr
	default:
		panic(fmt.Sprintf("unary operation %s can not use on %T", opMinus, expr))
	}
}

func unaryExclamation(exprValue interface{}) interface{} {
	switch expr := exprValue.(type) {
	case bool:
		return !expr
	default:
		panic(fmt.Sprintf("unary operation %s can not use on %T", opExclamation, expr))
	}
}
