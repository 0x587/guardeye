grammar Rosmsg;

options {
    caseInsensitive = true;
}

parse
    : msgStat
    | srvStat
;

msgStat
    : (field)*
;

srvStat
    : msgStat MSG_SPLIT msgStat
;

field
    : fieldType fieldName
;

fieldType
    : type
    | arrayType
;

type
    : buildinType
    | customType
;

arrayType
    : (type OPEN_BRA CLOSE_BRA)
    | (type OPEN_BRA NUMERIC_LITERAL CLOSE_BRA)
    | (type OPEN_BRA LT_EQ NUMERIC_LITERAL CLOSE_BRA)
    | (STRING LT_EQ NUMERIC_LITERAL)
    | (STRING OPEN_BRA LT_EQ NUMERIC_LITERAL CLOSE_BRA)
    | (STRING LT_EQ NUMERIC_LITERAL OPEN_BRA CLOSE_BRA)
    | (STRING LT_EQ NUMERIC_LITERAL OPEN_BRA LT_EQ NUMERIC_LITERAL CLOSE_BRA)
;

customType
    : (IDENTIFIER) |
      (IDENTIFIER '/' customType)
;

fieldName
    : IDENTIFIER
;

buildinType
    : BOOL | BYTES | CHAR | FLOAT32 | FLOAT64 |
    INT8 | UINT8 | INT16 | UINT16 | INT32 | UINT32 |
    INT64 | UINT64 | STRING | WSTRING
;

OPEN_BRA    : '[';
CLOSE_BRA   : ']';
EQ          : '=';
LT_EQ       : '<=';

MSG_SPLIT: '---';

// buildin types
BOOL    : 'bool';
BYTES   : 'bytes';
CHAR    : 'char';
FLOAT32 : 'float32';
FLOAT64 : 'float64';
INT8    : 'int8';
UINT8   : 'uint8';
INT16   : 'int16';
UINT16  : 'uint16';
INT32   : 'int32';
UINT32  : 'uint32';
INT64   : 'int64';
UINT64  : 'uint64';
STRING  : 'string';
WSTRING : 'wstring';

// keywords

IDENTIFIER: [A-Z_\u007F-\uFFFF] [A-Z_0-9\u007F-\uFFFF]*;
NUMERIC_LITERAL: ((DIGIT+ ('.' DIGIT*)?) | ('.' DIGIT+)) ('E' [-+]? DIGIT+)? | '0x' HEX_DIGIT+;
SINGLE_LINE_COMMENT: '#' ~[\r\n\fCtx]* -> channel(HIDDEN);
SPACES: [ \u000B\t\r\n] -> channel(HIDDEN);
UNEXPECTED_CHAR: .;
fragment HEX_DIGIT : [0-9A-F];
fragment DIGIT     : [0-9];
