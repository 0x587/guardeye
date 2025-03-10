// Code generated from Rossrv.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Rossrv
import "github.com/antlr4-go/antlr/v4"

// BaseRossrvListener is a complete listener for a parse tree produced by RossrvParser.
type BaseRossrvListener struct{}

var _ RossrvListener = &BaseRossrvListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRossrvListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRossrvListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRossrvListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRossrvListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseRossrvListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseRossrvListener) ExitParse(ctx *ParseContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseRossrvListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseRossrvListener) ExitRoot(ctx *RootContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseRossrvListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseRossrvListener) ExitEntry(ctx *EntryContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseRossrvListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseRossrvListener) ExitObj(ctx *ObjContext) {}

// EnterField is called when production field is entered.
func (s *BaseRossrvListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseRossrvListener) ExitField(ctx *FieldContext) {}

// EnterList is called when production list is entered.
func (s *BaseRossrvListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseRossrvListener) ExitList(ctx *ListContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseRossrvListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseRossrvListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseRossrvListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseRossrvListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseRossrvListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseRossrvListener) ExitInteger(ctx *IntegerContext) {}

// EnterString is called when production string is entered.
func (s *BaseRossrvListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseRossrvListener) ExitString(ctx *StringContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseRossrvListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseRossrvListener) ExitIdentifier(ctx *IdentifierContext) {}
