package fate

import "testing"

func TestBool(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = true
	`
	expect = true
	code, err := Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err := Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got := mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = true
	number = number == true
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = true
	number = !number
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = true
	number = number != true
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1
	number = number == 1
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1
	number = number >= 1
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1
	number = number <= 1
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1
	number = number < 1
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1
	number = number > 1
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1.1
	number = number == 1.1
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1.1
	number = number != 1.1
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1.1
	number = number >= 1.1
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1.1
	number = number <= 1.1
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1.1
	number = number < 1.1
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 1.1
	number = number > 1.1
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "test"
	number = number == "test"
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "test"
	number = number != "test"
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "tesu"
	number = number >= "test"
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "tess"
	number = number >= "test"
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "tesu"
	number = number <= "test"
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "tess"
	number = number <= "test"
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "tesu"
	number = number > "test"
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "tess"
	number = number > "test"
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "tesu"
	number = number < "test"
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "tess"
	number = number < "test"
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "test"
	number = number.count() == 4
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = number.avg() >= 3
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = 3 > number.max()
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = 0 < number.min()
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = number contains 1
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = number !contains 1
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = number contains [1,3]
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = number !contains [1,3]
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = number contains [1,4]
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = [1,2,3]
	number = number !contains [1,4]
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}
}


func TestBoolAnd(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = true && true
	`
	expect = true
	code, err := Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err := Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got := mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = true && false
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = true && false
	number = number && true
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 2 > 1 && 1 < 2
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 2 > 1 && 1 < 2 && 2 > 3
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 2.5 > 1 && 1 < 2 && 3.1 > 3
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "b" > "a" && 1 < 2 && 3.1 > 3
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = !false && 1 < 2 && 3.1 > 3
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}
}

func TestBoolOr(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = true || true
	`
	expect = true
	code, err := Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err := Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got := mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = true || false
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = false || false
	number = number || false
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 2 > 1 || 1 > 2
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 2 < 1 || 1 > 2 || 2 > 3
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = 2.5 > 1 || 1 < 2 || 3.1 > 3
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = "b" > "a" || 1 > 2 || 3.1 > 3
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = !true || 1 > 2 || 3.1 <= 3
	`
	expect = false
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}
}

func TestBoolComplicate(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = false || true && false
	`
	expect = false
	code, err := Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err := Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got := mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = false || true && false || true
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = false || true && (false || true)
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	number = true && ((true || false) && (false || true))
	`
	expect = true
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}
}