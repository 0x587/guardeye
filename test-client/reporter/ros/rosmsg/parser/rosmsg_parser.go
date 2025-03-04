// Code generated from Rosmsg.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Rosmsg
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

type RosmsgParser struct {
	*antlr.BaseParser
}

var RosmsgParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rosmsgParserInit() {
	staticData := &RosmsgParserStaticData
	staticData.LiteralNames = []string{
		"", "'/'", "'['", "']'", "'='", "'<='", "'---'", "'bool'", "'bytes'",
		"'char'", "'float32'", "'float64'", "'int8'", "'uint8'", "'int16'",
		"'uint16'", "'int32'", "'uint32'", "'int64'", "'uint64'", "'string'",
		"'wstring'",
	}
	staticData.SymbolicNames = []string{
		"", "", "OPEN_BRA", "CLOSE_BRA", "EQ", "LT_EQ", "MSG_SPLIT", "BOOL",
		"BYTES", "CHAR", "FLOAT32", "FLOAT64", "INT8", "UINT8", "INT16", "UINT16",
		"INT32", "UINT32", "INT64", "UINT64", "STRING", "WSTRING", "IDENTIFIER",
		"NUMERIC_LITERAL", "SINGLE_LINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
	}
	staticData.RuleNames = []string{
		"parse", "msgStat", "srvStat", "field", "fieldType", "type", "arrayType",
		"customType", "fieldName", "buildinType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 93, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 3, 0, 23, 8, 0, 1, 1, 5, 1, 26, 8, 1, 10, 1, 12, 1, 29, 9, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 40, 8, 4, 1, 5, 1,
		5, 3, 5, 44, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 81, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 87, 8, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		0, 1, 1, 0, 7, 21, 93, 0, 22, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 30, 1,
		0, 0, 0, 6, 34, 1, 0, 0, 0, 8, 39, 1, 0, 0, 0, 10, 43, 1, 0, 0, 0, 12,
		80, 1, 0, 0, 0, 14, 86, 1, 0, 0, 0, 16, 88, 1, 0, 0, 0, 18, 90, 1, 0, 0,
		0, 20, 23, 3, 2, 1, 0, 21, 23, 3, 4, 2, 0, 22, 20, 1, 0, 0, 0, 22, 21,
		1, 0, 0, 0, 23, 1, 1, 0, 0, 0, 24, 26, 3, 6, 3, 0, 25, 24, 1, 0, 0, 0,
		26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 3, 1, 0,
		0, 0, 29, 27, 1, 0, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 6, 0, 0, 32, 33,
		3, 2, 1, 0, 33, 5, 1, 0, 0, 0, 34, 35, 3, 8, 4, 0, 35, 36, 3, 16, 8, 0,
		36, 7, 1, 0, 0, 0, 37, 40, 3, 10, 5, 0, 38, 40, 3, 12, 6, 0, 39, 37, 1,
		0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 9, 1, 0, 0, 0, 41, 44, 3, 18, 9, 0, 42,
		44, 3, 14, 7, 0, 43, 41, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 11, 1, 0,
		0, 0, 45, 46, 3, 10, 5, 0, 46, 47, 5, 2, 0, 0, 47, 48, 5, 3, 0, 0, 48,
		81, 1, 0, 0, 0, 49, 50, 3, 10, 5, 0, 50, 51, 5, 2, 0, 0, 51, 52, 5, 23,
		0, 0, 52, 53, 5, 3, 0, 0, 53, 81, 1, 0, 0, 0, 54, 55, 3, 10, 5, 0, 55,
		56, 5, 2, 0, 0, 56, 57, 5, 5, 0, 0, 57, 58, 5, 23, 0, 0, 58, 59, 5, 3,
		0, 0, 59, 81, 1, 0, 0, 0, 60, 61, 5, 20, 0, 0, 61, 62, 5, 5, 0, 0, 62,
		81, 5, 23, 0, 0, 63, 64, 5, 20, 0, 0, 64, 65, 5, 2, 0, 0, 65, 66, 5, 5,
		0, 0, 66, 67, 5, 23, 0, 0, 67, 81, 5, 3, 0, 0, 68, 69, 5, 20, 0, 0, 69,
		70, 5, 5, 0, 0, 70, 71, 5, 23, 0, 0, 71, 72, 5, 2, 0, 0, 72, 81, 5, 3,
		0, 0, 73, 74, 5, 20, 0, 0, 74, 75, 5, 5, 0, 0, 75, 76, 5, 23, 0, 0, 76,
		77, 5, 2, 0, 0, 77, 78, 5, 5, 0, 0, 78, 79, 5, 23, 0, 0, 79, 81, 5, 3,
		0, 0, 80, 45, 1, 0, 0, 0, 80, 49, 1, 0, 0, 0, 80, 54, 1, 0, 0, 0, 80, 60,
		1, 0, 0, 0, 80, 63, 1, 0, 0, 0, 80, 68, 1, 0, 0, 0, 80, 73, 1, 0, 0, 0,
		81, 13, 1, 0, 0, 0, 82, 87, 5, 22, 0, 0, 83, 84, 5, 22, 0, 0, 84, 85, 5,
		1, 0, 0, 85, 87, 3, 14, 7, 0, 86, 82, 1, 0, 0, 0, 86, 83, 1, 0, 0, 0, 87,
		15, 1, 0, 0, 0, 88, 89, 5, 22, 0, 0, 89, 17, 1, 0, 0, 0, 90, 91, 7, 0,
		0, 0, 91, 19, 1, 0, 0, 0, 6, 22, 27, 39, 43, 80, 86,
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

// RosmsgParserInit initializes any static state used to implement RosmsgParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRosmsgParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RosmsgParserInit() {
	staticData := &RosmsgParserStaticData
	staticData.once.Do(rosmsgParserInit)
}

// NewRosmsgParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRosmsgParser(input antlr.TokenStream) *RosmsgParser {
	RosmsgParserInit()
	this := new(RosmsgParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RosmsgParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Rosmsg.g4"

	return this
}

// RosmsgParser tokens.
const (
	RosmsgParserEOF                 = antlr.TokenEOF
	RosmsgParserT__0                = 1
	RosmsgParserOPEN_BRA            = 2
	RosmsgParserCLOSE_BRA           = 3
	RosmsgParserEQ                  = 4
	RosmsgParserLT_EQ               = 5
	RosmsgParserMSG_SPLIT           = 6
	RosmsgParserBOOL                = 7
	RosmsgParserBYTES               = 8
	RosmsgParserCHAR                = 9
	RosmsgParserFLOAT32             = 10
	RosmsgParserFLOAT64             = 11
	RosmsgParserINT8                = 12
	RosmsgParserUINT8               = 13
	RosmsgParserINT16               = 14
	RosmsgParserUINT16              = 15
	RosmsgParserINT32               = 16
	RosmsgParserUINT32              = 17
	RosmsgParserINT64               = 18
	RosmsgParserUINT64              = 19
	RosmsgParserSTRING              = 20
	RosmsgParserWSTRING             = 21
	RosmsgParserIDENTIFIER          = 22
	RosmsgParserNUMERIC_LITERAL     = 23
	RosmsgParserSINGLE_LINE_COMMENT = 24
	RosmsgParserSPACES              = 25
	RosmsgParserUNEXPECTED_CHAR     = 26
)

// RosmsgParser rules.
const (
	RosmsgParserRULE_parse       = 0
	RosmsgParserRULE_msgStat     = 1
	RosmsgParserRULE_srvStat     = 2
	RosmsgParserRULE_field       = 3
	RosmsgParserRULE_fieldType   = 4
	RosmsgParserRULE_type        = 5
	RosmsgParserRULE_arrayType   = 6
	RosmsgParserRULE_customType  = 7
	RosmsgParserRULE_fieldName   = 8
	RosmsgParserRULE_buildinType = 9
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MsgStat() IMsgStatContext
	SrvStat() ISrvStatContext

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
	p.RuleIndex = RosmsgParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) MsgStat() IMsgStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMsgStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMsgStatContext)
}

