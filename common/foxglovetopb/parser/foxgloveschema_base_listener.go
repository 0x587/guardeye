// Code generated from FoxgloveSchema.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FoxgloveSchema
import "github.com/antlr4-go/antlr/v4"

// BaseFoxgloveSchemaListener is a complete listener for a parse tree produced by FoxgloveSchemaParser.
type BaseFoxgloveSchemaListener struct{}

var _ FoxgloveSchemaListener = &BaseFoxgloveSchemaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFoxgloveSchemaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFoxgloveSchemaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFoxgloveSchemaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFoxgloveSchemaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseFoxgloveSchemaListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseFoxgloveSchemaListener) ExitParse(ctx *ParseContext) {}

// EnterMainSchema is called when production mainSchema is entered.
func (s *BaseFoxgloveSchemaListener) EnterMainSchema(ctx *MainSchemaContext) {}

// ExitMainSchema is called when production mainSchema is exited.
func (s *BaseFoxgloveSchemaListener) ExitMainSchema(ctx *MainSchemaContext) {}

// EnterSubSchema is called when production subSchema is entered.
func (s *BaseFoxgloveSchemaListener) EnterSubSchema(ctx *SubSchemaContext) {}

// ExitSubSchema is called when production subSchema is exited.
func (s *BaseFoxgloveSchemaListener) ExitSubSchema(ctx *SubSchemaContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseFoxgloveSchemaListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseFoxgloveSchemaListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterSchema is called when production schema is entered.
func (s *BaseFoxgloveSchemaListener) EnterSchema(ctx *SchemaContext) {}

// ExitSchema is called when production schema is exited.
func (s *BaseFoxgloveSchemaListener) ExitSchema(ctx *SchemaContext) {}

// EnterConstance is called when production constance is entered.
func (s *BaseFoxgloveSchemaListener) EnterConstance(ctx *ConstanceContext) {}

// ExitConstance is called when production constance is exited.
func (s *BaseFoxgloveSchemaListener) ExitConstance(ctx *ConstanceContext) {}

// EnterField is called when production field is entered.
func (s *BaseFoxgloveSchemaListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseFoxgloveSchemaListener) ExitField(ctx *FieldContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseFoxgloveSchemaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseFoxgloveSchemaListener) ExitExpr(ctx *ExprContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseFoxgloveSchemaListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseFoxgloveSchemaListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterType is called when production type is entered.
func (s *BaseFoxgloveSchemaListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseFoxgloveSchemaListener) ExitType(ctx *TypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseFoxgloveSchemaListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseFoxgloveSchemaListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterCustomType is called when production customType is entered.
func (s *BaseFoxgloveSchemaListener) EnterCustomType(ctx *CustomTypeContext) {}

// ExitCustomType is called when production customType is exited.
func (s *BaseFoxgloveSchemaListener) ExitCustomType(ctx *CustomTypeContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseFoxgloveSchemaListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseFoxgloveSchemaListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterBuildinType is called when production buildinType is entered.
func (s *BaseFoxgloveSchemaListener) EnterBuildinType(ctx *BuildinTypeContext) {}

// ExitBuildinType is called when production buildinType is exited.
func (s *BaseFoxgloveSchemaListener) ExitBuildinType(ctx *BuildinTypeContext) {}
