package parser

import (
	"github.com/honeycombio/honeycomb-derived-column-validator/pkg/api"
)

type ParseError struct {
	Msg string
}

func (e *ParseError) Error() string {
	return e.Msg
}

func (e *ParseError) Status() int {
	return 400
}

func (e *ParseError) CustomerError() string {
	return e.Msg
}

// Canonicalize takes a derived column expression and returns the canonicalized version
// of it (removes all whitespace, unnecessary characters)
// func Canonicalize(expression string) (string, error) {
// 	var l antlrListener

// 	lexer := lexerPool.Get().(*HCDCLexer)
// 	defer lexerPool.Put(lexer)

// 	lexer.SetInputStream(antlr.NewInputStream(expression))
// 	lexer.RemoveErrorListeners()
// 	lexer.AddErrorListener(&l)

// 	var strBuilder strings.Builder
// 	for tok := lexer.NextToken(); tok.GetTokenType() != antlr.TokenEOF; tok = lexer.NextToken() {
// 		strBuilder.WriteString(tok.GetText())
// 	}
// 	if l.parseErr != nil {
// 		// On error, return the original expression, having failed to canonicalize.
// 		return expression, l.parseErr
// 	}
// 	return strBuilder.String(), nil
// }

func lookupOp(str string) api.DeriveOp {
	switch str {
	case "IF":
		return api.DeriveOp_D_IF
	case "COALESCE":
		return api.DeriveOp_D_COALESCE
	case "NOT", "!":
		return api.DeriveOp_D_NOT
	case "LT", "<":
		return api.DeriveOp_D_LT
	case "LTE", "<=":
		return api.DeriveOp_D_LTE
	case "GT", ">":
		return api.DeriveOp_D_GT
	case "GTE", ">=":
		return api.DeriveOp_D_GTE
	case "EQUALS", "=":
		return api.DeriveOp_D_IN
	case "!=":
		return api.DeriveOp_D_NOT_IN
	case "IN":
		return api.DeriveOp_D_IN
	case "AND":
		return api.DeriveOp_D_AND
	case "OR":
		return api.DeriveOp_D_OR
	case "MIN":
		return api.DeriveOp_D_MIN
	case "MAX":
		return api.DeriveOp_D_MAX
	case "CONCAT":
		return api.DeriveOp_D_CONCAT
	case "EXISTS":
		return api.DeriveOp_D_HAS_VALUE
	case "STARTS_WITH":
		return api.DeriveOp_D_PREFIX
	case "ENDS_WITH":
		return api.DeriveOp_D_SUFFIX
	case "SUM", "+":
		return api.DeriveOp_D_SUM
	case "SUB", "-":
		return api.DeriveOp_D_SUB
	case "MUL", "*":
		return api.DeriveOp_D_MUL
	case "DIV", "/":
		return api.DeriveOp_D_DIV
	case "INT":
		return api.DeriveOp_D_INT
	case "LOG10":
		return api.DeriveOp_D_LOG10
	case "CONTAINS":
		return api.DeriveOp_D_CONTAINS
	case "REG_MATCH":
		return api.DeriveOp_D_REG_MATCH
	case "REG_VALUE":
		return api.DeriveOp_D_REG_VALUE
	case "REG_COUNT":
		return api.DeriveOp_D_REG_COUNT
	case "UNIX_TIMESTAMP":
		return api.DeriveOp_D_UNIX_TIMESTAMP
	case "FLOAT":
		return api.DeriveOp_D_FLOAT
	case "BOOL":
		return api.DeriveOp_D_BOOL
	case "STRING":
		return api.DeriveOp_D_STRING
	case "BUCKET":
		return api.DeriveOp_D_BUCKET
	case "EVENT_TIMESTAMP":
		return api.DeriveOp_D_EVENT_TIMESTAMP
	case "INGEST_TIMESTAMP":
		return api.DeriveOp_D_INGEST_TIMESTAMP
	case "MOD", "%":
		return api.DeriveOp_D_MOD
	case "FORMAT_TIME":
		return api.DeriveOp_D_FORMAT_TIME
	case "LENGTH":
		return api.DeriveOp_D_LENGTH
	case "SWITCH":
		return api.DeriveOp_D_SWITCH
	case "TO_LOWER":
		return api.DeriveOp_D_TO_LOWER
	default:
		return api.DeriveOp_D_NONE
	}
}

type TimeseriesOp int

const (
	TimeseriesOpNone TimeseriesOp = iota
	TimeseriesOpRawMetric
	TimeseriesOpLast
	TimeseriesOpSummarize
	TimeseriesOpIncrease
	TimeseriesOpRate

	TimeseriesOps_Len // I am a placeholder to allow for iterating over all ops
	// DO NOT ADD ANY OPS BELOW THE PLACEHOLDER
)

func lookupTimeseriesOp(str string) TimeseriesOp {
	switch str {
	case "RAW_METRIC":
		return TimeseriesOpRawMetric
	case "LAST":
		return TimeseriesOpLast
	case "RATE":
		return TimeseriesOpRate
	case "INCREASE":
		return TimeseriesOpIncrease
	case "SUMMARIZE":
		return TimeseriesOpSummarize
	default:
		return TimeseriesOpNone
	}
}
