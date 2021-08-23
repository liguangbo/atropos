package fate

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/oleiade/lane"
	"strconv"
)

type StackErrorListener struct {
	stack *lane.Stack
	*antlr.DefaultErrorListener
}

func NewStackErrorListener() *StackErrorListener {
	return &StackErrorListener{stack: lane.NewStack()}
}

func (l *StackErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.stack.Push(fmt.Sprintln("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg))
}

func (l *StackErrorListener) Error() error {
	pop := l.stack.Pop()
	if pop == nil {
		return nil
	}
	return errors.New(pop.(string))
}
