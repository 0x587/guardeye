grammar Rossrv;

options {
    caseInsensitive = true;
}

parse : root EOF ;

root: obj;

entry
    : obj
    | list
    | string
    | integer
;

obj: typeName '(' (field)? (',' field)* ')';

field: fieldName '=' entry;

list: '[' (entry)? (',' entry)* ']';

fieldName: identifier;

typeName: identifier ('.' identifier)*;

integer : NUMERIC_LITERAL ;
string : STRING_LITERAL;
identifier : IDENTIFIER ;

NUMERIC_LITERAL: ((DIGIT+ ('.' DIGIT*)?) | ('.' DIGIT+)) ('E' [-+]? DIGIT+)?;
STRING_LITERAL: '\'' ( ~'\'' | '\'\'')* '\'';
IDENTIFIER: [A-Z_\u007F-\uFFFF] [A-Z_0-9\u007F-\uFFFF]*;
WS : [ \t\r\n]+ -> skip ;
fragment DIGIT     : [0-9];