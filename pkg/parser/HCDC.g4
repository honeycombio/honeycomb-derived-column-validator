// compile with:
// antlr -Dlanguage=Go HCDC.g4

grammar HCDC ;

derived: expr EOF;

// List operators in precedence order.
expr
    : op=OP_NOT expr
    | expr op=OP_MATH_HIGH expr
    | expr op=OP_MATH_MINUS expr
    | expr op=OP_MATH_PLUS expr
    | expr op=OP_COMPARE expr
    | expr op=OP_EQUAL expr
    | expr op=OP_AND expr
    | expr op=OP_OR expr
    | '(' expr ')'
    | fun
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
    : OP_MATH_MINUS? INT
    | OP_MATH_MINUS? FLOAT
    | RAWSTRING
    | STRING
    | TRUE
    | FALSE
    | NULL
    ;

// AND and OR will be parsed as operators instead of FUNCNAME.
funcname
    : FUNCNAME
    | OP_AND
    | OP_OR
    ;

// Order here is important: prefer specific tokens over function names.
// Note this uses some unicode property identifiers, enumerated here:
// http://unicode.org/reports/tr44/#Properties
TRUE: 'true';
FALSE: 'false';
NULL: 'null';
COLUMN: '$' COLRUNE+;

OP_NOT: '!';

OP_MATH_HIGH
    : '*'
    | '/'
    | '%'
    ;

OP_MATH_PLUS : '+';
OP_MATH_MINUS : '-';

OP_COMPARE
    : '<'
    | '<='
    | '>'
    | '>='
    ;

OP_EQUAL
    : '='
    | '!='
    ;

OP_AND: 'AND';
OP_OR: 'OR';

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


INT: DIGITS;

// Support leading decimal, because the original parser did.
FLOAT
    : DIGITS ('.' DIGITS)? ([eE] [+-]? DIGITS)?
    | '.' DIGITS ([eE] [+-]? DIGITS)?
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
