package fate

import (
	"fmt"
	"runtime/debug"
	"testing"
)

type cmpTestSuit struct {
	leftValue   interface{}
	rightValue  interface{}
	op          string
	shouldPanic bool
	expect      interface{}
}

func (suit *cmpTestSuit) String() string {
	return fmt.Sprintf("%T(%v) %s %T(%v) = %T(%v)", suit.leftValue, suit.leftValue, suit.op, suit.rightValue, suit.rightValue, suit.expect, suit.expect)
}

func testCmp(suit *cmpTestSuit, t *testing.T) {
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

	got := compare(suit.leftValue, suit.rightValue, suit.op)
	if got != suit.expect {
		t.Error(fmt.Sprintf("test %s but got %T(%v)", suit.String(), got, got))
	}
	return
}

func TestIntEqual(t *testing.T) {
	var equalInt_1 = &cmpTestSuit{int(1), int(1), opEqual, false, true}
	var equalInt_2 = &cmpTestSuit{int(1), int(2), opEqual, false, false}
	var equalInt64_1 = &cmpTestSuit{int(1), int64(1), opEqual, false, true}
	var equalInt64_2 = &cmpTestSuit{int(1), int64(2), opEqual, false, false}
	var equalFloat64_1 = &cmpTestSuit{int(1), float64(1), opEqual, false, true}
	var equalFloat64_2 = &cmpTestSuit{int(1), float64(0.99), opEqual, false, false}
	var equalNil_1 = &cmpTestSuit{int(1), nil, opEqual, false, false}
	var equalBool_1 = &cmpTestSuit{int(1), true, opEqual, false, true}
	var equalBool_2 = &cmpTestSuit{int(0), true, opEqual, false, false}

	var equalString = &cmpTestSuit{int(1), "test", opEqual, true, nil}
	var equalInt8 = &cmpTestSuit{int(1), int8(1), opEqual, true, nil}
	var equalInt16 = &cmpTestSuit{int(1), int16(1), opEqual, true, nil}
	var equalInt32 = &cmpTestSuit{int(1), int32(1), opEqual, true, nil}
	var equalFloat32 = &cmpTestSuit{int(1), float32(0.5), opEqual, true, nil}
	var equalByte = &cmpTestSuit{int(1), byte('a'), opEqual, true, nil}

	testCmp(equalInt_1, t)
	testCmp(equalInt_2, t)
	testCmp(equalInt64_1, t)
	testCmp(equalInt64_2, t)
	testCmp(equalFloat64_1, t)
	testCmp(equalFloat64_2, t)
	testCmp(equalNil_1, t)
	testCmp(equalString, t)
	testCmp(equalBool_1, t)
	testCmp(equalBool_2, t)
	testCmp(equalInt8, t)
	testCmp(equalInt16, t)
	testCmp(equalInt32, t)
	testCmp(equalFloat32, t)
	testCmp(equalByte, t)
}

func TestInt64Equal(t *testing.T) {
	var equalInt_1 = &cmpTestSuit{int64(1), int(1), opEqual, false, true}
	var equalInt_2 = &cmpTestSuit{int64(1), int(2), opEqual, false, false}
	var equalInt64_1 = &cmpTestSuit{int64(1), int64(1), opEqual, false, true}
	var equalInt64_2 = &cmpTestSuit{int64(1), int64(2), opEqual, false, false}
	var equalFloat64_1 = &cmpTestSuit{int64(1), float64(1), opEqual, false, true}
	var equalFloat64_2 = &cmpTestSuit{int64(1), float64(0.99), opEqual, false, false}
	var equalNil_1 = &cmpTestSuit{int64(1), nil, opEqual, false, false}
	var equalBool_1 = &cmpTestSuit{int64(1), true, opEqual, false, true}
	var equalBool_2 = &cmpTestSuit{int64(0), true, opEqual, false, false}

	var equalString = &cmpTestSuit{int64(1), "test", opEqual, true, nil}
	var equalInt8 = &cmpTestSuit{int64(1), int8(1), opEqual, true, nil}
	var equalInt16 = &cmpTestSuit{int64(1), int16(1), opEqual, true, nil}
	var equalInt32 = &cmpTestSuit{int64(1), int32(1), opEqual, true, nil}
	var equalFloat32 = &cmpTestSuit{int64(1), float32(0.5), opEqual, true, nil}
	var equalByte = &cmpTestSuit{int64(1), byte('a'), opEqual, true, nil}

	testCmp(equalInt_1, t)
	testCmp(equalInt_2, t)
	testCmp(equalInt64_1, t)
	testCmp(equalInt64_2, t)
	testCmp(equalFloat64_1, t)
	testCmp(equalFloat64_2, t)
	testCmp(equalNil_1, t)
	testCmp(equalString, t)
	testCmp(equalBool_1, t)
	testCmp(equalBool_2, t)
	testCmp(equalInt8, t)
	testCmp(equalInt16, t)
	testCmp(equalInt32, t)
	testCmp(equalFloat32, t)
	testCmp(equalByte, t)
}

func TestFloat64Equal(t *testing.T) {
	var equalInt_1 = &cmpTestSuit{float64(1), int(1), opEqual, false, true}
	var equalInt_2 = &cmpTestSuit{float64(1.5), int(2), opEqual, false, false}
	var equalInt64_1 = &cmpTestSuit{float64(1), int64(1), opEqual, false, true}
	var equalInt64_2 = &cmpTestSuit{float64(1.5), int64(2), opEqual, false, false}
	var equalFloat64_1 = &cmpTestSuit{float64(1.3), float64(1.3), opEqual, false, true}
	var equalFloat64_2 = &cmpTestSuit{float64(1.2), float64(0.99), opEqual, false, false}
	var equalNil_1 = &cmpTestSuit{float64(1.1), nil, opEqual, false, false}
	var equalBool_1 = &cmpTestSuit{float64(1), true, opEqual, false, true}
	var equalBool_2 = &cmpTestSuit{float64(0), true, opEqual, false, false}

	var equalString = &cmpTestSuit{float64(1), "test", opEqual, true, nil}
	var equalInt8 = &cmpTestSuit{float64(1), int8(1), opEqual, true, nil}
	var equalInt16 = &cmpTestSuit{float64(1), int16(1), opEqual, true, nil}
	var equalInt32 = &cmpTestSuit{float64(1), int32(1), opEqual, true, nil}
	var equalFloat32 = &cmpTestSuit{float64(1), float32(0.5), opEqual, true, nil}
	var equalByte = &cmpTestSuit{float64(1), byte('a'), opEqual, true, nil}

	testCmp(equalInt_1, t)
	testCmp(equalInt_2, t)
	testCmp(equalInt64_1, t)
	testCmp(equalInt64_2, t)
	testCmp(equalFloat64_1, t)
	testCmp(equalFloat64_2, t)
	testCmp(equalNil_1, t)
	testCmp(equalString, t)
	testCmp(equalBool_1, t)
	testCmp(equalBool_2, t)
	testCmp(equalInt8, t)
	testCmp(equalInt16, t)
	testCmp(equalInt32, t)
	testCmp(equalFloat32, t)
	testCmp(equalByte, t)
}

func TestStringEqual(t *testing.T) {
	var equalNil = &cmpTestSuit{"test", nil, opEqual, false, false}
	var equalString_1 = &cmpTestSuit{"test", "test", opEqual, false, true}
	var equalString_2 = &cmpTestSuit{"test", "test1", opEqual, false, false}
	var equalInt_1 = &cmpTestSuit{"1", int(1), opEqual, false, true}
	var equalInt_2 = &cmpTestSuit{"test", int(1), opEqual, false, false}
	var equalInt64_1 = &cmpTestSuit{"1", int64(1), opEqual, false, true}
	var equalInt64_2 = &cmpTestSuit{"test", int64(1), opEqual, false, false}
	var equalFloat64_1 = &cmpTestSuit{"1.3", float64(1.3), opEqual, false, true}
	var equalFloat64_2 = &cmpTestSuit{"test", float64(1.3), opEqual, false, false}
	var equalBool_1 = &cmpTestSuit{"true", true, opEqual, false, true}
	var equalBool_2 = &cmpTestSuit{"test", true, opEqual, false, false}

	var equalInt8 = &cmpTestSuit{"test", int8(1), opEqual, true, nil}
	var equalInt16 = &cmpTestSuit{"test", int16(1), opEqual, true, nil}
	var equalInt32 = &cmpTestSuit{"test", int32(1), opEqual, true, nil}
	var equalFloat32 = &cmpTestSuit{"test", float32(0.5), opEqual, true, nil}
	var equalByte = &cmpTestSuit{"test", byte('a'), opEqual, true, nil}

	testCmp(equalNil, t)
	testCmp(equalString_1, t)
	testCmp(equalString_2, t)
	testCmp(equalInt_1, t)
	testCmp(equalInt_2, t)
	testCmp(equalInt64_1, t)
	testCmp(equalInt64_2, t)
	testCmp(equalFloat64_1, t)
	testCmp(equalFloat64_2, t)
	testCmp(equalBool_1, t)
	testCmp(equalBool_2, t)
	testCmp(equalInt8, t)
	testCmp(equalInt16, t)
	testCmp(equalInt32, t)
	testCmp(equalFloat32, t)
	testCmp(equalByte, t)
}

func TestBoolEqual(t *testing.T) {
	var equalInt_1 = &cmpTestSuit{true, int(1), opEqual, false, true}
	var equalInt_2 = &cmpTestSuit{true, int(0), opEqual, false, false}
	var equalInt64_1 = &cmpTestSuit{true, int64(1), opEqual, false, true}
	var equalInt64_2 = &cmpTestSuit{true, int64(0), opEqual, false, false}
	var equalFloat64_1 = &cmpTestSuit{true, float64(1.3), opEqual, false, true}
	var equalFloat64_2 = &cmpTestSuit{true, float64(0), opEqual, false, false}
	var equalString_1 = &cmpTestSuit{true, "true", opEqual, false, true}
	var equalString_2 = &cmpTestSuit{true, "test", opEqual, false, false}
	var equalBool_1 = &cmpTestSuit{true, true, opEqual, false, true}
	var equalBool_2 = &cmpTestSuit{true, false, opEqual, false, false}
	var equalNil = &cmpTestSuit{true, nil, opEqual, false, false}

	var equalInt8 = &cmpTestSuit{true, int8(1), opEqual, true, nil}
	var equalInt16 = &cmpTestSuit{true, int16(1), opEqual, true, nil}
	var equalInt32 = &cmpTestSuit{true, int32(1), opEqual, true, nil}
	var equalFloat32 = &cmpTestSuit{true, float32(0.5), opEqual, true, nil}
	var equalByte = &cmpTestSuit{true, byte('a'), opEqual, true, nil}

	testCmp(equalInt_1, t)
	testCmp(equalInt_2, t)
	testCmp(equalInt64_1, t)
	testCmp(equalInt64_2, t)
	testCmp(equalFloat64_1, t)
	testCmp(equalFloat64_2, t)
	testCmp(equalString_1, t)
	testCmp(equalString_2, t)
	testCmp(equalBool_1, t)
	testCmp(equalBool_2, t)
	testCmp(equalNil, t)
	testCmp(equalInt8, t)
	testCmp(equalInt16, t)
	testCmp(equalInt32, t)
	testCmp(equalFloat32, t)
	testCmp(equalByte, t)
}

