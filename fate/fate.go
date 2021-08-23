package fate

import (
	"encoding/json"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/liguangbo/atropos/fate/grammar"
)

func Compile(script string) (grammar.IFateContext, error) {
	input := antlr.NewInputStream(script)
	lexer := grammar.NewFateLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewFateParser(tokens)
	parser.RemoveErrorListeners()
	errorListener := NewStackErrorListener()
	parser.AddErrorListener(errorListener)
	tree := parser.Fate()
	if err := errorListener.Error(); err != nil {
		return nil, err
	}
	return tree, nil
}

func Run(code grammar.IFateContext, data []byte) (mem map[string]interface{}, err error) {
	mem = map[string]interface{}{}
	err = json.Unmarshal(data, &mem)
	if err != nil {
		return nil, err
	}
	v := NewFateVisitor(mem)
	v.Visit(code)
	return v.Mem(), v.Error()
}
