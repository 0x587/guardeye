package listener

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func NewErrListener() *ErrListener {
	return &ErrListener{}
}

type ErrListener struct {
	Errs []string
}

func (i *ErrListener) GetErrs() []string {
	return i.Errs
}

func (i *ErrListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	s := fmt.Sprintf("%d:%d %s", line, column, msg)
	i.Errs = append(i.Errs, s)
}

func (i *ErrListener) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (i *ErrListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (i *ErrListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}