func TestNilEqual(t *testing.T) {
	var equalNil = &cmpTestSuit{nil, nil, opEqual, false, true}
	var equalBool = &cmpTestSuit{nil, true, opEqual, false, false}
	var equalInt = &cmpTestSuit{nil, int(1), opEqual, false, false}
	var equalInt64 = &cmpTestSuit{nil, int64(1), opEqual, false, false}
	var equalFloat64 = &cmpTestSuit{nil, float64(1.3), opEqual, false, false}
	var equalString = &cmpTestSuit{nil, "test", opEqual, false, false}
	var equalInt8 = &cmpTestSuit{nil, int8(1), opEqual, false, false}
	var equalInt16 = &cmpTestSuit{nil, int16(1), opEqual, false, false}
	var equalInt32 = &cmpTestSuit{nil, int32(1), opEqual, false, false}
	var equalFloat32 = &cmpTestSuit{nil, float32(0.5), opEqual, false, false}
	var equalByte = &cmpTestSuit{nil, byte('a'), opEqual, false, false}

	testCmp(equalNil, t)
	testCmp(equalBool, t)
	testCmp(equalInt, t)
	testCmp(equalInt64, t)
	testCmp(equalFloat64, t)
	testCmp(equalString, t)
	testCmp(equalInt8, t)
	testCmp(equalInt16, t)
	testCmp(equalInt32, t)
	testCmp(equalFloat32, t)
	testCmp(equalByte, t)
}

func TestIntNotEqual(t *testing.T) {
	var notEqualInt_1 = &cmpTestSuit{int(1), int(1), opNotEqual, false, false}
	var notEqualInt_2 = &cmpTestSuit{int(1), int(2), opNotEqual, false, true}
	var notEqualInt64_1 = &cmpTestSuit{int(1), int64(1), opNotEqual, false, false}
	var notEqualInt64_2 = &cmpTestSuit{int(1), int64(2), opNotEqual, false, true}
	var notEqualFloat64_1 = &cmpTestSuit{int(1), float64(1), opNotEqual, false, false}
	var notEqualFloat64_2 = &cmpTestSuit{int(1), float64(0.99), opNotEqual, false, true}
	var notEqualNil_1 = &cmpTestSuit{int(1), nil, opNotEqual, false, true}
	var notEqualBool_1 = &cmpTestSuit{int(1), true, opNotEqual, false, false}
	var notEqualBool_2 = &cmpTestSuit{int(0), true, opNotEqual, false, true}

	var notEqualString = &cmpTestSuit{int(1), "test", opNotEqual, true, nil}
	var notEqualInt8 = &cmpTestSuit{int(1), int8(1), opNotEqual, true, nil}
	var notEqualInt16 = &cmpTestSuit{int(1), int16(1), opNotEqual, true, nil}
	var notEqualInt32 = &cmpTestSuit{int(1), int32(1), opNotEqual, true, nil}
	var notEqualFloat32 = &cmpTestSuit{int(1), float32(0.5), opNotEqual, true, nil}
	var notEqualByte = &cmpTestSuit{int(1), byte('a'), opNotEqual, true, nil}

	testCmp(notEqualInt_1, t)
	testCmp(notEqualInt_2, t)
	testCmp(notEqualInt64_1, t)
	testCmp(notEqualInt64_2, t)
	testCmp(notEqualFloat64_1, t)
	testCmp(notEqualFloat64_2, t)
	testCmp(notEqualNil_1, t)
	testCmp(notEqualString, t)
	testCmp(notEqualBool_1, t)
	testCmp(notEqualBool_2, t)
	testCmp(notEqualInt8, t)
	testCmp(notEqualInt16, t)
	testCmp(notEqualInt32, t)
	testCmp(notEqualFloat32, t)
	testCmp(notEqualByte, t)
}

func TestInt64NotEqual(t *testing.T) {
	var notEqualInt_1 = &cmpTestSuit{int64(1), int(1), opNotEqual, false, false}
	var notEqualInt_2 = &cmpTestSuit{int64(1), int(2), opNotEqual, false, true}
	var notEqualInt64_1 = &cmpTestSuit{int64(1), int64(1), opNotEqual, false, false}
	var notEqualInt64_2 = &cmpTestSuit{int64(1), int64(2), opNotEqual, false, true}
	var notEqualFloat64_1 = &cmpTestSuit{int64(1), float64(1), opNotEqual, false, false}
	var notEqualFloat64_2 = &cmpTestSuit{int64(1), float64(0.99), opNotEqual, false, true}
	var notEqualNil_1 = &cmpTestSuit{int64(1), nil, opNotEqual, false, true}
	var notEqualBool_1 = &cmpTestSuit{int64(1), true, opNotEqual, false, false}
	var notEqualBool_2 = &cmpTestSuit{int64(0), true, opNotEqual, false, true}

	var notEqualString = &cmpTestSuit{int64(1), "test", opNotEqual, true, nil}
	var notEqualInt8 = &cmpTestSuit{int64(1), int8(1), opNotEqual, true, nil}
	var notEqualInt16 = &cmpTestSuit{int64(1), int16(1), opNotEqual, true, nil}
	var notEqualInt32 = &cmpTestSuit{int64(1), int32(1), opNotEqual, true, nil}
	var notEqualFloat32 = &cmpTestSuit{int64(1), float32(0.5), opNotEqual, true, nil}
	var notEqualByte = &cmpTestSuit{int64(1), byte('a'), opNotEqual, true, nil}

	testCmp(notEqualInt_1, t)
	testCmp(notEqualInt_2, t)
	testCmp(notEqualInt64_1, t)
	testCmp(notEqualInt64_2, t)
	testCmp(notEqualFloat64_1, t)
	testCmp(notEqualFloat64_2, t)
	testCmp(notEqualNil_1, t)
	testCmp(notEqualString, t)
	testCmp(notEqualBool_1, t)
	testCmp(notEqualBool_2, t)
	testCmp(notEqualInt8, t)
	testCmp(notEqualInt16, t)
	testCmp(notEqualInt32, t)
	testCmp(notEqualFloat32, t)
	testCmp(notEqualByte, t)
}

func TestFloat64NotEqual(t *testing.T) {
	var notEqualInt_1 = &cmpTestSuit{float64(1), int(1), opNotEqual, false, false}
	var notEqualInt_2 = &cmpTestSuit{float64(1.5), int(2), opNotEqual, false, true}
	var notEqualInt64_1 = &cmpTestSuit{float64(1), int64(1), opNotEqual, false, false}
	var notEqualInt64_2 = &cmpTestSuit{float64(1.5), int64(2), opNotEqual, false, true}
	var notEqualFloat64_1 = &cmpTestSuit{float64(1.3), float64(1.3), opNotEqual, false, false}
	var notEqualFloat64_2 = &cmpTestSuit{float64(1.2), float64(0.99), opNotEqual, false, true}
	var notEqualNil_1 = &cmpTestSuit{float64(1.1), nil, opNotEqual, false, true}
	var notEqualBool_1 = &cmpTestSuit{float64(1), true, opNotEqual, false, false}
	var notEqualBool_2 = &cmpTestSuit{float64(0), true, opNotEqual, false, true}

	var notEqualString = &cmpTestSuit{float64(1), "test", opNotEqual, true, nil}
	var notEqualInt8 = &cmpTestSuit{float64(1), int8(1), opNotEqual, true, nil}
	var notEqualInt16 = &cmpTestSuit{float64(1), int16(1), opNotEqual, true, nil}
	var notEqualInt32 = &cmpTestSuit{float64(1), int32(1), opNotEqual, true, nil}
	var notEqualFloat32 = &cmpTestSuit{float64(1), float32(0.5), opNotEqual, true, nil}
	var notEqualByte = &cmpTestSuit{float64(1), byte('a'), opNotEqual, true, nil}

	testCmp(notEqualInt_1, t)
	testCmp(notEqualInt_2, t)
	testCmp(notEqualInt64_1, t)
	testCmp(notEqualInt64_2, t)
	testCmp(notEqualFloat64_1, t)
	testCmp(notEqualFloat64_2, t)
	testCmp(notEqualNil_1, t)
	testCmp(notEqualString, t)
	testCmp(notEqualBool_1, t)
	testCmp(notEqualBool_2, t)
	testCmp(notEqualInt8, t)
	testCmp(notEqualInt16, t)
	testCmp(notEqualInt32, t)
	testCmp(notEqualFloat32, t)
	testCmp(notEqualByte, t)
}

func TestStringNotEqual(t *testing.T) {
	var notEqualNil = &cmpTestSuit{"test", nil, opNotEqual, false, true}
	var notEqualString_1 = &cmpTestSuit{"test", "test", opNotEqual, false, false}
	var notEqualString_2 = &cmpTestSuit{"test", "test1", opNotEqual, false, true}
	var notEqualInt_1 = &cmpTestSuit{"1", int(1), opNotEqual, false, false}
	var notEqualInt_2 = &cmpTestSuit{"test", int(1), opNotEqual, false, true}
	var notEqualInt64_1 = &cmpTestSuit{"1", int64(1), opNotEqual, false, false}
	var notEqualInt64_2 = &cmpTestSuit{"test", int64(1), opNotEqual, false, true}
	var notEqualFloat64_1 = &cmpTestSuit{"1.3", float64(1.3), opNotEqual, false, false}
	var notEqualFloat64_2 = &cmpTestSuit{"test", float64(1.3), opNotEqual, false, true}
	var notEqualBool_1 = &cmpTestSuit{"true", true, opNotEqual, false, false}
	var notEqualBool_2 = &cmpTestSuit{"test", true, opNotEqual, false, true}

	var notEqualInt8 = &cmpTestSuit{"test", int8(1), opNotEqual, true, nil}
	var notEqualInt16 = &cmpTestSuit{"test", int16(1), opNotEqual, true, nil}
	var notEqualInt32 = &cmpTestSuit{"test", int32(1), opNotEqual, true, nil}
	var notEqualFloat32 = &cmpTestSuit{"test", float32(0.5), opNotEqual, true, nil}
	var notEqualByte = &cmpTestSuit{"test", byte('a'), opNotEqual, true, nil}

	testCmp(notEqualNil, t)
	testCmp(notEqualString_1, t)
	testCmp(notEqualString_2, t)
	testCmp(notEqualInt_1, t)
	testCmp(notEqualInt_2, t)
	testCmp(notEqualInt64_1, t)
	testCmp(notEqualInt64_2, t)
	testCmp(notEqualFloat64_1, t)
	testCmp(notEqualFloat64_2, t)
	testCmp(notEqualBool_1, t)
	testCmp(notEqualBool_2, t)
	testCmp(notEqualInt8, t)
	testCmp(notEqualInt16, t)
	testCmp(notEqualInt32, t)
	testCmp(notEqualFloat32, t)
	testCmp(notEqualByte, t)
}

func TestBoolNotEqual(t *testing.T) {
	var notEqualInt_1 = &cmpTestSuit{true, int(1), opNotEqual, false, false}
	var notEqualInt_2 = &cmpTestSuit{true, int(0), opNotEqual, false, true}
	var notEqualInt64_1 = &cmpTestSuit{true, int64(1), opNotEqual, false, false}
	var notEqualInt64_2 = &cmpTestSuit{true, int64(0), opNotEqual, false, true}
	var notEqualFloat64_1 = &cmpTestSuit{true, float64(1.3), opNotEqual, false, false}
	var notEqualFloat64_2 = &cmpTestSuit{true, float64(0), opNotEqual, false, true}
	var notEqualString_1 = &cmpTestSuit{true, "true", opNotEqual, false, false}
	var notEqualString_2 = &cmpTestSuit{true, "test", opNotEqual, false, true}
	var notEqualBool_1 = &cmpTestSuit{true, true, opNotEqual, false, false}
	var notEqualBool_2 = &cmpTestSuit{true, false, opNotEqual, false, true}
	var notEqualNil = &cmpTestSuit{true, nil, opNotEqual, false, true}

	var notEqualInt8 = &cmpTestSuit{true, int8(1), opNotEqual, true, nil}
	var notEqualInt16 = &cmpTestSuit{true, int16(1), opNotEqual, true, nil}
	var notEqualInt32 = &cmpTestSuit{true, int32(1), opNotEqual, true, nil}
	var notEqualFloat32 = &cmpTestSuit{true, float32(0.5), opNotEqual, true, nil}
	var notEqualByte = &cmpTestSuit{true, byte('a'), opNotEqual, true, nil}

	testCmp(notEqualInt_1, t)
	testCmp(notEqualInt_2, t)
	testCmp(notEqualInt64_1, t)
	testCmp(notEqualInt64_2, t)
	testCmp(notEqualFloat64_1, t)
	testCmp(notEqualFloat64_2, t)
	testCmp(notEqualString_1, t)
	testCmp(notEqualString_2, t)
	testCmp(notEqualBool_1, t)
	testCmp(notEqualBool_2, t)
	testCmp(notEqualNil, t)
	testCmp(notEqualInt8, t)
	testCmp(notEqualInt16, t)
	testCmp(notEqualInt32, t)
	testCmp(notEqualFloat32, t)
	testCmp(notEqualByte, t)
}

