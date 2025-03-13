package foxglovetopb

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func newErrListener() *errListener {
	return &errListener{}
}

type errListener struct {
	Errs []error
}

func (i *errListener) GetErrs() []error {
	return i.Errs
}

func (i *errListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	s := fmt.Sprintf("%d:%d %s", line, column, msg)
	i.Errs = append(i.Errs, errors.New(s))
}

func (i *errListener) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (i *errListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (i *errListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}
