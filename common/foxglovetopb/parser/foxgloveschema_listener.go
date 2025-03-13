// Code generated from FoxgloveSchema.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FoxgloveSchema
import "github.com/antlr4-go/antlr/v4"

// FoxgloveSchemaListener is a complete listener for a parse tree produced by FoxgloveSchemaParser.
type FoxgloveSchemaListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterMainSchema is called when entering the mainSchema production.
	EnterMainSchema(c *MainSchemaContext)

	// EnterSubSchema is called when entering the subSchema production.
	EnterSubSchema(c *SubSchemaContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterSchema is called when entering the schema production.
	EnterSchema(c *SchemaContext)

	// EnterConstance is called when entering the constance production.
	EnterConstance(c *ConstanceContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterCustomType is called when entering the customType production.
	EnterCustomType(c *CustomTypeContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterBuildinType is called when entering the buildinType production.
	EnterBuildinType(c *BuildinTypeContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitMainSchema is called when exiting the mainSchema production.
	ExitMainSchema(c *MainSchemaContext)

	// ExitSubSchema is called when exiting the subSchema production.
	ExitSubSchema(c *SubSchemaContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitSchema is called when exiting the schema production.
	ExitSchema(c *SchemaContext)

	// ExitConstance is called when exiting the constance production.
	ExitConstance(c *ConstanceContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitCustomType is called when exiting the customType production.
	ExitCustomType(c *CustomTypeContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitBuildinType is called when exiting the buildinType production.
	ExitBuildinType(c *BuildinTypeContext)
}
