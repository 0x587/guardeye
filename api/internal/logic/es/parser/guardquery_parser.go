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
		"", "'$msg'", "'CURRENT_TIME'", "'CURRENT_DATE'", "'CURRENT_TIMESTAMP'",
		"'json'", "'yaml'", "';'", "'.'", "'('", "')'", "'['", "']'", "','",
		"'='", "'*'", "'+'", "'-'", "'/'", "'<'", "'<='", "'>'", "'>='", "'=='",
		"'!='", "'<>'", "'NODE'", "'PROVIDER'", "'AND'", "'OR'", "'NOT'", "'SELECT'",
		"'FROM'", "'AS'", "'WHERE'", "'LIMIT'", "'OFFSET'", "'ORDER'", "'BY'",
		"'ASC'", "'DESC'", "'TRUE'", "'FALSE'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "SCOL", "DOT", "OPEN_PAR", "CLOSE_PAR",
		"OPEN_BRA", "CLOSE_BRA", "COMMA", "ASSIGN", "STAR", "PLUS", "MINUS",
		"DIV", "LT", "LT_EQ", "GT", "GT_EQ", "EQ", "NOT_EQ1", "NOT_EQ2", "NODE_",
		"PROVIDER_", "AND_", "OR_", "NOT_", "SELECT_", "FROM_", "AS_", "WHERE_",
		"LIMIT_", "OFFSET_", "ORDER_", "BY_", "ASC_", "DESC_", "TRUE_", "FALSE_",
		"IDENTIFIER", "NUMERIC_LITERAL", "STRING_LITERAL", "SINGLE_LINE_COMMENT",
		"MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
	}
	staticData.RuleNames = []string{
		"parse", "select", "whereStmt", "sourceOrSubquery", "source", "node",
		"provider", "providerArg", "resultColumn", "resultAlias", "anyName",
		"boolExpr", "valueExpr", "buildinSource", "buildinFunction", "literalValue",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 49, 189, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 40, 8, 1, 10, 1, 12, 1,
		43, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 49, 8, 1, 10, 1, 12, 1, 52, 9,
		1, 1, 1, 1, 1, 3, 1, 56, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 66, 8, 4, 1, 4, 1, 4, 5, 4, 70, 8, 4, 10, 4, 12, 4, 73, 9,
		4, 1, 5, 1, 5, 3, 5, 77, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6,
		85, 8, 6, 10, 6, 12, 6, 88, 9, 6, 1, 6, 1, 6, 3, 6, 92, 8, 6, 3, 6, 94,
		8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 100, 8, 8, 1, 8, 3, 8, 103, 8, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 113, 8, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 132, 8, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 140, 8, 11, 10, 11, 12, 11, 143,
		9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 153,
		8, 12, 1, 12, 1, 12, 5, 12, 157, 8, 12, 10, 12, 12, 12, 160, 9, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 168, 8, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 178, 8, 12, 10, 12, 12,
		12, 181, 9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 0, 2,
		22, 24, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		0, 6, 1, 0, 44, 45, 2, 0, 43, 43, 45, 45, 1, 0, 19, 22, 2, 0, 14, 14, 23,
		25, 1, 0, 1, 4, 1, 0, 5, 6, 201, 0, 32, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0,
		4, 57, 1, 0, 0, 0, 6, 59, 1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 76, 1, 0,
		0, 0, 12, 93, 1, 0, 0, 0, 14, 95, 1, 0, 0, 0, 16, 97, 1, 0, 0, 0, 18, 104,
		1, 0, 0, 0, 20, 112, 1, 0, 0, 0, 22, 131, 1, 0, 0, 0, 24, 167, 1, 0, 0,
		0, 26, 182, 1, 0, 0, 0, 28, 184, 1, 0, 0, 0, 30, 186, 1, 0, 0, 0, 32, 33,
		3, 2, 1, 0, 33, 34, 5, 7, 0, 0, 34, 1, 1, 0, 0, 0, 35, 36, 5, 31, 0, 0,
		36, 41, 3, 16, 8, 0, 37, 38, 5, 13, 0, 0, 38, 40, 3, 16, 8, 0, 39, 37,
		1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0,
		42, 44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 5, 32, 0, 0, 45, 50, 3,
		6, 3, 0, 46, 47, 5, 13, 0, 0, 47, 49, 3, 6, 3, 0, 48, 46, 1, 0, 0, 0, 49,
		52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 55, 1, 0, 0,
		0, 52, 50, 1, 0, 0, 0, 53, 54, 5, 34, 0, 0, 54, 56, 3, 4, 2, 0, 55, 53,
		1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 3, 1, 0, 0, 0, 57, 58, 3, 22, 11, 0,
		58, 5, 1, 0, 0, 0, 59, 60, 3, 8, 4, 0, 60, 7, 1, 0, 0, 0, 61, 62, 5, 26,
		0, 0, 62, 63, 3, 10, 5, 0, 63, 65, 5, 27, 0, 0, 64, 66, 3, 12, 6, 0, 65,
		64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 71, 1, 0, 0, 0, 67, 68, 5, 13,
		0, 0, 68, 70, 3, 12, 6, 0, 69, 67, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71,
		69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 9, 1, 0, 0, 0, 73, 71, 1, 0, 0,
		0, 74, 77, 5, 15, 0, 0, 75, 77, 3, 20, 10, 0, 76, 74, 1, 0, 0, 0, 76, 75,
		1, 0, 0, 0, 77, 11, 1, 0, 0, 0, 78, 94, 5, 15, 0, 0, 79, 91, 3, 20, 10,
		0, 80, 81, 5, 9, 0, 0, 81, 86, 3, 14, 7, 0, 82, 83, 5, 13, 0, 0, 83, 85,
		3, 14, 7, 0, 84, 82, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		86, 87, 1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 5,
		10, 0, 0, 90, 92, 1, 0, 0, 0, 91, 80, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		94, 1, 0, 0, 0, 93, 78, 1, 0, 0, 0, 93, 79, 1, 0, 0, 0, 94, 13, 1, 0, 0,
		0, 95, 96, 7, 0, 0, 0, 96, 15, 1, 0, 0, 0, 97, 102, 3, 24, 12, 0, 98, 100,
		5, 33, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0,
		0, 101, 103, 3, 18, 9, 0, 102, 99, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103,
		17, 1, 0, 0, 0, 104, 105, 7, 1, 0, 0, 105, 19, 1, 0, 0, 0, 106, 113, 5,
		43, 0, 0, 107, 113, 5, 45, 0, 0, 108, 109, 5, 9, 0, 0, 109, 110, 3, 20,
		10, 0, 110, 111, 5, 10, 0, 0, 111, 113, 1, 0, 0, 0, 112, 106, 1, 0, 0,
		0, 112, 107, 1, 0, 0, 0, 112, 108, 1, 0, 0, 0, 113, 21, 1, 0, 0, 0, 114,
		115, 6, 11, -1, 0, 115, 132, 5, 41, 0, 0, 116, 132, 5, 42, 0, 0, 117, 118,
		3, 24, 12, 0, 118, 119, 7, 2, 0, 0, 119, 120, 3, 24, 12, 0, 120, 132, 1,
		0, 0, 0, 121, 122, 3, 24, 12, 0, 122, 123, 7, 3, 0, 0, 123, 124, 3, 24,
		12, 0, 124, 132, 1, 0, 0, 0, 125, 126, 5, 30, 0, 0, 126, 132, 3, 22, 11,
		2, 127, 128, 5, 9, 0, 0, 128, 129, 3, 22, 11, 0, 129, 130, 5, 10, 0, 0,
		130, 132, 1, 0, 0, 0, 131, 114, 1, 0, 0, 0, 131, 116, 1, 0, 0, 0, 131,
		117, 1, 0, 0, 0, 131, 121, 1, 0, 0, 0, 131, 125, 1, 0, 0, 0, 131, 127,
		1, 0, 0, 0, 132, 141, 1, 0, 0, 0, 133, 134, 10, 4, 0, 0, 134, 135, 5, 28,
		0, 0, 135, 140, 3, 22, 11, 5, 136, 137, 10, 3, 0, 0, 137, 138, 5, 29, 0,
		0, 138, 140, 3, 22, 11, 4, 139, 133, 1, 0, 0, 0, 139, 136, 1, 0, 0, 0,
		140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		23, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 6, 12, -1, 0, 145, 168,
		5, 44, 0, 0, 146, 168, 5, 45, 0, 0, 147, 168, 5, 43, 0, 0, 148, 168, 3,
		26, 13, 0, 149, 150, 3, 28, 14, 0, 150, 152, 5, 9, 0, 0, 151, 153, 3, 24,
		12, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 158, 1, 0, 0, 0,
		154, 155, 5, 13, 0, 0, 155, 157, 3, 24, 12, 0, 156, 154, 1, 0, 0, 0, 157,
		160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 161,
		1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 161, 162, 5, 10, 0, 0, 162, 168, 1, 0,
		0, 0, 163, 164, 5, 9, 0, 0, 164, 165, 3, 24, 12, 0, 165, 166, 5, 10, 0,
		0, 166, 168, 1, 0, 0, 0, 167, 144, 1, 0, 0, 0, 167, 146, 1, 0, 0, 0, 167,
		147, 1, 0, 0, 0, 167, 148, 1, 0, 0, 0, 167, 149, 1, 0, 0, 0, 167, 163,
		1, 0, 0, 0, 168, 179, 1, 0, 0, 0, 169, 170, 10, 4, 0, 0, 170, 171, 5, 8,
		0, 0, 171, 178, 3, 24, 12, 5, 172, 173, 10, 3, 0, 0, 173, 174, 5, 11, 0,
		0, 174, 175, 3, 24, 12, 0, 175, 176, 5, 12, 0, 0, 176, 178, 1, 0, 0, 0,
		177, 169, 1, 0, 0, 0, 177, 172, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179,
		177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 25, 1, 0, 0, 0, 181, 179, 1,
		0, 0, 0, 182, 183, 7, 4, 0, 0, 183, 27, 1, 0, 0, 0, 184, 185, 7, 5, 0,
		0, 185, 29, 1, 0, 0, 0, 186, 187, 7, 0, 0, 0, 187, 31, 1, 0, 0, 0, 20,
		41, 50, 55, 65, 71, 76, 86, 91, 93, 99, 102, 112, 131, 139, 141, 152, 158,
		167, 177, 179,
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
	GuardQueryParserSCOL                = 7
	GuardQueryParserDOT                 = 8
	GuardQueryParserOPEN_PAR            = 9
	GuardQueryParserCLOSE_PAR           = 10
	GuardQueryParserOPEN_BRA            = 11
	GuardQueryParserCLOSE_BRA           = 12
	GuardQueryParserCOMMA               = 13
	GuardQueryParserASSIGN              = 14
	GuardQueryParserSTAR                = 15
	GuardQueryParserPLUS                = 16
	GuardQueryParserMINUS               = 17
	GuardQueryParserDIV                 = 18
	GuardQueryParserLT                  = 19
	GuardQueryParserLT_EQ               = 20
	GuardQueryParserGT                  = 21
	GuardQueryParserGT_EQ               = 22
	GuardQueryParserEQ                  = 23
	GuardQueryParserNOT_EQ1             = 24
	GuardQueryParserNOT_EQ2             = 25
	GuardQueryParserNODE_               = 26
	GuardQueryParserPROVIDER_           = 27
	GuardQueryParserAND_                = 28
	GuardQueryParserOR_                 = 29
	GuardQueryParserNOT_                = 30
	GuardQueryParserSELECT_             = 31
	GuardQueryParserFROM_               = 32
	GuardQueryParserAS_                 = 33
	GuardQueryParserWHERE_              = 34
	GuardQueryParserLIMIT_              = 35
	GuardQueryParserOFFSET_             = 36
	GuardQueryParserORDER_              = 37
	GuardQueryParserBY_                 = 38
	GuardQueryParserASC_                = 39
	GuardQueryParserDESC_               = 40
	GuardQueryParserTRUE_               = 41
	GuardQueryParserFALSE_              = 42
	GuardQueryParserIDENTIFIER          = 43
	GuardQueryParserNUMERIC_LITERAL     = 44
	GuardQueryParserSTRING_LITERAL      = 45
	GuardQueryParserSINGLE_LINE_COMMENT = 46
	GuardQueryParserMULTILINE_COMMENT   = 47
	GuardQueryParserSPACES              = 48
	GuardQueryParserUNEXPECTED_CHAR     = 49
)

