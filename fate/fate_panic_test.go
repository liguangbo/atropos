package fate

import (
	"fmt"
	"math"
	"testing"
)

func TestPanic(t *testing.T) {
	data := []byte("{}")

	script := `
	number = [1,2,3]
	number[` + fmt.Sprintf("1%d", math.MaxInt64) + `] = 1
	`
	code, err := Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = [1,2,3]
	number[3] = 1
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = [1,2,3]
	test = number[3]
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = [1,2,3]
	test = number[` + fmt.Sprintf("1%d", math.MaxInt64) + `]
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = 1
	test = number.var
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = 1
	number.var = 2
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number.var = 2
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number.var.var1 = 2
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, []byte(`{"number":{}}`))
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = [1,2,3,nil]
	test = number[` + fmt.Sprintf("1%d", math.MaxInt64) + `].var
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = [1,2,3,nil]
	test = number[5].var
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = [1,2,3,nil]
	test = number[3].var
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	number = 1
	test = number.var.var1
	`
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
		return
	}
	_, err = Run(code, data)
	if err == nil {
		t.Errorf("run %s expect err != nil, but got %v", script, err)
	}

	script = `
	if true then
	`
	code, err = Compile(script)
	if err == nil {
		t.Errorf("compile %s expect err != nil, but got %v", script, err)
		return
	}
}
