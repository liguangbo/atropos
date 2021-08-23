package fate

import (
	"testing"
)

func TestSimpleAssignment(t *testing.T) {
	var expect interface{}
	data := []byte("{}")
	// int64 assign to simple var
	script := `
	number = 1
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

	// float64 assign to simple var
	script = `
	number = 1.5
	`
	expect = float64(1.5)
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

	// bool assign to simple var
	script = `
	number = false
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

	// string assign to simple var
	script = `
	number = "test"
	`
	expect = "test"
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

	// string assign to simple var
	script = `
	number = ""
	`
	expect = ""
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

	// nil assign to simple var
	script = `
	number = nil
	`
	expect = nil
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

	// array assign to simple var
	script = `
	number = [1,1.1,true,"test",nil,[1]]
	`
	expect = nil
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got1, ok := mem["number"].([]interface{})
	if !ok || len(got1) != 6 {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got1, got1)
	}

	// func count result assign to simple var
	script = `
	number = [1,1.1,true,"test",nil,[1]]
	number = number.count()
	`
	expect = int(6)
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

	// func sum result assign to simple var
	script = `
	number = [1,1.1,2]
	number = number.sum()
	`
	expect = float64(4.1)
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

	// func avg result assign to simple var
	script = `
	number = [1,1.5,2]
	number = number.avg()
	`
	expect = float64(1.5)
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

	// func max result assign to simple var
	script = `
	number = [1,1.5,2]
	number = number.max()
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

	// func min result assign to simple var
	script = `
	number = [1,1.5,2]
	number = number.min()
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
}

func TestObjectAssignment(t *testing.T) {
	var expect interface{}
	data := []byte(`{"object1":{"object2":{}}}`)
	// int64 assign to simple var
	script := `
	object1.object2.number = 1
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
	got := mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// float64 assign to simple var
	script = `
	object1.object2.number = 1.5
	`
	expect = float64(1.5)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// bool assign to simple var
	script = `
	object1.object2.number = false
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
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// string assign to simple var
	script = `
	object1.object2.number = "test"
	`
	expect = "test"
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// nil assign to simple var
	script = `
	object1.object2.number = nil
	`
	expect = nil
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// array assign to simple var
	script = `
	object1.object2.number = [1,1.1,true,"test",nil,[1]]
	`
	expect = nil
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got1, ok := mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"].([]interface{})
	if !ok || len(got1) != 6 {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got1, got1)
	}

	// func count result assign to simple var
	script = `
	object1.object2.number = [1,1.1,true,"test",nil,[1]]
	object1.object2.number = object1.object2.number.count()
	`
	expect = int(6)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func sum result assign to simple var
	script = `
	object1.object2.number = [1,1.1,2]
	object1.object2.number = object1.object2.number.sum()
	`
	expect = float64(4.1)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func avg result assign to simple var
	script = `
	object1.object2.number = [1,1.5,2]
	object1.object2.number = object1.object2.number.avg()
	`
	expect = float64(1.5)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func max result assign to simple var
	script = `
	object1.object2.number = [1,1.5,2]
	object1.object2.number = object1.object2.number.max()
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
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func min result assign to simple var
	script = `
	object1.object2.number = [1,1.5,2]
	object1.object2.number = object1.object2.number.min()
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
	got = mem["object1"].(map[string]interface{})["object2"].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

}

