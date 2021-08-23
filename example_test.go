package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("assignment", TestAssignment)
	t.Run("expression", TestExpression)
	t.Run("ifstatement", TestIfStatement)
	t.Run("return", TestReturn)
	t.Run("func", TestFunc)
	t.Run("comment", TestComment)
}

func TestAssignment(t *testing.T) {
	mem := `
{
    "a": 0,
    "input": {
        "a": 1,
        "b": [
            "test",
            true,
            2,
            2.1
        ],
        "c": [
            [
                "test1",
                false,
                3
            ],
            "test2"
        ]
    },
    "output": {
        "a": 1,
        "b": [
            "test",
            true,
            2,
            2.1
        ],
        "c": [
            [
                "test1",
                false,
                3
            ],
            "test2"
        ]
    }
}
`
	fate := `
a = 12345
input.a = a
input.b = ['foo', 'bar']
input.c[1] = 'foobar'
input.c[0][0] = 'ping'
input.c[0][1] = true

output.a = input.a
output.b = input.b
output.c[1] = input.c[1]
output.c[0][1] = input.c[0][1]
output.c[0][2] = nil
`
	err := runExample(fate, mem)
	if err != nil {
		t.Fatalf("run example error, %v", err)
	}
}

func TestExpression(t *testing.T) {
	mem := `
{
    "b": {
        "t": true,
        "f": false
    },
    "i": 128,
    "f": 10.5,
    "s": "foo",
	"n": null,
	"array": ["hello", "world"],
    "output": {}
}`
	fate := `
output.equal = i == 128
output.is_null = n == nil
output.not_equal = i != 127
output.greater = i > 100
output.greater_equal = i >= 128
output.less = i < 130
output.less_equal = i <= 128
output.regexp = s regexpMatch "^f.*"
output.add = i + i
output.sub = i - i
output.multi = i * 8
output.divide = i / 7
output.mod = i % 7
output.not = !b.f
output.minus = -i
output.parentheses = i + (i + i)
output.contains_object = array contains "hello"
output.contains_array = array contains ["hello"]
output.not_contains_object = array !contains "foo"
output.not_contains_array = array !contains ["hello","world","foo"]
output.and = b.t && i > 100
output.or = b.f || i > 1
output.and_or = b.t && (!b.f || array contains "bar")
`
	err := runExample(fate, mem)
	if err != nil {
		t.Fatalf("run example error, %v", err)
	}
}

func TestIfStatement(t *testing.T) {
	mem := `
{
    "b": {
        "t": true,
        "f": false
    },
    "i": 128,
    "f": 10.5,
    "s": "foo",
	"array": ["hello", "world"],
    "output": {}
}`
	fate := `
if b.t then
	output.if1 = "foo"
end

if b.t && b.f then
	output.if2 = "foo"
elseif i < 256 && f < 10.6
	output.if2 = "bar"
end

if array contains "foo" then
	output.if3 = "bar"
else
	output.if3 = "foo"
end
`
	err := runExample(fate, mem)
	if err != nil {
		t.Fatalf("run example error, %v", err)
	}
}

func TestReturn(t *testing.T) {
	mem := `
{
    "b": {
        "t": true,
        "f": false
    },
    "i": 128,
    "f": 10.5,
    "s": "foo",
	"array": ["hello", "world"],
    "output": {}
}`
	fate := `
if b.t then
	output.ret = "foo"
	return
end
output.ret = "bar"
`
	err := runExample(fate, mem)
	if err != nil {
		t.Fatalf("run example error, %v", err)
	}
}

func TestFunc(t *testing.T) {
	mem := `
{
    "b": {
        "t": true,
        "f": false
    },
    "i": 128,
    "f": 10.5,
    "s": "foo",
	"array": ["hello", "world"],
	"int_array": [1,2,3],
	"int_array_with_float": [1,2,3.1],
    "output": {}
}`
	fate := `
output.array_count = array.count()
output.object_count = b.count()
output.str_count = s.count()

output.int_array_sum = int_array.sum()
output.int_array_with_float_sum = int_array_with_float.sum()

output.int_array_avg = int_array.avg()
output.int_array_with_float_avg = int_array_with_float.avg()

output.int_array_max = int_array.max()
output.int_array_with_float_max = int_array_with_float.max()

output.int_array_min = int_array.min()
output.int_array_with_float_min = int_array_with_float.min()
`
	err := runExample(fate, mem)
	if err != nil {
		t.Fatalf("run example error, %v", err)
	}
}

func TestComment(t *testing.T) {
	mem := `
{
    "i": 128
}`
	fate := `
i = 64 
// comment line
/*
comment block
i = 256
*/
`
	err := runExample(fate, mem)
	if err != nil {
		t.Fatalf("run example error, %v", err)
	}
}
