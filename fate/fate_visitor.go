package fate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/liguangbo/atropos/fate/grammar"
	"github.com/oleiade/lane"
)

type needReturn struct{}

/*
	此文件实现访问者模式的语法树遍历
	需要参照语法定义逐个函数分析
*/

func NewFateVisitor(mem map[string]interface{}) *FateVisitor {
	return &FateVisitor{
		memory:     mem,
		errorStack: lane.NewStack(),
	}
}

type FateVisitor struct {
	*grammar.BaseFateVisitor
	errorStack *lane.Stack

	memory                   map[string]interface{}
	currentVariable          interface{}
	currentVariableName      string
	currentVariableFieldName string
}

func (v *FateVisitor) Mem() map[string]interface{} {
	return v.memory
}

func (v *FateVisitor) SetMem(mem map[string]interface{}) {
	v.memory = mem
}

func (v *FateVisitor) Error() error {
	err := v.errorStack.Pop()
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

// 访问者模式的接口，由于框架暂时没有提供。增加这个默认实现
func (v *FateVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// 脚本语言入口
func (v *FateVisitor) VisitFate(ctx *grammar.FateContext) interface{} {
	defer func() {
		err := recover()
		if err != nil {
			v.errorStack.Push(err)
		}
	}()
	return v.Visit(ctx.Block())
}

func (v *FateVisitor) VisitBlock(ctx *grammar.BlockContext) interface{} {
	stats := ctx.AllStat()
	for _, stat := range stats {
		ret := v.Visit(stat)
		if _, ok := ret.(*needReturn); ok {
			return ret
		}
	}
	return nil
}

func (v *FateVisitor) VisitStat(ctx *grammar.StatContext) interface{} {
	assignStat := ctx.Assignstat()
	if assignStat != nil {
		return v.Visit(assignStat)
	}
	returnStat := ctx.Returnstat()
	if returnStat != nil {
		return v.Visit(returnStat)
	}
	return v.Visit(ctx.Ifstat())
}

func (v *FateVisitor) VisitAssignstat(ctx *grammar.AssignstatContext) interface{} {
	// 先计算出右值，然后在赋值给左值
	value := v.Visit(ctx.Expr())
	v.Visit(ctx.Variable())
	v.assign(value)
	return nil
}

func (v *FateVisitor) VisitReturnstat(ctx *grammar.ReturnstatContext) interface{} {
	return &needReturn{}
}

// 赋值操作，将value赋值给当前变量
func (v *FateVisitor) assign(value interface{}) {
	switch variable := v.currentVariable.(type) {
	case nil:
		// 如果 currentVariable 是nil，在内存中创建变量
		v.memory[v.currentVariableFieldName] = value
		return
	case map[string]interface{}:
		// 如果 currentVariable map[string]interface{}，在map中创建变量
		variable[v.currentVariableFieldName] = value
	case []interface{}:
		// 如果 currentVariable []interface{}，给数组的某个index赋值，检验index合法
		index, err := strconv.Atoi(v.currentVariableFieldName)
		if err != nil {
			panic(v.currentVariableName + " is array, invalid index " + v.currentVariableFieldName)
		}
		if index >= len(variable) || index < 0 {
			panic(fmt.Sprintf("%s is array length is %d, %d index out of range", v.currentVariableName, len(variable), index))
		}
		variable[index] = value
	default:
		// currentVariable不是object，无法赋值
		panic(fmt.Sprintf("%s is %T, can not do assignment to field %s", v.currentVariableName, v.currentVariable, v.currentVariableFieldName))
	}
}

func (v *FateVisitor) VisitIfstat(ctx *grammar.IfstatContext) interface{} {
	exprs := ctx.AllExpr()
	blocks := ctx.AllBlock()
	finalElse := false
	// 如果执行块数量大于表达式数量，最后有一个else
	if len(blocks) > len(exprs) {
		finalElse = true
	} else if len(blocks) < len(exprs) {
		panic(fmt.Sprintf("ifstat has %d expr, but got %d stat", len(exprs), len(blocks)))
	}
	for i, expr := range exprs {
		if boolExpr(v.Visit(expr)) {
			return v.Visit(ctx.Block(i))
		}
	}
	if finalElse {
		return v.Visit(ctx.Block(len(blocks) - 1))
	}
	return nil
}

// 访问到括号的时候，先计算括号内的内容
func (v *FateVisitor) VisitParens(ctx *grammar.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *FateVisitor) VisitValueExpr(ctx *grammar.ValueExprContext) interface{} {
	return v.Visit(ctx.Value())
}

func (v *FateVisitor) VisitUnaryExpr(ctx *grammar.UnaryExprContext) interface{} {
	return unaryExpr(v.Visit(ctx.Expr()), ctx.GetOp().GetText())
}

// 简单的布尔值比较
func (v *FateVisitor) VisitSimpleBooleanExpr(ctx *grammar.SimpleBooleanExprContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	return compare(left, right, ctx.GetOp().GetText())
}

// 创造数组
func (v *FateVisitor) VisitArrayExpr(ctx *grammar.ArrayExprContext) interface{} {
	exprs := ctx.AllExpr()
	values := []interface{}{}
	for _, expr := range exprs {
		values = append(values, v.Visit(expr))
	}
	return values
}

// 变量以及变量的函数操作

func (v *FateVisitor) VisitOrExpr(ctx *grammar.OrExprContext) interface{} {
	return boolExpr(v.Visit(ctx.Expr(0))) || boolExpr(v.Visit(ctx.Expr(1)))
}

func (v *FateVisitor) VisitAddSub(ctx *grammar.AddSubContext) interface{} {
	return addSub(v.Visit(ctx.Expr(0)), v.Visit(ctx.Expr(1)), ctx.GetOp().GetText())
}

func (v *FateVisitor) VisitVarFuncExpr(ctx *grammar.VarFuncExprContext) interface{} {
	v.Visit(ctx.Variable())
	value := v.getVariable()
	functions := ctx.AllFUNC()
	for _, function := range functions {
		value = funcRunner(value, strings.TrimSuffix(function.GetText(), "()"))
	}
	return value
}

func (v *FateVisitor) VisitSetBooleanExpr(ctx *grammar.SetBooleanExprContext) interface{} {
	return compareSet(v.Visit(ctx.Expr(0)), v.Visit(ctx.Expr(1)), ctx.GetOp().GetText())
}

func (v *FateVisitor) VisitMulDiv(ctx *grammar.MulDivContext) interface{} {
	return mulDiv(v.Visit(ctx.Expr(0)), v.Visit(ctx.Expr(1)), ctx.GetOp().GetText())
}

func (v *FateVisitor) VisitAndExpr(ctx *grammar.AndExprContext) interface{} {
	return boolExpr(v.Visit(ctx.Expr(0))) && boolExpr(v.Visit(ctx.Expr(1)))
}

func (v *FateVisitor) getVariable() interface{} {
	switch variable := v.currentVariable.(type) {
	case nil:
		// 如果 currentVariable 是nil，返回内从中变量
		return v.memory[v.currentVariableFieldName]
	case map[string]interface{}:
		// 如果 currentVariable map[string]interface{}，返回map变量
		return variable[v.currentVariableFieldName]
	case []interface{}:
		// 如果 currentVariable []interface{}，返回数组某个index赋值，检验index合法
		index, err := strconv.Atoi(v.currentVariableFieldName)
		if err != nil {
			panic(v.currentVariableName + " is array, invalid index " + v.currentVariableFieldName)
		}
		if index >= len(variable) || index < 0 {
			panic(fmt.Sprintf("%s is array length is %d, %d index out of range", v.currentVariableName, len(variable), index))
		}
		return variable[index]
	default:
		// currentVariable不是object，无法返回值
		panic(fmt.Sprintf("%s is %T, can not do assignment to field %s", v.currentVariableName, v.currentVariable, v.currentVariableFieldName))
	}
}

func (v *FateVisitor) VisitVariable(ctx *grammar.VariableContext) interface{} {
	// 开始的时候清空变量查找数据
	v.currentVariableFieldName = ""
	v.currentVariableName = ""
	v.currentVariable = nil

	// 一个变量由多段field组成，任意一段查找变量出现错误，直接返回
	fields := ctx.AllField()
	for _, field := range fields {
		v.Visit(field)
	}
	return nil
}

func (v *FateVisitor) VisitField(ctx *grammar.FieldContext) interface{} {
	v.visitVariable(ctx.NAME().GetText())
	indexes := ctx.AllIndex()
	for _, index := range indexes {
		v.Visit(index)
	}
	return nil
}

func (v *FateVisitor) VisitString(ctx *grammar.StringContext) interface{} {
	str := ctx.STRING().GetText()
	strlen := len(str)
	if strlen <= 2 {
		return ""
	}
	// 去除单引号或者双引号
	return str[1 : strlen-1]
}

func (v *FateVisitor) VisitInt(ctx *grammar.IntContext) interface{} {
	value, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return value
}

func (v *FateVisitor) VisitFloat(ctx *grammar.FloatContext) interface{} {
	value, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return value
}

func (v *FateVisitor) VisitTrue(ctx *grammar.TrueContext) interface{} {
	return true
}

func (v *FateVisitor) VisitFalse(ctx *grammar.FalseContext) interface{} {
	return false
}

func (v *FateVisitor) VisitNil(ctx *grammar.NilContext) interface{} {
	return nil
}

func (v *FateVisitor) VisitIndex(ctx *grammar.IndexContext) interface{} {
	v.visitVariable(ctx.INT().GetText())
	return nil
}

func (v *FateVisitor) visitVariable(nextName string) {
	// 如果没有 fieldName ，证明并没有开始查找变量，赋值给 currentVariableFieldName 用户下次查找时使用
	if v.currentVariableFieldName == "" {
		v.currentVariableFieldName = nextName
		return
	}
	// 存在 currentVariableFieldName ，根据 currentVariable 的类型决定如何索引 currentVariableFieldName
	switch variable := v.currentVariable.(type) {
	case nil:
		// currentVariable 是nil，需要从内存中查找变量
		tempVariable, ok := v.memory[v.currentVariableFieldName]
		if !ok {
			// 内存中不存在 currentVariableFieldName，抛出错误。不能再nil中继续寻找 nextName
			panic(v.currentVariableFieldName + " is nil in memory, can not index " + nextName)
		}
		// 更新 currentVariable ，并记录 currentVariable 的名字
		v.currentVariable = tempVariable
		v.currentVariableName = v.currentVariableFieldName
		v.currentVariableFieldName = nextName
	case map[string]interface{}:
		// currentVariable 是map[string]interface{}，需要从这个map中索引出变量
		tempVariable, ok := variable[v.currentVariableFieldName]
		if !ok {
			// currentVariable 中不存在 currentVariableFieldName，抛出错误。不能再nil中继续寻找 nextName
			panic(fmt.Sprintf("%s.%s is nil, can not index field %s", v.currentVariableName, v.currentVariableFieldName, nextName))
		}
		// 更新 currentVariable ，并记录 currentVariable 的名字
		v.currentVariable = tempVariable
		v.currentVariableName += "." + v.currentVariableFieldName
		v.currentVariableFieldName = nextName
	case []interface{}:
		// currentVariable 是[]interface{}，需要从这个数组中索引出变量
		index, err := strconv.Atoi(v.currentVariableFieldName)
		if err != nil {
			// index转换数字的时候出现错误，抛出错误
			panic(v.currentVariableName + " is array, invalid index " + v.currentVariableFieldName + " " + err.Error())
		}
		if index >= len(variable) || index < 0 {
			// 非法的index，抛出错误
			panic(fmt.Sprintf("%s is array length is %d, %d index out of range", v.currentVariableName, len(variable), index))
		}
		if variable[index] == nil {
			// 这个index所指向的变量是nil，不能再nil中继续寻找 nextName
			panic(fmt.Sprintf("%s[%d] is nil, can not index field %s", v.currentVariableName, index, nextName))
		}
		// 更新 currentVariable ，并记录 currentVariable 的名字
		v.currentVariable = variable[index]
		v.currentVariableName += "[" + v.currentVariableFieldName + "]"
		v.currentVariableFieldName = nextName
	default:
		// currentVariable的类型为其他非object，直接抛错
		panic(fmt.Sprintf("%s is %T, can not do index to field %s", v.currentVariableName, v.currentVariable, v.currentVariableFieldName))
	}
}
