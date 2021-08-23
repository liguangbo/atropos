package fate

import (
	"fmt"
	"math"
	"runtime/debug"
	"testing"
)

type calcTestSuit struct {
	leftValue   interface{}
	rightValue  interface{}
	op          string
	shouldPanic bool
	expect      interface{}
}

func (suit *calcTestSuit) String() string {
	return fmt.Sprintf("%T(%v) %s %T(%v) = %T(%v)", suit.leftValue, suit.leftValue, suit.op, suit.rightValue, suit.rightValue, suit.expect, suit.expect)
}

func testCalc(suit *calcTestSuit, t *testing.T) {
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
	var got interface{}
	if suit.op == opAdd || suit.op == opSub {
		got = addSub(suit.leftValue, suit.rightValue, suit.op)
	} else {
		got = mulDiv(suit.leftValue, suit.rightValue, suit.op)
	}

	if got != suit.expect {
		t.Error(fmt.Sprintf("test %s but got %T(%v)", suit.String(), got, got))
	}
	return
}

func TestIntAdd(t *testing.T) {
	var addInt = &calcTestSuit{int(1), int(1), opAdd, false, int64(2)}
	var addNegativeInt = &calcTestSuit{int(1), int(-2), opAdd, false, int64(-1)}
	var addIntOverflow = &calcTestSuit{math.MaxInt64, int(1), opAdd, false, int64(math.MinInt64)}
	var addInt64 = &calcTestSuit{int(1), int64(1), opAdd, false, int64(2)}
	var addNegativeInt64 = &calcTestSuit{int(1), int64(-2), opAdd, false, int64(-1)}
	var addInt64Overflow = &calcTestSuit{math.MaxInt64, int64(1), opAdd, false, int64(math.MinInt64)}
	var addFloat64 = &calcTestSuit{int(1), float64(0.5), opAdd, false, float64(1.5)}
	var addNegativeFloat64 = &calcTestSuit{int(1), float64(-0.5), opAdd, false, float64(0.5)}
	var addString = &calcTestSuit{int(1), "test", opAdd, false, "1test"}

	var addBool = &calcTestSuit{int(1), true, opAdd, true, nil}
	var addNil = &calcTestSuit{int(1), nil, opAdd, true, nil}
	var addInt8 = &calcTestSuit{int(1), int8(1), opAdd, true, nil}
	var addInt16 = &calcTestSuit{int(1), int16(1), opAdd, true, nil}
	var addInt32 = &calcTestSuit{int(1), int32(1), opAdd, true, nil}
	var addFloat32 = &calcTestSuit{int(1), float32(0.5), opAdd, true, nil}
	var addByte = &calcTestSuit{int(1), byte('a'), opAdd, true, nil}

	testCalc(addInt, t)
	testCalc(addNegativeInt, t)
	testCalc(addIntOverflow, t)
	testCalc(addInt64, t)
	testCalc(addNegativeInt64, t)
	testCalc(addInt64Overflow, t)
	testCalc(addFloat64, t)
	testCalc(addNegativeFloat64, t)
	testCalc(addString, t)
	testCalc(addBool, t)
	testCalc(addNil, t)
	testCalc(addInt8, t)
	testCalc(addInt16, t)
	testCalc(addInt32, t)
	testCalc(addFloat32, t)
	testCalc(addByte, t)
}

func TestInt64Add(t *testing.T) {
	var addInt = &calcTestSuit{int64(1), int(1), opAdd, false, int64(2)}
	var addNegativeInt = &calcTestSuit{int64(1), int(-2), opAdd, false, int64(-1)}
	var addIntOverflow = &calcTestSuit{int64(math.MaxInt64), int(1), opAdd, false, int64(math.MinInt64)}
	var addInt64 = &calcTestSuit{int64(1), int64(1), opAdd, false, int64(2)}
	var addNegativeInt64 = &calcTestSuit{int64(1), int64(-2), opAdd, false, int64(-1)}
	var addInt64Overflow = &calcTestSuit{int64(math.MaxInt64), int64(1), opAdd, false, int64(math.MinInt64)}
	var addFloat64 = &calcTestSuit{int64(1), float64(0.5), opAdd, false, float64(1.5)}
	var addNegativeFloat64 = &calcTestSuit{int64(1), float64(-0.5), opAdd, false, float64(0.5)}
	var addString = &calcTestSuit{int64(1), "test", opAdd, false, "1test"}

	var addBool = &calcTestSuit{int64(1), true, opAdd, true, nil}
	var addNil = &calcTestSuit{int64(1), nil, opAdd, true, nil}
	var addInt8 = &calcTestSuit{int64(1), int8(1), opAdd, true, nil}
	var addInt16 = &calcTestSuit{int64(1), int16(1), opAdd, true, nil}
	var addInt32 = &calcTestSuit{int64(1), int32(1), opAdd, true, nil}
	var addFloat32 = &calcTestSuit{int64(1), float32(0.5), opAdd, true, nil}
	var addByte = &calcTestSuit{int64(1), byte('a'), opAdd, true, nil}

	testCalc(addInt, t)
	testCalc(addNegativeInt, t)
	testCalc(addIntOverflow, t)
	testCalc(addInt64, t)
	testCalc(addNegativeInt64, t)
	testCalc(addInt64Overflow, t)
	testCalc(addFloat64, t)
	testCalc(addNegativeFloat64, t)
	testCalc(addString, t)
	testCalc(addBool, t)
	testCalc(addNil, t)
	testCalc(addInt8, t)
	testCalc(addInt16, t)
	testCalc(addInt32, t)
	testCalc(addFloat32, t)
	testCalc(addByte, t)
}

