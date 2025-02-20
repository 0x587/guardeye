// Code generated from GuardQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GuardQuery
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

type GuardQueryParser struct {
	*antlr.BaseParser
}

var GuardQueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func guardqueryParserInit() {
	staticData := &GuardQueryParserStaticData
	staticData.LiteralNames = []string{
		"", "'m'", "'h'", "'d'", "'avg'", "'$msg'", "'json'", "'yaml'", "';'",
		"'.'", "'('", "')'", "'['", "']'", "','", "'='", "'*'", "'+'", "'-'",
		"'/'", "'<'", "'<='", "'>'", "'>='", "'=='", "'!='", "'<>'", "'NODE'",
		"'PROVIDER'", "'AND'", "'OR'", "'NOT'", "'IN'", "'TO'", "'SELECT'",
		"'FROM'", "'AS'", "'WHERE'", "'LIMIT'", "'OFFSET'", "'ORDER'", "'BY'",
		"'WITH'", "'WINDOW'", "'ASC'", "'DESC'", "'TRUE'", "'FALSE'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "SCOL", "DOT", "OPEN_PAR", "CLOSE_PAR",
		"OPEN_BRA", "CLOSE_BRA", "COMMA", "ASSIGN", "STAR", "PLUS", "MINUS",
		"DIV", "LT", "LT_EQ", "GT", "GT_EQ", "EQ", "NOT_EQ1", "NOT_EQ2", "NODE_",
		"PROVIDER_", "AND_", "OR_", "NOT_", "IN_", "TO_", "SELECT_", "FROM_",
		"AS_", "WHERE_", "LIMIT_", "OFFSET_", "ORDER_", "BY_", "WITH_", "WINDOW_",
		"ASC_", "DESC_", "TRUE_", "FALSE_", "IDENTIFIER", "NUMERIC_LITERAL",
		"STRING_LITERAL", "SINGLE_LINE_COMMENT", "MULTILINE_COMMENT", "SPACES",
		"UNEXPECTED_CHAR",
	}
	staticData.RuleNames = []string{
		"parse", "select", "windowSizeStmt", "timeStmt", "absTimeStmt", "relatTimeStmt",
		"timeUnit", "whereStmt", "sourceOrSubquery", "source", "node", "provider",
		"providerArg", "resultColumn", "windowFunction", "resultAlias", "anyName",
		"boolExpr", "valueExpr", "buildinSource", "buildinFunction", "literalValue",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 238, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 52, 8, 1, 10,
		1, 12, 1, 55, 9, 1, 1, 1, 1, 1, 3, 1, 59, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 65, 8, 1, 10, 1, 12, 1, 68, 9, 1, 1, 1, 1, 1, 3, 1, 72, 8, 1, 1,
		1, 1, 1, 3, 1, 76, 8, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 83, 8, 3,
		1, 4, 1, 4, 1, 4, 3, 4, 88, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 97, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 3, 9, 109, 8, 9, 1, 9, 1, 9, 5, 9, 113, 8, 9, 10, 9, 12, 9, 116,
		9, 9, 1, 10, 1, 10, 3, 10, 120, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 5, 11, 128, 8, 11, 10, 11, 12, 11, 131, 9, 11, 1, 11, 1, 11, 3,
		11, 135, 8, 11, 3, 11, 137, 8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3,
		13, 144, 8, 13, 1, 13, 3, 13, 147, 8, 13, 1, 13, 3, 13, 150, 8, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16,
		162, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 181,
		8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 189, 8, 17, 10,
		17, 12, 17, 192, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 3, 18, 202, 8, 18, 1, 18, 1, 18, 5, 18, 206, 8, 18, 10, 18, 12,
		18, 209, 9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 217, 8,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 227,
		8, 18, 10, 18, 12, 18, 230, 9, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 0, 2, 34, 36, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 0, 6, 1, 0, 1, 3, 1, 0, 49, 50,
		2, 0, 48, 48, 50, 50, 1, 0, 20, 23, 2, 0, 15, 15, 24, 26, 1, 0, 6, 7, 250,
		0, 44, 1, 0, 0, 0, 2, 47, 1, 0, 0, 0, 4, 77, 1, 0, 0, 0, 6, 82, 1, 0, 0,
		0, 8, 84, 1, 0, 0, 0, 10, 89, 1, 0, 0, 0, 12, 98, 1, 0, 0, 0, 14, 100,
		1, 0, 0, 0, 16, 102, 1, 0, 0, 0, 18, 104, 1, 0, 0, 0, 20, 119, 1, 0, 0,
		0, 22, 136, 1, 0, 0, 0, 24, 138, 1, 0, 0, 0, 26, 140, 1, 0, 0, 0, 28, 151,
		1, 0, 0, 0, 30, 153, 1, 0, 0, 0, 32, 161, 1, 0, 0, 0, 34, 180, 1, 0, 0,
		0, 36, 216, 1, 0, 0, 0, 38, 231, 1, 0, 0, 0, 40, 233, 1, 0, 0, 0, 42, 235,
		1, 0, 0, 0, 44, 45, 3, 2, 1, 0, 45, 46, 5, 8, 0, 0, 46, 1, 1, 0, 0, 0,
		47, 48, 5, 34, 0, 0, 48, 53, 3, 26, 13, 0, 49, 50, 5, 14, 0, 0, 50, 52,
		3, 26, 13, 0, 51, 49, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0,
		0, 53, 54, 1, 0, 0, 0, 54, 58, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57,
		5, 43, 0, 0, 57, 59, 3, 4, 2, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 61, 5, 35, 0, 0, 61, 66, 3, 16, 8, 0, 62, 63, 5,
		14, 0, 0, 63, 65, 3, 16, 8, 0, 64, 62, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0,
		66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 71, 1, 0, 0, 0, 68, 66, 1,
		0, 0, 0, 69, 70, 5, 32, 0, 0, 70, 72, 3, 6, 3, 0, 71, 69, 1, 0, 0, 0, 71,
		72, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 74, 5, 37, 0, 0, 74, 76, 3, 14,
		7, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 3, 1, 0, 0, 0, 77, 78,
		5, 49, 0, 0, 78, 79, 3, 12, 6, 0, 79, 5, 1, 0, 0, 0, 80, 83, 3, 8, 4, 0,
		81, 83, 3, 10, 5, 0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 7, 1,
		0, 0, 0, 84, 87, 5, 50, 0, 0, 85, 86, 5, 33, 0, 0, 86, 88, 5, 50, 0, 0,
		87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 9, 1, 0, 0, 0, 89, 90, 5, 18,
		0, 0, 90, 91, 5, 49, 0, 0, 91, 96, 3, 12, 6, 0, 92, 93, 5, 33, 0, 0, 93,
		94, 5, 18, 0, 0, 94, 95, 5, 49, 0, 0, 95, 97, 3, 12, 6, 0, 96, 92, 1, 0,
		0, 0, 96, 97, 1, 0, 0, 0, 97, 11, 1, 0, 0, 0, 98, 99, 7, 0, 0, 0, 99, 13,
		1, 0, 0, 0, 100, 101, 3, 34, 17, 0, 101, 15, 1, 0, 0, 0, 102, 103, 3, 18,
		9, 0, 103, 17, 1, 0, 0, 0, 104, 105, 5, 27, 0, 0, 105, 106, 3, 20, 10,
		0, 106, 108, 5, 28, 0, 0, 107, 109, 3, 22, 11, 0, 108, 107, 1, 0, 0, 0,
		108, 109, 1, 0, 0, 0, 109, 114, 1, 0, 0, 0, 110, 111, 5, 14, 0, 0, 111,
		113, 3, 22, 11, 0, 112, 110, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 19, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 117, 120, 5, 16, 0, 0, 118, 120, 3, 32, 16, 0, 119, 117, 1, 0, 0,
		0, 119, 118, 1, 0, 0, 0, 120, 21, 1, 0, 0, 0, 121, 137, 5, 16, 0, 0, 122,
		134, 3, 32, 16, 0, 123, 124, 5, 10, 0, 0, 124, 129, 3, 24, 12, 0, 125,
		126, 5, 14, 0, 0, 126, 128, 3, 24, 12, 0, 127, 125, 1, 0, 0, 0, 128, 131,
		1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 1, 0,
		0, 0, 131, 129, 1, 0, 0, 0, 132, 133, 5, 11, 0, 0, 133, 135, 1, 0, 0, 0,
		134, 123, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 1, 0, 0, 0, 136,
		121, 1, 0, 0, 0, 136, 122, 1, 0, 0, 0, 137, 23, 1, 0, 0, 0, 138, 139, 7,
		1, 0, 0, 139, 25, 1, 0, 0, 0, 140, 143, 3, 36, 18, 0, 141, 142, 5, 42,
		0, 0, 142, 144, 3, 28, 14, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0,
		0, 144, 149, 1, 0, 0, 0, 145, 147, 5, 36, 0, 0, 146, 145, 1, 0, 0, 0, 146,
		147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 150, 3, 30, 15, 0, 149, 146,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 27, 1, 0, 0, 0, 151, 152, 5, 4,
		0, 0, 152, 29, 1, 0, 0, 0, 153, 154, 7, 2, 0, 0, 154, 31, 1, 0, 0, 0, 155,
		162, 5, 48, 0, 0, 156, 162, 5, 50, 0, 0, 157, 158, 5, 10, 0, 0, 158, 159,
		3, 32, 16, 0, 159, 160, 5, 11, 0, 0, 160, 162, 1, 0, 0, 0, 161, 155, 1,
		0, 0, 0, 161, 156, 1, 0, 0, 0, 161, 157, 1, 0, 0, 0, 162, 33, 1, 0, 0,
		0, 163, 164, 6, 17, -1, 0, 164, 181, 5, 46, 0, 0, 165, 181, 5, 47, 0, 0,
		166, 167, 3, 36, 18, 0, 167, 168, 7, 3, 0, 0, 168, 169, 3, 36, 18, 0, 169,
		181, 1, 0, 0, 0, 170, 171, 3, 36, 18, 0, 171, 172, 7, 4, 0, 0, 172, 173,
		3, 36, 18, 0, 173, 181, 1, 0, 0, 0, 174, 175, 5, 31, 0, 0, 175, 181, 3,
		34, 17, 2, 176, 177, 5, 10, 0, 0, 177, 178, 3, 34, 17, 0, 178, 179, 5,
		11, 0, 0, 179, 181, 1, 0, 0, 0, 180, 163, 1, 0, 0, 0, 180, 165, 1, 0, 0,
		0, 180, 166, 1, 0, 0, 0, 180, 170, 1, 0, 0, 0, 180, 174, 1, 0, 0, 0, 180,
		176, 1, 0, 0, 0, 181, 190, 1, 0, 0, 0, 182, 183, 10, 4, 0, 0, 183, 184,
		5, 29, 0, 0, 184, 189, 3, 34, 17, 5, 185, 186, 10, 3, 0, 0, 186, 187, 5,
		30, 0, 0, 187, 189, 3, 34, 17, 4, 188, 182, 1, 0, 0, 0, 188, 185, 1, 0,
		0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0,
		191, 35, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 6, 18, -1, 0, 194,
		217, 5, 49, 0, 0, 195, 217, 5, 50, 0, 0, 196, 217, 5, 48, 0, 0, 197, 217,
		3, 38, 19, 0, 198, 199, 3, 40, 20, 0, 199, 201, 5, 10, 0, 0, 200, 202,
		3, 36, 18, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 207, 1,
		0, 0, 0, 203, 204, 5, 14, 0, 0, 204, 206, 3, 36, 18, 0, 205, 203, 1, 0,
		0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0,
		208, 210, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 211, 5, 11, 0, 0, 211,
		217, 1, 0, 0, 0, 212, 213, 5, 10, 0, 0, 213, 214, 3, 36, 18, 0, 214, 215,
		5, 11, 0, 0, 215, 217, 1, 0, 0, 0, 216, 193, 1, 0, 0, 0, 216, 195, 1, 0,
		0, 0, 216, 196, 1, 0, 0, 0, 216, 197, 1, 0, 0, 0, 216, 198, 1, 0, 0, 0,
		216, 212, 1, 0, 0, 0, 217, 228, 1, 0, 0, 0, 218, 219, 10, 4, 0, 0, 219,
		220, 5, 9, 0, 0, 220, 227, 3, 36, 18, 5, 221, 222, 10, 3, 0, 0, 222, 223,
		5, 12, 0, 0, 223, 224, 3, 36, 18, 0, 224, 225, 5, 13, 0, 0, 225, 227, 1,
		0, 0, 0, 226, 218, 1, 0, 0, 0, 226, 221, 1, 0, 0, 0, 227, 230, 1, 0, 0,
		0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 37, 1, 0, 0, 0, 230,
		228, 1, 0, 0, 0, 231, 232, 5, 5, 0, 0, 232, 39, 1, 0, 0, 0, 233, 234, 7,
		5, 0, 0, 234, 41, 1, 0, 0, 0, 235, 236, 7, 1, 0, 0, 236, 43, 1, 0, 0, 0,
		26, 53, 58, 66, 71, 75, 82, 87, 96, 108, 114, 119, 129, 134, 136, 143,
		146, 149, 161, 180, 188, 190, 201, 207, 216, 226, 228,
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

// GuardQueryParserInit initializes any static state used to implement GuardQueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGuardQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GuardQueryParserInit() {
	staticData := &GuardQueryParserStaticData
	staticData.once.Do(guardqueryParserInit)
}

// NewGuardQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGuardQueryParser(input antlr.TokenStream) *GuardQueryParser {
	GuardQueryParserInit()
	this := new(GuardQueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GuardQueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GuardQuery.g4"

	return this
}

// GuardQueryParser tokens.
const (
	GuardQueryParserEOF                 = antlr.TokenEOF
	GuardQueryParserT__0                = 1
	GuardQueryParserT__1                = 2
	GuardQueryParserT__2                = 3
	GuardQueryParserT__3                = 4
	GuardQueryParserT__4                = 5
	GuardQueryParserT__5                = 6
	GuardQueryParserT__6                = 7
	GuardQueryParserSCOL                = 8
	GuardQueryParserDOT                 = 9
	GuardQueryParserOPEN_PAR            = 10
	GuardQueryParserCLOSE_PAR           = 11
	GuardQueryParserOPEN_BRA            = 12
	GuardQueryParserCLOSE_BRA           = 13
	GuardQueryParserCOMMA               = 14
	GuardQueryParserASSIGN              = 15
	GuardQueryParserSTAR                = 16
	GuardQueryParserPLUS                = 17
	GuardQueryParserMINUS               = 18
	GuardQueryParserDIV                 = 19
	GuardQueryParserLT                  = 20
	GuardQueryParserLT_EQ               = 21
	GuardQueryParserGT                  = 22
	GuardQueryParserGT_EQ               = 23
	GuardQueryParserEQ                  = 24
	GuardQueryParserNOT_EQ1             = 25
	GuardQueryParserNOT_EQ2             = 26
	GuardQueryParserNODE_               = 27
	GuardQueryParserPROVIDER_           = 28
	GuardQueryParserAND_                = 29
	GuardQueryParserOR_                 = 30
	GuardQueryParserNOT_                = 31
	GuardQueryParserIN_                 = 32
	GuardQueryParserTO_                 = 33
	GuardQueryParserSELECT_             = 34
	GuardQueryParserFROM_               = 35
	GuardQueryParserAS_                 = 36
	GuardQueryParserWHERE_              = 37
	GuardQueryParserLIMIT_              = 38
	GuardQueryParserOFFSET_             = 39
	GuardQueryParserORDER_              = 40
	GuardQueryParserBY_                 = 41
	GuardQueryParserWITH_               = 42
	GuardQueryParserWINDOW_             = 43
	GuardQueryParserASC_                = 44
	GuardQueryParserDESC_               = 45
	GuardQueryParserTRUE_               = 46
	GuardQueryParserFALSE_              = 47
	GuardQueryParserIDENTIFIER          = 48
	GuardQueryParserNUMERIC_LITERAL     = 49
	GuardQueryParserSTRING_LITERAL      = 50
	GuardQueryParserSINGLE_LINE_COMMENT = 51
	GuardQueryParserMULTILINE_COMMENT   = 52
	GuardQueryParserSPACES              = 53
	GuardQueryParserUNEXPECTED_CHAR     = 54
)

// GuardQueryParser rules.
const (
	GuardQueryParserRULE_parse            = 0
	GuardQueryParserRULE_select           = 1
	GuardQueryParserRULE_windowSizeStmt   = 2
	GuardQueryParserRULE_timeStmt         = 3
	GuardQueryParserRULE_absTimeStmt      = 4
	GuardQueryParserRULE_relatTimeStmt    = 5
	GuardQueryParserRULE_timeUnit         = 6
	GuardQueryParserRULE_whereStmt        = 7
	GuardQueryParserRULE_sourceOrSubquery = 8
	GuardQueryParserRULE_source           = 9
	GuardQueryParserRULE_node             = 10
	GuardQueryParserRULE_provider         = 11
	GuardQueryParserRULE_providerArg      = 12
	GuardQueryParserRULE_resultColumn     = 13
	GuardQueryParserRULE_windowFunction   = 14
	GuardQueryParserRULE_resultAlias      = 15
	GuardQueryParserRULE_anyName          = 16
	GuardQueryParserRULE_boolExpr         = 17
	GuardQueryParserRULE_valueExpr        = 18
	GuardQueryParserRULE_buildinSource    = 19
	GuardQueryParserRULE_buildinFunction  = 20
	GuardQueryParserRULE_literalValue     = 21
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SCOL() antlr.TerminalNode
	Select_() ISelectContext

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) SCOL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSCOL, 0)
}