func TestNilNotEqual(t *testing.T) {
	var notEqualNil = &cmpTestSuit{nil, nil, opNotEqual, false, false}
	var notEqualBool = &cmpTestSuit{nil, true, opNotEqual, false, true}
	var notEqualInt = &cmpTestSuit{nil, int(1), opNotEqual, false, true}
	var notEqualInt64 = &cmpTestSuit{nil, int64(1), opNotEqual, false, true}
	var notEqualFloat64 = &cmpTestSuit{nil, float64(1.3), opNotEqual, false, true}
	var notEqualString = &cmpTestSuit{nil, "test", opNotEqual, false, true}
	var notEqualInt8 = &cmpTestSuit{nil, int8(1), opNotEqual, false, true}
	var notEqualInt16 = &cmpTestSuit{nil, int16(1), opNotEqual, false, true}
	var notEqualInt32 = &cmpTestSuit{nil, int32(1), opNotEqual, false, true}
	var notEqualFloat32 = &cmpTestSuit{nil, float32(0.5), opNotEqual, false, true}
	var notEqualByte = &cmpTestSuit{nil, byte('a'), opNotEqual, false, true}

	testCmp(notEqualNil, t)
	testCmp(notEqualBool, t)
	testCmp(notEqualInt, t)
	testCmp(notEqualInt64, t)
	testCmp(notEqualFloat64, t)
	testCmp(notEqualString, t)
	testCmp(notEqualInt8, t)
	testCmp(notEqualInt16, t)
	testCmp(notEqualInt32, t)
	testCmp(notEqualFloat32, t)
	testCmp(notEqualByte, t)
}

func TestIntLarger(t *testing.T) {
	var largerInt_1 = &cmpTestSuit{int(1), int(0), opLarger, false, true}
	var largerInt_2 = &cmpTestSuit{int(1), int(1), opLarger, false, false}
	var largerInt_3 = &cmpTestSuit{int(1), int(2), opLarger, false, false}
	var largerInt64_1 = &cmpTestSuit{int(1), int64(0), opLarger, false, true}
	var largerInt64_2 = &cmpTestSuit{int(1), int64(1), opLarger, false, false}
	var largerInt64_3 = &cmpTestSuit{int(1), int64(2), opLarger, false, false}
	var largerFloat64_1 = &cmpTestSuit{int(1), float64(0.99), opLarger, false, true}
	var largerFloat64_2 = &cmpTestSuit{int(1), float64(1), opLarger, false, false}
	var largerFloat64_3 = &cmpTestSuit{int(1), float64(1.01), opLarger, false, false}

	var largerBool = &cmpTestSuit{int(1), true, opLarger, true, nil}
	var largerNil = &cmpTestSuit{int(1), nil, opLarger, true, nil}
	var largerString = &cmpTestSuit{int(1), "test", opLarger, true, nil}
	var largerInt8 = &cmpTestSuit{int(1), int8(1), opLarger, true, nil}
	var largerInt16 = &cmpTestSuit{int(1), int16(1), opLarger, true, nil}
	var largerInt32 = &cmpTestSuit{int(1), int32(1), opLarger, true, nil}
	var largerFloat32 = &cmpTestSuit{int(1), float32(0.5), opLarger, true, nil}
	var largerByte = &cmpTestSuit{int(1), byte('a'), opLarger, true, nil}

	testCmp(largerInt_1, t)
	testCmp(largerInt_2, t)
	testCmp(largerInt_3, t)
	testCmp(largerInt64_1, t)
	testCmp(largerInt64_2, t)
	testCmp(largerInt64_3, t)
	testCmp(largerFloat64_1, t)
	testCmp(largerFloat64_2, t)
	testCmp(largerFloat64_3, t)
	testCmp(largerString, t)
	testCmp(largerBool, t)
	testCmp(largerNil, t)
	testCmp(largerInt8, t)
	testCmp(largerInt16, t)
	testCmp(largerInt32, t)
	testCmp(largerFloat32, t)
	testCmp(largerByte, t)
}

func TestInt64Larger(t *testing.T) {
	var largerInt_1 = &cmpTestSuit{int64(1), int(0), opLarger, false, true}
	var largerInt_2 = &cmpTestSuit{int64(1), int(1), opLarger, false, false}
	var largerInt_3 = &cmpTestSuit{int64(1), int(2), opLarger, false, false}
	var largerInt64_1 = &cmpTestSuit{int64(1), int64(0), opLarger, false, true}
	var largerInt64_2 = &cmpTestSuit{int64(1), int64(1), opLarger, false, false}
	var largerInt64_3 = &cmpTestSuit{int64(1), int64(2), opLarger, false, false}
	var largerFloat64_1 = &cmpTestSuit{int64(1), float64(0.99), opLarger, false, true}
	var largerFloat64_2 = &cmpTestSuit{int64(1), float64(1), opLarger, false, false}
	var largerFloat64_3 = &cmpTestSuit{int64(1), float64(1.01), opLarger, false, false}

	var largerBool = &cmpTestSuit{int64(1), true, opLarger, true, nil}
	var largerNil = &cmpTestSuit{int64(1), nil, opLarger, true, nil}
	var largerString = &cmpTestSuit{int64(1), "test", opLarger, true, nil}
	var largerInt8 = &cmpTestSuit{int64(1), int8(1), opLarger, true, nil}
	var largerInt16 = &cmpTestSuit{int64(1), int16(1), opLarger, true, nil}
	var largerInt32 = &cmpTestSuit{int64(1), int32(1), opLarger, true, nil}
	var largerFloat32 = &cmpTestSuit{int64(1), float32(0.5), opLarger, true, nil}
	var largerByte = &cmpTestSuit{int64(1), byte('a'), opLarger, true, nil}

	testCmp(largerInt_1, t)
	testCmp(largerInt_2, t)
	testCmp(largerInt_3, t)
	testCmp(largerInt64_1, t)
	testCmp(largerInt64_2, t)
	testCmp(largerInt64_3, t)
	testCmp(largerFloat64_1, t)
	testCmp(largerFloat64_2, t)
	testCmp(largerFloat64_3, t)
	testCmp(largerString, t)
	testCmp(largerBool, t)
	testCmp(largerNil, t)
	testCmp(largerInt8, t)
	testCmp(largerInt16, t)
	testCmp(largerInt32, t)
	testCmp(largerFloat32, t)
	testCmp(largerByte, t)
}

func TestFloat64Larger(t *testing.T) {
	var largerInt_1 = &cmpTestSuit{float64(1.5), int(1), opLarger, false, true}
	var largerInt_2 = &cmpTestSuit{float64(2), int(2), opLarger, false, false}
	var largerInt_3 = &cmpTestSuit{float64(2.5), int(3), opLarger, false, false}
	var largerInt64_1 = &cmpTestSuit{float64(1.5), int64(1), opLarger, false, true}
	var largerInt64_2 = &cmpTestSuit{float64(2), int64(2), opLarger, false, false}
	var largerInt64_3 = &cmpTestSuit{float64(2.5), int64(3), opLarger, false, false}
	var largerFloat64_1 = &cmpTestSuit{float64(1.3), float64(1.2), opLarger, false, true}
	var largerFloat64_2 = &cmpTestSuit{float64(1.3), float64(1.3), opLarger, false, false}
	var largerFloat64_3 = &cmpTestSuit{float64(1.3), float64(1.4), opLarger, false, false}

	var largerBool = &cmpTestSuit{float64(1), true, opLarger, true, nil}
	var largerNil = &cmpTestSuit{float64(1.1), nil, opLarger, true, nil}
	var largerString = &cmpTestSuit{float64(1), "test", opLarger, true, nil}
	var largerInt8 = &cmpTestSuit{float64(1), int8(1), opLarger, true, nil}
	var largerInt16 = &cmpTestSuit{float64(1), int16(1), opLarger, true, nil}
	var largerInt32 = &cmpTestSuit{float64(1), int32(1), opLarger, true, nil}
	var largerFloat32 = &cmpTestSuit{float64(1), float32(0.5), opLarger, true, nil}
	var largerByte = &cmpTestSuit{float64(1), byte('a'), opLarger, true, nil}

	testCmp(largerInt_1, t)
	testCmp(largerInt_2, t)
	testCmp(largerInt_3, t)
	testCmp(largerInt64_1, t)
	testCmp(largerInt64_2, t)
	testCmp(largerInt64_3, t)
	testCmp(largerFloat64_1, t)
	testCmp(largerFloat64_2, t)
	testCmp(largerFloat64_3, t)
	testCmp(largerBool, t)
	testCmp(largerNil, t)
	testCmp(largerString, t)
	testCmp(largerInt8, t)
	testCmp(largerInt16, t)
	testCmp(largerInt32, t)
	testCmp(largerFloat32, t)
	testCmp(largerByte, t)
}

func TestStringLarger(t *testing.T) {
	var largerString_1 = &cmpTestSuit{"test1", "test0", opLarger, false, true}
	var largerString_2 = &cmpTestSuit{"test1", "test1", opLarger, false, false}
	var largerString_3 = &cmpTestSuit{"test1", "test2", opLarger, false, false}
	var largerInt_1 = &cmpTestSuit{"1", int(0), opLarger, false, true}
	var largerInt_2 = &cmpTestSuit{"1", int(1), opLarger, false, false}
	var largerInt_3 = &cmpTestSuit{"1", int(2), opLarger, false, false}
	var largerInt64_1 = &cmpTestSuit{"1", int64(0), opLarger, false, true}
	var largerInt64_2 = &cmpTestSuit{"1", int64(1), opLarger, false, false}
	var largerInt64_3 = &cmpTestSuit{"1", int64(2), opLarger, false, false}
	var largerFloat64_1 = &cmpTestSuit{"1.3", float64(1.2), opLarger, false, true}
	var largerFloat64_2 = &cmpTestSuit{"1.3", float64(1.3), opLarger, false, false}
	var largerFloat64_3 = &cmpTestSuit{"1.3", float64(1.4), opLarger, false, false}

	var largerBool = &cmpTestSuit{"true", true, opLarger, true, nil}
	var largerNil = &cmpTestSuit{"test", nil, opLarger, true, nil}
	var largerInt8 = &cmpTestSuit{"test", int8(1), opLarger, true, nil}
	var largerInt16 = &cmpTestSuit{"test", int16(1), opLarger, true, nil}
	var largerInt32 = &cmpTestSuit{"test", int32(1), opLarger, true, nil}
	var largerFloat32 = &cmpTestSuit{"test", float32(0.5), opLarger, true, nil}
	var largerByte = &cmpTestSuit{"test", byte('a'), opLarger, true, nil}

	testCmp(largerString_1, t)
	testCmp(largerString_2, t)
	testCmp(largerString_3, t)
	testCmp(largerInt_1, t)
	testCmp(largerInt_2, t)
	testCmp(largerInt_3, t)
	testCmp(largerInt64_1, t)
	testCmp(largerInt64_2, t)
	testCmp(largerInt64_3, t)
	testCmp(largerFloat64_1, t)
	testCmp(largerFloat64_2, t)
	testCmp(largerFloat64_3, t)
	testCmp(largerBool, t)
	testCmp(largerNil, t)
	testCmp(largerInt8, t)
	testCmp(largerInt16, t)
	testCmp(largerInt32, t)
	testCmp(largerFloat32, t)
	testCmp(largerByte, t)
}