func TestFloat64Add(t *testing.T) {
	var addInt = &calcTestSuit{float64(0.5), int(1), opAdd, false, float64(1.5)}
	var addNegativeInt = &calcTestSuit{float64(0.5), int(-2), opAdd, false, float64(-1.5)}
	var addInt64 = &calcTestSuit{float64(0.5), int64(1), opAdd, false, float64(1.5)}
	var addNegativeInt64 = &calcTestSuit{float64(0.5), int64(-2), opAdd, false, float64(-1.5)}
	var addFloat64 = &calcTestSuit{float64(1.4), float64(0.5), opAdd, false, float64(1.9)}
	var addFloatOverflow = &calcTestSuit{math.MaxFloat64, math.MaxFloat64, opAdd, false, math.Inf(1)}
	var addNegativeFloat64 = &calcTestSuit{float64(1.5), float64(-0.5), opAdd, false, float64(1)}
	var addString = &calcTestSuit{float64(0.5), "test", opAdd, false, "0.5test"}

	var addBool = &calcTestSuit{float64(1), true, opAdd, true, nil}
	var addNil = &calcTestSuit{float64(1), nil, opAdd, true, nil}
	var addInt8 = &calcTestSuit{float64(1), int8(1), opAdd, true, nil}
	var addInt16 = &calcTestSuit{float64(1), int16(1), opAdd, true, nil}
	var addInt32 = &calcTestSuit{float64(1), int32(1), opAdd, true, nil}
	var addFloat32 = &calcTestSuit{float64(1), float32(0.5), opAdd, true, nil}
	var addByte = &calcTestSuit{float64(1), byte('a'), opAdd, true, nil}

	testCalc(addInt, t)
	testCalc(addNegativeInt, t)
	testCalc(addInt64, t)
	testCalc(addNegativeInt64, t)
	testCalc(addFloat64, t)
	testCalc(addFloatOverflow, t)
	testCalc(addNegativeFloat64, t)
	testCalc(addString, t)
	testCalc(addBool, t)
	testCalc(addNil, t)
	testCalc(addInt8, t)
	testCalc(addInt16, t)
	testCalc(addInt32, t)
	testCalc(addFloat32, t)
	testCalc(addByte, t)
}

func TestString64Add(t *testing.T) {
	var addInt = &calcTestSuit{"test", int(1), opAdd, false, "test1"}
	var addNegativeInt = &calcTestSuit{"test", int(-1), opAdd, false, "test-1"}
	var addInt64 = &calcTestSuit{"test", int64(1), opAdd, false, "test1"}
	var addNegativeInt64 = &calcTestSuit{"test", int64(-1), opAdd, false, "test-1"}
	var addFloat64 = &calcTestSuit{"test", float64(0.5), opAdd, false, "test0.5"}
	var addFloatOverflow = &calcTestSuit{"test", math.Inf(1), opAdd, false, "test+Inf"}
	var addNegativeFloat64 = &calcTestSuit{"test", float64(-0.5), opAdd, false, "test-0.5"}
	var addString = &calcTestSuit{"test", "test", opAdd, false, "testtest"}
	var addBool = &calcTestSuit{"test", true, opAdd, false, "testtrue"}

	var addNil = &calcTestSuit{"test", nil, opAdd, true, nil}
	var addInt8 = &calcTestSuit{"test", int8(1), opAdd, true, nil}
	var addInt16 = &calcTestSuit{"test", int16(1), opAdd, true, nil}
	var addInt32 = &calcTestSuit{"test", int32(1), opAdd, true, nil}
	var addFloat32 = &calcTestSuit{"test", float32(0.5), opAdd, true, nil}
	var addByte = &calcTestSuit{"test", byte('a'), opAdd, true, nil}

	testCalc(addInt, t)
	testCalc(addNegativeInt, t)
	testCalc(addInt64, t)
	testCalc(addNegativeInt64, t)
	testCalc(addFloat64, t)
	testCalc(addFloatOverflow, t)
	testCalc(addNegativeFloat64, t)
	testCalc(addString, t)
	testCalc(addBool, t)
	testCalc(addNil, t)
	testCalc(addInt8, t)
	testCalc(addInt16, t)
	testCalc(addInt32, t)
	testCalc(addFloat32, t)
	testCalc(addByte, t)
}

func TestBoolAdd(t *testing.T) {
	var addString = &calcTestSuit{true, "test", opAdd, false, "truetest"}

	var addInt = &calcTestSuit{true, int(1), opAdd, true, nil}
	var addInt64 = &calcTestSuit{true, int64(1), opAdd, true, nil}
	var addFloat64 = &calcTestSuit{true, float64(0.5), opAdd, true, nil}
	var addBool = &calcTestSuit{true, true, opAdd, true, nil}
	var addNil = &calcTestSuit{true, nil, opAdd, true, nil}
	var addInt8 = &calcTestSuit{true, int8(1), opAdd, true, nil}
	var addInt16 = &calcTestSuit{true, int16(1), opAdd, true, nil}
	var addInt32 = &calcTestSuit{true, int32(1), opAdd, true, nil}
	var addFloat32 = &calcTestSuit{true, float32(0.5), opAdd, true, nil}
	var addByte = &calcTestSuit{true, byte('a'), opAdd, true, nil}

	testCalc(addInt, t)
	testCalc(addInt64, t)
	testCalc(addFloat64, t)
	testCalc(addString, t)
	testCalc(addBool, t)
	testCalc(addNil, t)
	testCalc(addInt8, t)
	testCalc(addInt16, t)
	testCalc(addInt32, t)
	testCalc(addFloat32, t)
	testCalc(addByte, t)
}

func TestOtherAdd(t *testing.T) {
	var nilAdd = &calcTestSuit{nil, 1, opAdd, true, nil}
	var int8Add = &calcTestSuit{int8(1), 1, opAdd, true, nil}
	var int16Add = &calcTestSuit{int16(1), 1, opAdd, true, nil}
	var int32Add = &calcTestSuit{int32(1), 1, opAdd, true, nil}
	var byteAdd = &calcTestSuit{byte('1'), 1, opAdd, true, nil}
	var float32Add = &calcTestSuit{float32(0.5), 1, opAdd, true, nil}

	testCalc(nilAdd, t)
	testCalc(int8Add, t)
	testCalc(int16Add, t)
	testCalc(int32Add, t)
	testCalc(byteAdd, t)
	testCalc(float32Add, t)
}

func TestIntSub(t *testing.T) {
	var subInt = &calcTestSuit{int(1), int(1), opSub, false, int64(0)}
	var subNegativeInt = &calcTestSuit{int(1), int(-2), opSub, false, int64(3)}
	var subIntOverflow = &calcTestSuit{math.MinInt64, int(1), opSub, false, int64(math.MaxInt64)}
	var subInt64 = &calcTestSuit{int(1), int64(1), opSub, false, int64(0)}
	var subNegativeInt64 = &calcTestSuit{int(1), int64(-2), opSub, false, int64(3)}
	var subInt64Overflow = &calcTestSuit{math.MinInt64, int64(1), opSub, false, int64(math.MaxInt64)}
	var subFloat64 = &calcTestSuit{int(1), float64(0.5), opSub, false, float64(0.5)}
	var subNegativeFloat64 = &calcTestSuit{int(1), float64(-0.5), opSub, false, float64(1.5)}

	var subString = &calcTestSuit{int(1), "test", opSub, true, nil}
	var subBool = &calcTestSuit{int(1), true, opSub, true, nil}
	var subNil = &calcTestSuit{int(1), nil, opSub, true, nil}
	var subInt8 = &calcTestSuit{int(1), int8(1), opSub, true, nil}
	var subInt16 = &calcTestSuit{int(1), int16(1), opSub, true, nil}
	var subInt32 = &calcTestSuit{int(1), int32(1), opSub, true, nil}
	var subFloat32 = &calcTestSuit{int(1), float32(0.5), opSub, true, nil}
	var subByte = &calcTestSuit{int(1), byte('a'), opSub, true, nil}

	testCalc(subInt, t)
	testCalc(subNegativeInt, t)
	testCalc(subIntOverflow, t)
	testCalc(subInt64, t)
	testCalc(subNegativeInt64, t)
	testCalc(subInt64Overflow, t)
	testCalc(subFloat64, t)
	testCalc(subNegativeFloat64, t)
	testCalc(subString, t)
	testCalc(subBool, t)
	testCalc(subNil, t)
	testCalc(subInt8, t)
	testCalc(subInt16, t)
	testCalc(subInt32, t)
	testCalc(subFloat32, t)
	testCalc(subByte, t)
}

