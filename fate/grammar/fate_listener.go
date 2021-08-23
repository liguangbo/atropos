// Generated from /Users/liguangbo/Documents/GOPATH/src/github.com/liguangbo/atropos/fate/grammar/Fate.g4 by ANTLR 4.7.

package grammar // Fate
import "github.com/antlr/antlr4/runtime/Go/antlr"

// FateListener is a complete listener for a parse tree produced by FateParser.
type FateListener interface {
	antlr.ParseTreeListener

	// EnterFate is called when entering the fate production.
	EnterFate(c *FateContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterAssignstat is called when entering the assignstat production.
	EnterAssignstat(c *AssignstatContext)

	// EnterReturnstat is called when entering the returnstat production.
	EnterReturnstat(c *ReturnstatContext)

	// EnterIfstat is called when entering the ifstat production.
	EnterIfstat(c *IfstatContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterValueExpr is called when entering the valueExpr production.
	EnterValueExpr(c *ValueExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterSimpleBooleanExpr is called when entering the simpleBooleanExpr production.
	EnterSimpleBooleanExpr(c *SimpleBooleanExprContext)

	// EnterArrayExpr is called when entering the arrayExpr production.
	EnterArrayExpr(c *ArrayExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAddSub is called when entering the addSub production.
	EnterAddSub(c *AddSubContext)

	// EnterVarFuncExpr is called when entering the varFuncExpr production.
	EnterVarFuncExpr(c *VarFuncExprContext)

	// EnterSetBooleanExpr is called when entering the setBooleanExpr production.
	EnterSetBooleanExpr(c *SetBooleanExprContext)

	// EnterMulDiv is called when entering the mulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// EnterFloat is called when entering the float production.
	EnterFloat(c *FloatContext)

	// EnterTrue is called when entering the true production.
	EnterTrue(c *TrueContext)

	// EnterFalse is called when entering the false production.
	EnterFalse(c *FalseContext)

	// EnterNil is called when entering the nil production.
	EnterNil(c *NilContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// ExitFate is called when exiting the fate production.
	ExitFate(c *FateContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitAssignstat is called when exiting the assignstat production.
	ExitAssignstat(c *AssignstatContext)

	// ExitReturnstat is called when exiting the returnstat production.
	ExitReturnstat(c *ReturnstatContext)

	// ExitIfstat is called when exiting the ifstat production.
	ExitIfstat(c *IfstatContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitValueExpr is called when exiting the valueExpr production.
	ExitValueExpr(c *ValueExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitSimpleBooleanExpr is called when exiting the simpleBooleanExpr production.
	ExitSimpleBooleanExpr(c *SimpleBooleanExprContext)

	// ExitArrayExpr is called when exiting the arrayExpr production.
	ExitArrayExpr(c *ArrayExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAddSub is called when exiting the addSub production.
	ExitAddSub(c *AddSubContext)

	// ExitVarFuncExpr is called when exiting the varFuncExpr production.
	ExitVarFuncExpr(c *VarFuncExprContext)

	// ExitSetBooleanExpr is called when exiting the setBooleanExpr production.
	ExitSetBooleanExpr(c *SetBooleanExprContext)

	// ExitMulDiv is called when exiting the mulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)

	// ExitFloat is called when exiting the float production.
	ExitFloat(c *FloatContext)

	// ExitTrue is called when exiting the true production.
	ExitTrue(c *TrueContext)

	// ExitFalse is called when exiting the false production.
	ExitFalse(c *FalseContext)

	// ExitNil is called when exiting the nil production.
	ExitNil(c *NilContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)
}
