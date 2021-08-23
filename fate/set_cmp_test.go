package fate

import (
	"fmt"
	"runtime/debug"
	"testing"
)

type setCmpTestSuit struct {
	leftValue   interface{}
	rightValue  interface{}
	op          string
	shouldPanic bool
	expect      interface{}
}

func (suit *setCmpTestSuit) String() string {
	return fmt.Sprintf("%T(%v) %s %T(%v) = %T(%v)", suit.leftValue, suit.leftValue, suit.op, suit.rightValue, suit.rightValue, suit.expect, suit.expect)
}

func testSetCmp(suit *setCmpTestSuit, t *testing.T) {
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

	got := compareSet(suit.leftValue, suit.rightValue, suit.op)
	if got != suit.expect {
		t.Error(fmt.Sprintf("test %s but got %T(%v)", suit.String(), got, got))
	}
	return
}

func TestContains(t *testing.T) {
	arr1 := []interface{}{"1", 1, true, nil, 1.1}
	suits := []setCmpTestSuit{
		{arr1, "1", opContains, false, true},
		{arr1, int(1), opContains, false, true},
		{arr1, true, opContains, false, true},
		{arr1, nil, opContains, false, true},
		{arr1, 1.1, opContains, false, true},
		{arr1, 1.2, opContains, false, false},
		{arr1, map[string]string{}, opContains, false, false},
		{arr1, []interface{}{"1", true}, opContains, false, true},
		{arr1, []interface{}{"1", false}, opContains, false, false},
		{arr1, arr1, opContains, false, true},
		{arr1, append(arr1, "test2"), opContains, false, false},
		{1, 1, opContains, true, nil},
	}

	for _, suit := range suits {
		testSetCmp(&suit, t)
	}
}

func TestNotContains(t *testing.T) {
	arr1 := []interface{}{"1", 1, true, nil, 1.1}
	suits := []setCmpTestSuit{
		{arr1, "1", opNotContains, false, false},
		{arr1, int(1), opNotContains, false, false},
		{arr1, true, opNotContains, false, false},
		{arr1, nil, opNotContains, false, false},
		{arr1, 1.1, opNotContains, false, false},
		{arr1, 1.2, opNotContains, false, true},
		{arr1, map[string]string{}, opNotContains, false, true},
		{arr1, []interface{}{"1", true}, opNotContains, false, false},
		{arr1, []interface{}{"1", false}, opNotContains, false, true},
		{arr1, arr1, opNotContains, false, false},
		{arr1, append(arr1, "test2"), opNotContains, false, true},
		{1, 1, opNotContains, true, nil},
	}

	for _, suit := range suits {
		testSetCmp(&suit, t)
	}
}
