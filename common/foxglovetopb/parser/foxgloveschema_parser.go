// Code generated from FoxgloveSchema.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FoxgloveSchema
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

type FoxgloveSchemaParser struct {
	*antlr.BaseParser
}

var FoxgloveSchemaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func foxgloveschemaParserInit() {
	staticData := &FoxgloveSchemaParserStaticData
	staticData.LiteralNames = []string{
		"", "'/'", "'false'", "'true'", "'['", "']'", "'='", "'<='", "'================================================================================'",
		"'MSG: '", "'bool'", "'bytes'", "'char'", "'float32'", "'float64'",
		"'int8'", "'uint8'", "'int16'", "'uint16'", "'int32'", "'uint32'", "'int64'",
		"'uint64'", "'string'", "'wstring'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "OPEN_BRA", "CLOSE_BRA", "EQ", "LT_EQ", "SCHEMA_SPLIT",
		"SUB_SCHEMA_HEADER", "BOOL", "BYTES", "CHAR", "FLOAT32", "FLOAT64",
		"INT8", "UINT8", "INT16", "UINT16", "INT32", "UINT32", "INT64", "UINT64",
		"STRING", "WSTRING", "IDENTIFIER", "NUMERIC_LITERAL", "BOOL_LITERAL",
		"STRING_LITERAL", "SKIP_",
	}
	staticData.RuleNames = []string{
		"parse", "mainSchema", "subSchema", "schemaName", "schema", "constance",
		"field", "expr", "fieldType", "type", "arrayType", "customType", "fieldName",
		"buildinType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 127, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 5, 0, 32,
		8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 48, 8, 3, 10, 3, 12, 3, 51, 9, 3, 1, 4, 1,
		4, 5, 4, 55, 8, 4, 10, 4, 12, 4, 58, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 3, 6, 68, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 74, 8,
		8, 1, 9, 1, 9, 3, 9, 78, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 115, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 121, 8, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 0, 0, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 0, 2, 3, 0, 2, 3, 26, 26, 28, 28, 1, 0, 10, 24, 126, 0, 28, 1, 0, 0,
		0, 2, 38, 1, 0, 0, 0, 4, 40, 1, 0, 0, 0, 6, 44, 1, 0, 0, 0, 8, 56, 1, 0,
		0, 0, 10, 59, 1, 0, 0, 0, 12, 64, 1, 0, 0, 0, 14, 69, 1, 0, 0, 0, 16, 73,
		1, 0, 0, 0, 18, 77, 1, 0, 0, 0, 20, 114, 1, 0, 0, 0, 22, 120, 1, 0, 0,
		0, 24, 122, 1, 0, 0, 0, 26, 124, 1, 0, 0, 0, 28, 33, 3, 2, 1, 0, 29, 30,
		5, 8, 0, 0, 30, 32, 3, 4, 2, 0, 31, 29, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0,
		33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 33, 1,
		0, 0, 0, 36, 37, 5, 0, 0, 1, 37, 1, 1, 0, 0, 0, 38, 39, 3, 8, 4, 0, 39,
		3, 1, 0, 0, 0, 40, 41, 5, 9, 0, 0, 41, 42, 3, 6, 3, 0, 42, 43, 3, 8, 4,
		0, 43, 5, 1, 0, 0, 0, 44, 49, 5, 25, 0, 0, 45, 46, 5, 1, 0, 0, 46, 48,
		5, 25, 0, 0, 47, 45, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0,
		49, 50, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 55, 3, 12,
		6, 0, 53, 55, 3, 10, 5, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55,
		58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 9, 1, 0, 0,
		0, 58, 56, 1, 0, 0, 0, 59, 60, 3, 16, 8, 0, 60, 61, 5, 25, 0, 0, 61, 62,
		5, 6, 0, 0, 62, 63, 3, 14, 7, 0, 63, 11, 1, 0, 0, 0, 64, 65, 3, 16, 8,
		0, 65, 67, 3, 24, 12, 0, 66, 68, 3, 14, 7, 0, 67, 66, 1, 0, 0, 0, 67, 68,
		1, 0, 0, 0, 68, 13, 1, 0, 0, 0, 69, 70, 7, 0, 0, 0, 70, 15, 1, 0, 0, 0,
		71, 74, 3, 18, 9, 0, 72, 74, 3, 20, 10, 0, 73, 71, 1, 0, 0, 0, 73, 72,
		1, 0, 0, 0, 74, 17, 1, 0, 0, 0, 75, 78, 3, 26, 13, 0, 76, 78, 3, 22, 11,
		0, 77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 19, 1, 0, 0, 0, 79, 80,
		3, 18, 9, 0, 80, 81, 5, 4, 0, 0, 81, 82, 5, 5, 0, 0, 82, 115, 1, 0, 0,
		0, 83, 84, 3, 18, 9, 0, 84, 85, 5, 4, 0, 0, 85, 86, 5, 26, 0, 0, 86, 87,
		5, 5, 0, 0, 87, 115, 1, 0, 0, 0, 88, 89, 3, 18, 9, 0, 89, 90, 5, 4, 0,
		0, 90, 91, 5, 7, 0, 0, 91, 92, 5, 26, 0, 0, 92, 93, 5, 5, 0, 0, 93, 115,
		1, 0, 0, 0, 94, 95, 5, 23, 0, 0, 95, 96, 5, 7, 0, 0, 96, 115, 5, 26, 0,
		0, 97, 98, 5, 23, 0, 0, 98, 99, 5, 4, 0, 0, 99, 100, 5, 7, 0, 0, 100, 101,
		5, 26, 0, 0, 101, 115, 5, 5, 0, 0, 102, 103, 5, 23, 0, 0, 103, 104, 5,
		7, 0, 0, 104, 105, 5, 26, 0, 0, 105, 106, 5, 4, 0, 0, 106, 115, 5, 5, 0,
		0, 107, 108, 5, 23, 0, 0, 108, 109, 5, 7, 0, 0, 109, 110, 5, 26, 0, 0,
		110, 111, 5, 4, 0, 0, 111, 112, 5, 7, 0, 0, 112, 113, 5, 26, 0, 0, 113,
		115, 5, 5, 0, 0, 114, 79, 1, 0, 0, 0, 114, 83, 1, 0, 0, 0, 114, 88, 1,
		0, 0, 0, 114, 94, 1, 0, 0, 0, 114, 97, 1, 0, 0, 0, 114, 102, 1, 0, 0, 0,
		114, 107, 1, 0, 0, 0, 115, 21, 1, 0, 0, 0, 116, 121, 5, 25, 0, 0, 117,
		118, 5, 25, 0, 0, 118, 119, 5, 1, 0, 0, 119, 121, 3, 22, 11, 0, 120, 116,
		1, 0, 0, 0, 120, 117, 1, 0, 0, 0, 121, 23, 1, 0, 0, 0, 122, 123, 5, 25,
		0, 0, 123, 25, 1, 0, 0, 0, 124, 125, 7, 1, 0, 0, 125, 27, 1, 0, 0, 0, 9,
		33, 49, 54, 56, 67, 73, 77, 114, 120,
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

// FoxgloveSchemaParserInit initializes any static state used to implement FoxgloveSchemaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFoxgloveSchemaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FoxgloveSchemaParserInit() {
	staticData := &FoxgloveSchemaParserStaticData
	staticData.once.Do(foxgloveschemaParserInit)
}

// NewFoxgloveSchemaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFoxgloveSchemaParser(input antlr.TokenStream) *FoxgloveSchemaParser {
	FoxgloveSchemaParserInit()
	this := new(FoxgloveSchemaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FoxgloveSchemaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "FoxgloveSchema.g4"

	return this
}

// FoxgloveSchemaParser tokens.
const (
	FoxgloveSchemaParserEOF               = antlr.TokenEOF
	FoxgloveSchemaParserT__0              = 1
	FoxgloveSchemaParserT__1              = 2
	FoxgloveSchemaParserT__2              = 3
	FoxgloveSchemaParserOPEN_BRA          = 4
	FoxgloveSchemaParserCLOSE_BRA         = 5
	FoxgloveSchemaParserEQ                = 6
	FoxgloveSchemaParserLT_EQ             = 7
	FoxgloveSchemaParserSCHEMA_SPLIT      = 8
	FoxgloveSchemaParserSUB_SCHEMA_HEADER = 9
	FoxgloveSchemaParserBOOL              = 10
	FoxgloveSchemaParserBYTES             = 11
	FoxgloveSchemaParserCHAR              = 12
	FoxgloveSchemaParserFLOAT32           = 13
	FoxgloveSchemaParserFLOAT64           = 14
	FoxgloveSchemaParserINT8              = 15
	FoxgloveSchemaParserUINT8             = 16
	FoxgloveSchemaParserINT16             = 17
	FoxgloveSchemaParserUINT16            = 18
	FoxgloveSchemaParserINT32             = 19
	FoxgloveSchemaParserUINT32            = 20
	FoxgloveSchemaParserINT64             = 21
	FoxgloveSchemaParserUINT64            = 22
	FoxgloveSchemaParserSTRING            = 23
	FoxgloveSchemaParserWSTRING           = 24
	FoxgloveSchemaParserIDENTIFIER        = 25
	FoxgloveSchemaParserNUMERIC_LITERAL   = 26
	FoxgloveSchemaParserBOOL_LITERAL      = 27
	FoxgloveSchemaParserSTRING_LITERAL    = 28
	FoxgloveSchemaParserSKIP_             = 29
)

// FoxgloveSchemaParser rules.
const (
	FoxgloveSchemaParserRULE_parse       = 0
	FoxgloveSchemaParserRULE_mainSchema  = 1
	FoxgloveSchemaParserRULE_subSchema   = 2
	FoxgloveSchemaParserRULE_schemaName  = 3
	FoxgloveSchemaParserRULE_schema      = 4
	FoxgloveSchemaParserRULE_constance   = 5
	FoxgloveSchemaParserRULE_field       = 6
	FoxgloveSchemaParserRULE_expr        = 7
	FoxgloveSchemaParserRULE_fieldType   = 8
	FoxgloveSchemaParserRULE_type        = 9
	FoxgloveSchemaParserRULE_arrayType   = 10
	FoxgloveSchemaParserRULE_customType  = 11
	FoxgloveSchemaParserRULE_fieldName   = 12
	FoxgloveSchemaParserRULE_buildinType = 13
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MainSchema() IMainSchemaContext
	EOF() antlr.TerminalNode
	AllSCHEMA_SPLIT() []antlr.TerminalNode
	SCHEMA_SPLIT(i int) antlr.TerminalNode
	AllSubSchema() []ISubSchemaContext
	SubSchema(i int) ISubSchemaContext

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
	p.RuleIndex = FoxgloveSchemaParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) MainSchema() IMainSchemaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainSchemaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainSchemaContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserEOF, 0)
}