func (s *ParseContext) SrvStat() ISrvStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISrvStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISrvStatContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *RosmsgParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RosmsgParserRULE_parse)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.MsgStat()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.SrvStat()
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

// IMsgStatContext is an interface to support dynamic dispatch.
type IMsgStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsMsgStatContext differentiates from other interfaces.
	IsMsgStatContext()
}

type MsgStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgStatContext() *MsgStatContext {
	var p = new(MsgStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_msgStat
	return p
}

func InitEmptyMsgStatContext(p *MsgStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_msgStat
}

func (*MsgStatContext) IsMsgStatContext() {}

func NewMsgStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgStatContext {
	var p = new(MsgStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_msgStat

	return p
}

func (s *MsgStatContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgStatContext) AllField() []IFieldContext {
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

func (s *MsgStatContext) Field(i int) IFieldContext {
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

func (s *MsgStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterMsgStat(s)
	}
}

func (s *MsgStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitMsgStat(s)
	}
}

func (p *RosmsgParser) MsgStat() (localctx IMsgStatContext) {
	localctx = NewMsgStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RosmsgParserRULE_msgStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8388480) != 0 {
		{
			p.SetState(24)
			p.Field()
		}

		p.SetState(29)
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

// ISrvStatContext is an interface to support dynamic dispatch.
type ISrvStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMsgStat() []IMsgStatContext
	MsgStat(i int) IMsgStatContext
	MSG_SPLIT() antlr.TerminalNode

	// IsSrvStatContext differentiates from other interfaces.
	IsSrvStatContext()
}

type SrvStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySrvStatContext() *SrvStatContext {
	var p = new(SrvStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_srvStat
	return p
}

func InitEmptySrvStatContext(p *SrvStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_srvStat
}

func (*SrvStatContext) IsSrvStatContext() {}

func NewSrvStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SrvStatContext {
	var p = new(SrvStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_srvStat

	return p
}

func (s *SrvStatContext) GetParser() antlr.Parser { return s.parser }

func (s *SrvStatContext) AllMsgStat() []IMsgStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMsgStatContext); ok {
			len++
		}
	}

	tst := make([]IMsgStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMsgStatContext); ok {
			tst[i] = t.(IMsgStatContext)
			i++
		}
	}

	return tst
}

