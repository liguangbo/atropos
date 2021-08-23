package fate

import (
	"fmt"
	"github.com/git-hulk/go-lru"
	"math"
	"regexp"
)

/*
	此文件提供在执行脚本中普遍使用的工具函数
*/

var regexpCache = lru.NewCache(100000)

func getRegexp(regexpStr string) (*regexp.Regexp, error) {
	regexpInterface, ok := regexpCache.Get(regexpStr)
	if !ok {
		reg, err := regexp.Compile(regexpStr)
		if err != nil {
			return nil, err
		}
		regexpCache.Set(regexpStr, reg)
		return reg, nil
	}
	reg, ok := regexpInterface.(*regexp.Regexp)
	if !ok {
		reg, err := regexp.Compile(regexpStr)
		if err != nil {
			return nil, err
		}
		regexpCache.Set(regexpStr, reg)
		return reg, nil
	}
	return reg, nil
}

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func regexpMatch(str, regexpStr string) (bool, error) {
	if regexpStr == "" {
		return true, nil
	}
	reg, err := getRegexp(regexpStr)
	if err != nil {
		return false, err
	}
	return reg.MatchString(str), nil
}

// 用于各个比较以及计算中的错误信息输出
func errorOperationMsg(op string, left, right interface{}) string {
	return fmt.Sprintf("operation %s can not excute between %T and %T", op, left, right)
}

func boolExpr(exprValue interface{}) bool {
	value := false
	switch expr := exprValue.(type) {
	case bool:
		value = expr
	case int:
		value = expr != 0
	case int64:
		value = expr != 0
	case float64:
		value = !floatEqual(expr, 0.0)
	case string:
		value = expr != ""
	case nil:
	default:
		// 其他类型的值
		panic(fmt.Sprintf("%T it not a valid boolean", exprValue))
	}
	return value
}
