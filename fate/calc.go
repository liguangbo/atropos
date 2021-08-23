package fate

import (
	"strconv"
)

/*
	本文件实现基础类型的加减乘除取模
*/

const (
	opAdd = "+"
	opSub = "-"
	opMul = "*"
	opDiv = "/"
	opMod = "%"
)

func addSub(leftValue, rightValue interface{}, op string) interface{} {
	switch left := leftValue.(type) {
	case int:
		return addSubInt(int64(left), rightValue, op)
	case int64:
		return addSubInt(left, rightValue, op)
	case float64:
		return addSubFloat(left, rightValue, op)
	case string:
		return addSubString(left, rightValue, op)
	case bool:
		return addSubBool(left, rightValue, op)
	default:
		panic(errorOperationMsg(op, left, rightValue))
	}
}

func addSubInt(left int64, rightValue interface{}, op string) interface{} {
	switch right := rightValue.(type) {
	case int:
		if op == opAdd {
			return left + int64(right)
		}
		return left - int64(right)
	case int64:
		if op == opAdd {
			return left + right
		}
		return left - right
	case float64:
		if op == opAdd {
			return float64(left) + right
		}
		return float64(left) - right
	case string:
		if op == opAdd {
			return strconv.FormatInt(left, 10) + right
		}
		panic(errorOperationMsg(op, left, right))
	default:
		panic(errorOperationMsg(op, left, right))
	}
}

func addSubFloat(left float64, rightValue interface{}, op string) interface{} {
	switch right := rightValue.(type) {
	case int:
		if op == opAdd {
			return left + float64(right)
		}
		return left - float64(right)
	case int64:
		if op == opAdd {
			return left + float64(right)
		}
		return left - float64(right)
	case float64:
		if op == opAdd {
			return left + right
		}
		return left - right
	case string:
		if op == opAdd {
			return strconv.FormatFloat(left, 'f', -1, 64) + right
		}
		panic(errorOperationMsg(op, left, right))
	default:
		panic(errorOperationMsg(op, left, right))
	}
}

func addSubString(left string, rightValue interface{}, op string) interface{} {
	switch right := rightValue.(type) {
	case int:
		if op == opAdd {
			return left + strconv.Itoa(right)
		}
		panic(errorOperationMsg(op, left, right))
	case int64:
		if op == opAdd {
			return left + strconv.FormatInt(right, 10)
		}
		panic(errorOperationMsg(op, left, right))
	case float64:
		if op == opAdd {
			return left + strconv.FormatFloat(right, 'f', -1, 64)
		}
		panic(errorOperationMsg(op, left, right))
	case string:
		if op == opAdd {
			return left + right
		}
		panic(errorOperationMsg(op, left, right))
	case bool:
		if op == opAdd {
			return left + strconv.FormatBool(right)
		}
		panic(errorOperationMsg(op, left, right))
	default:
		panic(errorOperationMsg(op, left, right))
	}
}

func addSubBool(left bool, rightValue interface{}, op string) interface{} {
	switch right := rightValue.(type) {
	case string:
		if op == opAdd {
			return strconv.FormatBool(left) + right
		}
		panic(errorOperationMsg(op, left, right))
	default:
		panic(errorOperationMsg(op, left, right))
	}
}

func mulDiv(leftValue, rightValue interface{}, op string) interface{} {
	switch left := leftValue.(type) {
	case int:
		return mulDivInt(int64(left), rightValue, op)
	case int64:
		return mulDivInt(left, rightValue, op)
	case float64:
		return mulDivFloat(left, rightValue, op)
	default:
		panic(errorOperationMsg(op, left, rightValue))
	}
}

func mulDivInt(left int64, rightValue interface{}, op string) interface{} {
	switch right := rightValue.(type) {
	case int:
		if op == opMul {
			return left * int64(right)
		} else if op == opDiv {
			return left / int64(right)
		}
		return left % int64(right)
	case int64:
		if op == opMul {
			return left * right
		} else if op == opDiv {
			return left / right
		}
		return left % right
	case float64:
		if op == opMul {
			return float64(left) * right
		} else if op == opDiv {
			return float64(left) / right
		}
		return left % int64(right)
	default:
		panic(errorOperationMsg(op, left, right))
	}
}

func mulDivFloat(left float64, rightValue interface{}, op string) interface{} {
	switch right := rightValue.(type) {
	case int:
		if op == opMul {
			return left * float64(right)
		} else if op == opDiv {
			return left / float64(right)
		}
		return int64(left) % int64(right)
	case int64:
		if op == opMul {
			return left * float64(right)
		} else if op == opDiv {
			return left / float64(right)
		}
		return int64(left) % right
	case float64:
		if op == opMul {
			return left * right
		} else if op == opDiv {
			return left / right
		}
		return int64(left) % int64(right)
	default:
		panic(errorOperationMsg(op, left, right))
	}
}
