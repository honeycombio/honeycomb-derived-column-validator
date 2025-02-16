// Code generated from HCDC.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // HCDC

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type HCDCParser struct {
	*antlr.BaseParser
}

var HCDCParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hcdcParserInit() {
	staticData := &HCDCParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "','", "')'", "'$'", "'true'", "'false'", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "TRUE", "FALSE", "NULL", "COLUMN", "FUNCNAME", "INT",
		"FLOAT", "RAWSTRING", "STRING", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"derived", "expr", "fun", "params", "column", "literal", "funcname",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 49, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 21,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 27, 8, 2, 3, 2, 29, 8, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 3, 3, 36, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		43, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12,
		0, 1, 2, 0, 5, 7, 10, 13, 48, 0, 14, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4,
		22, 1, 0, 0, 0, 6, 32, 1, 0, 0, 0, 8, 42, 1, 0, 0, 0, 10, 44, 1, 0, 0,
		0, 12, 46, 1, 0, 0, 0, 14, 15, 3, 4, 2, 0, 15, 16, 5, 0, 0, 1, 16, 1, 1,
		0, 0, 0, 17, 21, 3, 4, 2, 0, 18, 21, 3, 8, 4, 0, 19, 21, 3, 10, 5, 0, 20,
		17, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 19, 1, 0, 0, 0, 21, 3, 1, 0, 0,
		0, 22, 23, 3, 12, 6, 0, 23, 28, 5, 1, 0, 0, 24, 26, 3, 6, 3, 0, 25, 27,
		5, 2, 0, 0, 26, 25, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 29, 1, 0, 0, 0,
		28, 24, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 5,
		3, 0, 0, 31, 5, 1, 0, 0, 0, 32, 35, 3, 2, 1, 0, 33, 34, 5, 2, 0, 0, 34,
		36, 3, 6, 3, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 7, 1, 0, 0,
		0, 37, 43, 5, 8, 0, 0, 38, 39, 5, 4, 0, 0, 39, 43, 5, 13, 0, 0, 40, 41,
		5, 4, 0, 0, 41, 43, 5, 12, 0, 0, 42, 37, 1, 0, 0, 0, 42, 38, 1, 0, 0, 0,
		42, 40, 1, 0, 0, 0, 43, 9, 1, 0, 0, 0, 44, 45, 7, 0, 0, 0, 45, 11, 1, 0,
		0, 0, 46, 47, 5, 9, 0, 0, 47, 13, 1, 0, 0, 0, 5, 20, 26, 28, 35, 42,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// HCDCParserInit initializes any static state used to implement HCDCParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHCDCParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HCDCParserInit() {
	staticData := &HCDCParserStaticData
	staticData.once.Do(hcdcParserInit)
}

// NewHCDCParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHCDCParser(input antlr.TokenStream) *HCDCParser {
	HCDCParserInit()
	this := new(HCDCParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &HCDCParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "HCDC.g4"

	return this
}

// HCDCParser tokens.
const (
	HCDCParserEOF        = antlr.TokenEOF
	HCDCParserT__0       = 1
	HCDCParserT__1       = 2
	HCDCParserT__2       = 3
	HCDCParserT__3       = 4
	HCDCParserTRUE       = 5
	HCDCParserFALSE      = 6
	HCDCParserNULL       = 7
	HCDCParserCOLUMN     = 8
	HCDCParserFUNCNAME   = 9
	HCDCParserINT        = 10
	HCDCParserFLOAT      = 11
	HCDCParserRAWSTRING  = 12
	HCDCParserSTRING     = 13
	HCDCParserWHITESPACE = 14
)

// HCDCParser rules.
const (
	HCDCParserRULE_derived  = 0
	HCDCParserRULE_expr     = 1
	HCDCParserRULE_fun      = 2
	HCDCParserRULE_params   = 3
	HCDCParserRULE_column   = 4
	HCDCParserRULE_literal  = 5
	HCDCParserRULE_funcname = 6
)

// IDerivedContext is an interface to support dynamic dispatch.
type IDerivedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fun() IFunContext
	EOF() antlr.TerminalNode

	// IsDerivedContext differentiates from other interfaces.
	IsDerivedContext()
}

type DerivedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDerivedContext() *DerivedContext {
	var p = new(DerivedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_derived
	return p
}

func InitEmptyDerivedContext(p *DerivedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_derived
}

func (*DerivedContext) IsDerivedContext() {}

func NewDerivedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DerivedContext {
	var p = new(DerivedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCDCParserRULE_derived

	return p
}

func (s *DerivedContext) GetParser() antlr.Parser { return s.parser }

func (s *DerivedContext) Fun() IFunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *DerivedContext) EOF() antlr.TerminalNode {
	return s.GetToken(HCDCParserEOF, 0)
}

func (s *DerivedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerivedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DerivedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.EnterDerived(s)
	}
}

func (s *DerivedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.ExitDerived(s)
	}
}

func (p *HCDCParser) Derived() (localctx IDerivedContext) {
	localctx = NewDerivedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HCDCParserRULE_derived)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Fun()
	}
	{
		p.SetState(15)
		p.Match(HCDCParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fun() IFunContext
	Column() IColumnContext
	Literal() ILiteralContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCDCParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Fun() IFunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *ExprContext) Column() IColumnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnContext)
}

func (s *ExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *HCDCParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HCDCParserRULE_expr)
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HCDCParserFUNCNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.Fun()
		}

	case HCDCParserT__3, HCDCParserCOLUMN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)
			p.Column()
		}

	case HCDCParserTRUE, HCDCParserFALSE, HCDCParserNULL, HCDCParserINT, HCDCParserFLOAT, HCDCParserRAWSTRING, HCDCParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(19)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunContext is an interface to support dynamic dispatch.
type IFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Funcname() IFuncnameContext
	Params() IParamsContext

	// IsFunContext differentiates from other interfaces.
	IsFunContext()
}

type FunContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunContext() *FunContext {
	var p = new(FunContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_fun
	return p
}

func InitEmptyFunContext(p *FunContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_fun
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCDCParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) Funcname() IFuncnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *FunContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.ExitFun(s)
	}
}

func (p *HCDCParser) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HCDCParserRULE_fun)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Funcname()
	}
	{
		p.SetState(23)
		p.Match(HCDCParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16368) != 0 {
		{
			p.SetState(24)
			p.Params()
		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HCDCParserT__1 {
			{
				p.SetState(25)
				p.Match(HCDCParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(30)
		p.Match(HCDCParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Params() IParamsContext

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCDCParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParamsContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *HCDCParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HCDCParserRULE_params)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Expr()
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(33)
			p.Match(HCDCParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Params()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnContext is an interface to support dynamic dispatch.
type IColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLUMN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RAWSTRING() antlr.TerminalNode

	// IsColumnContext differentiates from other interfaces.
	IsColumnContext()
}

type ColumnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnContext() *ColumnContext {
	var p = new(ColumnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_column
	return p
}

func InitEmptyColumnContext(p *ColumnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_column
}

func (*ColumnContext) IsColumnContext() {}

func NewColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnContext {
	var p = new(ColumnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCDCParserRULE_column

	return p
}

func (s *ColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(HCDCParserCOLUMN, 0)
}

func (s *ColumnContext) STRING() antlr.TerminalNode {
	return s.GetToken(HCDCParserSTRING, 0)
}

func (s *ColumnContext) RAWSTRING() antlr.TerminalNode {
	return s.GetToken(HCDCParserRAWSTRING, 0)
}

func (s *ColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.EnterColumn(s)
	}
}

func (s *ColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.ExitColumn(s)
	}
}

func (p *HCDCParser) Column() (localctx IColumnContext) {
	localctx = NewColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HCDCParserRULE_column)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(HCDCParserCOLUMN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(HCDCParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(HCDCParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Match(HCDCParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(HCDCParserRAWSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	RAWSTRING() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NULL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCDCParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(HCDCParserINT, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(HCDCParserFLOAT, 0)
}

func (s *LiteralContext) RAWSTRING() antlr.TerminalNode {
	return s.GetToken(HCDCParserRAWSTRING, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(HCDCParserSTRING, 0)
}

func (s *LiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(HCDCParserTRUE, 0)
}

func (s *LiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(HCDCParserFALSE, 0)
}

func (s *LiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(HCDCParserNULL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *HCDCParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HCDCParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15584) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCNAME() antlr.TerminalNode

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_funcname
	return p
}

func InitEmptyFuncnameContext(p *FuncnameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HCDCParserRULE_funcname
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HCDCParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) FUNCNAME() antlr.TerminalNode {
	return s.GetToken(HCDCParserFUNCNAME, 0)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HCDCListener); ok {
		listenerT.ExitFuncname(s)
	}
}

func (p *HCDCParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HCDCParserRULE_funcname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(HCDCParserFUNCNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
