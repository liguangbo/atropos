// Generated from /Users/liguangbo/Documents/GOPATH/src/github.com/liguangbo/atropos/fate/grammar/Fate.g4 by ANTLR 4.7.

package grammar // Fate
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFateVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFateVisitor) VisitFate(ctx *FateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitAssignstat(ctx *AssignstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitReturnstat(ctx *ReturnstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitIfstat(ctx *IfstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitValueExpr(ctx *ValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitSimpleBooleanExpr(ctx *SimpleBooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitArrayExpr(ctx *ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitVarFuncExpr(ctx *VarFuncExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitSetBooleanExpr(ctx *SetBooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitTrue(ctx *TrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitFalse(ctx *FalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitNil(ctx *NilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFateVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}
