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
		"parse", "msg_stat", "srv_stat", "field", "field_type", "type", "customed_type",
		"field_name", "buildin_type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 3, 0, 21,
		8, 0, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 72, 8, 4, 1, 5, 1, 5, 3, 5, 76, 8, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 82, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 0, 0, 9,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 1, 1, 0, 7, 21, 90, 0, 20, 1, 0, 0, 0,
		2, 25, 1, 0, 0, 0, 4, 28, 1, 0, 0, 0, 6, 31, 1, 0, 0, 0, 8, 71, 1, 0, 0,
		0, 10, 75, 1, 0, 0, 0, 12, 81, 1, 0, 0, 0, 14, 83, 1, 0, 0, 0, 16, 85,
		1, 0, 0, 0, 18, 21, 3, 2, 1, 0, 19, 21, 3, 4, 2, 0, 20, 18, 1, 0, 0, 0,
		20, 19, 1, 0, 0, 0, 21, 1, 1, 0, 0, 0, 22, 24, 3, 6, 3, 0, 23, 22, 1, 0,
		0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 3,
		1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 3, 2, 1, 0, 29, 30, 5, 6, 0, 0,
		30, 5, 1, 0, 0, 0, 31, 32, 3, 8, 4, 0, 32, 33, 3, 14, 7, 0, 33, 7, 1, 0,
		0, 0, 34, 72, 3, 10, 5, 0, 35, 36, 3, 10, 5, 0, 36, 37, 5, 2, 0, 0, 37,
		38, 5, 3, 0, 0, 38, 72, 1, 0, 0, 0, 39, 40, 3, 10, 5, 0, 40, 41, 5, 2,
		0, 0, 41, 42, 5, 23, 0, 0, 42, 43, 5, 3, 0, 0, 43, 72, 1, 0, 0, 0, 44,
		45, 3, 10, 5, 0, 45, 46, 5, 2, 0, 0, 46, 47, 5, 5, 0, 0, 47, 48, 5, 23,
		0, 0, 48, 49, 5, 3, 0, 0, 49, 72, 1, 0, 0, 0, 50, 72, 5, 20, 0, 0, 51,
		52, 5, 20, 0, 0, 52, 53, 5, 5, 0, 0, 53, 72, 5, 23, 0, 0, 54, 55, 5, 20,
		0, 0, 55, 56, 5, 2, 0, 0, 56, 57, 5, 5, 0, 0, 57, 58, 5, 23, 0, 0, 58,
		72, 5, 3, 0, 0, 59, 60, 5, 20, 0, 0, 60, 61, 5, 5, 0, 0, 61, 62, 5, 23,
		0, 0, 62, 63, 5, 2, 0, 0, 63, 72, 5, 3, 0, 0, 64, 65, 5, 20, 0, 0, 65,
		66, 5, 5, 0, 0, 66, 67, 5, 23, 0, 0, 67, 68, 5, 2, 0, 0, 68, 69, 5, 5,
		0, 0, 69, 70, 5, 23, 0, 0, 70, 72, 5, 3, 0, 0, 71, 34, 1, 0, 0, 0, 71,
		35, 1, 0, 0, 0, 71, 39, 1, 0, 0, 0, 71, 44, 1, 0, 0, 0, 71, 50, 1, 0, 0,
		0, 71, 51, 1, 0, 0, 0, 71, 54, 1, 0, 0, 0, 71, 59, 1, 0, 0, 0, 71, 64,
		1, 0, 0, 0, 72, 9, 1, 0, 0, 0, 73, 76, 3, 16, 8, 0, 74, 76, 3, 12, 6, 0,
		75, 73, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0, 76, 11, 1, 0, 0, 0, 77, 82, 5,
		22, 0, 0, 78, 79, 5, 22, 0, 0, 79, 80, 5, 1, 0, 0, 80, 82, 3, 12, 6, 0,
		81, 77, 1, 0, 0, 0, 81, 78, 1, 0, 0, 0, 82, 13, 1, 0, 0, 0, 83, 84, 5,
		22, 0, 0, 84, 15, 1, 0, 0, 0, 85, 86, 7, 0, 0, 0, 86, 17, 1, 0, 0, 0, 5,
		20, 25, 71, 75, 81,
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
	RosmsgParserRULE_parse         = 0
	RosmsgParserRULE_msg_stat      = 1
	RosmsgParserRULE_srv_stat      = 2
	RosmsgParserRULE_field         = 3
	RosmsgParserRULE_field_type    = 4
	RosmsgParserRULE_type          = 5
	RosmsgParserRULE_customed_type = 6
	RosmsgParserRULE_field_name    = 7
	RosmsgParserRULE_buildin_type  = 8
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Msg_stat() IMsg_statContext
	Srv_stat() ISrv_statContext

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

