// Generated from /Users/liguangbo/Documents/GOPATH/src/github.com/liguangbo/atropos/fate/grammar/Fate.g4 by ANTLR 4.7.

package grammar // Fate
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by FateParser.
type FateVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FateParser#fate.
	VisitFate(ctx *FateContext) interface{}

	// Visit a parse tree produced by FateParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by FateParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by FateParser#assignstat.
	VisitAssignstat(ctx *AssignstatContext) interface{}

	// Visit a parse tree produced by FateParser#returnstat.
	VisitReturnstat(ctx *ReturnstatContext) interface{}

	// Visit a parse tree produced by FateParser#ifstat.
	VisitIfstat(ctx *IfstatContext) interface{}

	// Visit a parse tree produced by FateParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by FateParser#valueExpr.
	VisitValueExpr(ctx *ValueExprContext) interface{}

	// Visit a parse tree produced by FateParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by FateParser#simpleBooleanExpr.
	VisitSimpleBooleanExpr(ctx *SimpleBooleanExprContext) interface{}

	// Visit a parse tree produced by FateParser#arrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by FateParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by FateParser#addSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by FateParser#varFuncExpr.
	VisitVarFuncExpr(ctx *VarFuncExprContext) interface{}

	// Visit a parse tree produced by FateParser#setBooleanExpr.
	VisitSetBooleanExpr(ctx *SetBooleanExprContext) interface{}

	// Visit a parse tree produced by FateParser#mulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by FateParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by FateParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by FateParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by FateParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by FateParser#int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by FateParser#float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by FateParser#true.
	VisitTrue(ctx *TrueContext) interface{}

	// Visit a parse tree produced by FateParser#false.
	VisitFalse(ctx *FalseContext) interface{}

	// Visit a parse tree produced by FateParser#nil.
	VisitNil(ctx *NilContext) interface{}

	// Visit a parse tree produced by FateParser#index.
	VisitIndex(ctx *IndexContext) interface{}

}