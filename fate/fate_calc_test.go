package fate


import (
	"testing"
)

func TestAdd(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = 1 + 1
	`
	expect = int64(2)
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
	number = 1 + 1
	number = number + 1
	`
	expect = int64(3)
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
	number = 1 + 1
	number = number + number
	`
	expect = int64(4)
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
	number = 1 + 1 + 1
	number = number + number + number
	`
	expect = int64(9)
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
	number = -1
	number = number + 1
	`
	expect = int64(0)
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
	number = -1
	number = number + -1
	`
	expect = int64(-2)
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
	number = number + 1.5
	`
	expect = float64(2.5)
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
	number = number + "1.5"
	`
	expect = "11.5"
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
	number = "1"
	number = number + true
	`
	expect = "1true"
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

func TestSub(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = 1 - 1
	`
	expect = int64(0)
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
	number = 1 - 1
	number = number - 1
	`
	expect = int64(-1)
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
	number = 1 - 1
	number = number - number
	`
	expect = int64(0)
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
	number = 1 - 1 - 1
	number = number - number - number
	`
	expect = int64(1)
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
	number = -1
	number = number - 1
	`
	expect = int64(-2)
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
	number = -1
	number = number - 1
	`
	expect = int64(-2)
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
	number = number - 1.5
	`
	expect = float64(-0.5)
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

func TestMul(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = 2 * 2
	`
	expect = int64(4)
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
	number = 2 * 2
	number = number * 2
	`
	expect = int64(8)
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
	number = 2 * 2
	number = number * number
	`
	expect = int64(16)
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
	number = 2 * 2 * 2
	number = number * number * number
	`
	expect = int64(512)
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
	number = 2
	number = number * -1
	`
	expect = int64(-2)
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
	number = 2
	number = number * 1.5
	`
	expect = float64(3)
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
	number = -1.5
	number = number * -1.5
	`
	expect = float64(2.25)
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

func TestDiv(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = 2 / 2
	`
	expect = int64(1)
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
	number = 2 / 2
	number = number / 2
	`
	expect = int64(0)
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
	number = 2 / 2
	number = number / number
	`
	expect = int64(1)
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
	number = 2
	number = number / -1
	`
	expect = int64(-2)
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
	number = 2
	number = number / 0.5
	`
	expect = float64(4)
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
	number = -1.5
	number = number / 3
	`
	expect = float64(-0.5)
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

func TestMod(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = 2 % 2
	`
	expect = int64(0)
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
	number = 3 % 2
	number = number % 2
	`
	expect = int64(1)
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
	number = 3 % 2
	number = number % number
	`
	expect = int64(0)
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
	number = 2
	number = number % -3
	`
	expect = int64(2)
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
	number = 5
	number = number % -3
	`
	expect = int64(2)
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
	number = -5
	number = number % -3
	`
	expect = int64(-2)
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

func TestComplicateCalc(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = 1 + 2 * 3 + 4
	`
	expect = int64(11)
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
	number = 1 + 3 / 3 + 4
	`
	expect = int64(6)
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
	number = 1 + 3 / 1.5 + 4
	`
	expect = float64(7)
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
	number = 1 + 3 * 2 + 4 * 2 + 1 + 2 * 3 * 4 / 2
	`
	expect = int64(28)
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
	number = 2 * (3 + 2 * 3) + 2 * (2 * (3 + 3))
	`
	expect = int64(42)
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