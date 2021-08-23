// Generated from /Users/liguangbo/Documents/GOPATH/src/github.com/liguangbo/atropos/fate/grammar/Fate.g4 by ANTLR 4.7.

package grammar // Fate
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 146, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 
	3, 2, 3, 2, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 4, 3, 4, 
	3, 4, 5, 4, 37, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 54, 10, 7, 12, 7, 14, 7, 
	57, 11, 7, 3, 7, 3, 7, 5, 7, 61, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 
	8, 3, 8, 3, 8, 6, 8, 71, 10, 8, 13, 8, 14, 8, 72, 5, 8, 75, 10, 8, 3, 8, 
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 87, 10, 8, 
	12, 8, 14, 8, 90, 11, 8, 3, 8, 3, 8, 5, 8, 94, 10, 8, 3, 8, 3, 8, 3, 8, 
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 
	3, 8, 3, 8, 3, 8, 7, 8, 114, 10, 8, 12, 8, 14, 8, 117, 11, 8, 3, 9, 3, 
	9, 3, 9, 7, 9, 122, 10, 9, 12, 9, 14, 9, 125, 11, 9, 3, 10, 3, 10, 7, 10, 
	129, 10, 10, 12, 10, 14, 10, 132, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 
	11, 3, 11, 5, 11, 140, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 2, 3, 
	14, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 7, 3, 2, 12, 13, 3, 
	2, 19, 21, 4, 2, 13, 13, 22, 22, 3, 2, 23, 29, 3, 2, 30, 31, 2, 159, 2, 
	24, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2, 8, 38, 3, 2, 2, 2, 
	10, 42, 3, 2, 2, 2, 12, 44, 3, 2, 2, 2, 14, 93, 3, 2, 2, 2, 16, 118, 3, 
	2, 2, 2, 18, 126, 3, 2, 2, 2, 20, 139, 3, 2, 2, 2, 22, 141, 3, 2, 2, 2, 
	24, 25, 5, 4, 3, 2, 25, 26, 7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 29, 5, 6, 
	4, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 
	3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 37, 5, 8, 5, 2, 
	34, 37, 5, 12, 7, 2, 35, 37, 5, 10, 6, 2, 36, 33, 3, 2, 2, 2, 36, 34, 3, 
	2, 2, 2, 36, 35, 3, 2, 2, 2, 37, 7, 3, 2, 2, 2, 38, 39, 5, 16, 9, 2, 39, 
	40, 7, 3, 2, 2, 40, 41, 5, 14, 8, 2, 41, 9, 3, 2, 2, 2, 42, 43, 7, 4, 2, 
	2, 43, 11, 3, 2, 2, 2, 44, 45, 7, 5, 2, 2, 45, 46, 5, 14, 8, 2, 46, 47, 
	7, 6, 2, 2, 47, 55, 5, 4, 3, 2, 48, 49, 7, 7, 2, 2, 49, 50, 5, 14, 8, 2, 
	50, 51, 7, 6, 2, 2, 51, 52, 5, 4, 3, 2, 52, 54, 3, 2, 2, 2, 53, 48, 3, 
	2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 
	60, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 59, 7, 8, 2, 2, 59, 61, 5, 4, 3, 
	2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 
	7, 9, 2, 2, 63, 13, 3, 2, 2, 2, 64, 65, 8, 8, 1, 2, 65, 94, 5, 20, 11, 
	2, 66, 74, 5, 16, 9, 2, 67, 68, 7, 10, 2, 2, 68, 69, 7, 39, 2, 2, 69, 71, 
	7, 11, 2, 2, 70, 67, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 
	72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 70, 3, 2, 2, 2, 74, 75, 3, 
	2, 2, 2, 75, 94, 3, 2, 2, 2, 76, 77, 9, 2, 2, 2, 77, 94, 5, 14, 8, 11, 
	78, 79, 7, 14, 2, 2, 79, 80, 5, 14, 8, 2, 80, 81, 7, 15, 2, 2, 81, 94, 
	3, 2, 2, 2, 82, 83, 7, 16, 2, 2, 83, 88, 5, 14, 8, 2, 84, 85, 7, 17, 2, 
	2, 85, 87, 5, 14, 8, 2, 86, 84, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 
	3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 91, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 
	91, 92, 7, 18, 2, 2, 92, 94, 3, 2, 2, 2, 93, 64, 3, 2, 2, 2, 93, 66, 3, 
	2, 2, 2, 93, 76, 3, 2, 2, 2, 93, 78, 3, 2, 2, 2, 93, 82, 3, 2, 2, 2, 94, 
	115, 3, 2, 2, 2, 95, 96, 12, 8, 2, 2, 96, 97, 9, 3, 2, 2, 97, 114, 5, 14, 
	8, 9, 98, 99, 12, 7, 2, 2, 99, 100, 9, 4, 2, 2, 100, 114, 5, 14, 8, 8, 
	101, 102, 12, 6, 2, 2, 102, 103, 9, 5, 2, 2, 103, 114, 5, 14, 8, 7, 104, 
	105, 12, 5, 2, 2, 105, 106, 9, 6, 2, 2, 106, 114, 5, 14, 8, 6, 107, 108, 
	12, 4, 2, 2, 108, 109, 7, 32, 2, 2, 109, 114, 5, 14, 8, 5, 110, 111, 12, 
	3, 2, 2, 111, 112, 7, 33, 2, 2, 112, 114, 5, 14, 8, 4, 113, 95, 3, 2, 2, 
	2, 113, 98, 3, 2, 2, 2, 113, 101, 3, 2, 2, 2, 113, 104, 3, 2, 2, 2, 113, 
	107, 3, 2, 2, 2, 113, 110, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 
	3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 15, 3, 2, 2, 2, 117, 115, 3, 2, 
	2, 2, 118, 123, 5, 18, 10, 2, 119, 120, 7, 10, 2, 2, 120, 122, 5, 18, 10, 
	2, 121, 119, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 
	124, 3, 2, 2, 2, 124, 17, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 126, 130, 7, 
	40, 2, 2, 127, 129, 5, 22, 12, 2, 128, 127, 3, 2, 2, 2, 129, 132, 3, 2, 
	2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 19, 3, 2, 2, 2, 
	132, 130, 3, 2, 2, 2, 133, 140, 7, 41, 2, 2, 134, 140, 7, 42, 2, 2, 135, 
	140, 7, 43, 2, 2, 136, 140, 7, 36, 2, 2, 137, 140, 7, 37, 2, 2, 138, 140, 
	7, 38, 2, 2, 139, 133, 3, 2, 2, 2, 139, 134, 3, 2, 2, 2, 139, 135, 3, 2, 
	2, 2, 139, 136, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 138, 3, 2, 2, 2, 
	140, 21, 3, 2, 2, 2, 141, 142, 7, 16, 2, 2, 142, 143, 7, 42, 2, 2, 143, 
	144, 7, 18, 2, 2, 144, 23, 3, 2, 2, 2, 15, 30, 36, 55, 60, 72, 74, 88, 
	93, 113, 115, 123, 130, 139,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'return'", "'if'", "'then'", "'elseif'", "'else'", "'end'", 
	"'.'", "'()'", "'!'", "'-'", "'('", "')'", "'['", "','", "']'", "'*'", 
	"'/'", "'%'", "'+'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'regexpMatch'", 
	"'contains'", "'!contains'", "'&&'", "'||'", "", "", "'true'", "'false'", 
	"'nil'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "LINE_COMMENT", 
	"COMMENT", "TRUE", "FALSE", "NIL", "FUNC", "NAME", "STRING", "INT", "FLOAT", 
	"DIGIT", "WS",
}