func (s *SrvStatContext) MsgStat(i int) IMsgStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMsgStatContext); ok {
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

	return t.(IMsgStatContext)
}

func (s *SrvStatContext) MSG_SPLIT() antlr.TerminalNode {
	return s.GetToken(RosmsgParserMSG_SPLIT, 0)
}

func (s *SrvStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrvStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SrvStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterSrvStat(s)
	}
}

func (s *SrvStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitSrvStat(s)
	}
}

func (p *RosmsgParser) SrvStat() (localctx ISrvStatContext) {
	localctx = NewSrvStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RosmsgParserRULE_srvStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.MsgStat()
	}
	{
		p.SetState(31)
		p.Match(RosmsgParserMSG_SPLIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.MsgStat()
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
	p.RuleIndex = RosmsgParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_field

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

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *RosmsgParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RosmsgParserRULE_field)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.FieldType()
	}
	{
		p.SetState(35)
		p.FieldName()
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
	p.RuleIndex = RosmsgParserRULE_fieldType
	return p
}

func InitEmptyFieldTypeContext(p *FieldTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_fieldType
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_fieldType

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
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *RosmsgParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RosmsgParserRULE_fieldType)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Type_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
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
	p.RuleIndex = RosmsgParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_type

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
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *RosmsgParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RosmsgParserRULE_type)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RosmsgParserBOOL, RosmsgParserBYTES, RosmsgParserCHAR, RosmsgParserFLOAT32, RosmsgParserFLOAT64, RosmsgParserINT8, RosmsgParserUINT8, RosmsgParserINT16, RosmsgParserUINT16, RosmsgParserINT32, RosmsgParserUINT32, RosmsgParserINT64, RosmsgParserUINT64, RosmsgParserSTRING, RosmsgParserWSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.BuildinType()
		}

	case RosmsgParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
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
	p.RuleIndex = RosmsgParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_arrayType

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
	return s.GetToken(RosmsgParserOPEN_BRA, 0)
}

func (s *ArrayTypeContext) CLOSE_BRA() antlr.TerminalNode {
	return s.GetToken(RosmsgParserCLOSE_BRA, 0)
}