func TestBoolLarger(t *testing.T) {
	var largerInt = &cmpTestSuit{true, int(1), opLarger, true, nil}
	var largerInt64 = &cmpTestSuit{true, int64(1), opLarger, true, nil}
	var largerFloat64 = &cmpTestSuit{true, float64(1.3), opLarger, true, nil}
	var largerString = &cmpTestSuit{true, "true", opLarger, true, nil}
	var largerBool = &cmpTestSuit{true, true, opLarger, true, nil}
	var largerNil = &cmpTestSuit{true, nil, opLarger, true, nil}
	var largerInt8 = &cmpTestSuit{true, int8(1), opLarger, true, nil}
	var largerInt16 = &cmpTestSuit{true, int16(1), opLarger, true, nil}
	var largerInt32 = &cmpTestSuit{true, int32(1), opLarger, true, nil}
	var largerFloat32 = &cmpTestSuit{true, float32(0.5), opLarger, true, nil}
	var largerByte = &cmpTestSuit{true, byte('a'), opLarger, true, nil}

	testCmp(largerInt, t)
	testCmp(largerInt64, t)
	testCmp(largerFloat64, t)
	testCmp(largerString, t)
	testCmp(largerBool, t)
	testCmp(largerNil, t)
	testCmp(largerInt8, t)
	testCmp(largerInt16, t)
	testCmp(largerInt32, t)
	testCmp(largerFloat32, t)
	testCmp(largerByte, t)
}

func TestNilLarger(t *testing.T) {
	var largerNil = &cmpTestSuit{nil, nil, opLarger, true, nil}
	var largerBool = &cmpTestSuit{nil, true, opLarger, true, nil}
	var largerInt = &cmpTestSuit{nil, int(1), opLarger, true, nil}
	var largerInt64 = &cmpTestSuit{nil, int64(1), opLarger, true, nil}
	var largerFloat64 = &cmpTestSuit{nil, float64(1.3), opLarger, true, nil}
	var largerString = &cmpTestSuit{nil, "test", opLarger, true, nil}
	var largerInt8 = &cmpTestSuit{nil, int8(1), opLarger, true, nil}
	var largerInt16 = &cmpTestSuit{nil, int16(1), opLarger, true, nil}
	var largerInt32 = &cmpTestSuit{nil, int32(1), opLarger, true, nil}
	var largerFloat32 = &cmpTestSuit{nil, float32(0.5), opLarger, true, nil}
	var largerByte = &cmpTestSuit{nil, byte('a'), opLarger, true, nil}

	testCmp(largerNil, t)
	testCmp(largerBool, t)
	testCmp(largerInt, t)
	testCmp(largerInt64, t)
	testCmp(largerFloat64, t)
	testCmp(largerString, t)
	testCmp(largerInt8, t)
	testCmp(largerInt16, t)
	testCmp(largerInt32, t)
	testCmp(largerFloat32, t)
	testCmp(largerByte, t)
}

func TestIntLargerEqual(t *testing.T) {
	var largerEqualInt_1 = &cmpTestSuit{int(1), int(0), opLargerEqual, false, true}
	var largerEqualInt_2 = &cmpTestSuit{int(1), int(1), opLargerEqual, false, true}
	var largerEqualInt_3 = &cmpTestSuit{int(1), int(2), opLargerEqual, false, false}
	var largerEqualInt64_1 = &cmpTestSuit{int(1), int64(0), opLargerEqual, false, true}
	var largerEqualInt64_2 = &cmpTestSuit{int(1), int64(1), opLargerEqual, false, true}
	var largerEqualInt64_3 = &cmpTestSuit{int(1), int64(2), opLargerEqual, false, false}
	var largerEqualFloat64_1 = &cmpTestSuit{int(1), float64(0.99), opLargerEqual, false, true}
	var largerEqualFloat64_2 = &cmpTestSuit{int(1), float64(1), opLargerEqual, false, true}
	var largerEqualFloat64_3 = &cmpTestSuit{int(1), float64(1.01), opLargerEqual, false, false}

	var largerEqualBool = &cmpTestSuit{int(1), true, opLargerEqual, true, nil}
	var largerEqualNil = &cmpTestSuit{int(1), nil, opLargerEqual, true, nil}
	var largerEqualString = &cmpTestSuit{int(1), "test", opLargerEqual, true, nil}
	var largerEqualInt8 = &cmpTestSuit{int(1), int8(1), opLargerEqual, true, nil}
	var largerEqualInt16 = &cmpTestSuit{int(1), int16(1), opLargerEqual, true, nil}
	var largerEqualInt32 = &cmpTestSuit{int(1), int32(1), opLargerEqual, true, nil}
	var largerEqualFloat32 = &cmpTestSuit{int(1), float32(0.5), opLargerEqual, true, nil}
	var largerEqualByte = &cmpTestSuit{int(1), byte('a'), opLargerEqual, true, nil}

	testCmp(largerEqualInt_1, t)
	testCmp(largerEqualInt_2, t)
	testCmp(largerEqualInt_3, t)
	testCmp(largerEqualInt64_1, t)
	testCmp(largerEqualInt64_2, t)
	testCmp(largerEqualInt64_3, t)
	testCmp(largerEqualFloat64_1, t)
	testCmp(largerEqualFloat64_2, t)
	testCmp(largerEqualFloat64_3, t)
	testCmp(largerEqualString, t)
	testCmp(largerEqualBool, t)
	testCmp(largerEqualNil, t)
	testCmp(largerEqualInt8, t)
	testCmp(largerEqualInt16, t)
	testCmp(largerEqualInt32, t)
	testCmp(largerEqualFloat32, t)
	testCmp(largerEqualByte, t)
}

func TestInt64LargerEqual(t *testing.T) {
	var largerEqualInt_1 = &cmpTestSuit{int64(1), int(0), opLargerEqual, false, true}
	var largerEqualInt_2 = &cmpTestSuit{int64(1), int(1), opLargerEqual, false, true}
	var largerEqualInt_3 = &cmpTestSuit{int64(1), int(2), opLargerEqual, false, false}
	var largerEqualInt64_1 = &cmpTestSuit{int64(1), int64(0), opLargerEqual, false, true}
	var largerEqualInt64_2 = &cmpTestSuit{int64(1), int64(1), opLargerEqual, false, true}
	var largerEqualInt64_3 = &cmpTestSuit{int64(1), int64(2), opLargerEqual, false, false}
	var largerEqualFloat64_1 = &cmpTestSuit{int64(1), float64(0.99), opLargerEqual, false, true}
	var largerEqualFloat64_2 = &cmpTestSuit{int64(1), float64(1), opLargerEqual, false, true}
	var largerEqualFloat64_3 = &cmpTestSuit{int64(1), float64(1.01), opLargerEqual, false, false}

	var largerEqualBool = &cmpTestSuit{int64(1), true, opLargerEqual, true, nil}
	var largerEqualNil = &cmpTestSuit{int64(1), nil, opLargerEqual, true, nil}
	var largerEqualString = &cmpTestSuit{int64(1), "test", opLargerEqual, true, nil}
	var largerEqualInt8 = &cmpTestSuit{int64(1), int8(1), opLargerEqual, true, nil}
	var largerEqualInt16 = &cmpTestSuit{int64(1), int16(1), opLargerEqual, true, nil}
	var largerEqualInt32 = &cmpTestSuit{int64(1), int32(1), opLargerEqual, true, nil}
	var largerEqualFloat32 = &cmpTestSuit{int64(1), float32(0.5), opLargerEqual, true, nil}
	var largerEqualByte = &cmpTestSuit{int64(1), byte('a'), opLargerEqual, true, nil}

	testCmp(largerEqualInt_1, t)
	testCmp(largerEqualInt_2, t)
	testCmp(largerEqualInt_3, t)
	testCmp(largerEqualInt64_1, t)
	testCmp(largerEqualInt64_2, t)
	testCmp(largerEqualInt64_3, t)
	testCmp(largerEqualFloat64_1, t)
	testCmp(largerEqualFloat64_2, t)
	testCmp(largerEqualFloat64_3, t)
	testCmp(largerEqualString, t)
	testCmp(largerEqualBool, t)
	testCmp(largerEqualNil, t)
	testCmp(largerEqualInt8, t)
	testCmp(largerEqualInt16, t)
	testCmp(largerEqualInt32, t)
	testCmp(largerEqualFloat32, t)
	testCmp(largerEqualByte, t)
}

func TestFloat64LargerEqual(t *testing.T) {
	var largerEqualInt_1 = &cmpTestSuit{float64(1.5), int(1), opLargerEqual, false, true}
	var largerEqualInt_2 = &cmpTestSuit{float64(2), int(2), opLargerEqual, false, true}
	var largerEqualInt_3 = &cmpTestSuit{float64(2.5), int(3), opLargerEqual, false, false}
	var largerEqualInt64_1 = &cmpTestSuit{float64(1.5), int64(1), opLargerEqual, false, true}
	var largerEqualInt64_2 = &cmpTestSuit{float64(2), int64(2), opLargerEqual, false, true}
	var largerEqualInt64_3 = &cmpTestSuit{float64(2.5), int64(3), opLargerEqual, false, false}
	var largerEqualFloat64_1 = &cmpTestSuit{float64(1.3), float64(1.2), opLargerEqual, false, true}
	var largerEqualFloat64_2 = &cmpTestSuit{float64(1.3), float64(1.3), opLargerEqual, false, true}
	var largerEqualFloat64_3 = &cmpTestSuit{float64(1.3), float64(1.4), opLargerEqual, false, false}

	var largerEqualBool = &cmpTestSuit{float64(1), true, opLargerEqual, true, nil}
	var largerEqualNil = &cmpTestSuit{float64(1.1), nil, opLargerEqual, true, nil}
	var largerEqualString = &cmpTestSuit{float64(1), "test", opLargerEqual, true, nil}
	var largerEqualInt8 = &cmpTestSuit{float64(1), int8(1), opLargerEqual, true, nil}
	var largerEqualInt16 = &cmpTestSuit{float64(1), int16(1), opLargerEqual, true, nil}
	var largerEqualInt32 = &cmpTestSuit{float64(1), int32(1), opLargerEqual, true, nil}
	var largerEqualFloat32 = &cmpTestSuit{float64(1), float32(0.5), opLargerEqual, true, nil}
	var largerEqualByte = &cmpTestSuit{float64(1), byte('a'), opLargerEqual, true, nil}

	testCmp(largerEqualInt_1, t)
	testCmp(largerEqualInt_2, t)
	testCmp(largerEqualInt_3, t)
	testCmp(largerEqualInt64_1, t)
	testCmp(largerEqualInt64_2, t)
	testCmp(largerEqualInt64_3, t)
	testCmp(largerEqualFloat64_1, t)
	testCmp(largerEqualFloat64_2, t)
	testCmp(largerEqualFloat64_3, t)
	testCmp(largerEqualBool, t)
	testCmp(largerEqualNil, t)
	testCmp(largerEqualString, t)
	testCmp(largerEqualInt8, t)
	testCmp(largerEqualInt16, t)
	testCmp(largerEqualInt32, t)
	testCmp(largerEqualFloat32, t)
	testCmp(largerEqualByte, t)
}