func TestInt64Sub(t *testing.T) {
	var subInt = &calcTestSuit{int64(1), int(1), opSub, false, int64(0)}
	var subNegativeInt = &calcTestSuit{int64(1), int(-2), opSub, false, int64(3)}
	var subIntOverflow = &calcTestSuit{int64(math.MinInt64), int(1), opSub, false, int64(math.MaxInt64)}
	var subInt64 = &calcTestSuit{int64(1), int64(1), opSub, false, int64(0)}
	var subNegativeInt64 = &calcTestSuit{int64(1), int64(-2), opSub, false, int64(3)}
	var subInt64Overflow = &calcTestSuit{int64(math.MinInt64), int64(1), opSub, false, int64(math.MaxInt64)}
	var subFloat64 = &calcTestSuit{int64(1), float64(0.5), opSub, false, float64(0.5)}
	var subNegativeFloat64 = &calcTestSuit{int64(1), float64(-0.5), opSub, false, float64(1.5)}

	var subString = &calcTestSuit{int64(1), "test", opSub, true, nil}
	var subBool = &calcTestSuit{int64(1), true, opSub, true, nil}
	var subNil = &calcTestSuit{int64(1), nil, opSub, true, nil}
	var subInt8 = &calcTestSuit{int64(1), int8(1), opSub, true, nil}
	var subInt16 = &calcTestSuit{int64(1), int16(1), opSub, true, nil}
	var subInt32 = &calcTestSuit{int64(1), int32(1), opSub, true, nil}
	var subFloat32 = &calcTestSuit{int64(1), float32(0.5), opSub, true, nil}
	var subByte = &calcTestSuit{int64(1), byte('a'), opSub, true, nil}

	testCalc(subInt, t)
	testCalc(subNegativeInt, t)
	testCalc(subIntOverflow, t)
	testCalc(subInt64, t)
	testCalc(subNegativeInt64, t)
	testCalc(subInt64Overflow, t)
	testCalc(subFloat64, t)
	testCalc(subNegativeFloat64, t)
	testCalc(subString, t)
	testCalc(subBool, t)
	testCalc(subNil, t)
	testCalc(subInt8, t)
	testCalc(subInt16, t)
	testCalc(subInt32, t)
	testCalc(subFloat32, t)
	testCalc(subByte, t)
}

func TestFloat64Sub(t *testing.T) {
	var subInt = &calcTestSuit{float64(0.5), int(1), opSub, false, float64(-0.5)}
	var subNegativeInt = &calcTestSuit{float64(0.5), int(-2), opSub, false, float64(2.5)}
	var subInt64 = &calcTestSuit{float64(0.5), int64(1), opSub, false, float64(-0.5)}
	var subNegativeInt64 = &calcTestSuit{float64(0.5), int64(-2), opSub, false, float64(2.5)}
	var subFloat64 = &calcTestSuit{float64(1.5), float64(0.5), opSub, false, float64(1)}
	var subFloatOverflow = &calcTestSuit{-math.MaxFloat64, math.MaxFloat64, opSub, false, math.Inf(-1)}
	var subNegativeFloat64 = &calcTestSuit{float64(1.5), float64(-0.5), opSub, false, float64(2)}

	var subString = &calcTestSuit{float64(0.5), "test", opSub, true, nil}
	var subBool = &calcTestSuit{float64(1), true, opSub, true, nil}
	var subNil = &calcTestSuit{float64(1), nil, opSub, true, nil}
	var subInt8 = &calcTestSuit{float64(1), int8(1), opSub, true, nil}
	var subInt16 = &calcTestSuit{float64(1), int16(1), opSub, true, nil}
	var subInt32 = &calcTestSuit{float64(1), int32(1), opSub, true, nil}
	var subFloat32 = &calcTestSuit{float64(1), float32(0.5), opSub, true, nil}
	var subByte = &calcTestSuit{float64(1), byte('a'), opSub, true, nil}

	testCalc(subInt, t)
	testCalc(subNegativeInt, t)
	testCalc(subInt64, t)
	testCalc(subNegativeInt64, t)
	testCalc(subFloat64, t)
	testCalc(subFloatOverflow, t)
	testCalc(subNegativeFloat64, t)
	testCalc(subString, t)
	testCalc(subBool, t)
	testCalc(subNil, t)
	testCalc(subInt8, t)
	testCalc(subInt16, t)
	testCalc(subInt32, t)
	testCalc(subFloat32, t)
	testCalc(subByte, t)
}

func TestString64Sub(t *testing.T) {
	var subInt = &calcTestSuit{"test", int(1), opSub, true, nil}
	var subInt64 = &calcTestSuit{"test", int64(1), opSub, true, nil}
	var subFloat64 = &calcTestSuit{"test", float64(0.5), opSub, true, nil}
	var subString = &calcTestSuit{"test", "test", opSub, true, nil}
	var subBool = &calcTestSuit{"test", true, opSub, true, nil}
	var subNil = &calcTestSuit{"test", nil, opSub, true, nil}
	var subInt8 = &calcTestSuit{"test", int8(1), opSub, true, nil}
	var subInt16 = &calcTestSuit{"test", int16(1), opSub, true, nil}
	var subInt32 = &calcTestSuit{"test", int32(1), opSub, true, nil}
	var subFloat32 = &calcTestSuit{"test", float32(0.5), opSub, true, nil}
	var subByte = &calcTestSuit{"test", byte('a'), opSub, true, nil}

	testCalc(subInt, t)
	testCalc(subInt64, t)
	testCalc(subFloat64, t)
	testCalc(subString, t)
	testCalc(subBool, t)
	testCalc(subNil, t)
	testCalc(subInt8, t)
	testCalc(subInt16, t)
	testCalc(subInt32, t)
	testCalc(subFloat32, t)
	testCalc(subByte, t)
}

