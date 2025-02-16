// Code generated from HCDC.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // HCDC

import "github.com/antlr4-go/antlr/v4"

// HCDCListener is a complete listener for a parse tree produced by HCDCParser.
type HCDCListener interface {
	antlr.ParseTreeListener

	// EnterDerived is called when entering the derived production.
	EnterDerived(c *DerivedContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterColumn is called when entering the column production.
	EnterColumn(c *ColumnContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterFuncname is called when entering the funcname production.
	EnterFuncname(c *FuncnameContext)

	// ExitDerived is called when exiting the derived production.
	ExitDerived(c *DerivedContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitColumn is called when exiting the column production.
	ExitColumn(c *ColumnContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitFuncname is called when exiting the funcname production.
	ExitFuncname(c *FuncnameContext)
}
