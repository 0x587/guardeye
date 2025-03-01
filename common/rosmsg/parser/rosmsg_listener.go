// Code generated from Rosmsg.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Rosmsg
import "github.com/antlr4-go/antlr/v4"

// RosmsgListener is a complete listener for a parse tree produced by RosmsgParser.
type RosmsgListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterMsg_stat is called when entering the msg_stat production.
	EnterMsg_stat(c *Msg_statContext)

	// EnterSrv_stat is called when entering the srv_stat production.
	EnterSrv_stat(c *Srv_statContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterField_type is called when entering the field_type production.
	EnterField_type(c *Field_typeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterCustomed_type is called when entering the customed_type production.
	EnterCustomed_type(c *Customed_typeContext)

	// EnterField_name is called when entering the field_name production.
	EnterField_name(c *Field_nameContext)

	// EnterBuildin_type is called when entering the buildin_type production.
	EnterBuildin_type(c *Buildin_typeContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitMsg_stat is called when exiting the msg_stat production.
	ExitMsg_stat(c *Msg_statContext)

	// ExitSrv_stat is called when exiting the srv_stat production.
	ExitSrv_stat(c *Srv_statContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitField_type is called when exiting the field_type production.
	ExitField_type(c *Field_typeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitCustomed_type is called when exiting the customed_type production.
	ExitCustomed_type(c *Customed_typeContext)

	// ExitField_name is called when exiting the field_name production.
	ExitField_name(c *Field_nameContext)

	// ExitBuildin_type is called when exiting the buildin_type production.
	ExitBuildin_type(c *Buildin_typeContext)
}