func (s *ParseContext) AllSCHEMA_SPLIT() []antlr.TerminalNode {
	return s.GetTokens(FoxgloveSchemaParserSCHEMA_SPLIT)
}

func (s *ParseContext) SCHEMA_SPLIT(i int) antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserSCHEMA_SPLIT, i)
}

func (s *ParseContext) AllSubSchema() []ISubSchemaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubSchemaContext); ok {
			len++
		}
	}

	tst := make([]ISubSchemaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubSchemaContext); ok {
			tst[i] = t.(ISubSchemaContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) SubSchema(i int) ISubSchemaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubSchemaContext); ok {
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

	return t.(ISubSchemaContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *FoxgloveSchemaParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FoxgloveSchemaParserRULE_parse)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.MainSchema()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FoxgloveSchemaParserSCHEMA_SPLIT {
		{
			p.SetState(29)
			p.Match(FoxgloveSchemaParserSCHEMA_SPLIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.SubSchema()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Match(FoxgloveSchemaParserEOF)
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

// IMainSchemaContext is an interface to support dynamic dispatch.
type IMainSchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Schema() ISchemaContext

	// IsMainSchemaContext differentiates from other interfaces.
	IsMainSchemaContext()
}

type MainSchemaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainSchemaContext() *MainSchemaContext {
	var p = new(MainSchemaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_mainSchema
	return p
}

func InitEmptyMainSchemaContext(p *MainSchemaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_mainSchema
}

func (*MainSchemaContext) IsMainSchemaContext() {}

func NewMainSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainSchemaContext {
	var p = new(MainSchemaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_mainSchema

	return p
}

func (s *MainSchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *MainSchemaContext) Schema() ISchemaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaContext)
}

func (s *MainSchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainSchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainSchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterMainSchema(s)
	}
}

func (s *MainSchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitMainSchema(s)
	}
}

func (p *FoxgloveSchemaParser) MainSchema() (localctx IMainSchemaContext) {
	localctx = NewMainSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FoxgloveSchemaParserRULE_mainSchema)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Schema()
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

// ISubSchemaContext is an interface to support dynamic dispatch.
type ISubSchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUB_SCHEMA_HEADER() antlr.TerminalNode
	SchemaName() ISchemaNameContext
	Schema() ISchemaContext

	// IsSubSchemaContext differentiates from other interfaces.
	IsSubSchemaContext()
}

type SubSchemaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubSchemaContext() *SubSchemaContext {
	var p = new(SubSchemaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_subSchema
	return p
}

func InitEmptySubSchemaContext(p *SubSchemaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_subSchema
}

func (*SubSchemaContext) IsSubSchemaContext() {}

func NewSubSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubSchemaContext {
	var p = new(SubSchemaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_subSchema

	return p
}

func (s *SubSchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *SubSchemaContext) SUB_SCHEMA_HEADER() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserSUB_SCHEMA_HEADER, 0)
}

func (s *SubSchemaContext) SchemaName() ISchemaNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaNameContext)
}

func (s *SubSchemaContext) Schema() ISchemaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaContext)
}

func (s *SubSchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubSchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubSchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterSubSchema(s)
	}
}

func (s *SubSchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitSubSchema(s)
	}
}

func (p *FoxgloveSchemaParser) SubSchema() (localctx ISubSchemaContext) {
	localctx = NewSubSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FoxgloveSchemaParserRULE_subSchema)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(FoxgloveSchemaParserSUB_SCHEMA_HEADER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.SchemaName()
	}
	{
		p.SetState(42)
		p.Schema()
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

// ISchemaNameContext is an interface to support dynamic dispatch.
type ISchemaNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsSchemaNameContext differentiates from other interfaces.
	IsSchemaNameContext()
}

type SchemaNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaNameContext() *SchemaNameContext {
	var p = new(SchemaNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_schemaName
	return p
}

func InitEmptySchemaNameContext(p *SchemaNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_schemaName
}

func (*SchemaNameContext) IsSchemaNameContext() {}

func NewSchemaNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaNameContext {
	var p = new(SchemaNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_schemaName

	return p
}

func (s *SchemaNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(FoxgloveSchemaParserIDENTIFIER)
}

func (s *SchemaNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserIDENTIFIER, i)
}

func (s *SchemaNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterSchemaName(s)
	}
}

func (s *SchemaNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitSchemaName(s)
	}
}

func (p *FoxgloveSchemaParser) SchemaName() (localctx ISchemaNameContext) {
	localctx = NewSchemaNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FoxgloveSchemaParserRULE_schemaName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(FoxgloveSchemaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FoxgloveSchemaParserT__0 {
		{
			p.SetState(45)
			p.Match(FoxgloveSchemaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(FoxgloveSchemaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ISchemaContext is an interface to support dynamic dispatch.
type ISchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllConstance() []IConstanceContext
	Constance(i int) IConstanceContext

	// IsSchemaContext differentiates from other interfaces.
	IsSchemaContext()
}

type SchemaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaContext() *SchemaContext {
	var p = new(SchemaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_schema
	return p
}

func InitEmptySchemaContext(p *SchemaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_schema
}

func (*SchemaContext) IsSchemaContext() {}

func NewSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaContext {
	var p = new(SchemaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_schema

	return p
}

func (s *SchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *SchemaContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *SchemaContext) AllConstance() []IConstanceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstanceContext); ok {
			len++
		}
	}

	tst := make([]IConstanceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstanceContext); ok {
			tst[i] = t.(IConstanceContext)
			i++
		}
	}

	return tst
}

func (s *SchemaContext) Constance(i int) IConstanceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstanceContext); ok {
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

	return t.(IConstanceContext)
}

func (s *SchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterSchema(s)
	}
}

func (s *SchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitSchema(s)
	}
}

func (p *FoxgloveSchemaParser) Schema() (localctx ISchemaContext) {
	localctx = NewSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FoxgloveSchemaParserRULE_schema)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67107840) != 0 {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(52)
				p.Field()
			}

		case 2:
			{
				p.SetState(53)
				p.Constance()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IConstanceContext is an interface to support dynamic dispatch.
type IConstanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldType() IFieldTypeContext
	IDENTIFIER() antlr.TerminalNode
	EQ() antlr.TerminalNode
	Expr() IExprContext

	// IsConstanceContext differentiates from other interfaces.
	IsConstanceContext()
}

type ConstanceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstanceContext() *ConstanceContext {
	var p = new(ConstanceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_constance
	return p
}

func InitEmptyConstanceContext(p *ConstanceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_constance
}

func (*ConstanceContext) IsConstanceContext() {}

func NewConstanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstanceContext {
	var p = new(ConstanceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_constance

	return p
}

func (s *ConstanceContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstanceContext) FieldType() IFieldTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *ConstanceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserIDENTIFIER, 0)
}