func (s *ParseContext) Msg_stat() IMsg_statContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMsg_statContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMsg_statContext)
}

func (s *ParseContext) Srv_stat() ISrv_statContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISrv_statContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISrv_statContext)
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
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Msg_stat()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Srv_stat()
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

// IMsg_statContext is an interface to support dynamic dispatch.
type IMsg_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsMsg_statContext differentiates from other interfaces.
	IsMsg_statContext()
}

type Msg_statContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsg_statContext() *Msg_statContext {
	var p = new(Msg_statContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_msg_stat
	return p
}

func InitEmptyMsg_statContext(p *Msg_statContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_msg_stat
}

func (*Msg_statContext) IsMsg_statContext() {}

func NewMsg_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Msg_statContext {
	var p = new(Msg_statContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_msg_stat

	return p
}

func (s *Msg_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Msg_statContext) AllField() []IFieldContext {
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

func (s *Msg_statContext) Field(i int) IFieldContext {
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

func (s *Msg_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Msg_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Msg_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterMsg_stat(s)
	}
}

func (s *Msg_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitMsg_stat(s)
	}
}

func (p *RosmsgParser) Msg_stat() (localctx IMsg_statContext) {
	localctx = NewMsg_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RosmsgParserRULE_msg_stat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8388480) != 0 {
		{
			p.SetState(22)
			p.Field()
		}

		p.SetState(27)
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

// ISrv_statContext is an interface to support dynamic dispatch.
type ISrv_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Msg_stat() IMsg_statContext
	MSG_SPLIT() antlr.TerminalNode

	// IsSrv_statContext differentiates from other interfaces.
	IsSrv_statContext()
}

type Srv_statContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySrv_statContext() *Srv_statContext {
	var p = new(Srv_statContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_srv_stat
	return p
}

func InitEmptySrv_statContext(p *Srv_statContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_srv_stat
}

func (*Srv_statContext) IsSrv_statContext() {}

func NewSrv_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Srv_statContext {
	var p = new(Srv_statContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_srv_stat

	return p
}

func (s *Srv_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Srv_statContext) Msg_stat() IMsg_statContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMsg_statContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMsg_statContext)
}

func (s *Srv_statContext) MSG_SPLIT() antlr.TerminalNode {
	return s.GetToken(RosmsgParserMSG_SPLIT, 0)
}

func (s *Srv_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Srv_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Srv_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterSrv_stat(s)
	}
}

func (s *Srv_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitSrv_stat(s)
	}
}

