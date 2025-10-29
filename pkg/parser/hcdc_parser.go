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
		"", "'('", "')'", "','", "'$'", "'true'", "'false'", "'null'", "", "'!'",
		"", "'+'", "'-'", "", "", "'AND'", "'OR'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "TRUE", "FALSE", "NULL", "COLUMN", "OP_NOT", "OP_MATH_HIGH",
		"OP_MATH_PLUS", "OP_MATH_MINUS", "OP_COMPARE", "OP_EQUAL", "OP_AND",
		"OP_OR", "FUNCNAME", "INT", "FLOAT", "RAWSTRING", "STRING", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"derived", "expr", "fun", "params", "column", "literal", "funcname",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 95, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 51, 8, 1, 10, 1, 12, 1, 54, 9, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 3, 2, 62, 8, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 3, 3, 69, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 76, 8,
		4, 1, 5, 3, 5, 79, 8, 5, 1, 5, 1, 5, 3, 5, 83, 8, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 91, 8, 5, 1, 6, 1, 6, 1, 6, 0, 1, 2, 7, 0, 2, 4,
		6, 8, 10, 12, 0, 1, 1, 0, 15, 17, 111, 0, 14, 1, 0, 0, 0, 2, 27, 1, 0,
		0, 0, 4, 55, 1, 0, 0, 0, 6, 65, 1, 0, 0, 0, 8, 75, 1, 0, 0, 0, 10, 90,
		1, 0, 0, 0, 12, 92, 1, 0, 0, 0, 14, 15, 3, 2, 1, 0, 15, 16, 5, 0, 0, 1,
		16, 1, 1, 0, 0, 0, 17, 18, 6, 1, -1, 0, 18, 19, 5, 9, 0, 0, 19, 28, 3,
		2, 1, 12, 20, 21, 5, 1, 0, 0, 21, 22, 3, 2, 1, 0, 22, 23, 5, 2, 0, 0, 23,
		28, 1, 0, 0, 0, 24, 28, 3, 4, 2, 0, 25, 28, 3, 8, 4, 0, 26, 28, 3, 10,
		5, 0, 27, 17, 1, 0, 0, 0, 27, 20, 1, 0, 0, 0, 27, 24, 1, 0, 0, 0, 27, 25,
		1, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28, 52, 1, 0, 0, 0, 29, 30, 10, 11, 0,
		0, 30, 31, 5, 10, 0, 0, 31, 51, 3, 2, 1, 12, 32, 33, 10, 10, 0, 0, 33,
		34, 5, 12, 0, 0, 34, 51, 3, 2, 1, 11, 35, 36, 10, 9, 0, 0, 36, 37, 5, 11,
		0, 0, 37, 51, 3, 2, 1, 10, 38, 39, 10, 8, 0, 0, 39, 40, 5, 13, 0, 0, 40,
		51, 3, 2, 1, 9, 41, 42, 10, 7, 0, 0, 42, 43, 5, 14, 0, 0, 43, 51, 3, 2,
		1, 8, 44, 45, 10, 6, 0, 0, 45, 46, 5, 15, 0, 0, 46, 51, 3, 2, 1, 7, 47,
		48, 10, 5, 0, 0, 48, 49, 5, 16, 0, 0, 49, 51, 3, 2, 1, 6, 50, 29, 1, 0,
		0, 0, 50, 32, 1, 0, 0, 0, 50, 35, 1, 0, 0, 0, 50, 38, 1, 0, 0, 0, 50, 41,
		1, 0, 0, 0, 50, 44, 1, 0, 0, 0, 50, 47, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0,
		52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 3, 1, 0, 0, 0, 54, 52, 1, 0,
		0, 0, 55, 56, 3, 12, 6, 0, 56, 61, 5, 1, 0, 0, 57, 59, 3, 6, 3, 0, 58,
		60, 5, 3, 0, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0,
		0, 61, 57, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64,
		5, 2, 0, 0, 64, 5, 1, 0, 0, 0, 65, 68, 3, 2, 1, 0, 66, 67, 5, 3, 0, 0,
		67, 69, 3, 6, 3, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 7, 1, 0,
		0, 0, 70, 76, 5, 8, 0, 0, 71, 72, 5, 4, 0, 0, 72, 76, 5, 21, 0, 0, 73,
		74, 5, 4, 0, 0, 74, 76, 5, 20, 0, 0, 75, 70, 1, 0, 0, 0, 75, 71, 1, 0,
		0, 0, 75, 73, 1, 0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 79, 5, 12, 0, 0, 78, 77,
		1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 91, 5, 18, 0, 0,
		81, 83, 5, 12, 0, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 91, 5, 19, 0, 0, 85, 91, 5, 20, 0, 0, 86, 91, 5, 21, 0, 0,
		87, 91, 5, 5, 0, 0, 88, 91, 5, 6, 0, 0, 89, 91, 5, 7, 0, 0, 90, 78, 1,
		0, 0, 0, 90, 82, 1, 0, 0, 0, 90, 85, 1, 0, 0, 0, 90, 86, 1, 0, 0, 0, 90,
		87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 11, 1, 0, 0,
		0, 92, 93, 7, 0, 0, 0, 93, 13, 1, 0, 0, 0, 10, 27, 50, 52, 59, 61, 68,
		75, 78, 82, 90,
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
	HCDCParserEOF           = antlr.TokenEOF
	HCDCParserT__0          = 1
	HCDCParserT__1          = 2
	HCDCParserT__2          = 3
	HCDCParserT__3          = 4
	HCDCParserTRUE          = 5
	HCDCParserFALSE         = 6
	HCDCParserNULL          = 7
	HCDCParserCOLUMN        = 8
	HCDCParserOP_NOT        = 9
	HCDCParserOP_MATH_HIGH  = 10
	HCDCParserOP_MATH_PLUS  = 11
	HCDCParserOP_MATH_MINUS = 12
	HCDCParserOP_COMPARE    = 13
	HCDCParserOP_EQUAL      = 14
	HCDCParserOP_AND        = 15
	HCDCParserOP_OR         = 16
	HCDCParserFUNCNAME      = 17
	HCDCParserINT           = 18
	HCDCParserFLOAT         = 19
	HCDCParserRAWSTRING     = 20
	HCDCParserSTRING        = 21
	HCDCParserWHITESPACE    = 22
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
	Expr() IExprContext
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

