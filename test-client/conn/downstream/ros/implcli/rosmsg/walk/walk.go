package walk

import (
	"github.com/antlr4-go/antlr/v4"

	parser2 "github.com/0x587/guardeye/test-client/conn/downstream/ros/implcli/rosmsg/parser"
)

func Walk(tree antlr.Tree) []*StructDefine {
	l := &listener{}
	antlr.NewParseTreeWalker().Walk(l, tree)
	return l.Structs
}

type listener struct {
	*parser2.BaseRosmsgListener
	Structs []*StructDefine
}

type StructDefine struct {
	Fields []*FieldDefine
}

type FieldDefine struct {
	Name string
	Type *FieldTypeDefine
}

type FieldTypeDefine struct {
	IsArray  bool
	Name     string
	IsCustom bool
}

func (l *listener) EnterMsgStat(ctx *parser2.MsgStatContext) {
	struc := &StructDefine{}
	for _, fCtx := range ctx.AllField() {
		field := &FieldDefine{}
		field.Name = fCtx.FieldName().GetText()
		field.Type = &FieldTypeDefine{}

		var ft parser2.ITypeContext

		if fCtx.FieldType().Type_() != nil {
			ft = fCtx.FieldType().Type_()
		}
		if fCtx.FieldType().ArrayType() != nil {
			field.Type.IsArray = true
			ft = fCtx.FieldType().ArrayType().Type_()
		}

		if ft != nil {
			field.Type.Name = ft.GetText()
			field.Type.IsCustom = ft.CustomType() != nil
		} else {
			field.Type.Name = fCtx.FieldType().ArrayType().STRING().GetText()
		}
		struc.Fields = append(struc.Fields, field)
	}
	l.Structs = append(l.Structs, struc)
}
