package walk

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/0x587/guardeye/test-client/reporter/ros/rossrv/parser"
)

func Walk(tree antlr.Tree) map[string]any {
	l := &listener{}
	antlr.NewParseTreeWalker().Walk(l, tree)
	return l.Data
}

type listener struct {
	*parser.BaseRossrvListener
	Data map[string]any
}

func (l *listener) EnterRoot(ctx *parser.RootContext) {
	l.Data = l.procObj(ctx.Obj())
}

func (l *listener) procEntry(entry parser.IEntryContext) any {
	if entry.Obj() != nil {
		return l.procObj(entry.Obj())
	}
	if entry.List() != nil {
		return l.procList(entry.List())
	}
	if entry.String_() != nil {
		return strings.TrimFunc(entry.String_().GetText(), func(r rune) bool {
			return r == '\''
		})
	}
	if entry.Integer() != nil {
		res, _ := strconv.Atoi(entry.Integer().GetText())
		return res
	}
	return nil
}

func (l *listener) procObj(obj parser.IObjContext) map[string]any {
	res := make(map[string]any)
	for _, f := range obj.AllField() {
		entry := f.Entry()
		res[f.FieldName().GetText()] = l.procEntry(entry)
	}
	return res
}

func (l *listener) procList(obj parser.IListContext) []any {
	res := make([]any, 0, len(obj.AllEntry()))
	for _, entry := range obj.AllEntry() {
		res = append(res, l.procEntry(entry))
	}
	return res
}