func TestStringLargerEqual(t *testing.T) {
	var largerEqualString_1 = &cmpTestSuit{"test1", "test0", opLargerEqual, false, true}
	var largerEqualString_2 = &cmpTestSuit{"test1", "test1", opLargerEqual, false, true}
	var largerEqualString_3 = &cmpTestSuit{"test1", "test2", opLargerEqual, false, false}
	var largerEqualInt_1 = &cmpTestSuit{"1", int(0), opLargerEqual, false, true}
	var largerEqualInt_2 = &cmpTestSuit{"1", int(1), opLargerEqual, false, true}
	var largerEqualInt_3 = &cmpTestSuit{"1", int(2), opLargerEqual, false, false}
	var largerEqualInt64_1 = &cmpTestSuit{"1", int64(0), opLargerEqual, false, true}
	var largerEqualInt64_2 = &cmpTestSuit{"1", int64(1), opLargerEqual, false, true}
	var largerEqualInt64_3 = &cmpTestSuit{"1", int64(2), opLargerEqual, false, false}
	var largerEqualFloat64_1 = &cmpTestSuit{"1.3", float64(1.2), opLargerEqual, false, true}
	var largerEqualFloat64_2 = &cmpTestSuit{"1.3", float64(1.3), opLargerEqual, false, true}
	var largerEqualFloat64_3 = &cmpTestSuit{"1.3", float64(1.4), opLargerEqual, false, false}

	var largerEqualBool = &cmpTestSuit{"true", true, opLargerEqual, true, nil}
	var largerEqualNil = &cmpTestSuit{"test", nil, opLargerEqual, true, nil}
	var largerEqualInt8 = &cmpTestSuit{"test", int8(1), opLargerEqual, true, nil}
	var largerEqualInt16 = &cmpTestSuit{"test", int16(1), opLargerEqual, true, nil}
	var largerEqualInt32 = &cmpTestSuit{"test", int32(1), opLargerEqual, true, nil}
	var largerEqualFloat32 = &cmpTestSuit{"test", float32(0.5), opLargerEqual, true, nil}
	var largerEqualByte = &cmpTestSuit{"test", byte('a'), opLargerEqual, true, nil}

	testCmp(largerEqualString_1, t)
	testCmp(largerEqualString_2, t)
	testCmp(largerEqualString_3, t)
	testCmp(largerEqualInt_1, t)
	testCmp(largerEqualInt_2, t)
	testCmp(largerEqualInt_3, t)
	testCmp(largerEqualInt64_1, t)
	testCmp(largerEqualInt64_2, t)
	testCmp(largerEqualInt64_3, t)
	testCmp(largerEqualFloat64_1, t)
	testCmp(largerEqualFloat64_2, t)
	testCmp(largerEqualFloat64_3, t)
	testCmp(largerEqualBool, t)
	testCmp(largerEqualNil, t)
	testCmp(largerEqualInt8, t)
	testCmp(largerEqualInt16, t)
	testCmp(largerEqualInt32, t)
	testCmp(largerEqualFloat32, t)
	testCmp(largerEqualByte, t)
}

func TestBoolLargerEqual(t *testing.T) {
	var largerEqualInt = &cmpTestSuit{true, int(1), opLargerEqual, true, nil}
	var largerEqualInt64 = &cmpTestSuit{true, int64(1), opLargerEqual, true, nil}
	var largerEqualFloat64 = &cmpTestSuit{true, float64(1.3), opLargerEqual, true, nil}
	var largerEqualString = &cmpTestSuit{true, "true", opLargerEqual, true, nil}
	var largerEqualBool = &cmpTestSuit{true, true, opLargerEqual, true, nil}
	var largerEqualNil = &cmpTestSuit{true, nil, opLargerEqual, true, nil}
	var largerEqualInt8 = &cmpTestSuit{true, int8(1), opLargerEqual, true, nil}
	var largerEqualInt16 = &cmpTestSuit{true, int16(1), opLargerEqual, true, nil}
	var largerEqualInt32 = &cmpTestSuit{true, int32(1), opLargerEqual, true, nil}
	var largerEqualFloat32 = &cmpTestSuit{true, float32(0.5), opLargerEqual, true, nil}
	var largerEqualByte = &cmpTestSuit{true, byte('a'), opLargerEqual, true, nil}

	testCmp(largerEqualInt, t)
	testCmp(largerEqualInt64, t)
	testCmp(largerEqualFloat64, t)
	testCmp(largerEqualString, t)
	testCmp(largerEqualBool, t)
	testCmp(largerEqualNil, t)
	testCmp(largerEqualInt8, t)
	testCmp(largerEqualInt16, t)
	testCmp(largerEqualInt32, t)
	testCmp(largerEqualFloat32, t)
	testCmp(largerEqualByte, t)
}

func TestNilLargerEqual(t *testing.T) {
	var largerEqualNil = &cmpTestSuit{nil, nil, opLargerEqual, true, nil}
	var largerEqualBool = &cmpTestSuit{nil, true, opLargerEqual, true, nil}
	var largerEqualInt = &cmpTestSuit{nil, int(1), opLargerEqual, true, nil}
	var largerEqualInt64 = &cmpTestSuit{nil, int64(1), opLargerEqual, true, nil}
	var largerEqualFloat64 = &cmpTestSuit{nil, float64(1.3), opLargerEqual, true, nil}
	var largerEqualString = &cmpTestSuit{nil, "test", opLargerEqual, true, nil}
	var largerEqualInt8 = &cmpTestSuit{nil, int8(1), opLargerEqual, true, nil}
	var largerEqualInt16 = &cmpTestSuit{nil, int16(1), opLargerEqual, true, nil}
	var largerEqualInt32 = &cmpTestSuit{nil, int32(1), opLargerEqual, true, nil}
	var largerEqualFloat32 = &cmpTestSuit{nil, float32(0.5), opLargerEqual, true, nil}
	var largerEqualByte = &cmpTestSuit{nil, byte('a'), opLargerEqual, true, nil}

	testCmp(largerEqualNil, t)
	testCmp(largerEqualBool, t)
	testCmp(largerEqualInt, t)
	testCmp(largerEqualInt64, t)
	testCmp(largerEqualFloat64, t)
	testCmp(largerEqualString, t)
	testCmp(largerEqualInt8, t)
	testCmp(largerEqualInt16, t)
	testCmp(largerEqualInt32, t)
	testCmp(largerEqualFloat32, t)
	testCmp(largerEqualByte, t)
}

func TestIntLess(t *testing.T) {
	var lessInt_1 = &cmpTestSuit{int(1), int(0), opLess, false, false}
	var lessInt_2 = &cmpTestSuit{int(1), int(1), opLess, false, false}
	var lessInt_3 = &cmpTestSuit{int(1), int(2), opLess, false, true}
	var lessInt64_1 = &cmpTestSuit{int(1), int64(0), opLess, false, false}
	var lessInt64_2 = &cmpTestSuit{int(1), int64(1), opLess, false, false}
	var lessInt64_3 = &cmpTestSuit{int(1), int64(2), opLess, false, true}
	var lessFloat64_1 = &cmpTestSuit{int(1), float64(0.99), opLess, false, false}
	var lessFloat64_2 = &cmpTestSuit{int(1), float64(1), opLess, false, false}
	var lessFloat64_3 = &cmpTestSuit{int(1), float64(1.01), opLess, false, true}

	var lessBool = &cmpTestSuit{int(1), true, opLess, true, nil}
	var lessNil = &cmpTestSuit{int(1), nil, opLess, true, nil}
	var lessString = &cmpTestSuit{int(1), "test", opLess, true, nil}
	var lessInt8 = &cmpTestSuit{int(1), int8(1), opLess, true, nil}
	var lessInt16 = &cmpTestSuit{int(1), int16(1), opLess, true, nil}
	var lessInt32 = &cmpTestSuit{int(1), int32(1), opLess, true, nil}
	var lessFloat32 = &cmpTestSuit{int(1), float32(0.5), opLess, true, nil}
	var lessByte = &cmpTestSuit{int(1), byte('a'), opLess, true, nil}

	testCmp(lessInt_1, t)
	testCmp(lessInt_2, t)
	testCmp(lessInt_3, t)
	testCmp(lessInt64_1, t)
	testCmp(lessInt64_2, t)
	testCmp(lessInt64_3, t)
	testCmp(lessFloat64_1, t)
	testCmp(lessFloat64_2, t)
	testCmp(lessFloat64_3, t)
	testCmp(lessString, t)
	testCmp(lessBool, t)
	testCmp(lessNil, t)
	testCmp(lessInt8, t)
	testCmp(lessInt16, t)
	testCmp(lessInt32, t)
	testCmp(lessFloat32, t)
	testCmp(lessByte, t)
}

func TestInt64Less(t *testing.T) {
	var lessInt_1 = &cmpTestSuit{int64(1), int(0), opLess, false, false}
	var lessInt_2 = &cmpTestSuit{int64(1), int(1), opLess, false, false}
	var lessInt_3 = &cmpTestSuit{int64(1), int(2), opLess, false, true}
	var lessInt64_1 = &cmpTestSuit{int64(1), int64(0), opLess, false, false}
	var lessInt64_2 = &cmpTestSuit{int64(1), int64(1), opLess, false, false}
	var lessInt64_3 = &cmpTestSuit{int64(1), int64(2), opLess, false, true}
	var lessFloat64_1 = &cmpTestSuit{int64(1), float64(0.99), opLess, false, false}
	var lessFloat64_2 = &cmpTestSuit{int64(1), float64(1), opLess, false, false}
	var lessFloat64_3 = &cmpTestSuit{int64(1), float64(1.01), opLess, false, true}

	var lessBool = &cmpTestSuit{int64(1), true, opLess, true, nil}
	var lessNil = &cmpTestSuit{int64(1), nil, opLess, true, nil}
	var lessString = &cmpTestSuit{int64(1), "test", opLess, true, nil}
	var lessInt8 = &cmpTestSuit{int64(1), int8(1), opLess, true, nil}
	var lessInt16 = &cmpTestSuit{int64(1), int16(1), opLess, true, nil}
	var lessInt32 = &cmpTestSuit{int64(1), int32(1), opLess, true, nil}
	var lessFloat32 = &cmpTestSuit{int64(1), float32(0.5), opLess, true, nil}
	var lessByte = &cmpTestSuit{int64(1), byte('a'), opLess, true, nil}

	testCmp(lessInt_1, t)
	testCmp(lessInt_2, t)
	testCmp(lessInt_3, t)
	testCmp(lessInt64_1, t)
	testCmp(lessInt64_2, t)
	testCmp(lessInt64_3, t)
	testCmp(lessFloat64_1, t)
	testCmp(lessFloat64_2, t)
	testCmp(lessFloat64_3, t)
	testCmp(lessString, t)
	testCmp(lessBool, t)
	testCmp(lessNil, t)
	testCmp(lessInt8, t)
	testCmp(lessInt16, t)
	testCmp(lessInt32, t)
	testCmp(lessFloat32, t)
	testCmp(lessByte, t)
}

func TestFloat64Less(t *testing.T) {
	var lessInt_1 = &cmpTestSuit{float64(1.5), int(1), opLess, false, false}
	var lessInt_2 = &cmpTestSuit{float64(2), int(2), opLess, false, false}
	var lessInt_3 = &cmpTestSuit{float64(2.5), int(3), opLess, false, true}
	var lessInt64_1 = &cmpTestSuit{float64(1.5), int64(1), opLess, false, false}
	var lessInt64_2 = &cmpTestSuit{float64(2), int64(2), opLess, false, false}
	var lessInt64_3 = &cmpTestSuit{float64(2.5), int64(3), opLess, false, true}
	var lessFloat64_1 = &cmpTestSuit{float64(1.3), float64(1.2), opLess, false, false}
	var lessFloat64_2 = &cmpTestSuit{float64(1.3), float64(1.3), opLess, false, false}
	var lessFloat64_3 = &cmpTestSuit{float64(1.3), float64(1.4), opLess, false, true}

	var lessBool = &cmpTestSuit{float64(1), true, opLess, true, nil}
	var lessNil = &cmpTestSuit{float64(1.1), nil, opLess, true, nil}
	var lessString = &cmpTestSuit{float64(1), "test", opLess, true, nil}
	var lessInt8 = &cmpTestSuit{float64(1), int8(1), opLess, true, nil}
	var lessInt16 = &cmpTestSuit{float64(1), int16(1), opLess, true, nil}
	var lessInt32 = &cmpTestSuit{float64(1), int32(1), opLess, true, nil}
	var lessFloat32 = &cmpTestSuit{float64(1), float32(0.5), opLess, true, nil}
	var lessByte = &cmpTestSuit{float64(1), byte('a'), opLess, true, nil}

	testCmp(lessInt_1, t)
	testCmp(lessInt_2, t)
	testCmp(lessInt_3, t)
	testCmp(lessInt64_1, t)
	testCmp(lessInt64_2, t)
	testCmp(lessInt64_3, t)
	testCmp(lessFloat64_1, t)
	testCmp(lessFloat64_2, t)
	testCmp(lessFloat64_3, t)
	testCmp(lessBool, t)
	testCmp(lessNil, t)
	testCmp(lessString, t)
	testCmp(lessInt8, t)
	testCmp(lessInt16, t)
	testCmp(lessInt32, t)
	testCmp(lessFloat32, t)
	testCmp(lessByte, t)
}

