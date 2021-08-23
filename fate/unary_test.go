package fate

import (
	"fmt"
	"runtime/debug"
	"testing"
)

type unaryTestSuit struct {
	leftValue   interface{}
	op          string
	shouldPanic bool
	expect      interface{}
}

func (suit *unaryTestSuit) String() string {
	return fmt.Sprintf("%s%T(%v) = %T(%v)", suit.op, suit.leftValue, suit.leftValue, suit.expect, suit.expect)
}

func testUnary(suit *unaryTestSuit, t *testing.T) {
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
	got := unaryExpr(suit.leftValue, suit.op)

	if got != suit.expect {
		t.Error(fmt.Sprintf("test %s but got %T(%v)", suit.String(), got, got))
	}
	return
}

func TestMinus(t *testing.T) {
	suits := []unaryTestSuit{
		{int(1), opMinus, false, int(-1)},
		{int(-1), opMinus, false, int(1)},
		{int64(1), opMinus, false, int64(-1)},
		{int64(-1), opMinus, false, int64(1)},
		{float64(0.5), opMinus, false, float64(-0.5)},
		{float64(-0.5), opMinus, false, float64(0.5)},
		{int8(1), opMinus, true, nil},
		{int16(1), opMinus, true, nil},
		{int32(1), opMinus, true, nil},
		{"test", opMinus, true, nil},
		{true, opMinus, true, nil},
		{nil, opMinus, true, nil},
	}

	for _, suit := range suits {
		testUnary(&suit, t)
	}
}

func TestExclamation(t *testing.T) {
	suits := []unaryTestSuit{
		{true, opExclamation, false, false},
		{false, opExclamation, false, true},
		{int(1), opExclamation, true, nil},
		{int64(1), opExclamation, true, nil},
		{float64(0.5), opExclamation, true, nil},
		{int8(1), opExclamation, true, nil},
		{int16(1), opExclamation, true, nil},
		{int32(1), opExclamation, true, nil},
		{"test", opExclamation, true, nil},
		{nil, opExclamation, true, nil},
	}

	for _, suit := range suits {
		testUnary(&suit, t)
	}
}
