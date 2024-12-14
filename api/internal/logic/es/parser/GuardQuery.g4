grammar GuardQuery;

options {
    caseInsensitive = true;
}

parse
    : (select)';'
;

select
    : (
        SELECT_ resultColumn (COMMA resultColumn)*
        FROM_ sourceOrSubquery (COMMA sourceOrSubquery)*
        IN_ timeStmt
        (WHERE_ whereStmt)?
    )
//    | values_clause
;

timeStmt
    : absTimeStmt
    | relatTimeStmt
;

// yyyy-MM-dd HH:mm:ss to yyyy-MM-dd HH:mm:ss
absTimeStmt
    : STRING_LITERAL (TO_ STRING_LITERAL)?
;

relatTimeStmt
    : '-' NUMERIC_LITERAL timeUnit (TO_ '-' NUMERIC_LITERAL timeUnit)?
;

timeUnit
    : 'm'
    | 'h'
    | 'd'
;

whereStmt
    : boolExpr
;

sourceOrSubquery
    : source
//      | subquery
;

source
    : NODE_ node PROVIDER_ provider? (COMMA provider)*
;

node
    : STAR
    | (nodeId = anyName)
;

provider
    : STAR
    | (providerType = anyName) (OPEN_PAR (providerArg) (COMMA providerArg)* CLOSE_PAR)?
;

providerArg
    : STRING_LITERAL
    | NUMERIC_LITERAL
;

resultColumn
    : valueExpr ( AS_? resultAlias)?
;

resultAlias
    : IDENTIFIER
    | STRING_LITERAL
;

anyName
    : IDENTIFIER
    | STRING_LITERAL
    | OPEN_PAR anyName CLOSE_PAR
;

boolExpr
    : TRUE_
    | FALSE_
    | valueExpr ( LT | LT_EQ | GT | GT_EQ) valueExpr
    | valueExpr (
        ASSIGN
        | EQ
        | NOT_EQ1
        | NOT_EQ2
//        | IN_
//        | LIKE_
//        | GLOB_
//        | MATCH_
//        | REGEXP_
    ) valueExpr
    | boolExpr AND_ boolExpr
    | boolExpr OR_ boolExpr
    | NOT_ boolExpr
    | OPEN_PAR boolExpr CLOSE_PAR
;

valueExpr
    : NUMERIC_LITERAL
    | STRING_LITERAL
    | IDENTIFIER
    | buildinSource
    | valueExpr DOT valueExpr
    | valueExpr OPEN_BRA valueExpr CLOSE_BRA
    | buildinFunction OPEN_PAR valueExpr? (COMMA valueExpr)* CLOSE_PAR
    | OPEN_PAR valueExpr CLOSE_PAR
;

buildinSource
    : '$msg'
    | 'CURRENT_TIME'
    | 'CURRENT_DATE'
    | 'CURRENT_TIMESTAMP'
;

buildinFunction
    : 'json'
    | 'yaml'
;
//expr1
//    : literalValue
//    | anyName
//    | BIND_PARAMETER
//    | ((schema_name DOT)? table_name DOT)? column_name
//    | unary_operator expr
//    | expr PIPE2 expr
//    | expr ( STAR | DIV | MOD) expr
//    | expr ( PLUS | MINUS) expr
//    | expr ( LT2 | GT2 | AMP | PIPE) expr
//    | expr ( LT | LT_EQ | GT | GT_EQ) expr
//    | expr (
//        ASSIGN
//        | EQ
//        | NOT_EQ1
//        | NOT_EQ2
//        | IS_
//        | IS_ NOT_
//        | IS_ NOT_? DISTINCT_ FROM_
//        | IN_
//        | LIKE_
//        | GLOB_
//        | MATCH_
//        | REGEXP_
//    ) expr
//    | expr AND_ expr
//    | expr OR_ expr
//    | function_name OPEN_PAR ((DISTINCT_? expr ( COMMA expr)*) | STAR)? CLOSE_PAR filter_clause? over_clause?
//    | OPEN_PAR expr (COMMA expr)* CLOSE_PAR
//    | CAST_ OPEN_PAR expr AS_ type_name CLOSE_PAR
//    | expr COLLATE_ collation_name
//    | expr NOT_? (LIKE_ | GLOB_ | REGEXP_ | MATCH_) expr (ESCAPE_ expr)?
//    | expr ( ISNULL_ | NOTNULL_ | NOT_ NULL_)
//    | expr IS_ NOT_? expr
//    | expr NOT_? BETWEEN_ expr AND_ expr
//    | expr NOT_? IN_ (
//        OPEN_PAR (select_stmt | expr ( COMMA expr)*)? CLOSE_PAR
//        | ( schema_name DOT)? table_name
//        | (schema_name DOT)? table_function_name OPEN_PAR (expr (COMMA expr)*)? CLOSE_PAR
//    )
//    | ((NOT_)? EXISTS_)? OPEN_PAR select_stmt CLOSE_PAR
//    | CASE_ expr? (WHEN_ expr THEN_ expr)+ (ELSE_ expr)? END_
//    | raise_function
//;

literalValue
    : NUMERIC_LITERAL
    | STRING_LITERAL
;

SCOL      : ';';
DOT       : '.';
OPEN_PAR  : '(';
CLOSE_PAR : ')';
OPEN_BRA  : '[';
CLOSE_BRA : ']';
COMMA     : ',';
ASSIGN    : '=';
STAR      : '*';
PLUS      : '+';
MINUS     : '-';
//TILDE     : '~';
//PIPE2     : '||';
DIV       : '/';
//MOD       : '%';
//AMP       : '&';
//PIPE      : '|';
LT        : '<';
LT_EQ     : '<=';
GT        : '>';
GT_EQ     : '>=';
EQ        : '==';
NOT_EQ1   : '!=';
NOT_EQ2   : '<>';

// keywords

NODE_              : 'NODE';
PROVIDER_          : 'PROVIDER';

AND_               : 'AND';
OR_                : 'OR';
NOT_               : 'NOT';
IN_                : 'IN';
TO_                : 'TO';

SELECT_            : 'SELECT';
FROM_              : 'FROM';
AS_                : 'AS';
//IS_                : 'IS';
WHERE_             : 'WHERE';
LIMIT_             : 'LIMIT';
OFFSET_            : 'OFFSET';
ORDER_             : 'ORDER';
BY_                : 'BY';
ASC_               : 'ASC';
DESC_              : 'DESC';

TRUE_              : 'TRUE';
FALSE_             : 'FALSE';
//NULL_              : 'NULL';

IDENTIFIER: [A-Z_\u007F-\uFFFF] [A-Z_0-9\u007F-\uFFFF]*;
NUMERIC_LITERAL: ((DIGIT+ ('.' DIGIT*)?) | ('.' DIGIT+)) ('E' [-+]? DIGIT+)? | '0x' HEX_DIGIT+;
STRING_LITERAL: '\'' ( ~'\'' | '\'\'')* '\'';
SINGLE_LINE_COMMENT: '--' ~[\r\n]* (('\r'? '\n') | EOF) -> channel(HIDDEN);
MULTILINE_COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
SPACES: [ \u000B\t\r\n] -> channel(HIDDEN);
UNEXPECTED_CHAR: .;
fragment HEX_DIGIT : [0-9A-F];
fragment DIGIT     : [0-9];