func TestStringLess(t *testing.T) {
	var lessString_1 = &cmpTestSuit{"test1", "test0", opLess, false, false}
	var lessString_2 = &cmpTestSuit{"test1", "test1", opLess, false, false}
	var lessString_3 = &cmpTestSuit{"test1", "test2", opLess, false, true}
	var lessInt_1 = &cmpTestSuit{"1", int(0), opLess, false, false}
	var lessInt_2 = &cmpTestSuit{"1", int(1), opLess, false, false}
	var lessInt_3 = &cmpTestSuit{"1", int(2), opLess, false, true}
	var lessInt64_1 = &cmpTestSuit{"1", int64(0), opLess, false, false}
	var lessInt64_2 = &cmpTestSuit{"1", int64(1), opLess, false, false}
	var lessInt64_3 = &cmpTestSuit{"1", int64(2), opLess, false, true}
	var lessFloat64_1 = &cmpTestSuit{"1.3", float64(1.2), opLess, false, false}
	var lessFloat64_2 = &cmpTestSuit{"1.3", float64(1.3), opLess, false, false}
	var lessFloat64_3 = &cmpTestSuit{"1.3", float64(1.4), opLess, false, true}

	var lessBool = &cmpTestSuit{"true", true, opLess, true, nil}
	var lessNil = &cmpTestSuit{"test", nil, opLess, true, nil}
	var lessInt8 = &cmpTestSuit{"test", int8(1), opLess, true, nil}
	var lessInt16 = &cmpTestSuit{"test", int16(1), opLess, true, nil}
	var lessInt32 = &cmpTestSuit{"test", int32(1), opLess, true, nil}
	var lessFloat32 = &cmpTestSuit{"test", float32(0.5), opLess, true, nil}
	var lessByte = &cmpTestSuit{"test", byte('a'), opLess, true, nil}

	testCmp(lessString_1, t)
	testCmp(lessString_2, t)
	testCmp(lessString_3, t)
	testCmp(lessInt_1, t)
	testCmp(lessInt_2, t)
	testCmp(lessInt_3, t)
	testCmp(lessInt64_1, t)
	testCmp(lessInt64_2, t)
	testCmp(lessInt64_3, t)
	testCmp(lessFloat64_1, t)
	testCmp(lessFloat64_2, t)
	testCmp(lessFloat64_3, t)
	testCmp(lessBool, t)
	testCmp(lessNil, t)
	testCmp(lessInt8, t)
	testCmp(lessInt16, t)
	testCmp(lessInt32, t)
	testCmp(lessFloat32, t)
	testCmp(lessByte, t)
}

func TestBoolLess(t *testing.T) {
	var lessInt = &cmpTestSuit{true, int(1), opLess, true, nil}
	var lessInt64 = &cmpTestSuit{true, int64(1), opLess, true, nil}
	var lessFloat64 = &cmpTestSuit{true, float64(1.3), opLess, true, nil}
	var lessString = &cmpTestSuit{true, "true", opLess, true, nil}
	var lessBool = &cmpTestSuit{true, true, opLess, true, nil}
	var lessNil = &cmpTestSuit{true, nil, opLess, true, nil}
	var lessInt8 = &cmpTestSuit{true, int8(1), opLess, true, nil}
	var lessInt16 = &cmpTestSuit{true, int16(1), opLess, true, nil}
	var lessInt32 = &cmpTestSuit{true, int32(1), opLess, true, nil}
	var lessFloat32 = &cmpTestSuit{true, float32(0.5), opLess, true, nil}
	var lessByte = &cmpTestSuit{true, byte('a'), opLess, true, nil}

	testCmp(lessInt, t)
	testCmp(lessInt64, t)
	testCmp(lessFloat64, t)
	testCmp(lessString, t)
	testCmp(lessBool, t)
	testCmp(lessNil, t)
	testCmp(lessInt8, t)
	testCmp(lessInt16, t)
	testCmp(lessInt32, t)
	testCmp(lessFloat32, t)
	testCmp(lessByte, t)
}

func TestNilLess(t *testing.T) {
	var lessNil = &cmpTestSuit{nil, nil, opLess, true, nil}
	var lessBool = &cmpTestSuit{nil, true, opLess, true, nil}
	var lessInt = &cmpTestSuit{nil, int(1), opLess, true, nil}
	var lessInt64 = &cmpTestSuit{nil, int64(1), opLess, true, nil}
	var lessFloat64 = &cmpTestSuit{nil, float64(1.3), opLess, true, nil}
	var lessString = &cmpTestSuit{nil, "test", opLess, true, nil}
	var lessInt8 = &cmpTestSuit{nil, int8(1), opLess, true, nil}
	var lessInt16 = &cmpTestSuit{nil, int16(1), opLess, true, nil}
	var lessInt32 = &cmpTestSuit{nil, int32(1), opLess, true, nil}
	var lessFloat32 = &cmpTestSuit{nil, float32(0.5), opLess, true, nil}
	var lessByte = &cmpTestSuit{nil, byte('a'), opLess, true, nil}

	testCmp(lessNil, t)
	testCmp(lessBool, t)
	testCmp(lessInt, t)
	testCmp(lessInt64, t)
	testCmp(lessFloat64, t)
	testCmp(lessString, t)
	testCmp(lessInt8, t)
	testCmp(lessInt16, t)
	testCmp(lessInt32, t)
	testCmp(lessFloat32, t)
	testCmp(lessByte, t)
}

func TestIntLessEqual(t *testing.T) {
	var lessEqualInt_1 = &cmpTestSuit{int(1), int(0), opLessEqual, false, false}
	var lessEqualInt_2 = &cmpTestSuit{int(1), int(1), opLessEqual, false, true}
	var lessEqualInt_3 = &cmpTestSuit{int(1), int(2), opLessEqual, false, true}
	var lessEqualInt64_1 = &cmpTestSuit{int(1), int64(0), opLessEqual, false, false}
	var lessEqualInt64_2 = &cmpTestSuit{int(1), int64(1), opLessEqual, false, true}
	var lessEqualInt64_3 = &cmpTestSuit{int(1), int64(2), opLessEqual, false, true}
	var lessEqualFloat64_1 = &cmpTestSuit{int(1), float64(0.99), opLessEqual, false, false}
	var lessEqualFloat64_2 = &cmpTestSuit{int(1), float64(1), opLessEqual, false, true}
	var lessEqualFloat64_3 = &cmpTestSuit{int(1), float64(1.01), opLessEqual, false, true}

	var lessEqualBool = &cmpTestSuit{int(1), true, opLessEqual, true, nil}
	var lessEqualNil = &cmpTestSuit{int(1), nil, opLessEqual, true, nil}
	var lessEqualString = &cmpTestSuit{int(1), "test", opLessEqual, true, nil}
	var lessEqualInt8 = &cmpTestSuit{int(1), int8(1), opLessEqual, true, nil}
	var lessEqualInt16 = &cmpTestSuit{int(1), int16(1), opLessEqual, true, nil}
	var lessEqualInt32 = &cmpTestSuit{int(1), int32(1), opLessEqual, true, nil}
	var lessEqualFloat32 = &cmpTestSuit{int(1), float32(0.5), opLessEqual, true, nil}
	var lessEqualByte = &cmpTestSuit{int(1), byte('a'), opLessEqual, true, nil}

	testCmp(lessEqualInt_1, t)
	testCmp(lessEqualInt_2, t)
	testCmp(lessEqualInt_3, t)
	testCmp(lessEqualInt64_1, t)
	testCmp(lessEqualInt64_2, t)
	testCmp(lessEqualInt64_3, t)
	testCmp(lessEqualFloat64_1, t)
	testCmp(lessEqualFloat64_2, t)
	testCmp(lessEqualFloat64_3, t)
	testCmp(lessEqualString, t)
	testCmp(lessEqualBool, t)
	testCmp(lessEqualNil, t)
	testCmp(lessEqualInt8, t)
	testCmp(lessEqualInt16, t)
	testCmp(lessEqualInt32, t)
	testCmp(lessEqualFloat32, t)
	testCmp(lessEqualByte, t)
}

func TestInt64LessEqual(t *testing.T) {
	var lessEqualInt_1 = &cmpTestSuit{int64(1), int(0), opLessEqual, false, false}
	var lessEqualInt_2 = &cmpTestSuit{int64(1), int(1), opLessEqual, false, true}
	var lessEqualInt_3 = &cmpTestSuit{int64(1), int(2), opLessEqual, false, true}
	var lessEqualInt64_1 = &cmpTestSuit{int64(1), int64(0), opLessEqual, false, false}
	var lessEqualInt64_2 = &cmpTestSuit{int64(1), int64(1), opLessEqual, false, true}
	var lessEqualInt64_3 = &cmpTestSuit{int64(1), int64(2), opLessEqual, false, true}
	var lessEqualFloat64_1 = &cmpTestSuit{int64(1), float64(0.99), opLessEqual, false, false}
	var lessEqualFloat64_2 = &cmpTestSuit{int64(1), float64(1), opLessEqual, false, true}
	var lessEqualFloat64_3 = &cmpTestSuit{int64(1), float64(1.01), opLessEqual, false, true}

	var lessEqualBool = &cmpTestSuit{int64(1), true, opLessEqual, true, nil}
	var lessEqualNil = &cmpTestSuit{int64(1), nil, opLessEqual, true, nil}
	var lessEqualString = &cmpTestSuit{int64(1), "test", opLessEqual, true, nil}
	var lessEqualInt8 = &cmpTestSuit{int64(1), int8(1), opLessEqual, true, nil}
	var lessEqualInt16 = &cmpTestSuit{int64(1), int16(1), opLessEqual, true, nil}
	var lessEqualInt32 = &cmpTestSuit{int64(1), int32(1), opLessEqual, true, nil}
	var lessEqualFloat32 = &cmpTestSuit{int64(1), float32(0.5), opLessEqual, true, nil}
	var lessEqualByte = &cmpTestSuit{int64(1), byte('a'), opLessEqual, true, nil}

	testCmp(lessEqualInt_1, t)
	testCmp(lessEqualInt_2, t)
	testCmp(lessEqualInt_3, t)
	testCmp(lessEqualInt64_1, t)
	testCmp(lessEqualInt64_2, t)
	testCmp(lessEqualInt64_3, t)
	testCmp(lessEqualFloat64_1, t)
	testCmp(lessEqualFloat64_2, t)
	testCmp(lessEqualFloat64_3, t)
	testCmp(lessEqualString, t)
	testCmp(lessEqualBool, t)
	testCmp(lessEqualNil, t)
	testCmp(lessEqualInt8, t)
	testCmp(lessEqualInt16, t)
	testCmp(lessEqualInt32, t)
	testCmp(lessEqualFloat32, t)
	testCmp(lessEqualByte, t)
}