func (s *ParseContext) Select_() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *GuardQueryParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GuardQueryParserRULE_parse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Select_()
	}

	{
		p.SetState(45)
		p.Match(GuardQueryParserSCOL)
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

// ISelectContext is an interface to support dynamic dispatch.
type ISelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT_() antlr.TerminalNode
	AllResultColumn() []IResultColumnContext
	ResultColumn(i int) IResultColumnContext
	FROM_() antlr.TerminalNode
	AllSourceOrSubquery() []ISourceOrSubqueryContext
	SourceOrSubquery(i int) ISourceOrSubqueryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	WINDOW_() antlr.TerminalNode
	WindowSizeStmt() IWindowSizeStmtContext
	IN_() antlr.TerminalNode
	TimeStmt() ITimeStmtContext
	WHERE_() antlr.TerminalNode
	WhereStmt() IWhereStmtContext

	// IsSelectContext differentiates from other interfaces.
	IsSelectContext()
}

type SelectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectContext() *SelectContext {
	var p = new(SelectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_select
	return p
}

func InitEmptySelectContext(p *SelectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_select
}

func (*SelectContext) IsSelectContext() {}

func NewSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectContext {
	var p = new(SelectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_select

	return p
}

func (s *SelectContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectContext) SELECT_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSELECT_, 0)
}