func (p *RosmsgParser) Srv_stat() (localctx ISrv_statContext) {
	localctx = NewSrv_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RosmsgParserRULE_srv_stat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Msg_stat()
	}
	{
		p.SetState(29)
		p.Match(RosmsgParserMSG_SPLIT)
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext
	Field_name() IField_nameContext

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

func (s *FieldContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *FieldContext) Field_name() IField_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
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
		p.SetState(31)
		p.Field_type()
	}
	{
		p.SetState(32)
		p.Field_name()
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

// IField_typeContext is an interface to support dynamic dispatch.
type IField_typeContext interface {
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

	// IsField_typeContext differentiates from other interfaces.
	IsField_typeContext()
}

type Field_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_typeContext() *Field_typeContext {
	var p = new(Field_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_field_type
	return p
}

func InitEmptyField_typeContext(p *Field_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_field_type
}

func (*Field_typeContext) IsField_typeContext() {}

func NewField_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_typeContext {
	var p = new(Field_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_field_type

	return p
}

func (s *Field_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_typeContext) Type_() ITypeContext {
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

func (s *Field_typeContext) OPEN_BRA() antlr.TerminalNode {
	return s.GetToken(RosmsgParserOPEN_BRA, 0)
}

func (s *Field_typeContext) CLOSE_BRA() antlr.TerminalNode {
	return s.GetToken(RosmsgParserCLOSE_BRA, 0)
}

func (s *Field_typeContext) AllNUMERIC_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(RosmsgParserNUMERIC_LITERAL)
}

func (s *Field_typeContext) NUMERIC_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(RosmsgParserNUMERIC_LITERAL, i)
}

func (s *Field_typeContext) AllLT_EQ() []antlr.TerminalNode {
	return s.GetTokens(RosmsgParserLT_EQ)
}

func (s *Field_typeContext) LT_EQ(i int) antlr.TerminalNode {
	return s.GetToken(RosmsgParserLT_EQ, i)
}

func (s *Field_typeContext) STRING() antlr.TerminalNode {
	return s.GetToken(RosmsgParserSTRING, 0)
}

func (s *Field_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterField_type(s)
	}
}

func (s *Field_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitField_type(s)
	}
}