func TestFloat64LessEqual(t *testing.T) {
	var lessEqualInt_1 = &cmpTestSuit{float64(1.5), int(1), opLessEqual, false, false}
	var lessEqualInt_2 = &cmpTestSuit{float64(2), int(2), opLessEqual, false, true}
	var lessEqualInt_3 = &cmpTestSuit{float64(2.5), int(3), opLessEqual, false, true}
	var lessEqualInt64_1 = &cmpTestSuit{float64(1.5), int64(1), opLessEqual, false, false}
	var lessEqualInt64_2 = &cmpTestSuit{float64(2), int64(2), opLessEqual, false, true}
	var lessEqualInt64_3 = &cmpTestSuit{float64(2.5), int64(3), opLessEqual, false, true}
	var lessEqualFloat64_1 = &cmpTestSuit{float64(1.3), float64(1.2), opLessEqual, false, false}
	var lessEqualFloat64_2 = &cmpTestSuit{float64(1.3), float64(1.3), opLessEqual, false, true}
	var lessEqualFloat64_3 = &cmpTestSuit{float64(1.3), float64(1.4), opLessEqual, false, true}

	var lessEqualBool = &cmpTestSuit{float64(1), true, opLessEqual, true, nil}
	var lessEqualNil = &cmpTestSuit{float64(1.1), nil, opLessEqual, true, nil}
	var lessEqualString = &cmpTestSuit{float64(1), "test", opLessEqual, true, nil}
	var lessEqualInt8 = &cmpTestSuit{float64(1), int8(1), opLessEqual, true, nil}
	var lessEqualInt16 = &cmpTestSuit{float64(1), int16(1), opLessEqual, true, nil}
	var lessEqualInt32 = &cmpTestSuit{float64(1), int32(1), opLessEqual, true, nil}
	var lessEqualFloat32 = &cmpTestSuit{float64(1), float32(0.5), opLessEqual, true, nil}
	var lessEqualByte = &cmpTestSuit{float64(1), byte('a'), opLessEqual, true, nil}

	testCmp(lessEqualInt_1, t)
	testCmp(lessEqualInt_2, t)
	testCmp(lessEqualInt_3, t)
	testCmp(lessEqualInt64_1, t)
	testCmp(lessEqualInt64_2, t)
	testCmp(lessEqualInt64_3, t)
	testCmp(lessEqualFloat64_1, t)
	testCmp(lessEqualFloat64_2, t)
	testCmp(lessEqualFloat64_3, t)
	testCmp(lessEqualBool, t)
	testCmp(lessEqualNil, t)
	testCmp(lessEqualString, t)
	testCmp(lessEqualInt8, t)
	testCmp(lessEqualInt16, t)
	testCmp(lessEqualInt32, t)
	testCmp(lessEqualFloat32, t)
	testCmp(lessEqualByte, t)
}

func TestStringLessEqual(t *testing.T) {
	var lessEqualString_1 = &cmpTestSuit{"test1", "test0", opLessEqual, false, false}
	var lessEqualString_2 = &cmpTestSuit{"test1", "test1", opLessEqual, false, true}
	var lessEqualString_3 = &cmpTestSuit{"test1", "test2", opLessEqual, false, true}
	var lessEqualInt_1 = &cmpTestSuit{"1", int(0), opLessEqual, false, false}
	var lessEqualInt_2 = &cmpTestSuit{"1", int(1), opLessEqual, false, true}
	var lessEqualInt_3 = &cmpTestSuit{"1", int(2), opLessEqual, false, true}
	var lessEqualInt64_1 = &cmpTestSuit{"1", int64(0), opLessEqual, false, false}
	var lessEqualInt64_2 = &cmpTestSuit{"1", int64(1), opLessEqual, false, true}
	var lessEqualInt64_3 = &cmpTestSuit{"1", int64(2), opLessEqual, false, true}
	var lessEqualFloat64_1 = &cmpTestSuit{"1.3", float64(1.2), opLessEqual, false, false}
	var lessEqualFloat64_2 = &cmpTestSuit{"1.3", float64(1.3), opLessEqual, false, true}
	var lessEqualFloat64_3 = &cmpTestSuit{"1.3", float64(1.4), opLessEqual, false, true}

	var lessEqualBool = &cmpTestSuit{"true", true, opLessEqual, true, nil}
	var lessEqualNil = &cmpTestSuit{"test", nil, opLessEqual, true, nil}
	var lessEqualInt8 = &cmpTestSuit{"test", int8(1), opLessEqual, true, nil}
	var lessEqualInt16 = &cmpTestSuit{"test", int16(1), opLessEqual, true, nil}
	var lessEqualInt32 = &cmpTestSuit{"test", int32(1), opLessEqual, true, nil}
	var lessEqualFloat32 = &cmpTestSuit{"test", float32(0.5), opLessEqual, true, nil}
	var lessEqualByte = &cmpTestSuit{"test", byte('a'), opLessEqual, true, nil}

	testCmp(lessEqualString_1, t)
	testCmp(lessEqualString_2, t)
	testCmp(lessEqualString_3, t)
	testCmp(lessEqualInt_1, t)
	testCmp(lessEqualInt_2, t)
	testCmp(lessEqualInt_3, t)
	testCmp(lessEqualInt64_1, t)
	testCmp(lessEqualInt64_2, t)
	testCmp(lessEqualInt64_3, t)
	testCmp(lessEqualFloat64_1, t)
	testCmp(lessEqualFloat64_2, t)
	testCmp(lessEqualFloat64_3, t)
	testCmp(lessEqualBool, t)
	testCmp(lessEqualNil, t)
	testCmp(lessEqualInt8, t)
	testCmp(lessEqualInt16, t)
	testCmp(lessEqualInt32, t)
	testCmp(lessEqualFloat32, t)
	testCmp(lessEqualByte, t)
}

func TestBoolLessEqual(t *testing.T) {
	var lessEqualInt = &cmpTestSuit{true, int(1), opLessEqual, true, nil}
	var lessEqualInt64 = &cmpTestSuit{true, int64(1), opLessEqual, true, nil}
	var lessEqualFloat64 = &cmpTestSuit{true, float64(1.3), opLessEqual, true, nil}
	var lessEqualString = &cmpTestSuit{true, "true", opLessEqual, true, nil}
	var lessEqualBool = &cmpTestSuit{true, true, opLessEqual, true, nil}
	var lessEqualNil = &cmpTestSuit{true, nil, opLessEqual, true, nil}
	var lessEqualInt8 = &cmpTestSuit{true, int8(1), opLessEqual, true, nil}
	var lessEqualInt16 = &cmpTestSuit{true, int16(1), opLessEqual, true, nil}
	var lessEqualInt32 = &cmpTestSuit{true, int32(1), opLessEqual, true, nil}
	var lessEqualFloat32 = &cmpTestSuit{true, float32(0.5), opLessEqual, true, nil}
	var lessEqualByte = &cmpTestSuit{true, byte('a'), opLessEqual, true, nil}

	testCmp(lessEqualInt, t)
	testCmp(lessEqualInt64, t)
	testCmp(lessEqualFloat64, t)
	testCmp(lessEqualString, t)
	testCmp(lessEqualBool, t)
	testCmp(lessEqualNil, t)
	testCmp(lessEqualInt8, t)
	testCmp(lessEqualInt16, t)
	testCmp(lessEqualInt32, t)
	testCmp(lessEqualFloat32, t)
	testCmp(lessEqualByte, t)
}

func TestNilLessEqual(t *testing.T) {
	var lessEqualNil = &cmpTestSuit{nil, nil, opLessEqual, true, nil}
	var lessEqualBool = &cmpTestSuit{nil, true, opLessEqual, true, nil}
	var lessEqualInt = &cmpTestSuit{nil, int(1), opLessEqual, true, nil}
	var lessEqualInt64 = &cmpTestSuit{nil, int64(1), opLessEqual, true, nil}
	var lessEqualFloat64 = &cmpTestSuit{nil, float64(1.3), opLessEqual, true, nil}
	var lessEqualString = &cmpTestSuit{nil, "test", opLessEqual, true, nil}
	var lessEqualInt8 = &cmpTestSuit{nil, int8(1), opLessEqual, true, nil}
	var lessEqualInt16 = &cmpTestSuit{nil, int16(1), opLessEqual, true, nil}
	var lessEqualInt32 = &cmpTestSuit{nil, int32(1), opLessEqual, true, nil}
	var lessEqualFloat32 = &cmpTestSuit{nil, float32(0.5), opLessEqual, true, nil}
	var lessEqualByte = &cmpTestSuit{nil, byte('a'), opLessEqual, true, nil}

	testCmp(lessEqualNil, t)
	testCmp(lessEqualBool, t)
	testCmp(lessEqualInt, t)
	testCmp(lessEqualInt64, t)
	testCmp(lessEqualFloat64, t)
	testCmp(lessEqualString, t)
	testCmp(lessEqualInt8, t)
	testCmp(lessEqualInt16, t)
	testCmp(lessEqualInt32, t)
	testCmp(lessEqualFloat32, t)
	testCmp(lessEqualByte, t)
}

func TestIntRegexpMatch(t *testing.T) {
	var regexpMatchString_1 = &cmpTestSuit{int(1), "\\d", opRegexpMatch, false, true}
	var regexpMatchString_2 = &cmpTestSuit{int(1), "\\s", opRegexpMatch, false, false}
	var regexpMatchString_3 = &cmpTestSuit{int(1), "\\d{2}", opRegexpMatch, false, false}
	var regexpMatchString_4 = &cmpTestSuit{int(1), ".*", opRegexpMatch, false, true}
	var regexpMatchString_5 = &cmpTestSuit{int(1), "", opRegexpMatch, false, true}
	var regexpMatchString_6 = &cmpTestSuit{int(1), "\\1", opRegexpMatch, true, nil}

	var regexpMatchInt = &cmpTestSuit{int(1), int(0), opRegexpMatch, true, nil}
	var regexpMatchInt64 = &cmpTestSuit{int(1), int64(0), opRegexpMatch, true, nil}
	var regexpMatchFloat64 = &cmpTestSuit{int(1), float64(0.99), opRegexpMatch, true, nil}
	var regexpMatchBool = &cmpTestSuit{int(1), true, opRegexpMatch, true, nil}
	var regexpMatchNil = &cmpTestSuit{int(1), nil, opRegexpMatch, true, nil}
	var regexpMatchInt8 = &cmpTestSuit{int(1), int8(1), opRegexpMatch, true, nil}
	var regexpMatchInt16 = &cmpTestSuit{int(1), int16(1), opRegexpMatch, true, nil}
	var regexpMatchInt32 = &cmpTestSuit{int(1), int32(1), opRegexpMatch, true, nil}
	var regexpMatchFloat32 = &cmpTestSuit{int(1), float32(0.5), opRegexpMatch, true, nil}
	var regexpMatchByte = &cmpTestSuit{int(1), byte('a'), opRegexpMatch, true, nil}

	testCmp(regexpMatchString_1, t)
	testCmp(regexpMatchString_2, t)
	testCmp(regexpMatchString_3, t)
	testCmp(regexpMatchString_4, t)
	testCmp(regexpMatchString_5, t)
	testCmp(regexpMatchString_6, t)
	testCmp(regexpMatchInt, t)
	testCmp(regexpMatchInt64, t)
	testCmp(regexpMatchFloat64, t)
	testCmp(regexpMatchBool, t)
	testCmp(regexpMatchNil, t)
	testCmp(regexpMatchInt8, t)
	testCmp(regexpMatchInt16, t)
	testCmp(regexpMatchInt32, t)
	testCmp(regexpMatchFloat32, t)
	testCmp(regexpMatchByte, t)
}