func (s *SelectContext) AllResultColumn() []IResultColumnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResultColumnContext); ok {
			len++
		}
	}

	tst := make([]IResultColumnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResultColumnContext); ok {
			tst[i] = t.(IResultColumnContext)
			i++
		}
	}

	return tst
}

func (s *SelectContext) ResultColumn(i int) IResultColumnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultColumnContext); ok {
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

	return t.(IResultColumnContext)
}

func (s *SelectContext) FROM_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserFROM_, 0)
}

func (s *SelectContext) AllSourceOrSubquery() []ISourceOrSubqueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISourceOrSubqueryContext); ok {
			len++
		}
	}

	tst := make([]ISourceOrSubqueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISourceOrSubqueryContext); ok {
			tst[i] = t.(ISourceOrSubqueryContext)
			i++
		}
	}

	return tst
}

func (s *SelectContext) SourceOrSubquery(i int) ISourceOrSubqueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceOrSubqueryContext); ok {
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

	return t.(ISourceOrSubqueryContext)
}

func (s *SelectContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GuardQueryParserCOMMA)
}

func (s *SelectContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCOMMA, i)
}

func (s *SelectContext) WINDOW_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserWINDOW_, 0)
}

func (s *SelectContext) WindowSizeStmt() IWindowSizeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowSizeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindowSizeStmtContext)
}

func (s *SelectContext) IN_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserIN_, 0)
}

func (s *SelectContext) TimeStmt() ITimeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeStmtContext)
}

func (s *SelectContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserWHERE_, 0)
}

func (s *SelectContext) WhereStmt() IWhereStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereStmtContext)
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitSelect(s)
	}
}

