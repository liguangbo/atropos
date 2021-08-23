package fate


import "testing"

func TestReturn(t *testing.T) {
	var expect interface{}
	data := []byte("{}")

	script := `
	number = 1
	return
	number = 2
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
	if 2 > 3 then
		number = "test"
		number = 2
	elseif 2 < 3 then
		number = 3
	else
		number = 4
	end
	return
	number = 1
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
	number = 1
	if 2 > 3 && "a" <= "b"then
		if "a" == "b" then
			number = "b"
			return
		end
		return
	elseif 2 == 3 then
		number = 3
	else
		if "a" > "b" then
			number = 4
		elseif "a" < "b" then
			number = "a"
			return
		else
			number = 5
		end
	end
	number = "c"
	`
	expect = "a"
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