func TestInt64RegexpMatch(t *testing.T) {
	var regexpMatchString_1 = &cmpTestSuit{int64(1), "\\d", opRegexpMatch, false, true}
	var regexpMatchString_2 = &cmpTestSuit{int64(1), "\\s", opRegexpMatch, false, false}
	var regexpMatchString_3 = &cmpTestSuit{int64(1), "\\d{2}", opRegexpMatch, false, false}
	var regexpMatchString_4 = &cmpTestSuit{int64(1), ".*", opRegexpMatch, false, true}
	var regexpMatchString_5 = &cmpTestSuit{int64(1), "", opRegexpMatch, false, true}
	var regexpMatchString_6 = &cmpTestSuit{int64(1), "\\1", opRegexpMatch, true, nil}

	var regexpMatchInt = &cmpTestSuit{int64(1), int(0), opRegexpMatch, true, nil}
	var regexpMatchInt64 = &cmpTestSuit{int64(1), int64(0), opRegexpMatch, true, nil}
	var regexpMatchFloat64 = &cmpTestSuit{int64(1), float64(0.99), opRegexpMatch, true, nil}
	var regexpMatchBool = &cmpTestSuit{int64(1), true, opRegexpMatch, true, nil}
	var regexpMatchNil = &cmpTestSuit{int64(1), nil, opRegexpMatch, true, nil}
	var regexpMatchInt8 = &cmpTestSuit{int64(1), int8(1), opRegexpMatch, true, nil}
	var regexpMatchInt16 = &cmpTestSuit{int64(1), int16(1), opRegexpMatch, true, nil}
	var regexpMatchInt32 = &cmpTestSuit{int64(1), int32(1), opRegexpMatch, true, nil}
	var regexpMatchFloat32 = &cmpTestSuit{int64(1), float32(0.5), opRegexpMatch, true, nil}
	var regexpMatchByte = &cmpTestSuit{int64(1), byte('a'), opRegexpMatch, true, nil}

	testCmp(regexpMatchString_1, t)
	testCmp(regexpMatchString_2, t)
	testCmp(regexpMatchString_3, t)
	testCmp(regexpMatchString_4, t)
	testCmp(regexpMatchString_5, t)
	testCmp(regexpMatchString_6, t)
	testCmp(regexpMatchInt, t)
	testCmp(regexpMatchInt64, t)
	testCmp(regexpMatchFloat64, t)
	testCmp(regexpMatchBool, t)
	testCmp(regexpMatchNil, t)
	testCmp(regexpMatchInt8, t)
	testCmp(regexpMatchInt16, t)
	testCmp(regexpMatchInt32, t)
	testCmp(regexpMatchFloat32, t)
	testCmp(regexpMatchByte, t)
}

func TestFloat64RegexpMatch(t *testing.T) {
	var regexpMatchString_1 = &cmpTestSuit{float64(1.2), "\\d\\.\\d", opRegexpMatch, false, true}
	var regexpMatchString_2 = &cmpTestSuit{float64(1.2), "\\s+", opRegexpMatch, false, false}
	var regexpMatchString_3 = &cmpTestSuit{float64(1.2), "\\d{2}", opRegexpMatch, false, false}
	var regexpMatchString_4 = &cmpTestSuit{float64(1.2), ".*", opRegexpMatch, false, true}
	var regexpMatchString_5 = &cmpTestSuit{float64(1.2), "", opRegexpMatch, false, true}
	var regexpMatchString_6 = &cmpTestSuit{float64(1.2), "\\1", opRegexpMatch, true, nil}

	var regexpMatchInt = &cmpTestSuit{float64(1), int(0), opRegexpMatch, true, nil}
	var regexpMatchInt64 = &cmpTestSuit{float64(1), int64(0), opRegexpMatch, true, nil}
	var regexpMatchFloat64 = &cmpTestSuit{float64(1), float64(0.99), opRegexpMatch, true, nil}
	var regexpMatchBool = &cmpTestSuit{float64(1), true, opRegexpMatch, true, nil}
	var regexpMatchNil = &cmpTestSuit{float64(1), nil, opRegexpMatch, true, nil}
	var regexpMatchInt8 = &cmpTestSuit{float64(1), int8(1), opRegexpMatch, true, nil}
	var regexpMatchInt16 = &cmpTestSuit{float64(1), int16(1), opRegexpMatch, true, nil}
	var regexpMatchInt32 = &cmpTestSuit{float64(1), int32(1), opRegexpMatch, true, nil}
	var regexpMatchFloat32 = &cmpTestSuit{float64(1), float32(0.5), opRegexpMatch, true, nil}
	var regexpMatchByte = &cmpTestSuit{float64(1), byte('a'), opRegexpMatch, true, nil}

	testCmp(regexpMatchString_1, t)
	testCmp(regexpMatchString_2, t)
	testCmp(regexpMatchString_3, t)
	testCmp(regexpMatchString_4, t)
	testCmp(regexpMatchString_5, t)
	testCmp(regexpMatchString_6, t)
	testCmp(regexpMatchInt, t)
	testCmp(regexpMatchInt64, t)
	testCmp(regexpMatchFloat64, t)
	testCmp(regexpMatchBool, t)
	testCmp(regexpMatchNil, t)
	testCmp(regexpMatchInt8, t)
	testCmp(regexpMatchInt16, t)
	testCmp(regexpMatchInt32, t)
	testCmp(regexpMatchFloat32, t)
	testCmp(regexpMatchByte, t)
}

func TestStringRegexpMatch(t *testing.T) {
	var regexpMatchString_1 = &cmpTestSuit{"test", "\\d", opRegexpMatch, false, false}
	var regexpMatchString_2 = &cmpTestSuit{"test", "\\S+", opRegexpMatch, false, true}
	var regexpMatchString_3 = &cmpTestSuit{"test", "\\d{2}", opRegexpMatch, false, false}
	var regexpMatchString_4 = &cmpTestSuit{"test", ".*", opRegexpMatch, false, true}
	var regexpMatchString_5 = &cmpTestSuit{"test", "", opRegexpMatch, false, true}
	var regexpMatchString_6 = &cmpTestSuit{"test", "\\1", opRegexpMatch, true, nil}

	var regexpMatchInt = &cmpTestSuit{"test", int(0), opRegexpMatch, true, nil}
	var regexpMatchInt64 = &cmpTestSuit{"test", int64(0), opRegexpMatch, true, nil}
	var regexpMatchFloat64 = &cmpTestSuit{"test", float64(0.99), opRegexpMatch, true, nil}
	var regexpMatchBool = &cmpTestSuit{"test", true, opRegexpMatch, true, nil}
	var regexpMatchNil = &cmpTestSuit{"test", nil, opRegexpMatch, true, nil}
	var regexpMatchInt8 = &cmpTestSuit{"test", int8(1), opRegexpMatch, true, nil}
	var regexpMatchInt16 = &cmpTestSuit{"test", int16(1), opRegexpMatch, true, nil}
	var regexpMatchInt32 = &cmpTestSuit{"test", int32(1), opRegexpMatch, true, nil}
	var regexpMatchFloat32 = &cmpTestSuit{"test", float32(0.5), opRegexpMatch, true, nil}
	var regexpMatchByte = &cmpTestSuit{"test", byte('a'), opRegexpMatch, true, nil}

	testCmp(regexpMatchString_1, t)
	testCmp(regexpMatchString_2, t)
	testCmp(regexpMatchString_3, t)
	testCmp(regexpMatchString_4, t)
	testCmp(regexpMatchString_5, t)
	testCmp(regexpMatchString_6, t)
	testCmp(regexpMatchInt, t)
	testCmp(regexpMatchInt64, t)
	testCmp(regexpMatchFloat64, t)
	testCmp(regexpMatchBool, t)
	testCmp(regexpMatchNil, t)
	testCmp(regexpMatchInt8, t)
	testCmp(regexpMatchInt16, t)
	testCmp(regexpMatchInt32, t)
	testCmp(regexpMatchFloat32, t)
	testCmp(regexpMatchByte, t)
}

func TestBoolRegexpMatch(t *testing.T) {
	var regexpMatchInt = &cmpTestSuit{true, int(1), opRegexpMatch, true, nil}
	var regexpMatchInt64 = &cmpTestSuit{true, int64(1), opRegexpMatch, true, nil}
	var regexpMatchFloat64 = &cmpTestSuit{true, float64(1.3), opRegexpMatch, true, nil}
	var regexpMatchString = &cmpTestSuit{true, "true", opRegexpMatch, true, nil}
	var regexpMatchBool = &cmpTestSuit{true, true, opRegexpMatch, true, nil}
	var regexpMatchNil = &cmpTestSuit{true, nil, opRegexpMatch, true, nil}
	var regexpMatchInt8 = &cmpTestSuit{true, int8(1), opRegexpMatch, true, nil}
	var regexpMatchInt16 = &cmpTestSuit{true, int16(1), opRegexpMatch, true, nil}
	var regexpMatchInt32 = &cmpTestSuit{true, int32(1), opRegexpMatch, true, nil}
	var regexpMatchFloat32 = &cmpTestSuit{true, float32(0.5), opRegexpMatch, true, nil}
	var regexpMatchByte = &cmpTestSuit{true, byte('a'), opRegexpMatch, true, nil}

	testCmp(regexpMatchInt, t)
	testCmp(regexpMatchInt64, t)
	testCmp(regexpMatchFloat64, t)
	testCmp(regexpMatchString, t)
	testCmp(regexpMatchBool, t)
	testCmp(regexpMatchNil, t)
	testCmp(regexpMatchInt8, t)
	testCmp(regexpMatchInt16, t)
	testCmp(regexpMatchInt32, t)
	testCmp(regexpMatchFloat32, t)
	testCmp(regexpMatchByte, t)
}

func TestNilRegexpMatch(t *testing.T) {
	var regexpMatchNil = &cmpTestSuit{nil, nil, opRegexpMatch, true, nil}
	var regexpMatchBool = &cmpTestSuit{nil, true, opRegexpMatch, true, nil}
	var regexpMatchInt = &cmpTestSuit{nil, int(1), opRegexpMatch, true, nil}
	var regexpMatchInt64 = &cmpTestSuit{nil, int64(1), opRegexpMatch, true, nil}
	var regexpMatchFloat64 = &cmpTestSuit{nil, float64(1.3), opRegexpMatch, true, nil}
	var regexpMatchString = &cmpTestSuit{nil, "test", opRegexpMatch, true, nil}
	var regexpMatchInt8 = &cmpTestSuit{nil, int8(1), opRegexpMatch, true, nil}
	var regexpMatchInt16 = &cmpTestSuit{nil, int16(1), opRegexpMatch, true, nil}
	var regexpMatchInt32 = &cmpTestSuit{nil, int32(1), opRegexpMatch, true, nil}
	var regexpMatchFloat32 = &cmpTestSuit{nil, float32(0.5), opRegexpMatch, true, nil}
	var regexpMatchByte = &cmpTestSuit{nil, byte('a'), opRegexpMatch, true, nil}

	testCmp(regexpMatchNil, t)
	testCmp(regexpMatchBool, t)
	testCmp(regexpMatchInt, t)
	testCmp(regexpMatchInt64, t)
	testCmp(regexpMatchFloat64, t)
	testCmp(regexpMatchString, t)
	testCmp(regexpMatchInt8, t)
	testCmp(regexpMatchInt16, t)
	testCmp(regexpMatchInt32, t)
	testCmp(regexpMatchFloat32, t)
	testCmp(regexpMatchByte, t)
}

func TestOtherEqual(t *testing.T)  {
	var equalNil = &cmpTestSuit{[]string{"test"}, nil, opEqual, false, false}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, nil, opEqual, false, false}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int(1), opEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int64(1), opEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, float64(1), opEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, true, opEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, "test", opEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, float32(1), opEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int8(1), opEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int16(1), opEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int32(1), opEqual, true, nil}
	testCmp(equalNil, t)
}

func TestOtherNotEqual(t *testing.T)  {
	var equalNil = &cmpTestSuit{[]string{"test"}, nil, opNotEqual, false, true}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, nil, opNotEqual, false, true}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int(1), opNotEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int64(1), opNotEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, float64(1), opNotEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, true, opNotEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, "test", opNotEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, float32(1), opNotEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int8(1), opNotEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int16(1), opNotEqual, true, nil}
	testCmp(equalNil, t)

	equalNil = &cmpTestSuit{map[string]string{}, int32(1), opNotEqual, true, nil}
	testCmp(equalNil, t)
}