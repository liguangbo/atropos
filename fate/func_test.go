package fate

import (
	"fmt"
	"runtime/debug"
	"testing"
)

type funcTestSuit struct {
	value       interface{}
	funcName    string
	shouldPanic bool
	expect      interface{}
}

func (suit *funcTestSuit) String() string {
	return fmt.Sprintf("%T(%v).%s() = %T(%v)", suit.value, suit.value, suit.funcName, suit.expect, suit.expect)
}

func testFunc(suit *funcTestSuit, t *testing.T) {
	defer func() {
		err := recover()
		if !suit.shouldPanic {
			// 不应该panic的时候panic了
			if err != nil {
				t.Error(fmt.Sprintf("test %s should not panic\n panic:%s\n trace:%s", suit.String(), err, string(debug.Stack())))
			}
			return
		}
		// 应该panic的时候没有panic
		if err == nil {
			t.Error(fmt.Sprintf("test %s should panic but not panic\n panic:%s\n trace:%s", suit.String(), err, string(debug.Stack())))
		}
	}()
	got := funcRunner(suit.value, suit.funcName)
	if got != suit.expect {
		t.Error(fmt.Sprintf("test %s but got %T(%v)", suit.String(), got, got))
	}
	return
}

func TestCount(t *testing.T) {
	suits := []funcTestSuit{
		{[]interface{}{"1", 1, true, nil, 1.1}, "count", false, 5},
		{[]int{0, 1, 2, 3}, "count", false, 4},
		{[5]interface{}{"1", 1, true, nil, 1.1}, "count", false, 5},
		{map[string]interface{}{"a": "a", "b": true}, "count", false, 2},
		{[]byte("test"), "count", false, 4},
		{"test", "count", false, 4},
		{int(1), "count", true, nil},
		{int8(1), "count", true, nil},
		{int16(1), "count", true, nil},
		{int32(1), "count", true, nil},
		{int64(1), "count", true, nil},
		{float32(1), "count", true, nil},
		{float64(1), "count", true, nil},
		{true, "count", true, nil},
	}

	for _, suit := range suits {
		testFunc(&suit, t)
	}
}

func TestSum(t *testing.T) {
	suits := []funcTestSuit{
		{[]interface{}{0, 1, 2, int64(3)}, "sum", false, int64(6)},
		{[]interface{}{0, 1, 2, -3}, "sum", false, int64(0)},
		{[]interface{}{0, 1, 2, 3, 0.5}, "sum", false, float64(6.5)},
		{[]interface{}{"1", 1, true, nil, 1.1}, "sum", true, nil},
		{map[string]interface{}{"a": 1, "b": int64(3)}, "sum", false, int64(4)},
		{map[string]interface{}{"a": 1, "b": 3.5}, "sum", false, float64(4.5)},
		{map[string]interface{}{"a": "a", "b": true}, "sum", true, nil},
		{[]byte("test"), "sum", true, nil},
		{"test", "sum", true, nil},
		{int(1), "sum", true, nil},
		{int8(1), "sum", true, nil},
		{int16(1), "sum", true, nil},
		{int32(1), "sum", true, nil},
		{int64(1), "sum", true, nil},
		{float32(1), "sum", true, nil},
		{float64(1), "sum", true, nil},
		{true, "sum", true, nil},
	}

	for _, suit := range suits {
		testFunc(&suit, t)
	}
}

func TestAvg(t *testing.T) {
	suits := []funcTestSuit{
		{[]interface{}{}, "avg", false, float64(0)},
		{[]interface{}{0, 1, 2, int64(3)}, "avg", false, float64(1.5)},
		{[]interface{}{0, 1, 2, -3}, "avg", false, float64(0)},
		{[]interface{}{0, 1, 2, 3, 0.5}, "avg", false, float64(1.3)},
		{[]interface{}{"1", 1, true, nil, 1.1}, "avg", true, nil},
		{map[string]interface{}{}, "avg", false, float64(0)},
		{map[string]interface{}{"a": 1, "b": int64(3)}, "avg", false, float64(2)},
		{map[string]interface{}{"a": 1, "b": 3.5, "c": 0}, "avg", false, float64(1.5)},
		{map[string]interface{}{"a": "a", "b": true}, "avg", true, nil},
		{[]byte("test"), "avg", true, nil},
		{"test", "avg", true, nil},
		{int(1), "avg", true, nil},
		{int8(1), "avg", true, nil},
		{int16(1), "avg", true, nil},
		{int32(1), "avg", true, nil},
		{int64(1), "avg", true, nil},
		{float32(1), "avg", true, nil},
		{float64(1), "avg", true, nil},
		{true, "avg", true, nil},
	}

	for _, suit := range suits {
		testFunc(&suit, t)
	}
}

func TestMax(t *testing.T) {
	suits := []funcTestSuit{
		{[]interface{}{}, "max", true, nil},
		{[]interface{}{0, 1, 2, 3}, "max", false, 3},
		{[]interface{}{0, 1, int64(2), -3}, "max", false, int64(2)},
		{[]interface{}{0, 1, 2, 3, 3.5}, "max", false, float64(3.5)},
		{[]interface{}{"1", 1, true, nil, 1.1}, "max", true, nil},
		{map[string]interface{}{}, "max", true, nil},
		{map[string]interface{}{"a": 1, "b": 3}, "max", false, 3},
		{map[string]interface{}{"a": 1, "b": 3.5, "c": int64(4)}, "max", false, int64(4)},
		{map[string]interface{}{"a": 1, "b": 3.5, "c": 0}, "max", false, float64(3.5)},
		{map[string]interface{}{"a": "a", "b": true}, "max", true, nil},
		{[]byte("test"), "max", true, nil},
		{"test", "max", true, nil},
		{int(1), "max", true, nil},
		{int8(1), "max", true, nil},
		{int16(1), "max", true, nil},
		{int32(1), "max", true, nil},
		{int64(1), "max", true, nil},
		{float32(1), "max", true, nil},
		{float64(1), "max", true, nil},
		{true, "max", true, nil},
	}

	for _, suit := range suits {
		testFunc(&suit, t)
	}
}

func TestMin(t *testing.T) {
	suits := []funcTestSuit{
		{[]interface{}{}, "min", true, nil},
		{[]interface{}{0, 1, 2, 3}, "min", false, 0},
		{[]interface{}{0, 1, 2, int64(-3)}, "min", false, int64(-3)},
		{[]interface{}{0, 1, 2, 3, -3.5}, "min", false, float64(-3.5)},
		{[]interface{}{"1", 1, true, nil, 1.1}, "min", true, nil},
		{map[string]interface{}{}, "min", true, nil},
		{map[string]interface{}{"a": 1, "b": 3}, "min", false, 1},
		{map[string]interface{}{"a": int64(1), "b": 3.5, "c": 4}, "min", false, int64(1)},
		{map[string]interface{}{"a": 1, "b": -3.5, "c": 0}, "min", false, float64(-3.5)},
		{map[string]interface{}{"a": "a", "b": true}, "min", true, nil},
		{[]byte("test"), "min", true, nil},
		{"test", "min", true, nil},
		{int(1), "min", true, nil},
		{int8(1), "min", true, nil},
		{int16(1), "min", true, nil},
		{int32(1), "min", true, nil},
		{int64(1), "min", true, nil},
		{float32(1), "min", true, nil},
		{float64(1), "min", true, nil},
		{true, "min", true, nil},
	}

	for _, suit := range suits {
		testFunc(&suit, t)
	}
}