func (s *ConstanceContext) EQ() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserEQ, 0)
}

func (s *ConstanceContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterConstance(s)
	}
}

func (s *ConstanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitConstance(s)
	}
}

func (p *FoxgloveSchemaParser) Constance() (localctx IConstanceContext) {
	localctx = NewConstanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FoxgloveSchemaParserRULE_constance)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.FieldType()
	}
	{
		p.SetState(60)
		p.Match(FoxgloveSchemaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(FoxgloveSchemaParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(62)
		p.Expr()
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldType() IFieldTypeContext
	FieldName() IFieldNameContext
	Expr() IExprContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldType() IFieldTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *FoxgloveSchemaParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FoxgloveSchemaParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.FieldType()
	}
	{
		p.SetState(65)
		p.FieldName()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&335544332) != 0 {
		{
			p.SetState(66)
			p.Expr()
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserNUMERIC_LITERAL, 0)
}

func (s *ExprContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserSTRING_LITERAL, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *FoxgloveSchemaParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FoxgloveSchemaParserRULE_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&335544332) != 0) {
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

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ArrayType() IArrayTypeContext

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_fieldType
	return p
}

func InitEmptyFieldTypeContext(p *FieldTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_fieldType
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FieldTypeContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *FoxgloveSchemaParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FoxgloveSchemaParserRULE_fieldType)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Type_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.ArrayType()
		}

	case antlr.ATNInvalidAltNumber:
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BuildinType() IBuildinTypeContext
	CustomType() ICustomTypeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) BuildinType() IBuildinTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuildinTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuildinTypeContext)
}

func (s *TypeContext) CustomType() ICustomTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICustomTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICustomTypeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *FoxgloveSchemaParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FoxgloveSchemaParserRULE_type)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FoxgloveSchemaParserBOOL, FoxgloveSchemaParserBYTES, FoxgloveSchemaParserCHAR, FoxgloveSchemaParserFLOAT32, FoxgloveSchemaParserFLOAT64, FoxgloveSchemaParserINT8, FoxgloveSchemaParserUINT8, FoxgloveSchemaParserINT16, FoxgloveSchemaParserUINT16, FoxgloveSchemaParserINT32, FoxgloveSchemaParserUINT32, FoxgloveSchemaParserINT64, FoxgloveSchemaParserUINT64, FoxgloveSchemaParserSTRING, FoxgloveSchemaParserWSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.BuildinType()
		}

	case FoxgloveSchemaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.CustomType()
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

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	OPEN_BRA() antlr.TerminalNode
	CLOSE_BRA() antlr.TerminalNode
	AllNUMERIC_LITERAL() []antlr.TerminalNode
	NUMERIC_LITERAL(i int) antlr.TerminalNode
	AllLT_EQ() []antlr.TerminalNode
	LT_EQ(i int) antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArrayTypeContext) OPEN_BRA() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserOPEN_BRA, 0)
}

func (s *ArrayTypeContext) CLOSE_BRA() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserCLOSE_BRA, 0)
}

func (s *ArrayTypeContext) AllNUMERIC_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(FoxgloveSchemaParserNUMERIC_LITERAL)
}

func (s *ArrayTypeContext) NUMERIC_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserNUMERIC_LITERAL, i)
}

func (s *ArrayTypeContext) AllLT_EQ() []antlr.TerminalNode {
	return s.GetTokens(FoxgloveSchemaParserLT_EQ)
}

func (s *ArrayTypeContext) LT_EQ(i int) antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserLT_EQ, i)
}

