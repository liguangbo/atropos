package fate

import (
	"fmt"
	"math"
	"reflect"
)

/*
	此文件实现对于变量的函数操作
*/

var supportFuncHandler = map[string]func(interface{}) interface{}{
	"count": countRunner,
	"sum":   sumRunner,
	"avg":   avgRunner,
	"max":   maxRunner,
	"min":   minRunner,
}

func funcRunner(actual interface{}, funcName string) interface{} {
	fn, ok := supportFuncHandler[funcName]
	if !ok {
		panic(fmt.Sprintf("func %s not found", funcName))
	}
	return fn(actual)
}

func countRunner(in interface{}) interface{} {
	arr := reflect.ValueOf(in)
	switch arr.Type().Kind() {
	case reflect.Slice, reflect.Map, reflect.Array:
		return arr.Len()
	case reflect.String:
		return len(arr.String())
	}
	panic(fmt.Sprintf("can not do count on %T", in))
}

func sumRunner(in interface{}) interface{} {
	switch t := in.(type) {
	case []interface{}:
		return sumSlice(t)
	case map[string]interface{}:
		return sumMap(t)
	default:
		panic(fmt.Sprintf("can not do sum on %T", in))
	}
}

func sumSlice(arr []interface{}) interface{} {
	var res float64
	var convertInt64 = true
	for _, v := range arr {
		switch tv := v.(type) {
		case int:
			res += float64(tv)
		case int64:
			res += float64(tv)
		case float64:
			convertInt64 = false
			res += tv
		default:
			panic(fmt.Sprintf("array contain %T when do sum", v))
		}
	}
	if convertInt64 {
		return int64(res)
	}
	return res
}

func sumMap(m map[string]interface{}) interface{} {
	var res float64
	var convertInt64 = true
	for _, v := range m {
		switch tv := v.(type) {
		case int:
			res += float64(tv)
		case int64:
			res += float64(tv)
		case float64:
			convertInt64 = false
			res += tv
		default:
			panic(fmt.Sprintf("map contain %T when do sum", v))
		}
	}
	if convertInt64 {
		return int64(res)
	}
	return res
}

func maxRunner(in interface{}) interface{} {
	switch t := in.(type) {
	case []interface{}:
		return maxSlice(t)
	case map[string]interface{}:
		return maxMap(t)
	default:
		panic(fmt.Sprintf("can not do max on %T", in))
	}
}

func maxSlice(arr []interface{}) interface{} {
	if len(arr) == 0 {
		panic(fmt.Sprintf("array %T empty when do max", arr))
	}
	var cur = -math.MaxFloat64
	var maxIdx int
	for i, v := range arr {
		switch tv := v.(type) {
		case int:
			if cur <= float64(tv) {
				maxIdx = i
				cur = float64(tv)
			}
		case int64:
			if cur <= float64(tv) {
				maxIdx = i
				cur = float64(tv)
			}
		case float64:
			if cur <= tv {
				maxIdx = i
				cur = tv
			}
		default:
			panic(fmt.Sprintf("array contain %T when do max", v))
		}
	}
	return arr[maxIdx]
}

func maxMap(m map[string]interface{}) interface{} {
	if len(m) == 0 {
		panic(fmt.Sprintf("map %T empty when do max", m))
	}
	var cur = -math.MaxFloat64
	var maxKey string
	for k, v := range m {
		switch tv := v.(type) {
		case int:
			if cur <= float64(tv) {
				maxKey = k
				cur = float64(tv)
			}
		case int64:
			if cur <= float64(tv) {
				maxKey = k
				cur = float64(tv)
			}
		case float64:
			if cur <= tv {
				maxKey = k
				cur = tv
			}
		default:
			panic(fmt.Sprintf("map contain %T when do max", v))
		}
	}
	return m[maxKey]
}

func minRunner(in interface{}) interface{} {
	switch t := in.(type) {
	case []interface{}:
		return minSlice(t)
	case map[string]interface{}:
		return minMap(t)
	default:
		panic(fmt.Sprintf("can not do min on %T", in))
	}
}

func minSlice(arr []interface{}) interface{} {
	if len(arr) == 0 {
		panic(fmt.Sprintf("array %T empty when do min", arr))
	}
	var cur = math.MaxFloat64
	var minIdx int
	for i, v := range arr {
		switch tv := v.(type) {
		case int:
			if cur >= float64(tv) {
				minIdx = i
				cur = float64(tv)
			}
		case int64:
			if cur >= float64(tv) {
				minIdx = i
				cur = float64(tv)
			}
		case float64:
			if cur >= tv {
				minIdx = i
				cur = tv
			}
		default:
			panic(fmt.Sprintf("array contain %T when do min", v))
		}
	}
	return arr[minIdx]
}

func minMap(m map[string]interface{}) interface{} {
	if len(m) == 0 {
		panic(fmt.Sprintf("map %T empty when do min", m))
	}
	var cur = math.MaxFloat64
	var minKey string
	for k, v := range m {
		switch tv := v.(type) {
		case int:
			if cur >= float64(tv) {
				minKey = k
				cur = float64(tv)
			}
		case int64:
			if cur >= float64(tv) {
				minKey = k
				cur = float64(tv)
			}
		case float64:
			if cur >= tv {
				minKey = k
				cur = tv
			}
		default:
			panic(fmt.Sprintf("map contain %T when do min", v))
		}
	}
	return m[minKey]
}

func avgRunner(in interface{}) interface{} {
	switch t := in.(type) {
	case []interface{}:
		return avgSlice(t)
	case map[string]interface{}:
		return avgMap(t)
	default:
		panic(fmt.Sprintf("can not do avg on %T", in))
	}
}

func avgSlice(arr []interface{}) interface{} {
	var sum float64
	for _, v := range arr {
		switch tv := v.(type) {
		case int:
			sum += float64(tv)
		case int64:
			sum += float64(tv)
		case float64:
			sum += tv
		default:
			panic(fmt.Sprintf("array contain %T when do avg", v))
		}
	}
	if floatEqual(sum, 0.0) {
		return float64(0)
	}
	return sum / float64(len(arr))
}

func avgMap(m map[string]interface{}) interface{} {
	var sum float64
	for _, v := range m {
		switch tv := v.(type) {
		case int:
			sum += float64(tv)
		case int64:
			sum += float64(tv)
		case float64:
			sum += tv
		default:
			panic(fmt.Sprintf("map contain %T when do avg", v))
		}
	}
	if floatEqual(sum, 0.0) {
		return float64(0)
	}
	return sum / float64(len(m))
}