var ruleNames = []string{
	"fate", "block", "stat", "assignstat", "returnstat", "ifstat", "expr", 
	"variable", "field", "value", "index",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FateParser struct {
	*antlr.BaseParser
}

func NewFateParser(input antlr.TokenStream) *FateParser {
	this := new(FateParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Fate.g4"

	return this
}

// FateParser tokens.
const (
	FateParserEOF = antlr.TokenEOF
	FateParserT__0 = 1
	FateParserT__1 = 2
	FateParserT__2 = 3
	FateParserT__3 = 4
	FateParserT__4 = 5
	FateParserT__5 = 6
	FateParserT__6 = 7
	FateParserT__7 = 8
	FateParserT__8 = 9
	FateParserT__9 = 10
	FateParserT__10 = 11
	FateParserT__11 = 12
	FateParserT__12 = 13
	FateParserT__13 = 14
	FateParserT__14 = 15
	FateParserT__15 = 16
	FateParserT__16 = 17
	FateParserT__17 = 18
	FateParserT__18 = 19
	FateParserT__19 = 20
	FateParserT__20 = 21
	FateParserT__21 = 22
	FateParserT__22 = 23
	FateParserT__23 = 24
	FateParserT__24 = 25
	FateParserT__25 = 26
	FateParserT__26 = 27
	FateParserT__27 = 28
	FateParserT__28 = 29
	FateParserT__29 = 30
	FateParserT__30 = 31
	FateParserLINE_COMMENT = 32
	FateParserCOMMENT = 33
	FateParserTRUE = 34
	FateParserFALSE = 35
	FateParserNIL = 36
	FateParserFUNC = 37
	FateParserNAME = 38
	FateParserSTRING = 39
	FateParserINT = 40
	FateParserFLOAT = 41
	FateParserDIGIT = 42
	FateParserWS = 43
)

// FateParser rules.
const (
	FateParserRULE_fate = 0
	FateParserRULE_block = 1
	FateParserRULE_stat = 2
	FateParserRULE_assignstat = 3
	FateParserRULE_returnstat = 4
	FateParserRULE_ifstat = 5
	FateParserRULE_expr = 6
	FateParserRULE_variable = 7
	FateParserRULE_field = 8
	FateParserRULE_value = 9
	FateParserRULE_index = 10
)

// IFateContext is an interface to support dynamic dispatch.
type IFateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFateContext differentiates from other interfaces.
	IsFateContext()
}

type FateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFateContext() *FateContext {
	var p = new(FateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_fate
	return p
}

func (*FateContext) IsFateContext() {}

func NewFateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FateContext {
	var p = new(FateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_fate

	return p
}

func (s *FateContext) GetParser() antlr.Parser { return s.parser }

func (s *FateContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FateContext) EOF() antlr.TerminalNode {
	return s.GetToken(FateParserEOF, 0)
}

func (s *FateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterFate(s)
	}
}

func (s *FateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitFate(s)
	}
}

func (s *FateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitFate(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Fate() (localctx IFateContext) {
	localctx = NewFateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FateParserRULE_fate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Block()
	}
	{
		p.SetState(23)
		p.Match(FateParserEOF)
	}



	return localctx
}


// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FateParserRULE_block)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == FateParserT__1 || _la == FateParserT__2 || _la == FateParserNAME {
		{
			p.SetState(25)
			p.Stat()
		}


		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Assignstat() IAssignstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignstatContext)
}