func TestBoolSub(t *testing.T) {
	var subString = &calcTestSuit{true, "test", opSub, true, nil}
	var subInt = &calcTestSuit{true, int(1), opSub, true, nil}
	var subInt64 = &calcTestSuit{true, int64(1), opSub, true, nil}
	var subFloat64 = &calcTestSuit{true, float64(0.5), opSub, true, nil}
	var subBool = &calcTestSuit{true, true, opSub, true, nil}
	var subNil = &calcTestSuit{true, nil, opSub, true, nil}
	var subInt8 = &calcTestSuit{true, int8(1), opSub, true, nil}
	var subInt16 = &calcTestSuit{true, int16(1), opSub, true, nil}
	var subInt32 = &calcTestSuit{true, int32(1), opSub, true, nil}
	var subFloat32 = &calcTestSuit{true, float32(0.5), opSub, true, nil}
	var subByte = &calcTestSuit{true, byte('a'), opSub, true, nil}

	testCalc(subInt, t)
	testCalc(subInt64, t)
	testCalc(subFloat64, t)
	testCalc(subString, t)
	testCalc(subBool, t)
	testCalc(subNil, t)
	testCalc(subInt8, t)
	testCalc(subInt16, t)
	testCalc(subInt32, t)
	testCalc(subFloat32, t)
	testCalc(subByte, t)
}

func TestOtherSub(t *testing.T) {
	var nilSub = &calcTestSuit{nil, 1, opSub, true, nil}
	var int8Sub = &calcTestSuit{int8(1), 1, opSub, true, nil}
	var int16Sub = &calcTestSuit{int16(1), 1, opSub, true, nil}
	var int32Sub = &calcTestSuit{int32(1), 1, opSub, true, nil}
	var byteSub = &calcTestSuit{byte('1'), 1, opSub, true, nil}
	var float32Sub = &calcTestSuit{float32(0.5), 1, opSub, true, nil}

	testCalc(nilSub, t)
	testCalc(int8Sub, t)
	testCalc(int16Sub, t)
	testCalc(int32Sub, t)
	testCalc(byteSub, t)
	testCalc(float32Sub, t)
}

func TestIntMul(t *testing.T) {
	var mulInt = &calcTestSuit{int(2), int(2), opMul, false, int64(4)}
	var mulNegativeInt = &calcTestSuit{int(2), int(-2), opMul, false, int64(-4)}
	var mulIntOverflow = &calcTestSuit{math.MaxInt64, int(2), opMul, false, int64(-2)}
	var mulInt64 = &calcTestSuit{int(2), int64(2), opMul, false, int64(4)}
	var mulNegativeInt64 = &calcTestSuit{int(2), int64(-2), opMul, false, int64(-4)}
	var mulInt64Overflow = &calcTestSuit{math.MaxInt64, int64(2), opMul, false, int64(-2)}
	var mulFloat64 = &calcTestSuit{int(2), float64(0.5), opMul, false, float64(1)}
	var mulNegativeFloat64 = &calcTestSuit{int(2), float64(-0.5), opMul, false, float64(-1)}
	var mulFloat64Overflow = &calcTestSuit{int(2), math.MaxFloat64, opMul, false, math.Inf(1)}

	var mulString = &calcTestSuit{int(1), "test", opMul, true, nil}
	var mulBool = &calcTestSuit{int(1), true, opMul, true, nil}
	var mulNil = &calcTestSuit{int(1), nil, opMul, true, nil}
	var mulInt8 = &calcTestSuit{int(1), int8(1), opMul, true, nil}
	var mulInt16 = &calcTestSuit{int(1), int16(1), opMul, true, nil}
	var mulInt32 = &calcTestSuit{int(1), int32(1), opMul, true, nil}
	var mulFloat32 = &calcTestSuit{int(1), float32(0.5), opMul, true, nil}
	var mulByte = &calcTestSuit{int(1), byte('a'), opMul, true, nil}

	testCalc(mulInt, t)
	testCalc(mulNegativeInt, t)
	testCalc(mulIntOverflow, t)
	testCalc(mulInt64, t)
	testCalc(mulNegativeInt64, t)
	testCalc(mulInt64Overflow, t)
	testCalc(mulFloat64, t)
	testCalc(mulNegativeFloat64, t)
	testCalc(mulFloat64Overflow, t)
	testCalc(mulString, t)
	testCalc(mulBool, t)
	testCalc(mulNil, t)
	testCalc(mulInt8, t)
	testCalc(mulInt16, t)
	testCalc(mulInt32, t)
	testCalc(mulFloat32, t)
	testCalc(mulByte, t)
}

func TestInt64Mul(t *testing.T) {
	var mulInt = &calcTestSuit{int64(2), int(2), opMul, false, int64(4)}
	var mulNegativeInt = &calcTestSuit{int64(2), int(-2), opMul, false, int64(-4)}
	var mulIntOverflow = &calcTestSuit{int64(math.MaxInt64), int(2), opMul, false, int64(-2)}
	var mulInt64 = &calcTestSuit{int64(2), int64(2), opMul, false, int64(4)}
	var mulNegativeInt64 = &calcTestSuit{int64(2), int64(-2), opMul, false, int64(-4)}
	var mulInt64Overflow = &calcTestSuit{int64(math.MaxInt64), int64(2), opMul, false, int64(-2)}
	var mulFloat64 = &calcTestSuit{int64(2), float64(0.5), opMul, false, float64(1)}
	var mulNegativeFloat64 = &calcTestSuit{int64(2), float64(-0.5), opMul, false, float64(-1)}
	var mulFloat64Overflow = &calcTestSuit{int64(2), math.MaxFloat64, opMul, false, math.Inf(1)}

	var mulString = &calcTestSuit{int64(1), "test", opMul, true, nil}
	var mulBool = &calcTestSuit{int64(1), true, opMul, true, nil}
	var mulNil = &calcTestSuit{int64(1), nil, opMul, true, nil}
	var mulInt8 = &calcTestSuit{int64(1), int8(1), opMul, true, nil}
	var mulInt16 = &calcTestSuit{int64(1), int16(1), opMul, true, nil}
	var mulInt32 = &calcTestSuit{int64(1), int32(1), opMul, true, nil}
	var mulFloat32 = &calcTestSuit{int64(1), float32(0.5), opMul, true, nil}
	var mulByte = &calcTestSuit{int64(1), byte('a'), opMul, true, nil}

	testCalc(mulInt, t)
	testCalc(mulNegativeInt, t)
	testCalc(mulIntOverflow, t)
	testCalc(mulInt64, t)
	testCalc(mulNegativeInt64, t)
	testCalc(mulInt64Overflow, t)
	testCalc(mulFloat64, t)
	testCalc(mulNegativeFloat64, t)
	testCalc(mulFloat64Overflow, t)
	testCalc(mulString, t)
	testCalc(mulBool, t)
	testCalc(mulNil, t)
	testCalc(mulInt8, t)
	testCalc(mulInt16, t)
	testCalc(mulInt32, t)
	testCalc(mulFloat32, t)
	testCalc(mulByte, t)
}

