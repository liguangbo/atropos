package main

import (
	"encoding/json"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/liguangbo/atropos/fate"
	"github.com/liguangbo/atropos/fate/grammar"
)

var data = `
{
	"a": 3,
	"b": 100,
	"data": {
		"arr": [
			["1", "2", "3"],
			["4", "5", "6"],
			null
		],
		"cores": ["system", "tbs", "hahah"],
		"webview": "tbs",
		"system": {
			"out": {
				"println": "你好"
			},
			"1": 2
		}
	}
}
`

func main() {
	example := `
	if a == 4 then
		a = 5
		return
	elseif a == 3 then
		a = 6
	end

	arr = [a, "1", 1, "test", a + 1]

	if arr contains [a, a+1, a+2] then
		arr[0] = 7
	end

	if ["system", "tbs",  "hahah"] !contains data.cores then
		data.cores[0] = "happy"
	end

	if data.c == nil then
		b = "fuck"
	end

	if data.cores.count() > 2 then
		c = "aaaa"
	end

	if data.webview == "core" || data.webview == "tbs" && a == 6 then
		data.webview = "fuck"
	end

	data.cores[2] = data.webview
	data.arr[0][2] = "fuck you"
	if data.arr[0][2] regexpMatch ".*you.*" then
		data.system.out.println = "再见"
	end
	data.system[1] = 3

	testUnaryMinus = -1
	testUnaryMinus1 = -testUnaryMinus
	testUnaryExclamation = true
	testUnaryExclamation1 = !testUnaryExclamation
	`
	runExample(example, data)
}

func runExample(example, input string) error {
	stream := antlr.NewInputStream(example)
	lexer := grammar.NewFateLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewFateParser(tokens)
	tree := parser.Fate()
	mem := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &mem)
	if err != nil {
		return fmt.Errorf("unmarshal input error, %v", err)
	}
	v := fate.NewFateVisitor(mem)
	v.Visit(tree)
	result, _ := json.MarshalIndent(v.Mem(), "", "    ")
	fmt.Printf("Example:\n %s\n", example)
	fmt.Printf("Error: %v\n", v.Error())
	fmt.Printf("Output:\n %s\n", string(result))
	return nil
}