func (s *StatContext) Ifstat() IIfstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatContext)
}

func (s *StatContext) Returnstat() IReturnstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnstatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitStat(s)
	}
}

func (s *StatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitStat(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FateParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FateParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Assignstat()
		}


	case FateParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Ifstat()
		}


	case FateParserT__1:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Returnstat()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IAssignstatContext is an interface to support dynamic dispatch.
type IAssignstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignstatContext differentiates from other interfaces.
	IsAssignstatContext()
}

type AssignstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignstatContext() *AssignstatContext {
	var p = new(AssignstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_assignstat
	return p
}

func (*AssignstatContext) IsAssignstatContext() {}

func NewAssignstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignstatContext {
	var p = new(AssignstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_assignstat

	return p
}

func (s *AssignstatContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignstatContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignstatContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AssignstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterAssignstat(s)
	}
}

func (s *AssignstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitAssignstat(s)
	}
}

func (s *AssignstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitAssignstat(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Assignstat() (localctx IAssignstatContext) {
	localctx = NewAssignstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FateParserRULE_assignstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Variable()
	}
	{
		p.SetState(37)
		p.Match(FateParserT__0)
	}
	{
		p.SetState(38)
		p.expr(0)
	}



	return localctx
}


// IReturnstatContext is an interface to support dynamic dispatch.
type IReturnstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnstatContext differentiates from other interfaces.
	IsReturnstatContext()
}

type ReturnstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnstatContext() *ReturnstatContext {
	var p = new(ReturnstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_returnstat
	return p
}

func (*ReturnstatContext) IsReturnstatContext() {}

func NewReturnstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnstatContext {
	var p = new(ReturnstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_returnstat

	return p
}

func (s *ReturnstatContext) GetParser() antlr.Parser { return s.parser }
func (s *ReturnstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturnstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterReturnstat(s)
	}
}

func (s *ReturnstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitReturnstat(s)
	}
}

func (s *ReturnstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitReturnstat(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Returnstat() (localctx IReturnstatContext) {
	localctx = NewReturnstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FateParserRULE_returnstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(FateParserT__1)
	}



	return localctx
}


// IIfstatContext is an interface to support dynamic dispatch.
type IIfstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatContext differentiates from other interfaces.
	IsIfstatContext()
}

type IfstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatContext() *IfstatContext {
	var p = new(IfstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_ifstat
	return p
}

func (*IfstatContext) IsIfstatContext() {}

func NewIfstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatContext {
	var p = new(IfstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_ifstat

	return p
}

func (s *IfstatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IfstatContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstatContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfstatContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IfstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterIfstat(s)
	}
}

func (s *IfstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitIfstat(s)
	}
}