func TestFloat64Mul(t *testing.T) {
	var mulInt = &calcTestSuit{float64(0.5), int(3), opMul, false, float64(1.5)}
	var mulNegativeInt = &calcTestSuit{float64(0.5), int(-3), opMul, false, float64(-1.5)}
	var mulInt64 = &calcTestSuit{float64(0.5), int64(3), opMul, false, float64(1.5)}
	var mulNegativeInt64 = &calcTestSuit{float64(0.5), int64(-3), opMul, false, float64(-1.5)}
	var mulFloat64 = &calcTestSuit{float64(2.5), float64(0.4), opMul, false, float64(1)}
	var mulFloatOverflow = &calcTestSuit{-math.MaxFloat64, math.MaxFloat64, opMul, false, math.Inf(-1)}
	var mulNegativeFloat64 = &calcTestSuit{float64(2.5), float64(-0.4), opMul, false, float64(-1)}

	var mulString = &calcTestSuit{float64(0.5), "test", opMul, true, nil}
	var mulBool = &calcTestSuit{float64(1), true, opMul, true, nil}
	var mulNil = &calcTestSuit{float64(1), nil, opMul, true, nil}
	var mulInt8 = &calcTestSuit{float64(1), int8(1), opMul, true, nil}
	var mulInt16 = &calcTestSuit{float64(1), int16(1), opMul, true, nil}
	var mulInt32 = &calcTestSuit{float64(1), int32(1), opMul, true, nil}
	var mulFloat32 = &calcTestSuit{float64(1), float32(0.5), opMul, true, nil}
	var mulByte = &calcTestSuit{float64(1), byte('a'), opMul, true, nil}

	testCalc(mulInt, t)
	testCalc(mulNegativeInt, t)
	testCalc(mulInt64, t)
	testCalc(mulNegativeInt64, t)
	testCalc(mulFloat64, t)
	testCalc(mulFloatOverflow, t)
	testCalc(mulNegativeFloat64, t)
	testCalc(mulString, t)
	testCalc(mulBool, t)
	testCalc(mulNil, t)
	testCalc(mulInt8, t)
	testCalc(mulInt16, t)
	testCalc(mulInt32, t)
	testCalc(mulFloat32, t)
	testCalc(mulByte, t)
}

func TestString64Mul(t *testing.T) {
	var mulInt = &calcTestSuit{"test", int(1), opMul, true, nil}
	var mulInt64 = &calcTestSuit{"test", int64(1), opMul, true, nil}
	var mulFloat64 = &calcTestSuit{"test", float64(0.5), opMul, true, nil}
	var mulString = &calcTestSuit{"test", "test", opMul, true, nil}
	var mulBool = &calcTestSuit{"test", true, opMul, true, nil}
	var mulNil = &calcTestSuit{"test", nil, opMul, true, nil}
	var mulInt8 = &calcTestSuit{"test", int8(1), opMul, true, nil}
	var mulInt16 = &calcTestSuit{"test", int16(1), opMul, true, nil}
	var mulInt32 = &calcTestSuit{"test", int32(1), opMul, true, nil}
	var mulFloat32 = &calcTestSuit{"test", float32(0.5), opMul, true, nil}
	var mulByte = &calcTestSuit{"test", byte('a'), opMul, true, nil}

	testCalc(mulInt, t)
	testCalc(mulInt64, t)
	testCalc(mulFloat64, t)
	testCalc(mulString, t)
	testCalc(mulBool, t)
	testCalc(mulNil, t)
	testCalc(mulInt8, t)
	testCalc(mulInt16, t)
	testCalc(mulInt32, t)
	testCalc(mulFloat32, t)
	testCalc(mulByte, t)
}

func TestBoolMul(t *testing.T) {
	var mulString = &calcTestSuit{true, "test", opMul, true, nil}
	var mulInt = &calcTestSuit{true, int(1), opMul, true, nil}
	var mulInt64 = &calcTestSuit{true, int64(1), opMul, true, nil}
	var mulFloat64 = &calcTestSuit{true, float64(0.5), opMul, true, nil}
	var mulBool = &calcTestSuit{true, true, opMul, true, nil}
	var mulNil = &calcTestSuit{true, nil, opMul, true, nil}
	var mulInt8 = &calcTestSuit{true, int8(1), opMul, true, nil}
	var mulInt16 = &calcTestSuit{true, int16(1), opMul, true, nil}
	var mulInt32 = &calcTestSuit{true, int32(1), opMul, true, nil}
	var mulFloat32 = &calcTestSuit{true, float32(0.5), opMul, true, nil}
	var mulByte = &calcTestSuit{true, byte('a'), opMul, true, nil}

	testCalc(mulInt, t)
	testCalc(mulInt64, t)
	testCalc(mulFloat64, t)
	testCalc(mulString, t)
	testCalc(mulBool, t)
	testCalc(mulNil, t)
	testCalc(mulInt8, t)
	testCalc(mulInt16, t)
	testCalc(mulInt32, t)
	testCalc(mulFloat32, t)
	testCalc(mulByte, t)
}

func TestOtherMul(t *testing.T) {
	var nilMul = &calcTestSuit{nil, 1, opMul, true, nil}
	var int8Mul = &calcTestSuit{int8(1), 1, opMul, true, nil}
	var int16Mul = &calcTestSuit{int16(1), 1, opMul, true, nil}
	var int32Mul = &calcTestSuit{int32(1), 1, opMul, true, nil}
	var byteMul = &calcTestSuit{byte('1'), 1, opMul, true, nil}
	var float32Mul = &calcTestSuit{float32(0.5), 1, opMul, true, nil}

	testCalc(nilMul, t)
	testCalc(int8Mul, t)
	testCalc(int16Mul, t)
	testCalc(int32Mul, t)
	testCalc(byteMul, t)
	testCalc(float32Mul, t)
}

