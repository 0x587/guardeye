// Code generated from Rosmsg.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Rosmsg
import "github.com/antlr4-go/antlr/v4"

// RosmsgListener is a complete listener for a parse tree produced by RosmsgParser.
type RosmsgListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterMsgStat is called when entering the msgStat production.
	EnterMsgStat(c *MsgStatContext)

	// EnterSrvStat is called when entering the srvStat production.
	EnterSrvStat(c *SrvStatContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

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

	// ExitMsgStat is called when exiting the msgStat production.
	ExitMsgStat(c *MsgStatContext)

	// ExitSrvStat is called when exiting the srvStat production.
	ExitSrvStat(c *SrvStatContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

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