func (s *IfstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitIfstat(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Ifstat() (localctx IIfstatContext) {
	localctx = NewIfstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FateParserRULE_ifstat)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(FateParserT__2)
	}
	{
		p.SetState(43)
		p.expr(0)
	}
	{
		p.SetState(44)
		p.Match(FateParserT__3)
	}
	{
		p.SetState(45)
		p.Block()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == FateParserT__4 {
		{
			p.SetState(46)
			p.Match(FateParserT__4)
		}
		{
			p.SetState(47)
			p.expr(0)
		}
		{
			p.SetState(48)
			p.Match(FateParserT__3)
		}
		{
			p.SetState(49)
			p.Block()
		}


		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == FateParserT__5 {
		{
			p.SetState(56)
			p.Match(FateParserT__5)
		}
		{
			p.SetState(57)
			p.Block()
		}

	}
	{
		p.SetState(60)
		p.Match(FateParserT__6)
	}



	return localctx
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type ParensContext struct {
	*ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}


type ValueExprContext struct {
	*ExprContext
}

func NewValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExprContext {
	var p = new(ValueExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}


func (s *ValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterValueExpr(s)
	}
}

func (s *ValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitValueExpr(s)
	}
}

func (s *ValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type UnaryExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *UnaryExprContext) GetOp() antlr.Token { return s.op }


func (s *UnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type SimpleBooleanExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewSimpleBooleanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleBooleanExprContext {
	var p = new(SimpleBooleanExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *SimpleBooleanExprContext) GetOp() antlr.Token { return s.op }


func (s *SimpleBooleanExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *SimpleBooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleBooleanExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SimpleBooleanExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *SimpleBooleanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterSimpleBooleanExpr(s)
	}
}

func (s *SimpleBooleanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitSimpleBooleanExpr(s)
	}
}

func (s *SimpleBooleanExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitSimpleBooleanExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayExprContext struct {
	*ExprContext
}

func NewArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExprContext {
	var p = new(ArrayExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArrayExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ArrayExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterArrayExpr(s)
	}
}

func (s *ArrayExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitArrayExpr(s)
	}
}

func (s *ArrayExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitArrayExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type OrExprContext struct {
	*ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type AddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *AddSubContext) GetOp() antlr.Token { return s.op }


func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}


type VarFuncExprContext struct {
	*ExprContext
}

func NewVarFuncExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarFuncExprContext {
	var p = new(VarFuncExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VarFuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarFuncExprContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarFuncExprContext) AllFUNC() []antlr.TerminalNode {
	return s.GetTokens(FateParserFUNC)
}

func (s *VarFuncExprContext) FUNC(i int) antlr.TerminalNode {
	return s.GetToken(FateParserFUNC, i)
}


func (s *VarFuncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterVarFuncExpr(s)
	}
}

func (s *VarFuncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitVarFuncExpr(s)
	}
}

func (s *VarFuncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitVarFuncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type SetBooleanExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewSetBooleanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetBooleanExprContext {
	var p = new(SetBooleanExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *SetBooleanExprContext) GetOp() antlr.Token { return s.op }


func (s *SetBooleanExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *SetBooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetBooleanExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SetBooleanExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *SetBooleanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterSetBooleanExpr(s)
	}
}

func (s *SetBooleanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitSetBooleanExpr(s)
	}
}

func (s *SetBooleanExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitSetBooleanExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type MulDivContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *MulDivContext) GetOp() antlr.Token { return s.op }


func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}