func TestIntDiv(t *testing.T) {
	var divInt = &calcTestSuit{int(4), int(2), opDiv, false, int64(2)}
	var divNegativeInt = &calcTestSuit{int(4), int(-2), opDiv, false, int64(-2)}
	var divIntZero = &calcTestSuit{int(1), int(2), opDiv, false, int64(0)}
	var divIntPanic = &calcTestSuit{math.MaxInt64, int(0), opDiv, true, nil}
	var divInt64 = &calcTestSuit{int(4), int64(2), opDiv, false, int64(2)}
	var divNegativeInt64 = &calcTestSuit{int(4), int64(-2), opDiv, false, int64(-2)}
	var divInt64Panic = &calcTestSuit{math.MaxInt64, int64(0), opDiv, true, nil}
	var divFloat64 = &calcTestSuit{int(2), float64(0.5), opDiv, false, float64(4)}
	var divNegativeFloat64 = &calcTestSuit{int(2), float64(-0.5), opDiv, false, float64(-4)}
	var divFloat64Panic = &calcTestSuit{int(1), float64(0), opDiv, false, math.Inf(1)}
	var divFloat64Overflow = &calcTestSuit{int(1), math.Inf(1), opDiv, false, float64(0)}

	var divString = &calcTestSuit{int(1), "test", opDiv, true, nil}
	var divBool = &calcTestSuit{int(1), true, opDiv, true, nil}
	var divNil = &calcTestSuit{int(1), nil, opDiv, true, nil}
	var divInt8 = &calcTestSuit{int(1), int8(1), opDiv, true, nil}
	var divInt16 = &calcTestSuit{int(1), int16(1), opDiv, true, nil}
	var divInt32 = &calcTestSuit{int(1), int32(1), opDiv, true, nil}
	var divFloat32 = &calcTestSuit{int(1), float32(0.5), opDiv, true, nil}
	var divByte = &calcTestSuit{int(1), byte('a'), opDiv, true, nil}

	testCalc(divInt, t)
	testCalc(divNegativeInt, t)
	testCalc(divIntZero, t)
	testCalc(divIntPanic, t)
	testCalc(divInt64, t)
	testCalc(divNegativeInt64, t)
	testCalc(divInt64Panic, t)
	testCalc(divFloat64, t)
	testCalc(divNegativeFloat64, t)
	testCalc(divFloat64Panic, t)
	testCalc(divFloat64Overflow, t)
	testCalc(divString, t)
	testCalc(divBool, t)
	testCalc(divNil, t)
	testCalc(divInt8, t)
	testCalc(divInt16, t)
	testCalc(divInt32, t)
	testCalc(divFloat32, t)
	testCalc(divByte, t)
}

func TestInt64Div(t *testing.T) {
	var divInt = &calcTestSuit{int64(4), int(2), opDiv, false, int64(2)}
	var divNegativeInt = &calcTestSuit{int64(4), int(-2), opDiv, false, int64(-2)}
	var divIntZero = &calcTestSuit{int64(1), int(2), opDiv, false, int64(0)}
	var divIntPanic = &calcTestSuit{int64(math.MaxInt64), int(0), opDiv, true, nil}
	var divInt64 = &calcTestSuit{int64(4), int64(2), opDiv, false, int64(2)}
	var divNegativeInt64 = &calcTestSuit{int64(4), int64(-2), opDiv, false, int64(-2)}
	var divInt64Panic = &calcTestSuit{int64(math.MaxInt64), int64(0), opDiv, true, nil}
	var divFloat64 = &calcTestSuit{int64(2), float64(0.5), opDiv, false, float64(4)}
	var divNegativeFloat64 = &calcTestSuit{int64(2), float64(-0.5), opDiv, false, float64(-4)}
	var divFloat64Panic = &calcTestSuit{int64(1), float64(0), opDiv, false, math.Inf(1)}
	var divFloat64Overflow = &calcTestSuit{int64(1), math.Inf(1), opDiv, false, float64(0)}

	var divString = &calcTestSuit{int64(1), "test", opDiv, true, nil}
	var divBool = &calcTestSuit{int64(1), true, opDiv, true, nil}
	var divNil = &calcTestSuit{int64(1), nil, opDiv, true, nil}
	var divInt8 = &calcTestSuit{int64(1), int8(1), opDiv, true, nil}
	var divInt16 = &calcTestSuit{int64(1), int16(1), opDiv, true, nil}
	var divInt32 = &calcTestSuit{int64(1), int32(1), opDiv, true, nil}
	var divFloat32 = &calcTestSuit{int64(1), float32(0.5), opDiv, true, nil}
	var divByte = &calcTestSuit{int64(1), byte('a'), opDiv, true, nil}

	testCalc(divInt, t)
	testCalc(divNegativeInt, t)
	testCalc(divIntZero, t)
	testCalc(divIntPanic, t)
	testCalc(divInt64, t)
	testCalc(divNegativeInt64, t)
	testCalc(divInt64Panic, t)
	testCalc(divFloat64, t)
	testCalc(divNegativeFloat64, t)
	testCalc(divFloat64Panic, t)
	testCalc(divFloat64Overflow, t)
	testCalc(divString, t)
	testCalc(divBool, t)
	testCalc(divNil, t)
	testCalc(divInt8, t)
	testCalc(divInt16, t)
	testCalc(divInt32, t)
	testCalc(divFloat32, t)
	testCalc(divByte, t)
}

func TestFloat64Div(t *testing.T) {
	var divInt = &calcTestSuit{float64(1.5), int(3), opDiv, false, float64(0.5)}
	var divNegativeInt = &calcTestSuit{float64(1.5), int(-3), opDiv, false, float64(-0.5)}
	var divIntOverflow = &calcTestSuit{float64(1.5), int(0), opDiv, false, math.Inf(1)}
	var divInt64 = &calcTestSuit{float64(1.5), int64(3), opDiv, false, float64(0.5)}
	var divNegativeInt64 = &calcTestSuit{float64(1.5), int64(-3), opDiv, false, float64(-0.5)}
	var divInt64Overflow = &calcTestSuit{float64(1.5), int64(0), opDiv, false, math.Inf(1)}
	var divFloat64 = &calcTestSuit{float64(2.5), float64(0.5), opDiv, false, float64(5)}
	var divFloatOverflow = &calcTestSuit{-math.MaxFloat64, math.SmallestNonzeroFloat64, opDiv, false, math.Inf(-1)}
	var divNegativeFloat64 = &calcTestSuit{float64(2.5), float64(-0.5), opDiv, false, float64(-5)}

	var divString = &calcTestSuit{float64(0.5), "test", opDiv, true, nil}
	var divBool = &calcTestSuit{float64(1), true, opDiv, true, nil}
	var divNil = &calcTestSuit{float64(1), nil, opDiv, true, nil}
	var divInt8 = &calcTestSuit{float64(1), int8(1), opDiv, true, nil}
	var divInt16 = &calcTestSuit{float64(1), int16(1), opDiv, true, nil}
	var divInt32 = &calcTestSuit{float64(1), int32(1), opDiv, true, nil}
	var divFloat32 = &calcTestSuit{float64(1), float32(0.5), opDiv, true, nil}
	var divByte = &calcTestSuit{float64(1), byte('a'), opDiv, true, nil}

	testCalc(divInt, t)
	testCalc(divNegativeInt, t)
	testCalc(divIntOverflow, t)
	testCalc(divInt64, t)
	testCalc(divNegativeInt64, t)
	testCalc(divInt64Overflow, t)
	testCalc(divFloat64, t)
	testCalc(divFloatOverflow, t)
	testCalc(divNegativeFloat64, t)
	testCalc(divString, t)
	testCalc(divBool, t)
	testCalc(divNil, t)
	testCalc(divInt8, t)
	testCalc(divInt16, t)
	testCalc(divInt32, t)
	testCalc(divFloat32, t)
	testCalc(divByte, t)
}

