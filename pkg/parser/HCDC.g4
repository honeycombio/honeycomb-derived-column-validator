// compile with:
// antlr -Dlanguage=Go HCDC.g4

grammar HCDC ;

derived: fun EOF;

expr
    : fun
    | column
    | literal
    ;

// Support trailing comma, which is common in some coding standards.
fun
    : funcname '(' (params ','?)? ')'
    ;

params
    : expr (',' params)?
    ;

column
    : COLUMN
    | '$' STRING
    | '$' RAWSTRING
    ;

literal
    : INT
    | FLOAT
    | RAWSTRING
    | STRING
    | TRUE
    | FALSE
    | NULL
    ;

funcname: FUNCNAME;

// Order here is important: prefer specific tokens over function names.
// Note this uses some unicode property identifiers, enumerated here:
// http://unicode.org/reports/tr44/#Properties
TRUE: 'true';
FALSE: 'false';
NULL: 'null';
COLUMN: '$' COLRUNE+;
FUNCNAME: [a-zA-Z] [a-zA-Z0-9_]+;

fragment COLRUNE
    : [\p{Letter}]
    | [\p{Digit}]
    | '_'
    | '.'
    | '/'
    | ':'
    | '='
    | '+'
    | '?'
    | '-'
    ;


INT: '-'? DIGITS;

// Support leading decimal, because the original parser did.
FLOAT
    : '-'? DIGITS ('.' DIGITS)? ([eE] [+-]? DIGITS)?
    | '-'? '.' DIGITS ([eE] [+-]? DIGITS)?
    ;

fragment DIGITS: [0-9]+;

// String lexing is taken from the golang grammar
// https://github.com/antlr/grammars-v4/tree/master/golang
RAWSTRING: '`' ~'`'* '`';
STRING: '"' (~["\\] | ESCAPED_VALUE)* '"';

fragment ESCAPED_VALUE
    : '\\' ('u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | [abfnrtv\\'"]
           | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
           | 'x' HEX_DIGIT HEX_DIGIT)
    ;

fragment OCTAL_DIGIT
    : [0-7]
    ;

fragment HEX_DIGIT
    : [0-9a-fA-F]
    ;

WHITESPACE: [\p{White_Space}]+ -> skip;
