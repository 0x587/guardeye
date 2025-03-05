// Code generated from Rossrv.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Rossrv
import "github.com/antlr4-go/antlr/v4"

// RossrvListener is a complete listener for a parse tree produced by RossrvParser.
type RossrvListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
