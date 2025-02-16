// Code generated from HCDC.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // HCDC

import "github.com/antlr4-go/antlr/v4"

// BaseHCDCListener is a complete listener for a parse tree produced by HCDCParser.
type BaseHCDCListener struct{}

var _ HCDCListener = &BaseHCDCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHCDCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHCDCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHCDCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHCDCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDerived is called when production derived is entered.
func (s *BaseHCDCListener) EnterDerived(ctx *DerivedContext) {}

// ExitDerived is called when production derived is exited.
func (s *BaseHCDCListener) ExitDerived(ctx *DerivedContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseHCDCListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseHCDCListener) ExitExpr(ctx *ExprContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseHCDCListener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseHCDCListener) ExitFun(ctx *FunContext) {}

// EnterParams is called when production params is entered.
func (s *BaseHCDCListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseHCDCListener) ExitParams(ctx *ParamsContext) {}

// EnterColumn is called when production column is entered.
func (s *BaseHCDCListener) EnterColumn(ctx *ColumnContext) {}

// ExitColumn is called when production column is exited.
func (s *BaseHCDCListener) ExitColumn(ctx *ColumnContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseHCDCListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseHCDCListener) ExitLiteral(ctx *LiteralContext) {}

// EnterFuncname is called when production funcname is entered.
func (s *BaseHCDCListener) EnterFuncname(ctx *FuncnameContext) {}

// ExitFuncname is called when production funcname is exited.
func (s *BaseHCDCListener) ExitFuncname(ctx *FuncnameContext) {}
