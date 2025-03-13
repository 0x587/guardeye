package walk

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/0x587/guardeye/common/foxglovetopb/parser"
)

func Walk(tree antlr.Tree) (*StructDefine, map[string]*StructDefine) {
	l := &listener{Sub: make(map[string]*StructDefine)}
	antlr.NewParseTreeWalker().Walk(l, tree)
	return l.Main, l.Sub
}

type listener struct {
	*parser.BaseFoxgloveSchemaListener
	Main *StructDefine
	Sub  map[string]*StructDefine
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

func (l *listener) EnterParse(ctx *parser.ParseContext) {
	l.Main = l.parseStruct(ctx.MainSchema().Schema())
	for _, subSchema := range ctx.AllSubSchema() {
		l.Sub[subSchema.SchemaName().GetText()] = l.parseStruct(subSchema.Schema())
	}
}

func (l *listener) parseStruct(ctx parser.ISchemaContext) *StructDefine {
	struc := &StructDefine{}
	for _, fCtx := range ctx.AllField() {
		field := &FieldDefine{}
		field.Name = fCtx.FieldName().GetText()
		field.Type = &FieldTypeDefine{}

		var ft parser.ITypeContext

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
	return struc
}
