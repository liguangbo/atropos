// Generated from /Users/liguangbo/Documents/GOPATH/src/github.com/liguangbo/atropos/fate/grammar/Fate.g4 by ANTLR 4.7.

package grammar // Fate
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFateListener is a complete listener for a parse tree produced by FateParser.
type BaseFateListener struct{}

var _ FateListener = &BaseFateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFate is called when production fate is entered.
func (s *BaseFateListener) EnterFate(ctx *FateContext) {}

// ExitFate is called when production fate is exited.
func (s *BaseFateListener) ExitFate(ctx *FateContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseFateListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseFateListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseFateListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseFateListener) ExitStat(ctx *StatContext) {}

// EnterAssignstat is called when production assignstat is entered.
func (s *BaseFateListener) EnterAssignstat(ctx *AssignstatContext) {}

// ExitAssignstat is called when production assignstat is exited.
func (s *BaseFateListener) ExitAssignstat(ctx *AssignstatContext) {}

// EnterReturnstat is called when production returnstat is entered.
func (s *BaseFateListener) EnterReturnstat(ctx *ReturnstatContext) {}

// ExitReturnstat is called when production returnstat is exited.
func (s *BaseFateListener) ExitReturnstat(ctx *ReturnstatContext) {}

// EnterIfstat is called when production ifstat is entered.
func (s *BaseFateListener) EnterIfstat(ctx *IfstatContext) {}

// ExitIfstat is called when production ifstat is exited.
func (s *BaseFateListener) ExitIfstat(ctx *IfstatContext) {}

// EnterParens is called when production parens is entered.
func (s *BaseFateListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BaseFateListener) ExitParens(ctx *ParensContext) {}

// EnterValueExpr is called when production valueExpr is entered.
func (s *BaseFateListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BaseFateListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseFateListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseFateListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterSimpleBooleanExpr is called when production simpleBooleanExpr is entered.
func (s *BaseFateListener) EnterSimpleBooleanExpr(ctx *SimpleBooleanExprContext) {}

// ExitSimpleBooleanExpr is called when production simpleBooleanExpr is exited.
func (s *BaseFateListener) ExitSimpleBooleanExpr(ctx *SimpleBooleanExprContext) {}

// EnterArrayExpr is called when production arrayExpr is entered.
func (s *BaseFateListener) EnterArrayExpr(ctx *ArrayExprContext) {}

// ExitArrayExpr is called when production arrayExpr is exited.
func (s *BaseFateListener) ExitArrayExpr(ctx *ArrayExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseFateListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseFateListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAddSub is called when production addSub is entered.
func (s *BaseFateListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production addSub is exited.
func (s *BaseFateListener) ExitAddSub(ctx *AddSubContext) {}

// EnterVarFuncExpr is called when production varFuncExpr is entered.
func (s *BaseFateListener) EnterVarFuncExpr(ctx *VarFuncExprContext) {}

// ExitVarFuncExpr is called when production varFuncExpr is exited.
func (s *BaseFateListener) ExitVarFuncExpr(ctx *VarFuncExprContext) {}

// EnterSetBooleanExpr is called when production setBooleanExpr is entered.
func (s *BaseFateListener) EnterSetBooleanExpr(ctx *SetBooleanExprContext) {}

// ExitSetBooleanExpr is called when production setBooleanExpr is exited.
func (s *BaseFateListener) ExitSetBooleanExpr(ctx *SetBooleanExprContext) {}

// EnterMulDiv is called when production mulDiv is entered.
func (s *BaseFateListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production mulDiv is exited.
func (s *BaseFateListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseFateListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseFateListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseFateListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseFateListener) ExitVariable(ctx *VariableContext) {}

// EnterField is called when production field is entered.
func (s *BaseFateListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseFateListener) ExitField(ctx *FieldContext) {}

// EnterString is called when production string is entered.
func (s *BaseFateListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseFateListener) ExitString(ctx *StringContext) {}

// EnterInt is called when production int is entered.
func (s *BaseFateListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseFateListener) ExitInt(ctx *IntContext) {}

// EnterFloat is called when production float is entered.
func (s *BaseFateListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production float is exited.
func (s *BaseFateListener) ExitFloat(ctx *FloatContext) {}

// EnterTrue is called when production true is entered.
func (s *BaseFateListener) EnterTrue(ctx *TrueContext) {}

// ExitTrue is called when production true is exited.
func (s *BaseFateListener) ExitTrue(ctx *TrueContext) {}

// EnterFalse is called when production false is entered.
func (s *BaseFateListener) EnterFalse(ctx *FalseContext) {}

// ExitFalse is called when production false is exited.
func (s *BaseFateListener) ExitFalse(ctx *FalseContext) {}

// EnterNil is called when production nil is entered.
func (s *BaseFateListener) EnterNil(ctx *NilContext) {}

// ExitNil is called when production nil is exited.
func (s *BaseFateListener) ExitNil(ctx *NilContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseFateListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseFateListener) ExitIndex(ctx *IndexContext) {}