func (p *GuardQueryParser) Select_() (localctx ISelectContext) {
	localctx = NewSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GuardQueryParserRULE_select)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(GuardQueryParserSELECT_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.ResultColumn()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GuardQueryParserCOMMA {
		{
			p.SetState(49)
			p.Match(GuardQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.ResultColumn()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GuardQueryParserWINDOW_ {
		{
			p.SetState(56)
			p.Match(GuardQueryParserWINDOW_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.WindowSizeStmt()
		}

	}
	{
		p.SetState(60)
		p.Match(GuardQueryParserFROM_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.SourceOrSubquery()
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GuardQueryParserCOMMA {
		{
			p.SetState(62)
			p.Match(GuardQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.SourceOrSubquery()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GuardQueryParserIN_ {
		{
			p.SetState(69)
			p.Match(GuardQueryParserIN_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.TimeStmt()
		}

	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GuardQueryParserWHERE_ {
		{
			p.SetState(73)
			p.Match(GuardQueryParserWHERE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.WhereStmt()
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

// IWindowSizeStmtContext is an interface to support dynamic dispatch.
type IWindowSizeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode
	TimeUnit() ITimeUnitContext

	// IsWindowSizeStmtContext differentiates from other interfaces.
	IsWindowSizeStmtContext()
}

type WindowSizeStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowSizeStmtContext() *WindowSizeStmtContext {
	var p = new(WindowSizeStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_windowSizeStmt
	return p
}

func InitEmptyWindowSizeStmtContext(p *WindowSizeStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_windowSizeStmt
}

func (*WindowSizeStmtContext) IsWindowSizeStmtContext() {}

func NewWindowSizeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowSizeStmtContext {
	var p = new(WindowSizeStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_windowSizeStmt

	return p
}

func (s *WindowSizeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowSizeStmtContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNUMERIC_LITERAL, 0)
}

func (s *WindowSizeStmtContext) TimeUnit() ITimeUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeUnitContext)
}

func (s *WindowSizeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowSizeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowSizeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterWindowSizeStmt(s)
	}
}

func (s *WindowSizeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitWindowSizeStmt(s)
	}
}

func (p *GuardQueryParser) WindowSizeStmt() (localctx IWindowSizeStmtContext) {
	localctx = NewWindowSizeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GuardQueryParserRULE_windowSizeStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(GuardQueryParserNUMERIC_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.TimeUnit()
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

// ITimeStmtContext is an interface to support dynamic dispatch.
type ITimeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AbsTimeStmt() IAbsTimeStmtContext
	RelatTimeStmt() IRelatTimeStmtContext

	// IsTimeStmtContext differentiates from other interfaces.
	IsTimeStmtContext()
}

type TimeStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeStmtContext() *TimeStmtContext {
	var p = new(TimeStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_timeStmt
	return p
}

func InitEmptyTimeStmtContext(p *TimeStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_timeStmt
}

func (*TimeStmtContext) IsTimeStmtContext() {}

func NewTimeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeStmtContext {
	var p = new(TimeStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_timeStmt

	return p
}

func (s *TimeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeStmtContext) AbsTimeStmt() IAbsTimeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsTimeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsTimeStmtContext)
}

func (s *TimeStmtContext) RelatTimeStmt() IRelatTimeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelatTimeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelatTimeStmtContext)
}

func (s *TimeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterTimeStmt(s)
	}
}

func (s *TimeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitTimeStmt(s)
	}
}

func (p *GuardQueryParser) TimeStmt() (localctx ITimeStmtContext) {
	localctx = NewTimeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GuardQueryParserRULE_timeStmt)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.AbsTimeStmt()
		}

	case GuardQueryParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.RelatTimeStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IAbsTimeStmtContext is an interface to support dynamic dispatch.
type IAbsTimeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING_LITERAL() []antlr.TerminalNode
	STRING_LITERAL(i int) antlr.TerminalNode
	TO_() antlr.TerminalNode

	// IsAbsTimeStmtContext differentiates from other interfaces.
	IsAbsTimeStmtContext()
}

type AbsTimeStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsTimeStmtContext() *AbsTimeStmtContext {
	var p = new(AbsTimeStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_absTimeStmt
	return p
}

func InitEmptyAbsTimeStmtContext(p *AbsTimeStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_absTimeStmt
}

func (*AbsTimeStmtContext) IsAbsTimeStmtContext() {}

func NewAbsTimeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsTimeStmtContext {
	var p = new(AbsTimeStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_absTimeStmt

	return p
}

func (s *AbsTimeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsTimeStmtContext) AllSTRING_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(GuardQueryParserSTRING_LITERAL)
}

func (s *AbsTimeStmtContext) STRING_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSTRING_LITERAL, i)
}

func (s *AbsTimeStmtContext) TO_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserTO_, 0)
}

func (s *AbsTimeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsTimeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsTimeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterAbsTimeStmt(s)
	}
}

func (s *AbsTimeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitAbsTimeStmt(s)
	}
}

func (p *GuardQueryParser) AbsTimeStmt() (localctx IAbsTimeStmtContext) {
	localctx = NewAbsTimeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GuardQueryParserRULE_absTimeStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(GuardQueryParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GuardQueryParserTO_ {
		{
			p.SetState(85)
			p.Match(GuardQueryParserTO_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(GuardQueryParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IRelatTimeStmtContext is an interface to support dynamic dispatch.
type IRelatTimeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode
	AllNUMERIC_LITERAL() []antlr.TerminalNode
	NUMERIC_LITERAL(i int) antlr.TerminalNode
	AllTimeUnit() []ITimeUnitContext
	TimeUnit(i int) ITimeUnitContext
	TO_() antlr.TerminalNode

	// IsRelatTimeStmtContext differentiates from other interfaces.
	IsRelatTimeStmtContext()
}

type RelatTimeStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelatTimeStmtContext() *RelatTimeStmtContext {
	var p = new(RelatTimeStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_relatTimeStmt
	return p
}

func InitEmptyRelatTimeStmtContext(p *RelatTimeStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_relatTimeStmt
}

func (*RelatTimeStmtContext) IsRelatTimeStmtContext() {}

func NewRelatTimeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelatTimeStmtContext {
	var p = new(RelatTimeStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_relatTimeStmt

	return p
}

func (s *RelatTimeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RelatTimeStmtContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(GuardQueryParserMINUS)
}

func (s *RelatTimeStmtContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(GuardQueryParserMINUS, i)
}

func (s *RelatTimeStmtContext) AllNUMERIC_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(GuardQueryParserNUMERIC_LITERAL)
}

func (s *RelatTimeStmtContext) NUMERIC_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNUMERIC_LITERAL, i)
}

func (s *RelatTimeStmtContext) AllTimeUnit() []ITimeUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITimeUnitContext); ok {
			len++
		}
	}

	tst := make([]ITimeUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITimeUnitContext); ok {
			tst[i] = t.(ITimeUnitContext)
			i++
		}
	}

	return tst
}

func (s *RelatTimeStmtContext) TimeUnit(i int) ITimeUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeUnitContext); ok {
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

	return t.(ITimeUnitContext)
}