func (s *ArrayTypeContext) AllNUMERIC_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(RosmsgParserNUMERIC_LITERAL)
}

func (s *ArrayTypeContext) NUMERIC_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(RosmsgParserNUMERIC_LITERAL, i)
}

func (s *ArrayTypeContext) AllLT_EQ() []antlr.TerminalNode {
	return s.GetTokens(RosmsgParserLT_EQ)
}

func (s *ArrayTypeContext) LT_EQ(i int) antlr.TerminalNode {
	return s.GetToken(RosmsgParserLT_EQ, i)
}

func (s *ArrayTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(RosmsgParserSTRING, 0)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (p *RosmsgParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RosmsgParserRULE_arrayType)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Type_()
		}
		{
			p.SetState(46)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Type_()
		}
		{
			p.SetState(50)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Type_()
		}
		{
			p.SetState(55)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(60)
			p.Match(RosmsgParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Match(RosmsgParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(68)
			p.Match(RosmsgParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(73)
			p.Match(RosmsgParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(RosmsgParserCLOSE_BRA)
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
	p.RuleIndex = RosmsgParserRULE_customType
	return p
}

func InitEmptyCustomTypeContext(p *CustomTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_customType
}

func (*CustomTypeContext) IsCustomTypeContext() {}

func NewCustomTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CustomTypeContext {
	var p = new(CustomTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_customType

	return p
}

func (s *CustomTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CustomTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RosmsgParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterCustomType(s)
	}
}

func (s *CustomTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitCustomType(s)
	}
}

func (p *RosmsgParser) CustomType() (localctx ICustomTypeContext) {
	localctx = NewCustomTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RosmsgParserRULE_customType)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(RosmsgParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(RosmsgParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(RosmsgParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
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
	p.RuleIndex = RosmsgParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RosmsgParserIDENTIFIER, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *RosmsgParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RosmsgParserRULE_fieldName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(RosmsgParserIDENTIFIER)
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
	p.RuleIndex = RosmsgParserRULE_buildinType
	return p
}

func InitEmptyBuildinTypeContext(p *BuildinTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_buildinType
}

func (*BuildinTypeContext) IsBuildinTypeContext() {}

func NewBuildinTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuildinTypeContext {
	var p = new(BuildinTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_buildinType

	return p
}

func (s *BuildinTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuildinTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(RosmsgParserBOOL, 0)
}

func (s *BuildinTypeContext) BYTES() antlr.TerminalNode {
	return s.GetToken(RosmsgParserBYTES, 0)
}

func (s *BuildinTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(RosmsgParserCHAR, 0)
}

func (s *BuildinTypeContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(RosmsgParserFLOAT32, 0)
}

func (s *BuildinTypeContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(RosmsgParserFLOAT64, 0)
}

func (s *BuildinTypeContext) INT8() antlr.TerminalNode {
	return s.GetToken(RosmsgParserINT8, 0)
}

func (s *BuildinTypeContext) UINT8() antlr.TerminalNode {
	return s.GetToken(RosmsgParserUINT8, 0)
}

func (s *BuildinTypeContext) INT16() antlr.TerminalNode {
	return s.GetToken(RosmsgParserINT16, 0)
}

func (s *BuildinTypeContext) UINT16() antlr.TerminalNode {
	return s.GetToken(RosmsgParserUINT16, 0)
}

func (s *BuildinTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(RosmsgParserINT32, 0)
}

func (s *BuildinTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(RosmsgParserUINT32, 0)
}

func (s *BuildinTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(RosmsgParserINT64, 0)
}

func (s *BuildinTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(RosmsgParserUINT64, 0)
}

func (s *BuildinTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(RosmsgParserSTRING, 0)
}

func (s *BuildinTypeContext) WSTRING() antlr.TerminalNode {
	return s.GetToken(RosmsgParserWSTRING, 0)
}

func (s *BuildinTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildinTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuildinTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterBuildinType(s)
	}
}

func (s *BuildinTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitBuildinType(s)
	}
}

func (p *RosmsgParser) BuildinType() (localctx IBuildinTypeContext) {
	localctx = NewBuildinTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RosmsgParserRULE_buildinType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4194176) != 0) {
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
