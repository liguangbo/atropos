package fate

import (
	"fmt"
	"regexp"
	"runtime/debug"
	"testing"
)

func testBoolExpr(value interface{}, shouldPanic bool, expect interface{}, t *testing.T) {
	defer func() {
		err := recover()
		if !shouldPanic {
			// 不应该panic的时候panic了
			if err != nil {
				t.Error(fmt.Sprintf("test %T(%v) to bool should not panic\n panic:%s\n trace:%s", value, value, err, string(debug.Stack())))
			}
			return
		}
		// 应该panic的时候没有panic
		if err == nil {
			t.Error(fmt.Sprintf("test %T(%v) to bool should panic but not panic\n panic:%s\n trace:%s", value, value, err, string(debug.Stack())))
		}
	}()
	got := boolExpr(value)
	if got != expect {
		t.Error(fmt.Sprintf("test %T(%v) to bool but got %T(%v)", value, value, got, got))
	}
}

func TestBoolExpr(t *testing.T) {
	testBoolExpr(true, false, true, t)
	testBoolExpr(false, false, false, t)
	testBoolExpr(int(0), false, false, t)
	testBoolExpr(int(1), false, true, t)
	testBoolExpr(int64(0), false, false, t)
	testBoolExpr(int64(1), false, true, t)
	testBoolExpr(float64(0), false, false, t)
	testBoolExpr(float64(0.5), false, true, t)
	testBoolExpr("", false, false, t)
	testBoolExpr("test", false, true, t)
	testBoolExpr(nil, false, false, t)
	testBoolExpr(int8(0), true, nil, t)
	testBoolExpr(int16(0), true, nil, t)
	testBoolExpr(int32(0), true, nil, t)
	testBoolExpr(float32(0), true, nil, t)
	testBoolExpr(map[string]string{}, true, nil, t)
}

func TestGetRegexp(t *testing.T) {
	// compile
	_, err := getRegexp(".*")
	if err != nil {
		t.Error("get regexp error", err)
	}
	_, ok := regexpCache.Get(".*")
	if !ok {
		t.Error("save cache failed")
	}
	// getFromCache
	_, err = getRegexp(".*")
	if err != nil {
		t.Error("get regexp error", err)
	}
	regexpCache.Set(".*", "test")
	// getFromCache and compile again
	_, err = getRegexp(".*")
	if err != nil {
		t.Error("get regexp error", err)
	}
	regInterface, ok := regexpCache.Get(".*")
	if !ok {
		t.Error("save cache failed")
	}
	_, ok = regInterface.(*regexp.Regexp)
	if !ok {
		t.Error("save wrong regexp to cache")
	}

	regexpCache.Set("\\1", "test")
	_, err = getRegexp("\\1")
	if err == nil {
		t.Error("get regexp should error")
	}
}