func (s *RelatTimeStmtContext) TO_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserTO_, 0)
}

func (s *RelatTimeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelatTimeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelatTimeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterRelatTimeStmt(s)
	}
}

func (s *RelatTimeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitRelatTimeStmt(s)
	}
}

func (p *GuardQueryParser) RelatTimeStmt() (localctx IRelatTimeStmtContext) {
	localctx = NewRelatTimeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GuardQueryParserRULE_relatTimeStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(GuardQueryParserMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(GuardQueryParserNUMERIC_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.TimeUnit()
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GuardQueryParserTO_ {
		{
			p.SetState(92)
			p.Match(GuardQueryParserTO_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(GuardQueryParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(GuardQueryParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.TimeUnit()
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

// ITimeUnitContext is an interface to support dynamic dispatch.
type ITimeUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTimeUnitContext differentiates from other interfaces.
	IsTimeUnitContext()
}

type TimeUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeUnitContext() *TimeUnitContext {
	var p = new(TimeUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_timeUnit
	return p
}

func InitEmptyTimeUnitContext(p *TimeUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_timeUnit
}

func (*TimeUnitContext) IsTimeUnitContext() {}

func NewTimeUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeUnitContext {
	var p = new(TimeUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_timeUnit

	return p
}

func (s *TimeUnitContext) GetParser() antlr.Parser { return s.parser }
func (s *TimeUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterTimeUnit(s)
	}
}

func (s *TimeUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitTimeUnit(s)
	}
}

func (p *GuardQueryParser) TimeUnit() (localctx ITimeUnitContext) {
	localctx = NewTimeUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GuardQueryParserRULE_timeUnit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
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

// IWhereStmtContext is an interface to support dynamic dispatch.
type IWhereStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BoolExpr() IBoolExprContext

	// IsWhereStmtContext differentiates from other interfaces.
	IsWhereStmtContext()
}

type WhereStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereStmtContext() *WhereStmtContext {
	var p = new(WhereStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_whereStmt
	return p
}

func InitEmptyWhereStmtContext(p *WhereStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_whereStmt
}

func (*WhereStmtContext) IsWhereStmtContext() {}

func NewWhereStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereStmtContext {
	var p = new(WhereStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_whereStmt

	return p
}

func (s *WhereStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereStmtContext) BoolExpr() IBoolExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *WhereStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterWhereStmt(s)
	}
}

func (s *WhereStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitWhereStmt(s)
	}
}

func (p *GuardQueryParser) WhereStmt() (localctx IWhereStmtContext) {
	localctx = NewWhereStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GuardQueryParserRULE_whereStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.boolExpr(0)
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

// ISourceOrSubqueryContext is an interface to support dynamic dispatch.
type ISourceOrSubqueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Source() ISourceContext

	// IsSourceOrSubqueryContext differentiates from other interfaces.
	IsSourceOrSubqueryContext()
}

type SourceOrSubqueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceOrSubqueryContext() *SourceOrSubqueryContext {
	var p = new(SourceOrSubqueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_sourceOrSubquery
	return p
}

func InitEmptySourceOrSubqueryContext(p *SourceOrSubqueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_sourceOrSubquery
}

func (*SourceOrSubqueryContext) IsSourceOrSubqueryContext() {}

func NewSourceOrSubqueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceOrSubqueryContext {
	var p = new(SourceOrSubqueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_sourceOrSubquery

	return p
}

func (s *SourceOrSubqueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceOrSubqueryContext) Source() ISourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SourceOrSubqueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceOrSubqueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceOrSubqueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterSourceOrSubquery(s)
	}
}

func (s *SourceOrSubqueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitSourceOrSubquery(s)
	}
}

func (p *GuardQueryParser) SourceOrSubquery() (localctx ISourceOrSubqueryContext) {
	localctx = NewSourceOrSubqueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GuardQueryParserRULE_sourceOrSubquery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Source()
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

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NODE_() antlr.TerminalNode
	Node() INodeContext
	PROVIDER_() antlr.TerminalNode
	AllProvider() []IProviderContext
	Provider(i int) IProviderContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_source
	return p
}

func InitEmptySourceContext(p *SourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_source
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) NODE_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNODE_, 0)
}

func (s *SourceContext) Node() INodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *SourceContext) PROVIDER_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserPROVIDER_, 0)
}

func (s *SourceContext) AllProvider() []IProviderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProviderContext); ok {
			len++
		}
	}

	tst := make([]IProviderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProviderContext); ok {
			tst[i] = t.(IProviderContext)
			i++
		}
	}

	return tst
}

func (s *SourceContext) Provider(i int) IProviderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProviderContext); ok {
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

	return t.(IProviderContext)
}

func (s *SourceContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GuardQueryParserCOMMA)
}

func (s *SourceContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCOMMA, i)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterSource(s)
	}
}

func (s *SourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitSource(s)
	}
}

func (p *GuardQueryParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GuardQueryParserRULE_source)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(GuardQueryParserNODE_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Node()
	}
	{
		p.SetState(106)
		p.Match(GuardQueryParserPROVIDER_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1407374883619840) != 0 {
		{
			p.SetState(107)
			p.Provider()
		}

	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(110)
				p.Match(GuardQueryParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(111)
				p.Provider()
			}

		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNodeId returns the nodeId rule contexts.
	GetNodeId() IAnyNameContext

	// SetNodeId sets the nodeId rule contexts.
	SetNodeId(IAnyNameContext)

	// Getter signatures
	STAR() antlr.TerminalNode
	AnyName() IAnyNameContext

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	nodeId IAnyNameContext
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_node
	return p
}

func InitEmptyNodeContext(p *NodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_node
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) GetNodeId() IAnyNameContext { return s.nodeId }

func (s *NodeContext) SetNodeId(v IAnyNameContext) { s.nodeId = v }

func (s *NodeContext) STAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSTAR, 0)
}

func (s *NodeContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterNode(s)
	}
}

func (s *NodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitNode(s)
	}
}

func (p *GuardQueryParser) Node() (localctx INodeContext) {
	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GuardQueryParserRULE_node)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Match(GuardQueryParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserOPEN_PAR, GuardQueryParserIDENTIFIER, GuardQueryParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)

			var _x = p.AnyName()

			localctx.(*NodeContext).nodeId = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IProviderContext is an interface to support dynamic dispatch.
type IProviderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetProviderType returns the providerType rule contexts.
	GetProviderType() IAnyNameContext

	// SetProviderType sets the providerType rule contexts.
	SetProviderType(IAnyNameContext)

	// Getter signatures
	STAR() antlr.TerminalNode
	AnyName() IAnyNameContext
	OPEN_PAR() antlr.TerminalNode
	CLOSE_PAR() antlr.TerminalNode
	AllProviderArg() []IProviderArgContext
	ProviderArg(i int) IProviderArgContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsProviderContext differentiates from other interfaces.
	IsProviderContext()
}

type ProviderContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	providerType IAnyNameContext
}

func NewEmptyProviderContext() *ProviderContext {
	var p = new(ProviderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_provider
	return p
}

func InitEmptyProviderContext(p *ProviderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_provider
}

func (*ProviderContext) IsProviderContext() {}

func NewProviderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProviderContext {
	var p = new(ProviderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_provider

	return p
}

func (s *ProviderContext) GetParser() antlr.Parser { return s.parser }

func (s *ProviderContext) GetProviderType() IAnyNameContext { return s.providerType }

func (s *ProviderContext) SetProviderType(v IAnyNameContext) { s.providerType = v }

func (s *ProviderContext) STAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSTAR, 0)
}

func (s *ProviderContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *ProviderContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserOPEN_PAR, 0)
}

func (s *ProviderContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCLOSE_PAR, 0)
}

func (s *ProviderContext) AllProviderArg() []IProviderArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProviderArgContext); ok {
			len++
		}
	}

	tst := make([]IProviderArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProviderArgContext); ok {
			tst[i] = t.(IProviderArgContext)
			i++
		}
	}

	return tst
}

func (s *ProviderContext) ProviderArg(i int) IProviderArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProviderArgContext); ok {
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

	return t.(IProviderArgContext)
}

func (s *ProviderContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GuardQueryParserCOMMA)
}

func (s *ProviderContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCOMMA, i)
}

func (s *ProviderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProviderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProviderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterProvider(s)
	}
}

func (s *ProviderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitProvider(s)
	}
}

func (p *GuardQueryParser) Provider() (localctx IProviderContext) {
	localctx = NewProviderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GuardQueryParserRULE_provider)
	var _la int

	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(GuardQueryParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserOPEN_PAR, GuardQueryParserIDENTIFIER, GuardQueryParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)

			var _x = p.AnyName()

			localctx.(*ProviderContext).providerType = _x
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GuardQueryParserOPEN_PAR {
			{
				p.SetState(123)
				p.Match(GuardQueryParserOPEN_PAR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			{
				p.SetState(124)
				p.ProviderArg()
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GuardQueryParserCOMMA {
				{
					p.SetState(125)
					p.Match(GuardQueryParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(126)
					p.ProviderArg()
				}

				p.SetState(131)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(132)
				p.Match(GuardQueryParserCLOSE_PAR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IProviderArgContext is an interface to support dynamic dispatch.
type IProviderArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	NUMERIC_LITERAL() antlr.TerminalNode

	// IsProviderArgContext differentiates from other interfaces.
	IsProviderArgContext()
}

type ProviderArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProviderArgContext() *ProviderArgContext {
	var p = new(ProviderArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_providerArg
	return p
}

func InitEmptyProviderArgContext(p *ProviderArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_providerArg
}

func (*ProviderArgContext) IsProviderArgContext() {}

func NewProviderArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProviderArgContext {
	var p = new(ProviderArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_providerArg

	return p
}

func (s *ProviderArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProviderArgContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSTRING_LITERAL, 0)
}

func (s *ProviderArgContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNUMERIC_LITERAL, 0)
}

func (s *ProviderArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProviderArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProviderArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterProviderArg(s)
	}
}

func (s *ProviderArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitProviderArg(s)
	}
}

func (p *GuardQueryParser) ProviderArg() (localctx IProviderArgContext) {
	localctx = NewProviderArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GuardQueryParserRULE_providerArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GuardQueryParserNUMERIC_LITERAL || _la == GuardQueryParserSTRING_LITERAL) {
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

// IResultColumnContext is an interface to support dynamic dispatch.
type IResultColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValueExpr() IValueExprContext
	WITH_() antlr.TerminalNode
	WindowFunction() IWindowFunctionContext
	ResultAlias() IResultAliasContext
	AS_() antlr.TerminalNode

	// IsResultColumnContext differentiates from other interfaces.
	IsResultColumnContext()
}

type ResultColumnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultColumnContext() *ResultColumnContext {
	var p = new(ResultColumnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_resultColumn
	return p
}

func InitEmptyResultColumnContext(p *ResultColumnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_resultColumn
}

func (*ResultColumnContext) IsResultColumnContext() {}

func NewResultColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultColumnContext {
	var p = new(ResultColumnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_resultColumn

	return p
}

func (s *ResultColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultColumnContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *ResultColumnContext) WITH_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserWITH_, 0)
}

func (s *ResultColumnContext) WindowFunction() IWindowFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindowFunctionContext)
}

func (s *ResultColumnContext) ResultAlias() IResultAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultAliasContext)
}

func (s *ResultColumnContext) AS_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserAS_, 0)
}

func (s *ResultColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterResultColumn(s)
	}
}

func (s *ResultColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitResultColumn(s)
	}
}

func (p *GuardQueryParser) ResultColumn() (localctx IResultColumnContext) {
	localctx = NewResultColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GuardQueryParserRULE_resultColumn)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.valueExpr(0)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GuardQueryParserWITH_ {
		{
			p.SetState(141)
			p.Match(GuardQueryParserWITH_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.WindowFunction()
		}

	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1407443603030016) != 0 {
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GuardQueryParserAS_ {
			{
				p.SetState(145)
				p.Match(GuardQueryParserAS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(148)
			p.ResultAlias()
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

// IWindowFunctionContext is an interface to support dynamic dispatch.
type IWindowFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWindowFunctionContext differentiates from other interfaces.
	IsWindowFunctionContext()
}

type WindowFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowFunctionContext() *WindowFunctionContext {
	var p = new(WindowFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_windowFunction
	return p
}

func InitEmptyWindowFunctionContext(p *WindowFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_windowFunction
}

func (*WindowFunctionContext) IsWindowFunctionContext() {}

func NewWindowFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowFunctionContext {
	var p = new(WindowFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_windowFunction

	return p
}

func (s *WindowFunctionContext) GetParser() antlr.Parser { return s.parser }
func (s *WindowFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterWindowFunction(s)
	}
}

func (s *WindowFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitWindowFunction(s)
	}
}

func (p *GuardQueryParser) WindowFunction() (localctx IWindowFunctionContext) {
	localctx = NewWindowFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GuardQueryParserRULE_windowFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(GuardQueryParserT__3)
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

// IResultAliasContext is an interface to support dynamic dispatch.
type IResultAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode

	// IsResultAliasContext differentiates from other interfaces.
	IsResultAliasContext()
}

type ResultAliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultAliasContext() *ResultAliasContext {
	var p = new(ResultAliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_resultAlias
	return p
}

func InitEmptyResultAliasContext(p *ResultAliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_resultAlias
}

func (*ResultAliasContext) IsResultAliasContext() {}

func NewResultAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultAliasContext {
	var p = new(ResultAliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_resultAlias

	return p
}

func (s *ResultAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultAliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserIDENTIFIER, 0)
}

func (s *ResultAliasContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSTRING_LITERAL, 0)
}

func (s *ResultAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultAliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterResultAlias(s)
	}
}

func (s *ResultAliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitResultAlias(s)
	}
}

func (p *GuardQueryParser) ResultAlias() (localctx IResultAliasContext) {
	localctx = NewResultAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GuardQueryParserRULE_resultAlias)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GuardQueryParserIDENTIFIER || _la == GuardQueryParserSTRING_LITERAL) {
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

// IAnyNameContext is an interface to support dynamic dispatch.
type IAnyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	OPEN_PAR() antlr.TerminalNode
	AnyName() IAnyNameContext
	CLOSE_PAR() antlr.TerminalNode

	// IsAnyNameContext differentiates from other interfaces.
	IsAnyNameContext()
}

type AnyNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyNameContext() *AnyNameContext {
	var p = new(AnyNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_anyName
	return p
}

func InitEmptyAnyNameContext(p *AnyNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_anyName
}

func (*AnyNameContext) IsAnyNameContext() {}

func NewAnyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyNameContext {
	var p = new(AnyNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_anyName

	return p
}

func (s *AnyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserIDENTIFIER, 0)
}

func (s *AnyNameContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSTRING_LITERAL, 0)
}

func (s *AnyNameContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserOPEN_PAR, 0)
}