func TestString64Div(t *testing.T) {
	var divInt = &calcTestSuit{"test", int(1), opDiv, true, nil}
	var divInt64 = &calcTestSuit{"test", int64(1), opDiv, true, nil}
	var divFloat64 = &calcTestSuit{"test", float64(0.5), opDiv, true, nil}
	var divString = &calcTestSuit{"test", "test", opDiv, true, nil}
	var divBool = &calcTestSuit{"test", true, opDiv, true, nil}
	var divNil = &calcTestSuit{"test", nil, opDiv, true, nil}
	var divInt8 = &calcTestSuit{"test", int8(1), opDiv, true, nil}
	var divInt16 = &calcTestSuit{"test", int16(1), opDiv, true, nil}
	var divInt32 = &calcTestSuit{"test", int32(1), opDiv, true, nil}
	var divFloat32 = &calcTestSuit{"test", float32(0.5), opDiv, true, nil}
	var divByte = &calcTestSuit{"test", byte('a'), opDiv, true, nil}

	testCalc(divInt, t)
	testCalc(divInt64, t)
	testCalc(divFloat64, t)
	testCalc(divString, t)
	testCalc(divBool, t)
	testCalc(divNil, t)
	testCalc(divInt8, t)
	testCalc(divInt16, t)
	testCalc(divInt32, t)
	testCalc(divFloat32, t)
	testCalc(divByte, t)
}

func TestBoolDiv(t *testing.T) {
	var divString = &calcTestSuit{true, "test", opDiv, true, nil}
	var divInt = &calcTestSuit{true, int(1), opDiv, true, nil}
	var divInt64 = &calcTestSuit{true, int64(1), opDiv, true, nil}
	var divFloat64 = &calcTestSuit{true, float64(0.5), opDiv, true, nil}
	var divBool = &calcTestSuit{true, true, opDiv, true, nil}
	var divNil = &calcTestSuit{true, nil, opDiv, true, nil}
	var divInt8 = &calcTestSuit{true, int8(1), opDiv, true, nil}
	var divInt16 = &calcTestSuit{true, int16(1), opDiv, true, nil}
	var divInt32 = &calcTestSuit{true, int32(1), opDiv, true, nil}
	var divFloat32 = &calcTestSuit{true, float32(0.5), opDiv, true, nil}
	var divByte = &calcTestSuit{true, byte('a'), opDiv, true, nil}

	testCalc(divInt, t)
	testCalc(divInt64, t)
	testCalc(divFloat64, t)
	testCalc(divString, t)
	testCalc(divBool, t)
	testCalc(divNil, t)
	testCalc(divInt8, t)
	testCalc(divInt16, t)
	testCalc(divInt32, t)
	testCalc(divFloat32, t)
	testCalc(divByte, t)
}

func TestOtherDiv(t *testing.T) {
	var nilDiv = &calcTestSuit{nil, 1, opDiv, true, nil}
	var int8Div = &calcTestSuit{int8(1), 1, opDiv, true, nil}
	var int16Div = &calcTestSuit{int16(1), 1, opDiv, true, nil}
	var int32Div = &calcTestSuit{int32(1), 1, opDiv, true, nil}
	var byteDiv = &calcTestSuit{byte('1'), 1, opDiv, true, nil}
	var float32Div = &calcTestSuit{float32(0.5), 1, opDiv, true, nil}

	testCalc(nilDiv, t)
	testCalc(int8Div, t)
	testCalc(int16Div, t)
	testCalc(int32Div, t)
	testCalc(byteDiv, t)
	testCalc(float32Div, t)
}

func TestIntMod(t *testing.T) {
	var modInt = &calcTestSuit{int(5), int(3), opMod, false, int64(2)}
	var modNegativeInt = &calcTestSuit{int(5), int(-3), opMod, false, int64(2)}
	var modIntZero = &calcTestSuit{int(5), int(0), opMod, true, nil}
	var modInt64 = &calcTestSuit{int(5), int64(3), opMod, false, int64(2)}
	var modNegativeInt64 = &calcTestSuit{int(5), int64(-3), opMod, false, int64(2)}
	var modInt64Panic = &calcTestSuit{math.MaxInt64, int64(0), opMod, true, nil}
	var modFloat64 = &calcTestSuit{int(2), float64(0.5), opMod, true, nil}

	var modString = &calcTestSuit{int(1), "test", opMod, true, nil}
	var modBool = &calcTestSuit{int(1), true, opMod, true, nil}
	var modNil = &calcTestSuit{int(1), nil, opMod, true, nil}
	var modInt8 = &calcTestSuit{int(1), int8(1), opMod, true, nil}
	var modInt16 = &calcTestSuit{int(1), int16(1), opMod, true, nil}
	var modInt32 = &calcTestSuit{int(1), int32(1), opMod, true, nil}
	var modFloat32 = &calcTestSuit{int(1), float32(0.5), opMod, true, nil}
	var modByte = &calcTestSuit{int(1), byte('a'), opMod, true, nil}

	testCalc(modInt, t)
	testCalc(modNegativeInt, t)
	testCalc(modIntZero, t)
	testCalc(modInt64, t)
	testCalc(modNegativeInt64, t)
	testCalc(modInt64Panic, t)
	testCalc(modFloat64, t)
	testCalc(modString, t)
	testCalc(modBool, t)
	testCalc(modNil, t)
	testCalc(modInt8, t)
	testCalc(modInt16, t)
	testCalc(modInt32, t)
	testCalc(modFloat32, t)
	testCalc(modByte, t)
}