func (s *DerivedContext) Expr() IExprContext {
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
		p.expr(0)
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

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	OP_NOT() antlr.TerminalNode
	Fun() IFunContext
	Column() IColumnContext
	Literal() ILiteralContext
	OP_MATH_HIGH() antlr.TerminalNode
	OP_MATH_MINUS() antlr.TerminalNode
	OP_MATH_PLUS() antlr.TerminalNode
	OP_COMPARE() antlr.TerminalNode
	OP_EQUAL() antlr.TerminalNode
	OP_AND() antlr.TerminalNode
	OP_OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
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

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_NOT, 0)
}

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

func (s *ExprContext) OP_MATH_HIGH() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_MATH_HIGH, 0)
}

func (s *ExprContext) OP_MATH_MINUS() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_MATH_MINUS, 0)
}

func (s *ExprContext) OP_MATH_PLUS() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_MATH_PLUS, 0)
}

func (s *ExprContext) OP_COMPARE() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_COMPARE, 0)
}

func (s *ExprContext) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_EQUAL, 0)
}

func (s *ExprContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_AND, 0)
}

func (s *ExprContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_OR, 0)
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
	return p.expr(0)
}

func (p *HCDCParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, HCDCParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HCDCParserOP_NOT:
		{
			p.SetState(18)

			var _m = p.Match(HCDCParserOP_NOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.expr(12)
		}

	case HCDCParserT__0:
		{
			p.SetState(20)
			p.Match(HCDCParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(21)
			p.expr(0)
		}
		{
			p.SetState(22)
			p.Match(HCDCParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case HCDCParserOP_AND, HCDCParserOP_OR, HCDCParserFUNCNAME:
		{
			p.SetState(24)
			p.Fun()
		}

	case HCDCParserT__3, HCDCParserCOLUMN:
		{
			p.SetState(25)
			p.Column()
		}

	case HCDCParserTRUE, HCDCParserFALSE, HCDCParserNULL, HCDCParserOP_MATH_MINUS, HCDCParserINT, HCDCParserFLOAT, HCDCParserRAWSTRING, HCDCParserSTRING:
		{
			p.SetState(26)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, HCDCParserRULE_expr)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(30)

					var _m = p.Match(HCDCParserOP_MATH_HIGH)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(31)
					p.expr(12)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, HCDCParserRULE_expr)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(33)

					var _m = p.Match(HCDCParserOP_MATH_MINUS)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(34)
					p.expr(11)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, HCDCParserRULE_expr)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(36)

					var _m = p.Match(HCDCParserOP_MATH_PLUS)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(37)
					p.expr(10)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, HCDCParserRULE_expr)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(39)

					var _m = p.Match(HCDCParserOP_COMPARE)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(40)
					p.expr(9)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, HCDCParserRULE_expr)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

					var _m = p.Match(HCDCParserOP_EQUAL)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(43)
					p.expr(8)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, HCDCParserRULE_expr)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(45)

					var _m = p.Match(HCDCParserOP_AND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(46)
					p.expr(7)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, HCDCParserRULE_expr)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(48)

					var _m = p.Match(HCDCParserOP_OR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(49)
					p.expr(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(55)
		p.Funcname()
	}
	{
		p.SetState(56)
		p.Match(HCDCParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4166642) != 0 {
		{
			p.SetState(57)
			p.Params()
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HCDCParserT__2 {
			{
				p.SetState(58)
				p.Match(HCDCParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(63)
		p.Match(HCDCParserT__1)
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
		p.SetState(65)
		p.expr(0)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(66)
			p.Match(HCDCParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(HCDCParserCOLUMN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(HCDCParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(HCDCParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(HCDCParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
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
	OP_MATH_MINUS() antlr.TerminalNode
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

func (s *LiteralContext) OP_MATH_MINUS() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_MATH_MINUS, 0)
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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HCDCParserOP_MATH_MINUS {
			{
				p.SetState(77)
				p.Match(HCDCParserOP_MATH_MINUS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(80)
			p.Match(HCDCParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HCDCParserOP_MATH_MINUS {
			{
				p.SetState(81)
				p.Match(HCDCParserOP_MATH_MINUS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(84)
			p.Match(HCDCParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.Match(HCDCParserRAWSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(86)
			p.Match(HCDCParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(87)
			p.Match(HCDCParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(88)
			p.Match(HCDCParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(89)
			p.Match(HCDCParserNULL)
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

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCNAME() antlr.TerminalNode
	OP_AND() antlr.TerminalNode
	OP_OR() antlr.TerminalNode

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

func (s *FuncnameContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_AND, 0)
}

func (s *FuncnameContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(HCDCParserOP_OR, 0)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229376) != 0) {
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

func (p *HCDCParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *HCDCParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