func (s *AnyNameContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *AnyNameContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCLOSE_PAR, 0)
}

func (s *AnyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterAnyName(s)
	}
}

func (s *AnyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitAnyName(s)
	}
}

func (p *GuardQueryParser) AnyName() (localctx IAnyNameContext) {
	localctx = NewAnyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GuardQueryParserRULE_anyName)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Match(GuardQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(GuardQueryParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserOPEN_PAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.Match(GuardQueryParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.AnyName()
		}
		{
			p.SetState(159)
			p.Match(GuardQueryParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IBoolExprContext is an interface to support dynamic dispatch.
type IBoolExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE_() antlr.TerminalNode
	FALSE_() antlr.TerminalNode
	AllValueExpr() []IValueExprContext
	ValueExpr(i int) IValueExprContext
	LT() antlr.TerminalNode
	LT_EQ() antlr.TerminalNode
	GT() antlr.TerminalNode
	GT_EQ() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NOT_EQ1() antlr.TerminalNode
	NOT_EQ2() antlr.TerminalNode
	NOT_() antlr.TerminalNode
	AllBoolExpr() []IBoolExprContext
	BoolExpr(i int) IBoolExprContext
	OPEN_PAR() antlr.TerminalNode
	CLOSE_PAR() antlr.TerminalNode
	AND_() antlr.TerminalNode
	OR_() antlr.TerminalNode

	// IsBoolExprContext differentiates from other interfaces.
	IsBoolExprContext()
}

type BoolExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExprContext() *BoolExprContext {
	var p = new(BoolExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_boolExpr
	return p
}

func InitEmptyBoolExprContext(p *BoolExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_boolExpr
}

func (*BoolExprContext) IsBoolExprContext() {}

func NewBoolExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExprContext {
	var p = new(BoolExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_boolExpr

	return p
}

func (s *BoolExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExprContext) TRUE_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserTRUE_, 0)
}

func (s *BoolExprContext) FALSE_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserFALSE_, 0)
}

func (s *BoolExprContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *BoolExprContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
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

	return t.(IValueExprContext)
}

func (s *BoolExprContext) LT() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserLT, 0)
}

func (s *BoolExprContext) LT_EQ() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserLT_EQ, 0)
}

func (s *BoolExprContext) GT() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserGT, 0)
}

func (s *BoolExprContext) GT_EQ() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserGT_EQ, 0)
}

func (s *BoolExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserASSIGN, 0)
}

func (s *BoolExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserEQ, 0)
}

func (s *BoolExprContext) NOT_EQ1() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNOT_EQ1, 0)
}

func (s *BoolExprContext) NOT_EQ2() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNOT_EQ2, 0)
}

func (s *BoolExprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNOT_, 0)
}

func (s *BoolExprContext) AllBoolExpr() []IBoolExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolExprContext); ok {
			len++
		}
	}

	tst := make([]IBoolExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolExprContext); ok {
			tst[i] = t.(IBoolExprContext)
			i++
		}
	}

	return tst
}

func (s *BoolExprContext) BoolExpr(i int) IBoolExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExprContext); ok {
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

	return t.(IBoolExprContext)
}

func (s *BoolExprContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserOPEN_PAR, 0)
}

func (s *BoolExprContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCLOSE_PAR, 0)
}

func (s *BoolExprContext) AND_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserAND_, 0)
}

func (s *BoolExprContext) OR_() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserOR_, 0)
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (p *GuardQueryParser) BoolExpr() (localctx IBoolExprContext) {
	return p.boolExpr(0)
}

func (p *GuardQueryParser) boolExpr(_p int) (localctx IBoolExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBoolExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, GuardQueryParserRULE_boolExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(164)
			p.Match(GuardQueryParserTRUE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(165)
			p.Match(GuardQueryParserFALSE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(166)
			p.valueExpr(0)
		}
		{
			p.SetState(167)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728640) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(168)
			p.valueExpr(0)
		}

	case 4:
		{
			p.SetState(170)
			p.valueExpr(0)
		}
		{
			p.SetState(171)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117473280) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(172)
			p.valueExpr(0)
		}

	case 5:
		{
			p.SetState(174)
			p.Match(GuardQueryParserNOT_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.boolExpr(2)
		}

	case 6:
		{
			p.SetState(176)
			p.Match(GuardQueryParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.boolExpr(0)
		}
		{
			p.SetState(178)
			p.Match(GuardQueryParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBoolExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GuardQueryParserRULE_boolExpr)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(183)
					p.Match(GuardQueryParserAND_)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(184)
					p.boolExpr(5)
				}

			case 2:
				localctx = NewBoolExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GuardQueryParserRULE_boolExpr)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(186)
					p.Match(GuardQueryParserOR_)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(187)
					p.boolExpr(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
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

// IValueExprContext is an interface to support dynamic dispatch.
type IValueExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	BuildinSource() IBuildinSourceContext
	BuildinFunction() IBuildinFunctionContext
	OPEN_PAR() antlr.TerminalNode
	CLOSE_PAR() antlr.TerminalNode
	AllValueExpr() []IValueExprContext
	ValueExpr(i int) IValueExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode
	OPEN_BRA() antlr.TerminalNode
	CLOSE_BRA() antlr.TerminalNode

	// IsValueExprContext differentiates from other interfaces.
	IsValueExprContext()
}

type ValueExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExprContext() *ValueExprContext {
	var p = new(ValueExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_valueExpr
	return p
}

func InitEmptyValueExprContext(p *ValueExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_valueExpr
}

func (*ValueExprContext) IsValueExprContext() {}

func NewValueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExprContext {
	var p = new(ValueExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_valueExpr

	return p
}

func (s *ValueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExprContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNUMERIC_LITERAL, 0)
}

func (s *ValueExprContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSTRING_LITERAL, 0)
}

func (s *ValueExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserIDENTIFIER, 0)
}