func TestArrayAssignment(t *testing.T) {
	var expect interface{}
	data := []byte(`{"object1":{"object2":[0]}}`)
	// int64 assign to simple var
	script := `
	object1.object2[0] = 1
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
	got := mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// float64 assign to simple var
	script = `
	object1.object2[0] = 1.5
	`
	expect = float64(1.5)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// bool assign to simple var
	script = `
	object1.object2[0] = false
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
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// string assign to simple var
	script = `
	object1.object2[0] = "test"
	`
	expect = "test"
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// nil assign to simple var
	script = `
	object1.object2[0] = nil
	`
	expect = nil
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// array assign to simple var
	script = `
	object1.object2[0] = [1,1.1,true,"test",nil,[1]]
	`
	expect = nil
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got1, ok := mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].([]interface{})
	if !ok || len(got1) != 6 {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got1, got1)
	}

	// func count result assign to simple var
	script = `
	object1.object2[0] = [1,1.1,true,"test",nil,[1]]
	object1.object2[0] = object1.object2[0].count()
	`
	expect = int(6)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func sum result assign to simple var
	script = `
	object1.object2[0] = [1,1.1,2]
	object1.object2[0] = object1.object2[0].sum()
	`
	expect = float64(4.1)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func avg result assign to simple var
	script = `
	object1.object2[0] = [1,1.5,2]
	object1.object2[0] = object1.object2[0].avg()
	`
	expect = float64(1.5)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func max result assign to simple var
	script = `
	object1.object2[0] = [1,1.5,2]
	object1.object2[0] = object1.object2[0].max()
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
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func min result assign to simple var
	script = `
	object1.object2[0] = [1,1.5,2]
	object1.object2[0] = object1.object2[0].min()
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
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}
}

func TestObjectInArrayAssignment(t *testing.T) {
	var expect interface{}
	data := []byte(`{"object1":{"object2":[{}]}}`)
	// int64 assign to simple var
	script := `
	object1.object2[0].number = 1
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
	got := mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// float64 assign to simple var
	script = `
	object1.object2[0].number = 1.5
	`
	expect = float64(1.5)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// bool assign to simple var
	script = `
	object1.object2[0].number = false
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
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// string assign to simple var
	script = `
	object1.object2[0].number = "test"
	`
	expect = "test"
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// nil assign to simple var
	script = `
	object1.object2[0].number = nil
	`
	expect = nil
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// array assign to simple var
	script = `
	object1.object2[0].number = [1,1.1,true,"test",nil,[1]]
	`
	expect = nil
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got1, ok := mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"].([]interface{})
	if !ok || len(got1) != 6 {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got1, got1)
	}

	// func count result assign to simple var
	script = `
	object1.object2[0].number = [1,1.1,true,"test",nil,[1]]
	object1.object2[0].number = object1.object2[0].number.count()
	`
	expect = int(6)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func sum result assign to simple var
	script = `
	object1.object2[0].number = [1,1.1,2]
	object1.object2[0].number = object1.object2[0].number.sum()
	`
	expect = float64(4.1)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func avg result assign to simple var
	script = `
	object1.object2[0].number = [1,1.5,2]
	object1.object2[0].number = object1.object2[0].number.avg()
	`
	expect = float64(1.5)
	code, err = Compile(script)
	if err != nil {
		t.Errorf("compile %s expect err = nil, but got %v", script, err)
	}
	mem, err = Run(code, data)
	if err != nil {
		t.Errorf("run %s expect err = nil, but got %v", script, err)
	}
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func max result assign to simple var
	script = `
	object1.object2[0].number = [1,1.5,2]
	object1.object2[0].number = object1.object2[0].number.max()
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
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	// func min result assign to simple var
	script = `
	object1.object2[0].number = [1,1.5,2]
	object1.object2[0].number = object1.object2[0].number.min()
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
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0].(map[string]interface{})["number"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}
}

func TestComplicateAssignment(t *testing.T) {
	var expect interface{}
	data := []byte(`{"object1":{"object2":[{"var":true}]}}`)

	script := `
	number = object1.object2[0].var
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
	object1.object2[0] = object1.object2[0].var
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
	got = mem["object1"].(map[string]interface{})["object2"].([]interface{})[0]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	object1.object2 = object1.object2[0].var
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
	got = mem["object1"].(map[string]interface{})["object2"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	object1 = object1.object2[0].var
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
	got = mem["object1"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	object1 = object1.object2[0]
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
	got = mem["object1"].(map[string]interface{})["var"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}

	script = `
	object1 = object1.object2
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
	got = mem["object1"].([]interface{})[0].(map[string]interface{})["var"]
	if expect != got {
		t.Errorf("run %s expect %T(%d), but got %T(%d)", script, expect, expect, got, got)
	}
}