type AndExprContext struct {
	*ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *FateParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *FateParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, FateParserRULE_expr, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FateParserTRUE, FateParserFALSE, FateParserNIL, FateParserSTRING, FateParserINT, FateParserFLOAT:
		localctx = NewValueExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(63)
			p.Value()
		}


	case FateParserNAME:
		localctx = NewVarFuncExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.Variable()
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
						{
							p.SetState(65)
							p.Match(FateParserT__7)
						}
						{
							p.SetState(66)
							p.Match(FateParserFUNC)
						}
						{
							p.SetState(67)
							p.Match(FateParserT__8)
						}




				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(70)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
			}


		}


	case FateParserT__9, FateParserT__10:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(74)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*UnaryExprContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == FateParserT__9 || _la == FateParserT__10) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*UnaryExprContext).op = _ri
		} else {
		    p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(75)
			p.expr(9)
		}


	case FateParserT__11:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(76)
			p.Match(FateParserT__11)
		}
		{
			p.SetState(77)
			p.expr(0)
		}
		{
			p.SetState(78)
			p.Match(FateParserT__12)
		}


	case FateParserT__13:
		localctx = NewArrayExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.Match(FateParserT__13)
		}
		{
			p.SetState(81)
			p.expr(0)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == FateParserT__14 {
			{
				p.SetState(82)
				p.Match(FateParserT__14)
			}
			{
				p.SetState(83)
				p.expr(0)
			}


			p.SetState(88)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(89)
			p.Match(FateParserT__15)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FateParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				p.SetState(94)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MulDivContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << FateParserT__16) | (1 << FateParserT__17) | (1 << FateParserT__18))) != 0)) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MulDivContext).op = _ri
				} else {
				    p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(95)
					p.expr(7)
				}


			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FateParserRULE_expr)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(97)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AddSubContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == FateParserT__10 || _la == FateParserT__19) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AddSubContext).op = _ri
				} else {
				    p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(98)
					p.expr(6)
				}


			case 3:
				localctx = NewSimpleBooleanExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FateParserRULE_expr)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(100)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*SimpleBooleanExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << FateParserT__20) | (1 << FateParserT__21) | (1 << FateParserT__22) | (1 << FateParserT__23) | (1 << FateParserT__24) | (1 << FateParserT__25) | (1 << FateParserT__26))) != 0)) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*SimpleBooleanExprContext).op = _ri
				} else {
				    p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(101)
					p.expr(5)
				}


			case 4:
				localctx = NewSetBooleanExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FateParserRULE_expr)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(103)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*SetBooleanExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == FateParserT__27 || _la == FateParserT__28) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*SetBooleanExprContext).op = _ri
				} else {
				    p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(104)
					p.expr(4)
				}


			case 5:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FateParserRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(106)
					p.Match(FateParserT__29)
				}
				{
					p.SetState(107)
					p.expr(3)
				}


			case 6:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FateParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(109)
					p.Match(FateParserT__30)
				}
				{
					p.SetState(110)
					p.expr(2)
				}

			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}



	return localctx
}


// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *VariableContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FateParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Field()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(117)
				p.Match(FateParserT__7)
			}
			{
				p.SetState(118)
				p.Field()
			}


		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}



	return localctx
}


// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(FateParserNAME, 0)
}

func (s *FieldContext) AllIndex() []IIndexContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIndexContext)(nil)).Elem())
	var tst = make([]IIndexContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIndexContext)
		}
	}

	return tst
}

func (s *FieldContext) Index(i int) IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FateParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(FateParserNAME)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(125)
				p.Index()
			}


		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}



	return localctx
}


// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type NilContext struct {
	*ValueContext
}

func NewNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilContext {
	var p = new(NilContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilContext) NIL() antlr.TerminalNode {
	return s.GetToken(FateParserNIL, 0)
}


func (s *NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterNil(s)
	}
}

func (s *NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitNil(s)
	}
}

func (s *NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitNil(s)

	default:
		return t.VisitChildren(s)
	}
}


type StringContext struct {
	*ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(FateParserSTRING, 0)
}


func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}


type TrueContext struct {
	*ValueContext
}

func NewTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueContext {
	var p = new(TrueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(FateParserTRUE, 0)
}


func (s *TrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterTrue(s)
	}
}

func (s *TrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitTrue(s)
	}
}

func (s *TrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitTrue(s)

	default:
		return t.VisitChildren(s)
	}
}


type FalseContext struct {
	*ValueContext
}

func NewFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseContext {
	var p = new(FalseContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(FateParserFALSE, 0)
}


func (s *FalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterFalse(s)
	}
}

func (s *FalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitFalse(s)
	}
}

func (s *FalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitFalse(s)

	default:
		return t.VisitChildren(s)
	}
}


type FloatContext struct {
	*ValueContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(FateParserFLOAT, 0)
}


func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}


type IntContext struct {
	*ValueContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(FateParserINT, 0)
}


func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *FateParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FateParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FateParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(FateParserSTRING)
		}


	case FateParserINT:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(FateParserINT)
		}


	case FateParserFLOAT:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Match(FateParserFLOAT)
		}


	case FateParserTRUE:
		localctx = NewTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)
			p.Match(FateParserTRUE)
		}


	case FateParserFALSE:
		localctx = NewFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(135)
			p.Match(FateParserFALSE)
		}


	case FateParserNIL:
		localctx = NewNilContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(136)
			p.Match(FateParserNIL)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FateParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FateParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) INT() antlr.TerminalNode {
	return s.GetToken(FateParserINT, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FateListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FateVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *FateParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FateParserRULE_index)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(FateParserT__13)
	}
	{
		p.SetState(140)
		p.Match(FateParserINT)
	}
	{
		p.SetState(141)
		p.Match(FateParserT__15)
	}



	return localctx
}


func (p *FateParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FateParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