func (p *RosmsgParser) Field_type() (localctx IField_typeContext) {
	localctx = NewField_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RosmsgParserRULE_field_type)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Type_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Type_()
		}
		{
			p.SetState(36)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Type_()
		}
		{
			p.SetState(40)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(44)
			p.Type_()
		}
		{
			p.SetState(45)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.Match(RosmsgParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(51)
			p.Match(RosmsgParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(54)
			p.Match(RosmsgParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(59)
			p.Match(RosmsgParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(RosmsgParserCLOSE_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(64)
			p.Match(RosmsgParserSTRING)
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
			p.Match(RosmsgParserOPEN_BRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(RosmsgParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(RosmsgParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Buildin_type() IBuildin_typeContext
	Customed_type() ICustomed_typeContext

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

func (s *TypeContext) Buildin_type() IBuildin_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuildin_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuildin_typeContext)
}

func (s *TypeContext) Customed_type() ICustomed_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICustomed_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICustomed_typeContext)
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RosmsgParserBOOL, RosmsgParserBYTES, RosmsgParserCHAR, RosmsgParserFLOAT32, RosmsgParserFLOAT64, RosmsgParserINT8, RosmsgParserUINT8, RosmsgParserINT16, RosmsgParserUINT16, RosmsgParserINT32, RosmsgParserUINT32, RosmsgParserINT64, RosmsgParserUINT64, RosmsgParserSTRING, RosmsgParserWSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Buildin_type()
		}

	case RosmsgParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Customed_type()
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

// ICustomed_typeContext is an interface to support dynamic dispatch.
type ICustomed_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Customed_type() ICustomed_typeContext

	// IsCustomed_typeContext differentiates from other interfaces.
	IsCustomed_typeContext()
}

type Customed_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCustomed_typeContext() *Customed_typeContext {
	var p = new(Customed_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_customed_type
	return p
}

func InitEmptyCustomed_typeContext(p *Customed_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_customed_type
}

func (*Customed_typeContext) IsCustomed_typeContext() {}

func NewCustomed_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Customed_typeContext {
	var p = new(Customed_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_customed_type

	return p
}

func (s *Customed_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Customed_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RosmsgParserIDENTIFIER, 0)
}

func (s *Customed_typeContext) Customed_type() ICustomed_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICustomed_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICustomed_typeContext)
}

func (s *Customed_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Customed_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Customed_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterCustomed_type(s)
	}
}

func (s *Customed_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitCustomed_type(s)
	}
}

func (p *RosmsgParser) Customed_type() (localctx ICustomed_typeContext) {
	localctx = NewCustomed_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RosmsgParserRULE_customed_type)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(RosmsgParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Match(RosmsgParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(RosmsgParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Customed_type()
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

// IField_nameContext is an interface to support dynamic dispatch.
type IField_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsField_nameContext differentiates from other interfaces.
	IsField_nameContext()
}

type Field_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_nameContext() *Field_nameContext {
	var p = new(Field_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_field_name
	return p
}

func InitEmptyField_nameContext(p *Field_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_field_name
}

func (*Field_nameContext) IsField_nameContext() {}

func NewField_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_nameContext {
	var p = new(Field_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_field_name

	return p
}

func (s *Field_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RosmsgParserIDENTIFIER, 0)
}

func (s *Field_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterField_name(s)
	}
}

func (s *Field_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitField_name(s)
	}
}

func (p *RosmsgParser) Field_name() (localctx IField_nameContext) {
	localctx = NewField_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RosmsgParserRULE_field_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
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

// IBuildin_typeContext is an interface to support dynamic dispatch.
type IBuildin_typeContext interface {
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

	// IsBuildin_typeContext differentiates from other interfaces.
	IsBuildin_typeContext()
}

type Buildin_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuildin_typeContext() *Buildin_typeContext {
	var p = new(Buildin_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_buildin_type
	return p
}

func InitEmptyBuildin_typeContext(p *Buildin_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RosmsgParserRULE_buildin_type
}

func (*Buildin_typeContext) IsBuildin_typeContext() {}

func NewBuildin_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Buildin_typeContext {
	var p = new(Buildin_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RosmsgParserRULE_buildin_type

	return p
}

func (s *Buildin_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Buildin_typeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(RosmsgParserBOOL, 0)
}

func (s *Buildin_typeContext) BYTES() antlr.TerminalNode {
	return s.GetToken(RosmsgParserBYTES, 0)
}

func (s *Buildin_typeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(RosmsgParserCHAR, 0)
}

func (s *Buildin_typeContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(RosmsgParserFLOAT32, 0)
}

func (s *Buildin_typeContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(RosmsgParserFLOAT64, 0)
}

func (s *Buildin_typeContext) INT8() antlr.TerminalNode {
	return s.GetToken(RosmsgParserINT8, 0)
}

func (s *Buildin_typeContext) UINT8() antlr.TerminalNode {
	return s.GetToken(RosmsgParserUINT8, 0)
}

func (s *Buildin_typeContext) INT16() antlr.TerminalNode {
	return s.GetToken(RosmsgParserINT16, 0)
}

func (s *Buildin_typeContext) UINT16() antlr.TerminalNode {
	return s.GetToken(RosmsgParserUINT16, 0)
}

func (s *Buildin_typeContext) INT32() antlr.TerminalNode {
	return s.GetToken(RosmsgParserINT32, 0)
}

func (s *Buildin_typeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(RosmsgParserUINT32, 0)
}

func (s *Buildin_typeContext) INT64() antlr.TerminalNode {
	return s.GetToken(RosmsgParserINT64, 0)
}

func (s *Buildin_typeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(RosmsgParserUINT64, 0)
}

func (s *Buildin_typeContext) STRING() antlr.TerminalNode {
	return s.GetToken(RosmsgParserSTRING, 0)
}

func (s *Buildin_typeContext) WSTRING() antlr.TerminalNode {
	return s.GetToken(RosmsgParserWSTRING, 0)
}

func (s *Buildin_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Buildin_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Buildin_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.EnterBuildin_type(s)
	}
}

func (s *Buildin_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RosmsgListener); ok {
		listenerT.ExitBuildin_type(s)
	}
}

func (p *RosmsgParser) Buildin_type() (localctx IBuildin_typeContext) {
	localctx = NewBuildin_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RosmsgParserRULE_buildin_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
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