func TestInt64Mod(t *testing.T) {
	var modInt = &calcTestSuit{int64(5), int(3), opMod, false, int64(2)}
	var modNegativeInt = &calcTestSuit{int64(5), int(-3), opMod, false, int64(2)}
	var modIntZero = &calcTestSuit{int64(5), int(0), opMod, true, nil}
	var modInt64 = &calcTestSuit{int64(5), int64(3), opMod, false, int64(2)}
	var modNegativeInt64 = &calcTestSuit{int64(5), int64(-3), opMod, false, int64(2)}
	var modInt64Panic = &calcTestSuit{int64(math.MaxInt64), int64(0), opMod, true, nil}
	var modFloat64 = &calcTestSuit{int64(2), float64(0.5), opMod, true, nil}

	var modString = &calcTestSuit{int64(1), "test", opMod, true, nil}
	var modBool = &calcTestSuit{int64(1), true, opMod, true, nil}
	var modNil = &calcTestSuit{int64(1), nil, opMod, true, nil}
	var modInt8 = &calcTestSuit{int64(1), int8(1), opMod, true, nil}
	var modInt16 = &calcTestSuit{int64(1), int16(1), opMod, true, nil}
	var modInt32 = &calcTestSuit{int64(1), int32(1), opMod, true, nil}
	var modFloat32 = &calcTestSuit{int64(1), float32(0.5), opMod, true, nil}
	var modByte = &calcTestSuit{int64(1), byte('a'), opMod, true, nil}

	testCalc(modInt, t)
	testCalc(modNegativeInt, t)
	testCalc(modIntZero, t)
	testCalc(modInt64, t)
	testCalc(modNegativeInt64, t)
	testCalc(modInt64Panic, t)
	testCalc(modFloat64, t)
	testCalc(modString, t)
	testCalc(modBool, t)
	testCalc(modNil, t)
	testCalc(modInt8, t)
	testCalc(modInt16, t)
	testCalc(modInt32, t)
	testCalc(modFloat32, t)
	testCalc(modByte, t)
}

func TestFloat64Mod(t *testing.T) {
	var modInt = &calcTestSuit{float64(1.5), int(3), opMod, true, nil}
	var modInt64 = &calcTestSuit{float64(1.5), int64(3), opMod, true, nil}
	var modFloat64 = &calcTestSuit{float64(2.5), float64(0.5), opMod, true, nil}
	var modString = &calcTestSuit{float64(0.5), "test", opMod, true, nil}
	var modBool = &calcTestSuit{float64(1), true, opMod, true, nil}
	var modNil = &calcTestSuit{float64(1), nil, opMod, true, nil}
	var modInt8 = &calcTestSuit{float64(1), int8(1), opMod, true, nil}
	var modInt16 = &calcTestSuit{float64(1), int16(1), opMod, true, nil}
	var modInt32 = &calcTestSuit{float64(1), int32(1), opMod, true, nil}
	var modFloat32 = &calcTestSuit{float64(1), float32(0.5), opMod, true, nil}
	var modByte = &calcTestSuit{float64(1), byte('a'), opMod, true, nil}

	testCalc(modInt, t)
	testCalc(modInt64, t)
	testCalc(modFloat64, t)
	testCalc(modString, t)
	testCalc(modBool, t)
	testCalc(modNil, t)
	testCalc(modInt8, t)
	testCalc(modInt16, t)
	testCalc(modInt32, t)
	testCalc(modFloat32, t)
	testCalc(modByte, t)
}

func TestString64Mod(t *testing.T) {
	var modInt = &calcTestSuit{"test", int(1), opMod, true, nil}
	var modInt64 = &calcTestSuit{"test", int64(1), opMod, true, nil}
	var modFloat64 = &calcTestSuit{"test", float64(0.5), opMod, true, nil}
	var modString = &calcTestSuit{"test", "test", opMod, true, nil}
	var modBool = &calcTestSuit{"test", true, opMod, true, nil}
	var modNil = &calcTestSuit{"test", nil, opMod, true, nil}
	var modInt8 = &calcTestSuit{"test", int8(1), opMod, true, nil}
	var modInt16 = &calcTestSuit{"test", int16(1), opMod, true, nil}
	var modInt32 = &calcTestSuit{"test", int32(1), opMod, true, nil}
	var modFloat32 = &calcTestSuit{"test", float32(0.5), opMod, true, nil}
	var modByte = &calcTestSuit{"test", byte('a'), opMod, true, nil}

	testCalc(modInt, t)
	testCalc(modInt64, t)
	testCalc(modFloat64, t)
	testCalc(modString, t)
	testCalc(modBool, t)
	testCalc(modNil, t)
	testCalc(modInt8, t)
	testCalc(modInt16, t)
	testCalc(modInt32, t)
	testCalc(modFloat32, t)
	testCalc(modByte, t)
}

func TestBoolMod(t *testing.T) {
	var modString = &calcTestSuit{true, "test", opMod, true, nil}
	var modInt = &calcTestSuit{true, int(1), opMod, true, nil}
	var modInt64 = &calcTestSuit{true, int64(1), opMod, true, nil}
	var modFloat64 = &calcTestSuit{true, float64(0.5), opMod, true, nil}
	var modBool = &calcTestSuit{true, true, opMod, true, nil}
	var modNil = &calcTestSuit{true, nil, opMod, true, nil}
	var modInt8 = &calcTestSuit{true, int8(1), opMod, true, nil}
	var modInt16 = &calcTestSuit{true, int16(1), opMod, true, nil}
	var modInt32 = &calcTestSuit{true, int32(1), opMod, true, nil}
	var modFloat32 = &calcTestSuit{true, float32(0.5), opMod, true, nil}
	var modByte = &calcTestSuit{true, byte('a'), opMod, true, nil}

	testCalc(modInt, t)
	testCalc(modInt64, t)
	testCalc(modFloat64, t)
	testCalc(modString, t)
	testCalc(modBool, t)
	testCalc(modNil, t)
	testCalc(modInt8, t)
	testCalc(modInt16, t)
	testCalc(modInt32, t)
	testCalc(modFloat32, t)
	testCalc(modByte, t)
}

func TestOtherMod(t *testing.T) {
	var nilMod = &calcTestSuit{nil, 1, opMod, true, nil}
	var int8Mod = &calcTestSuit{int8(1), 1, opMod, true, nil}
	var int16Mod = &calcTestSuit{int16(1), 1, opMod, true, nil}
	var int32Mod = &calcTestSuit{int32(1), 1, opMod, true, nil}
	var byteMod = &calcTestSuit{byte('1'), 1, opMod, true, nil}
	var float32Mod = &calcTestSuit{float32(0.5), 1, opMod, true, nil}

	testCalc(nilMod, t)
	testCalc(int8Mod, t)
	testCalc(int16Mod, t)
	testCalc(int32Mod, t)
	testCalc(byteMod, t)
	testCalc(float32Mod, t)
}
