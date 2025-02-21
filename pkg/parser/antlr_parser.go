package parser

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
	"github.com/honeycombio/honeycomb-derived-column-validator/pkg/api"
)

// Debug controls whether debug logging is printed.
var Debug bool // set to true to enable debug output
var logger *log.Logger

var lexerPool sync.Pool
var parserPool sync.Pool

func init() {
	logger = log.New(os.Stderr, "", 0)

	lexerPool.New = func() any {
		return NewHCDCLexer(nil)
	}
	parserPool.New = func() any {
		return NewHCDCParser(nil)
	}
}

func ANTLRParse(expression string, debug bool) (*api.DerivedValue, error) {
	if expression == "" {
		return nil, &ParseError{
			Msg: "empty expression",
		}
	}

	if debug {
		Debug = true
	}

	lexer := lexerPool.Get().(*HCDCLexer)
	defer lexerPool.Put(lexer)
	lexer.SetInputStream(antlr.NewInputStream(expression))

	l := antlrListener{
		// colMapper:          colMapper,
		// missingColumnAsNil: missingColumnAsNil,
	}

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&l)

	p := parserPool.Get().(*HCDCParser)
	defer func() {
		// Reset the input stream so this parser doesn't retain a reference to the lexer.
		p.SetInputStream(nil)
		parserPool.Put(p)
	}()
	p.SetInputStream(antlr.NewCommonTokenStream(lexer, 0))

	p.RemoveErrorListeners()
	p.AddErrorListener(&l)

	antlr.ParseTreeWalkerDefault.Walk(&l, p.Derived())

	return l.result()
}

type ColumnMapper interface{}

type antlrListener struct {
	antlr.DefaultErrorListener
	BaseHCDCListener

	colMapper          ColumnMapper
	missingColumnAsNil bool

	vals     []*api.DerivedValue
	parseErr error
}

func (l *antlrListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol any,
	line int,
	column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.err(msg)
}

func (l *antlrListener) result() (*api.DerivedValue, error) {
	if l.parseErr != nil {
		return nil, l.parseErr
	}
	if len(l.vals) != 1 {
		return nil, errors.New("nonsensical expression")
	}
	return l.vals[0], nil
}

// A derived column is a nested series of functions, with lists of parameters.
// So, when we enter a function (the top level is always a function), push it
// onto the stack, and attach all successive expressions as parameters.
// When we exit the function, pop it off the stack and turn it into a param for
// the underlying function.
// Treat operators exactly like functions, pushing the relevant DC value onto
// the stack - just need to unify the exit function, which is the same for all.
func (l *antlrListener) EnterFun(c *FunContext) {
	name := c.Funcname().GetText()

	op := lookupOp(name)
	if op == api.DeriveOp_D_NONE {
		l.err("invalid function:", name)
	}

	l.vals = append(l.vals, &api.DerivedValue{
		Op:     op,
		Params: make([]*api.DerivedParam, 0, 4),
	})

	if Debug {
		logger.Printf("operator: %s", name)
	}
}

func (l *antlrListener) ExitFun(c *FunContext) {
	if len(l.vals) <= 1 {
		return
	}

	// pop the top val, as a param, off the stack
	param := l.vals[len(l.vals)-1]
	l.vals = l.vals[:len(l.vals)-1]

	val := l.vals[len(l.vals)-1]
	val.Params = append(val.Params, &api.DerivedParam{Value: &api.DerivedParam_Derived{Derived: param}})
}

func (l *antlrListener) ExitColumn(c *ColumnContext) {
	if len(l.vals) == 0 || l.parseErr != nil {
		return
	}

	// // Get the token and strip off the leading $
	var ident string
	var err error
	switch {
	case c.COLUMN() != nil:
		ident = c.COLUMN().GetText()[1:]
	case c.STRING() != nil:
		ident, err = strconv.Unquote(c.STRING().GetText())
	case c.RAWSTRING() != nil:
		ident = c.RAWSTRING().GetText()[1 : len(c.RAWSTRING().GetText())-1]
	default:
		panic("unhandled column token" + c.GetText())
	}
	if err != nil {
		l.err(err.Error())
	}

	if Debug {
		logger.Printf("column: %s", ident)
	}
}

func (l *antlrListener) ExitLiteral(c *LiteralContext) {
	if len(l.vals) == 0 || l.parseErr != nil {
		return
	}

	var static any
	var err error
	switch {
	case c.INT() != nil:
		static, err = strconv.ParseInt(c.INT().GetText(), 10, 64)
	case c.FLOAT() != nil:
		static, err = strconv.ParseFloat(c.FLOAT().GetText(), 64)
	case c.TRUE() != nil:
		static = true
	case c.FALSE() != nil:
		static = false
	case c.NULL() != nil:
		static = nil
	case c.STRING() != nil:
		static, err = strconv.Unquote(c.STRING().GetText())
	case c.RAWSTRING() != nil:
		static = c.RAWSTRING().GetText()[1 : len(c.RAWSTRING().GetText())-1]
	default:
		panic("unhandled literal token" + c.GetText())
	}
	if err != nil {
		l.err(err.Error())
	} else {
		if Debug {
			logger.Printf("literal: %v", static)
		}
	}
}

func (l *antlrListener) err(msg ...string) {
	if l.parseErr == nil {
		l.parseErr = &ParseError{
			Msg: strings.Join(msg, " "),
		}
	}
}