// GuardQueryParser rules.
const (
	GuardQueryParserRULE_parse            = 0
	GuardQueryParserRULE_select           = 1
	GuardQueryParserRULE_whereStmt        = 2
	GuardQueryParserRULE_sourceOrSubquery = 3
	GuardQueryParserRULE_source           = 4
	GuardQueryParserRULE_node             = 5
	GuardQueryParserRULE_provider         = 6
	GuardQueryParserRULE_providerArg      = 7
	GuardQueryParserRULE_resultColumn     = 8
	GuardQueryParserRULE_resultAlias      = 9
	GuardQueryParserRULE_anyName          = 10
	GuardQueryParserRULE_boolExpr         = 11
	GuardQueryParserRULE_valueExpr        = 12
	GuardQueryParserRULE_buildinSource    = 13
	GuardQueryParserRULE_buildinFunction  = 14
	GuardQueryParserRULE_literalValue     = 15
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
		p.SetState(32)
		p.Select_()
	}

	{
		p.SetState(33)
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
		p.SetState(35)
		p.Match(GuardQueryParserSELECT_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.ResultColumn()
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GuardQueryParserCOMMA {
		{
			p.SetState(37)
			p.Match(GuardQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.ResultColumn()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
		p.Match(GuardQueryParserFROM_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.SourceOrSubquery()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GuardQueryParserCOMMA {
		{
			p.SetState(46)
			p.Match(GuardQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.SourceOrSubquery()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GuardQueryParserWHERE_ {
		{
			p.SetState(53)
			p.Match(GuardQueryParserWHERE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
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
	p.EnterRule(localctx, 4, GuardQueryParserRULE_whereStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
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
	p.EnterRule(localctx, 6, GuardQueryParserRULE_sourceOrSubquery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
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
	p.EnterRule(localctx, 8, GuardQueryParserRULE_source)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(GuardQueryParserNODE_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(62)
		p.Node()
	}
	{
		p.SetState(63)
		p.Match(GuardQueryParserPROVIDER_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&43980465144320) != 0 {
		{
			p.SetState(64)
			p.Provider()
		}

	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(67)
				p.Match(GuardQueryParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(68)
				p.Provider()
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 10, GuardQueryParserRULE_node)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(GuardQueryParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserOPEN_PAR, GuardQueryParserIDENTIFIER, GuardQueryParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)

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
	p.EnterRule(localctx, 12, GuardQueryParserRULE_provider)
	var _la int

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(GuardQueryParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserOPEN_PAR, GuardQueryParserIDENTIFIER, GuardQueryParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)

			var _x = p.AnyName()

			localctx.(*ProviderContext).providerType = _x
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GuardQueryParserOPEN_PAR {
			{
				p.SetState(80)
				p.Match(GuardQueryParserOPEN_PAR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			{
				p.SetState(81)
				p.ProviderArg()
			}

			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GuardQueryParserCOMMA {
				{
					p.SetState(82)
					p.Match(GuardQueryParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(83)
					p.ProviderArg()
				}

				p.SetState(88)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(89)
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
	p.EnterRule(localctx, 14, GuardQueryParserRULE_providerArg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
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
	p.EnterRule(localctx, 16, GuardQueryParserRULE_resultColumn)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.valueExpr(0)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&43989055045632) != 0 {
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GuardQueryParserAS_ {
			{
				p.SetState(98)
				p.Match(GuardQueryParserAS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(101)
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
	p.EnterRule(localctx, 18, GuardQueryParserRULE_resultAlias)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
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
	p.EnterRule(localctx, 20, GuardQueryParserRULE_anyName)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(GuardQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Match(GuardQueryParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserOPEN_PAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Match(GuardQueryParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.AnyName()
		}
		{
			p.SetState(110)
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
	_startState := 22
	p.EnterRecursionRule(localctx, 22, GuardQueryParserRULE_boolExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(115)
			p.Match(GuardQueryParserTRUE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(116)
			p.Match(GuardQueryParserFALSE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(117)
			p.valueExpr(0)
		}
		{
			p.SetState(118)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7864320) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(119)
			p.valueExpr(0)
		}

	case 4:
		{
			p.SetState(121)
			p.valueExpr(0)
		}
		{
			p.SetState(122)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&58736640) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(123)
			p.valueExpr(0)
		}

	case 5:
		{
			p.SetState(125)
			p.Match(GuardQueryParserNOT_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.boolExpr(2)
		}

	case 6:
		{
			p.SetState(127)
			p.Match(GuardQueryParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.boolExpr(0)
		}
		{
			p.SetState(129)
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBoolExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GuardQueryParserRULE_boolExpr)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(134)
					p.Match(GuardQueryParserAND_)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(135)
					p.boolExpr(5)
				}

			case 2:
				localctx = NewBoolExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GuardQueryParserRULE_boolExpr)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(137)
					p.Match(GuardQueryParserOR_)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(138)
					p.boolExpr(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, GuardQueryParserRULE_valueExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GuardQueryParserNUMERIC_LITERAL:
		{
			p.SetState(145)
			p.Match(GuardQueryParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserSTRING_LITERAL:
		{
			p.SetState(146)
			p.Match(GuardQueryParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserIDENTIFIER:
		{
			p.SetState(147)
			p.Match(GuardQueryParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserT__0, GuardQueryParserT__1, GuardQueryParserT__2, GuardQueryParserT__3:
		{
			p.SetState(148)
			p.BuildinSource()
		}

	case GuardQueryParserT__4, GuardQueryParserT__5:
		{
			p.SetState(149)
			p.BuildinFunction()
		}
		{
			p.SetState(150)
			p.Match(GuardQueryParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&61572651156094) != 0 {
			{
				p.SetState(151)
				p.valueExpr(0)
			}

		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GuardQueryParserCOMMA {
			{
				p.SetState(154)
				p.Match(GuardQueryParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(155)
				p.valueExpr(0)
			}

			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(161)
			p.Match(GuardQueryParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GuardQueryParserOPEN_PAR:
		{
			p.SetState(163)
			p.Match(GuardQueryParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.valueExpr(0)
		}
		{
			p.SetState(165)
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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewValueExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GuardQueryParserRULE_valueExpr)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(170)
					p.Match(GuardQueryParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(171)
					p.valueExpr(5)
				}

			case 2:
				localctx = NewValueExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GuardQueryParserRULE_valueExpr)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(173)
					p.Match(GuardQueryParserOPEN_BRA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(174)
					p.valueExpr(0)
				}
				{
					p.SetState(175)
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
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, GuardQueryParserRULE_buildinSource)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
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
	p.EnterRule(localctx, 28, GuardQueryParserRULE_buildinFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GuardQueryParserT__4 || _la == GuardQueryParserT__5) {
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
	p.EnterRule(localctx, 30, GuardQueryParserRULE_literalValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
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
	case 11:
		var t *BoolExprContext = nil
		if localctx != nil {
			t = localctx.(*BoolExprContext)
		}
		return p.BoolExpr_Sempred(t, predIndex)

	case 12:
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
