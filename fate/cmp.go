package fate

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	本文件实现基础类型的比较
*/

const (
	opEqual       = "=="
	opNotEqual    = "!="
	opLarger      = ">"
	opLargerEqual = ">="
	opLess        = "<"
	opLessEqual   = "<="
	opRegexpMatch = "regexpMatch"
)

const (
	epsilon = float64(1e-10)
)

func compare(leftValue interface{}, rightValue interface{}, op string) bool {
	switch left := leftValue.(type) {
	case int:
		return cmpInt(int64(left), rightValue, op)
	case int64:
		return cmpInt(left, rightValue, op)
	case float64:
		return cmpFloat(left, rightValue, op)
	case string:
		return cmpStr(left, rightValue, op)
	case bool:
		return cmpBool(left, rightValue, op)
	case nil:
		return cmpNil(left, rightValue, op)
	default:
		switch rightValue {
		case nil:
			switch op {
			case opEqual:
				return false
			case opNotEqual:
				return true
			}
		}
		panic(errorOperationMsg(op, left, rightValue))
	}
}

func cmpInt(left int64, rightValue interface{}, op string) bool {
	switch right := rightValue.(type) {
	case int:
		switch op {
		case opEqual:
			return left == int64(right)
		case opNotEqual:
			return left != int64(right)
		case opLarger:
			return left > int64(right)
		case opLargerEqual:
			return left >= int64(right)
		case opLess:
			return left < int64(right)
		case opLessEqual:
			return left <= int64(right)
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case int64:
		switch op {
		case opEqual:
			return left == int64(right)
		case opNotEqual:
			return left != int64(right)
		case opLarger:
			return left > int64(right)
		case opLargerEqual:
			return left >= int64(right)
		case opLess:
			return left < int64(right)
		case opLessEqual:
			return left <= int64(right)
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case float64:
		switch op {
		case opEqual:
			return floatEqual(float64(left), right)
		case opNotEqual:
			return !floatEqual(float64(left), right)
		case opLarger:
			return float64(left) > right
		case opLargerEqual:
			return float64(left) > right || floatEqual(float64(left), right)
		case opLess:
			return float64(left) < right
		case opLessEqual:
			return float64(left) < right || floatEqual(float64(left), right)
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case string:
		switch op {
		case opRegexpMatch:
			ok, err := regexpMatch(strconv.FormatInt(left, 10), right)
			if err != nil {
				panic(fmt.Sprintf("compile regexp error when do regexpMatch, %s", err.Error()))
			}
			return ok
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case bool:
		switch op {
		case opEqual:
			return (left != 0) == right
		case opNotEqual:
			return (left != 0) != right
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case nil:
		switch op {
		case opEqual:
			return false
		case opNotEqual:
			return true
		default:
			panic(errorOperationMsg(op, left, right))
		}
	default:
		panic(fmt.Sprintf("%T can not compare with %T", left, right))
	}
}

func cmpFloat(left float64, rightValue interface{}, op string) bool {
	switch right := rightValue.(type) {
	case int:
		switch op {
		case opEqual:
			return floatEqual(left, float64(right))
		case opNotEqual:
			return !floatEqual(left, float64(right))
		case opLarger:
			return left > float64(right)
		case opLargerEqual:
			return left >= float64(right) || floatEqual(left, float64(right))
		case opLess:
			return left < float64(right)
		case opLessEqual:
			return left <= float64(right) || floatEqual(left, float64(right))
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case int64:
		switch op {
		case opEqual:
			return floatEqual(left, float64(right))
		case opNotEqual:
			return !floatEqual(left, float64(right))
		case opLarger:
			return left > float64(right)
		case opLargerEqual:
			return left >= float64(right) || floatEqual(left, float64(right))
		case opLess:
			return left < float64(right)
		case opLessEqual:
			return left <= float64(right) || floatEqual(left, float64(right))
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case float64:
		switch op {
		case opEqual:
			return floatEqual(left, right)
		case opNotEqual:
			return !floatEqual(left, right)
		case opLarger:
			return left > right
		case opLargerEqual:
			return left > right || floatEqual(left, right)
		case opLess:
			return left < right
		case opLessEqual:
			return left < right || floatEqual(left, right)
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case string:
		switch op {
		case opRegexpMatch:
			ok, err := regexpMatch(strconv.FormatFloat(left, 'f', -1, 64), right)
			if err != nil {
				panic(fmt.Sprintf("compile regexp error when do regexpMatch, %s", err.Error()))
			}
			return ok
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case bool:
		switch op {
		case opEqual:
			return !floatEqual(left, 0.0) == right
		case opNotEqual:
			return !floatEqual(left, 0.0) != right
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case nil:
		switch op {
		case opEqual:
			return false
		case opNotEqual:
			return true
		default:
			panic(errorOperationMsg(op, left, right))
		}
	default:
		panic(fmt.Sprintf("%T can not compare with %T", left, right))
	}
}

func cmpStr(left string, rightValue interface{}, op string) bool {
	switch right := rightValue.(type) {
	case int, int64:
		cmp := strings.Compare(left, fmt.Sprintf("%d", right))
		switch op {
		case opEqual:
			return cmp == 0
		case opNotEqual:
			return cmp != 0
		case opLarger:
			return cmp == +1
		case opLargerEqual:
			return cmp == +1 || cmp == 0
		case opLess:
			return cmp == -1
		case opLessEqual:
			return cmp == -1 || cmp == 0
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case float64:
		cmp := strings.Compare(left, strconv.FormatFloat(right, 'f', -1, 64))
		switch op {
		case opEqual:
			return cmp == 0
		case opNotEqual:
			return cmp != 0
		case opLarger:
			return cmp == +1
		case opLargerEqual:
			return cmp == +1 || cmp == 0
		case opLess:
			return cmp == -1
		case opLessEqual:
			return cmp == -1 || cmp == 0
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case string:
		cmp := strings.Compare(left, right)
		switch op {
		case opEqual:
			return cmp == 0
		case opNotEqual:
			return cmp != 0
		case opLarger:
			return cmp == +1
		case opLargerEqual:
			return cmp == +1 || cmp == 0
		case opLess:
			return cmp == -1
		case opLessEqual:
			return cmp == -1 || cmp == 0
		case opRegexpMatch:
			ok, err := regexpMatch(left, right)
			if err != nil {
				panic(fmt.Sprintf("compile regexp error when do regexpMatch, %s", err.Error()))
			}
			return ok
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case bool:
		b, _ := strconv.ParseBool(left)
		switch op {
		case opEqual:
			return b == right
		case opNotEqual:
			return b != right
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case nil:
		switch op {
		case opEqual:
			return false
		case opNotEqual:
			return true
		default:
			panic(errorOperationMsg(op, left, right))
		}
	default:
		panic(fmt.Sprintf("%T can not compare with %T", left, right))
	}
}

func cmpBool(left bool, rightValue interface{}, op string) bool {
	switch right := rightValue.(type) {
	case int:
		switch op {
		case opEqual:
			return left == (right != 0)
		case opNotEqual:
			return left != (right != 0)
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case int64:
		switch op {
		case opEqual:
			return left == (right != 0)
		case opNotEqual:
			return left != (right != 0)
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case float64:
		switch op {
		case opEqual:
			return left == !floatEqual(right, 0.0)
		case opNotEqual:
			return left != !floatEqual(right, 0.0)
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case string:
		b, _ := strconv.ParseBool(right)
		switch op {
		case opEqual:
			return left == b
		case opNotEqual:
			return left != b
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case bool:
		switch op {
		case opEqual:
			return left == right
		case opNotEqual:
			return left != right
		default:
			panic(errorOperationMsg(op, left, right))
		}
	case nil:
		switch op {
		case opEqual:
			return false
		case opNotEqual:
			return true
		default:
			panic(errorOperationMsg(op, left, right))
		}
	default:
		panic(fmt.Sprintf("%T can not compare with %T", left, right))
	}
}

func cmpNil(left interface{}, rightValue interface{}, op string) bool {
	switch op {
	case opEqual:
		return left == rightValue
	case opNotEqual:
		return left != rightValue
	default:
		panic(errorOperationMsg(op, left, rightValue))
	}
}
