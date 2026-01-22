package main

import (
	"strings"
	"testing"
)

func TestParseDerived(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty expression",
			input:   "",
			wantErr: true,
			errMsg:  "empty expression",
		},
		{
			name:    "invalid expression (mismatched parenthesis)",
			input:   "IF(1,2", // missing closing parenthesis
			wantErr: true,
		},
		{
			name:    "valid expression",
			input:   "IF(1,2)",
			wantErr: false,
		},
		{
			name:    "infix expression",
			input:   "$a + 1",
			wantErr: false,
		},
		{
			name:    "degenerate infix expression",
			input:   `$"0"-5-$"1"--5`,
			wantErr: false,
		},
		{
			name:    "and is a func and an op",
			input:   `$foo AND AND($bar AND $baz, $quux)`,
			wantErr: false,
		},
		{
			name:    "timseries function",
			input:   "MAX(LAST(1))",
			wantErr: false,
		},
		{
			name:    "nested timeseries function",
			input:   "RATE(LAST(1))",
			wantErr: true,
			errMsg:  "timeseries operations cannot be nested",
		},
		{
			// Not "nested" per se but we aren't allowed more than one TS op
			// per expression.
			name:    "multiple timeseries functions",
			input:   "COALESCE(RATE(1),RATE(2))",
			wantErr: true,
			errMsg:  "timeseries operations cannot be nested",
		},
		{
			name:    "unknown function",
			input:   "DECREASE(1)",
			wantErr: true,
			errMsg:  "invalid function: DECREASE",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ParseDerived(tc.input, false)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error for input %q; got nil", tc.input)
				} else if tc.errMsg != "" && err.Error() != tc.errMsg {
					t.Errorf("expected error message %q, got %q", tc.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error for input %q, got %v", tc.input, err)
				}
			}
		})
	}
}

func TestProcessInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		verbose     bool
		wantExpr    string
		wantJSON    bool
		wantErr     bool
		errContains string
	}{
		{
			name:     "plain text input",
			input:    "IF(1,2)",
			wantExpr: "IF(1,2)",
			wantJSON: false,
			wantErr:  false,
		},
		{
			name:     "valid JSON input",
			input:    `{"expression": "IF(1,2)"}`,
			wantExpr: "IF(1,2)",
			wantJSON: true,
			wantErr:  false,
		},
		{
			name:        "invalid JSON - missing expression",
			input:       `{"wrong_key": "IF(1,2)"}`,
			wantExpr:    "",
			wantJSON:    false,
			wantErr:     true,
			errContains: "missing 'expression' value",
		},
		{
			name:     "malformed JSON falls back to plain text",
			input:    `{"expression": "IF(1,2)" bad json`,
			wantExpr: `{"expression": "IF(1,2)" bad json`,
			wantJSON: false,
			wantErr:  false,
		},
		{
			name:     "complex valid expression",
			input:    `IF(AND(NOT(EXISTS($trace.parent_id)),EXISTS($duration_ms)),LTE($duration_ms,300))`,
			wantExpr: `IF(AND(NOT(EXISTS($trace.parent_id)),EXISTS($duration_ms)),LTE($duration_ms,300))`,
			wantJSON: false,
			wantErr:  false,
		},
		{
			name:     "valid JSON expression with a quoted string",
			input:    `{"expression": "EQUALS($request.foo, \"POST\")"}`,
			wantExpr: `EQUALS($request.foo, "POST")`,
			wantJSON: true,
			wantErr:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := &Processor{config: Config{Verbose: tc.verbose}}
			expr, isJSON, err := p.processInput(tc.input)

			if tc.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
				} else if !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("error %q does not contain %q", err.Error(), tc.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if expr != tc.wantExpr {
				t.Errorf("got expression %q, want %q", expr, tc.wantExpr)
			}
			if isJSON != tc.wantJSON {
				t.Errorf("got JSON mode %v, want %v", isJSON, tc.wantJSON)
			}
		})
	}
}
