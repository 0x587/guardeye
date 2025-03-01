grammar Rosmsg;

options {
    caseInsensitive = true;
}

parse
    : msg_stat |
      srv_stat
;

msg_stat
    : (field)*
;

srv_stat
    : msg_stat MSG_SPLIT
;

field
    : field_type field_name
;

//int32[] unbounded_integer_array
//int32[5] five_integers_array
//int32[<=5] up_to_five_integers_array
//
//string string_of_unbounded_size
//string<=10 up_to_ten_characters_string
//
//string[<=5] up_to_five_unbounded_strings
//string<=10[] unbounded_array_of_strings_up_to_ten_characters_each
//string<=10[<=5] up_to_five_strings_up_to_ten_characters_each

field_type
    : (type) |
      (type OPEN_BRA CLOSE_BRA) |
      (type OPEN_BRA NUMERIC_LITERAL CLOSE_BRA) |
      (type OPEN_BRA LT_EQ NUMERIC_LITERAL CLOSE_BRA) |
      (STRING) |
      (STRING LT_EQ NUMERIC_LITERAL) |
      (STRING OPEN_BRA LT_EQ NUMERIC_LITERAL CLOSE_BRA) |
      (STRING LT_EQ NUMERIC_LITERAL OPEN_BRA CLOSE_BRA) |
      (STRING LT_EQ NUMERIC_LITERAL OPEN_BRA LT_EQ NUMERIC_LITERAL CLOSE_BRA)
;

type
    : buildin_type
    | customed_type
;

customed_type
    : (IDENTIFIER) |
      (IDENTIFIER '/' customed_type)
;

field_name
    : IDENTIFIER
;

buildin_type
    : BOOL | BYTES | CHAR | FLOAT32 | FLOAT64 |
    INT8 | UINT8 | INT16 | UINT16 | INT32 | UINT32 |
    INT64 | UINT64 | STRING | WSTRING
;

OPEN_BRA  : '[';
CLOSE_BRA : ']';
EQ    : '=';
LT_EQ     : '<=';

MSG_SPLIT: '---';

// types
BOOL : 'bool';
BYTES : 'bytes';
CHAR : 'char';
FLOAT32 : 'float32';
FLOAT64 : 'float64';
INT8 : 'int8';
UINT8 : 'uint8';
INT16 : 'int16';
UINT16 : 'uint16';
INT32 : 'int32';
UINT32 : 'uint32';
INT64 : 'int64';
UINT64 : 'uint64';
STRING : 'string';
WSTRING : 'wstring';

// keywords

IDENTIFIER: [A-Z_\u007F-\uFFFF] [A-Z_0-9\u007F-\uFFFF]*;
NUMERIC_LITERAL: ((DIGIT+ ('.' DIGIT*)?) | ('.' DIGIT+)) ('E' [-+]? DIGIT+)? | '0x' HEX_DIGIT+;
//STRING_LITERAL: '\'' ( ~'\'' | '\'\'')* '\'';
SINGLE_LINE_COMMENT: '#' ~[\r\n\f]* -> channel(HIDDEN);
SPACES: [ \u000B\t\r\n] -> channel(HIDDEN);
UNEXPECTED_CHAR: .;
fragment HEX_DIGIT : [0-9A-F];
fragment DIGIT     : [0-9];