func (s *ValueExprContext) BuildinSource() IBuildinSourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuildinSourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuildinSourceContext)
}

func (s *ValueExprContext) BuildinFunction() IBuildinFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuildinFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuildinFunctionContext)
}

func (s *ValueExprContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserOPEN_PAR, 0)
}

func (s *ValueExprContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCLOSE_PAR, 0)
}

func (s *ValueExprContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *ValueExprContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
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

	return t.(IValueExprContext)
}

func (s *ValueExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GuardQueryParserCOMMA)
}

func (s *ValueExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCOMMA, i)
}

func (s *ValueExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserDOT, 0)
}

func (s *ValueExprContext) OPEN_BRA() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserOPEN_BRA, 0)
}

func (s *ValueExprContext) CLOSE_BRA() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserCLOSE_BRA, 0)
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterValueExpr(s)
	}
}

func (s *ValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitValueExpr(s)
	}
}

func (p *GuardQueryParser) ValueExpr() (localctx IValueExprContext) {
	return p.valueExpr(0)
}

func (p *GuardQueryParser) valueExpr(_p int) (localctx IValueExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewValueExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, GuardQueryParserRULE_valueExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserNUMERIC_LITERAL:
		{
			p.SetState(194)
			p.Match(GuardQueryParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserSTRING_LITERAL:
		{
			p.SetState(195)
			p.Match(GuardQueryParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserIDENTIFIER:
		{
			p.SetState(196)
			p.Match(GuardQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserT__4:
		{
			p.SetState(197)
			p.BuildinSource()
		}

	case GuardQueryParserT__5, GuardQueryParserT__6:
		{
			p.SetState(198)
			p.BuildinFunction()
		}
		{
			p.SetState(199)
			p.Match(GuardQueryParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1970324836975840) != 0 {
			{
				p.SetState(200)
				p.valueExpr(0)
			}

		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GuardQueryParserCOMMA {
			{
				p.SetState(203)
				p.Match(GuardQueryParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(204)
				p.valueExpr(0)
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(210)
			p.Match(GuardQueryParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserOPEN_PAR:
		{
			p.SetState(212)
			p.Match(GuardQueryParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.valueExpr(0)
		}
		{
			p.SetState(214)
			p.Match(GuardQueryParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewValueExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GuardQueryParserRULE_valueExpr)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(219)
					p.Match(GuardQueryParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(220)
					p.valueExpr(5)
				}

			case 2:
				localctx = NewValueExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GuardQueryParserRULE_valueExpr)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(222)
					p.Match(GuardQueryParserOPEN_BRA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(223)
					p.valueExpr(0)
				}
				{
					p.SetState(224)
					p.Match(GuardQueryParserCLOSE_BRA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
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

// IBuildinSourceContext is an interface to support dynamic dispatch.
type IBuildinSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBuildinSourceContext differentiates from other interfaces.
	IsBuildinSourceContext()
}

type BuildinSourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildinSourceContext() *BuildinSourceContext {
	var p = new(BuildinSourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_buildinSource
	return p
}

func InitEmptyBuildinSourceContext(p *BuildinSourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_buildinSource
}

func (*BuildinSourceContext) IsBuildinSourceContext() {}

func NewBuildinSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildinSourceContext {
	var p = new(BuildinSourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_buildinSource

	return p
}

func (s *BuildinSourceContext) GetParser() antlr.Parser { return s.parser }
func (s *BuildinSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildinSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildinSourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterBuildinSource(s)
	}
}

func (s *BuildinSourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitBuildinSource(s)
	}
}

func (p *GuardQueryParser) BuildinSource() (localctx IBuildinSourceContext) {
	localctx = NewBuildinSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GuardQueryParserRULE_buildinSource)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(GuardQueryParserT__4)
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

// IBuildinFunctionContext is an interface to support dynamic dispatch.
type IBuildinFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBuildinFunctionContext differentiates from other interfaces.
	IsBuildinFunctionContext()
}

type BuildinFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildinFunctionContext() *BuildinFunctionContext {
	var p = new(BuildinFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_buildinFunction
	return p
}

func InitEmptyBuildinFunctionContext(p *BuildinFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_buildinFunction
}

func (*BuildinFunctionContext) IsBuildinFunctionContext() {}

func NewBuildinFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildinFunctionContext {
	var p = new(BuildinFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_buildinFunction

	return p
}

func (s *BuildinFunctionContext) GetParser() antlr.Parser { return s.parser }
func (s *BuildinFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildinFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildinFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterBuildinFunction(s)
	}
}

func (s *BuildinFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitBuildinFunction(s)
	}
}

func (p *GuardQueryParser) BuildinFunction() (localctx IBuildinFunctionContext) {
	localctx = NewBuildinFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GuardQueryParserRULE_buildinFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GuardQueryParserT__5 || _la == GuardQueryParserT__6) {
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

// ILiteralValueContext is an interface to support dynamic dispatch.
type ILiteralValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode

	// IsLiteralValueContext differentiates from other interfaces.
	IsLiteralValueContext()
}

type LiteralValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralValueContext() *LiteralValueContext {
	var p = new(LiteralValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_literalValue
	return p
}

func InitEmptyLiteralValueContext(p *LiteralValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GuardQueryParserRULE_literalValue
}

func (*LiteralValueContext) IsLiteralValueContext() {}

func NewLiteralValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralValueContext {
	var p = new(LiteralValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GuardQueryParserRULE_literalValue

	return p
}

func (s *LiteralValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralValueContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserNUMERIC_LITERAL, 0)
}

func (s *LiteralValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(GuardQueryParserSTRING_LITERAL, 0)
}

func (s *LiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.EnterLiteralValue(s)
	}
}

func (s *LiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GuardQueryListener); ok {
		listenerT.ExitLiteralValue(s)
	}
}

func (p *GuardQueryParser) LiteralValue() (localctx ILiteralValueContext) {
	localctx = NewLiteralValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GuardQueryParserRULE_literalValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GuardQueryParserNUMERIC_LITERAL || _la == GuardQueryParserSTRING_LITERAL) {
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

func (p *GuardQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *BoolExprContext = nil
		if localctx != nil {
			t = localctx.(*BoolExprContext)
		}
		return p.BoolExpr_Sempred(t, predIndex)

	case 18:
		var t *ValueExprContext = nil
		if localctx != nil {
			t = localctx.(*ValueExprContext)
		}
		return p.ValueExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GuardQueryParser) BoolExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GuardQueryParser) ValueExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