func (s *ArrayTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserSTRING, 0)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (p *FoxgloveSchemaParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FoxgloveSchemaParserRULE_arrayType)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Type_()
		}
		{
			p.SetState(80)
			p.Match(FoxgloveSchemaParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(FoxgloveSchemaParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Type_()
		}
		{
			p.SetState(84)
			p.Match(FoxgloveSchemaParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(FoxgloveSchemaParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(FoxgloveSchemaParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Type_()
		}
		{
			p.SetState(89)
			p.Match(FoxgloveSchemaParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(FoxgloveSchemaParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(FoxgloveSchemaParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(FoxgloveSchemaParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)
			p.Match(FoxgloveSchemaParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(FoxgloveSchemaParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Match(FoxgloveSchemaParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)
			p.Match(FoxgloveSchemaParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(FoxgloveSchemaParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(FoxgloveSchemaParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(FoxgloveSchemaParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(FoxgloveSchemaParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(102)
			p.Match(FoxgloveSchemaParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(FoxgloveSchemaParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(FoxgloveSchemaParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(FoxgloveSchemaParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(FoxgloveSchemaParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(107)
			p.Match(FoxgloveSchemaParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(FoxgloveSchemaParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(FoxgloveSchemaParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(FoxgloveSchemaParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(FoxgloveSchemaParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(FoxgloveSchemaParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(FoxgloveSchemaParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// ICustomTypeContext is an interface to support dynamic dispatch.
type ICustomTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	CustomType() ICustomTypeContext

	// IsCustomTypeContext differentiates from other interfaces.
	IsCustomTypeContext()
}

type CustomTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCustomTypeContext() *CustomTypeContext {
	var p = new(CustomTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_customType
	return p
}

func InitEmptyCustomTypeContext(p *CustomTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_customType
}

func (*CustomTypeContext) IsCustomTypeContext() {}

func NewCustomTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CustomTypeContext {
	var p = new(CustomTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_customType

	return p
}

func (s *CustomTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CustomTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserIDENTIFIER, 0)
}

func (s *CustomTypeContext) CustomType() ICustomTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICustomTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICustomTypeContext)
}

func (s *CustomTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CustomTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CustomTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterCustomType(s)
	}
}

func (s *CustomTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitCustomType(s)
	}
}

func (p *FoxgloveSchemaParser) CustomType() (localctx ICustomTypeContext) {
	localctx = NewCustomTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FoxgloveSchemaParserRULE_customType)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(FoxgloveSchemaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(FoxgloveSchemaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(FoxgloveSchemaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.CustomType()
		}

	case antlr.ATNInvalidAltNumber:
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

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserIDENTIFIER, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *FoxgloveSchemaParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FoxgloveSchemaParserRULE_fieldName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(FoxgloveSchemaParserIDENTIFIER)
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

// IBuildinTypeContext is an interface to support dynamic dispatch.
type IBuildinTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	FLOAT32() antlr.TerminalNode
	FLOAT64() antlr.TerminalNode
	INT8() antlr.TerminalNode
	UINT8() antlr.TerminalNode
	INT16() antlr.TerminalNode
	UINT16() antlr.TerminalNode
	INT32() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	STRING() antlr.TerminalNode
	WSTRING() antlr.TerminalNode

	// IsBuildinTypeContext differentiates from other interfaces.
	IsBuildinTypeContext()
}

type BuildinTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildinTypeContext() *BuildinTypeContext {
	var p = new(BuildinTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_buildinType
	return p
}

func InitEmptyBuildinTypeContext(p *BuildinTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FoxgloveSchemaParserRULE_buildinType
}

func (*BuildinTypeContext) IsBuildinTypeContext() {}

func NewBuildinTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildinTypeContext {
	var p = new(BuildinTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxgloveSchemaParserRULE_buildinType

	return p
}

func (s *BuildinTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildinTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserBOOL, 0)
}

func (s *BuildinTypeContext) BYTES() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserBYTES, 0)
}

func (s *BuildinTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserCHAR, 0)
}

func (s *BuildinTypeContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserFLOAT32, 0)
}

func (s *BuildinTypeContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserFLOAT64, 0)
}

func (s *BuildinTypeContext) INT8() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserINT8, 0)
}

func (s *BuildinTypeContext) UINT8() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserUINT8, 0)
}

func (s *BuildinTypeContext) INT16() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserINT16, 0)
}

func (s *BuildinTypeContext) UINT16() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserUINT16, 0)
}

func (s *BuildinTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserINT32, 0)
}

func (s *BuildinTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserUINT32, 0)
}

func (s *BuildinTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserINT64, 0)
}

func (s *BuildinTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserUINT64, 0)
}

func (s *BuildinTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserSTRING, 0)
}

func (s *BuildinTypeContext) WSTRING() antlr.TerminalNode {
	return s.GetToken(FoxgloveSchemaParserWSTRING, 0)
}

func (s *BuildinTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildinTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildinTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.EnterBuildinType(s)
	}
}

func (s *BuildinTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxgloveSchemaListener); ok {
		listenerT.ExitBuildinType(s)
	}
}

func (p *FoxgloveSchemaParser) BuildinType() (localctx IBuildinTypeContext) {
	localctx = NewBuildinTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FoxgloveSchemaParserRULE_buildinType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33553408) != 0) {
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
