// Code generated from Rossrv.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RossrvLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var RossrvLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rossrvlexerLexerInit() {
	staticData := &RossrvLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "','", "')'", "'='", "'['", "']'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "NUMERIC_LITERAL", "STRING_LITERAL",
		"IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "NUMERIC_LITERAL",
		"STRING_LITERAL", "IDENTIFIER", "WS", "DIGIT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 99, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 4, 7, 41, 8, 7, 11, 7, 12, 7,
		42, 1, 7, 1, 7, 5, 7, 47, 8, 7, 10, 7, 12, 7, 50, 9, 7, 3, 7, 52, 8, 7,
		1, 7, 1, 7, 4, 7, 56, 8, 7, 11, 7, 12, 7, 57, 3, 7, 60, 8, 7, 1, 7, 1,
		7, 3, 7, 64, 8, 7, 1, 7, 4, 7, 67, 8, 7, 11, 7, 12, 7, 68, 3, 7, 71, 8,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 77, 8, 8, 10, 8, 12, 8, 80, 9, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 5, 9, 86, 8, 9, 10, 9, 12, 9, 89, 9, 9, 1, 10, 4, 10,
		92, 8, 10, 11, 10, 12, 10, 93, 1, 10, 1, 10, 1, 11, 1, 11, 0, 0, 12, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 0, 1, 0, 7, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 1, 0, 39,
		39, 4, 0, 65, 90, 95, 95, 97, 122, 127, 65535, 5, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 127, 65535, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 109,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 1, 25, 1, 0,
		0, 0, 3, 27, 1, 0, 0, 0, 5, 29, 1, 0, 0, 0, 7, 31, 1, 0, 0, 0, 9, 33, 1,
		0, 0, 0, 11, 35, 1, 0, 0, 0, 13, 37, 1, 0, 0, 0, 15, 59, 1, 0, 0, 0, 17,
		72, 1, 0, 0, 0, 19, 83, 1, 0, 0, 0, 21, 91, 1, 0, 0, 0, 23, 97, 1, 0, 0,
		0, 25, 26, 5, 40, 0, 0, 26, 2, 1, 0, 0, 0, 27, 28, 5, 44, 0, 0, 28, 4,
		1, 0, 0, 0, 29, 30, 5, 41, 0, 0, 30, 6, 1, 0, 0, 0, 31, 32, 5, 61, 0, 0,
		32, 8, 1, 0, 0, 0, 33, 34, 5, 91, 0, 0, 34, 10, 1, 0, 0, 0, 35, 36, 5,
		93, 0, 0, 36, 12, 1, 0, 0, 0, 37, 38, 5, 46, 0, 0, 38, 14, 1, 0, 0, 0,
		39, 41, 3, 23, 11, 0, 40, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 40, 1,
		0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 51, 1, 0, 0, 0, 44, 48, 5, 46, 0, 0, 45,
		47, 3, 23, 11, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0,
		0, 0, 48, 49, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 44,
		1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 60, 1, 0, 0, 0, 53, 55, 5, 46, 0, 0,
		54, 56, 3, 23, 11, 0, 55, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1,
		0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 40, 1, 0, 0, 0, 59,
		53, 1, 0, 0, 0, 60, 70, 1, 0, 0, 0, 61, 63, 7, 0, 0, 0, 62, 64, 7, 1, 0,
		0, 63, 62, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 67,
		3, 23, 11, 0, 66, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0,
		0, 68, 69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 61, 1, 0, 0, 0, 70, 71,
		1, 0, 0, 0, 71, 16, 1, 0, 0, 0, 72, 78, 5, 39, 0, 0, 73, 77, 8, 2, 0, 0,
		74, 75, 5, 39, 0, 0, 75, 77, 5, 39, 0, 0, 76, 73, 1, 0, 0, 0, 76, 74, 1,
		0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79,
		81, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 5, 39, 0, 0, 82, 18, 1, 0,
		0, 0, 83, 87, 7, 3, 0, 0, 84, 86, 7, 4, 0, 0, 85, 84, 1, 0, 0, 0, 86, 89,
		1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 20, 1, 0, 0, 0,
		89, 87, 1, 0, 0, 0, 90, 92, 7, 5, 0, 0, 91, 90, 1, 0, 0, 0, 92, 93, 1,
		0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95,
		96, 6, 10, 0, 0, 96, 22, 1, 0, 0, 0, 97, 98, 7, 6, 0, 0, 98, 24, 1, 0,
		0, 0, 13, 0, 42, 48, 51, 57, 59, 63, 68, 70, 76, 78, 87, 93, 1, 6, 0, 0,
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

// RossrvLexerInit initializes any static state used to implement RossrvLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRossrvLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RossrvLexerInit() {
	staticData := &RossrvLexerLexerStaticData
	staticData.once.Do(rossrvlexerLexerInit)
}

// NewRossrvLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRossrvLexer(input antlr.CharStream) *RossrvLexer {
	RossrvLexerInit()
	l := new(RossrvLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &RossrvLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Rossrv.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RossrvLexer tokens.
const (
	RossrvLexerT__0            = 1
	RossrvLexerT__1            = 2
	RossrvLexerT__2            = 3
	RossrvLexerT__3            = 4
	RossrvLexerT__4            = 5
	RossrvLexerT__5            = 6
	RossrvLexerT__6            = 7
	RossrvLexerNUMERIC_LITERAL = 8
	RossrvLexerSTRING_LITERAL  = 9
	RossrvLexerIDENTIFIER      = 10
	RossrvLexerWS              = 11
)
