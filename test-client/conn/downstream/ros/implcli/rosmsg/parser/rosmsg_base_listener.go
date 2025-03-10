// Code generated from Rosmsg.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Rosmsg
import "github.com/antlr4-go/antlr/v4"

// BaseRosmsgListener is a complete listener for a parse tree produced by RosmsgParser.
type BaseRosmsgListener struct{}

var _ RosmsgListener = &BaseRosmsgListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRosmsgListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRosmsgListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRosmsgListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRosmsgListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseRosmsgListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseRosmsgListener) ExitParse(ctx *ParseContext) {}

// EnterMsgStat is called when production msgStat is entered.
func (s *BaseRosmsgListener) EnterMsgStat(ctx *MsgStatContext) {}

// ExitMsgStat is called when production msgStat is exited.
func (s *BaseRosmsgListener) ExitMsgStat(ctx *MsgStatContext) {}

// EnterSrvStat is called when production srvStat is entered.
func (s *BaseRosmsgListener) EnterSrvStat(ctx *SrvStatContext) {}

// ExitSrvStat is called when production srvStat is exited.
func (s *BaseRosmsgListener) ExitSrvStat(ctx *SrvStatContext) {}

// EnterField is called when production field is entered.
func (s *BaseRosmsgListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseRosmsgListener) ExitField(ctx *FieldContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseRosmsgListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseRosmsgListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterType is called when production type is entered.
func (s *BaseRosmsgListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseRosmsgListener) ExitType(ctx *TypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseRosmsgListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseRosmsgListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterCustomType is called when production customType is entered.
func (s *BaseRosmsgListener) EnterCustomType(ctx *CustomTypeContext) {}

// ExitCustomType is called when production customType is exited.
func (s *BaseRosmsgListener) ExitCustomType(ctx *CustomTypeContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseRosmsgListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseRosmsgListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterBuildinType is called when production buildinType is entered.
func (s *BaseRosmsgListener) EnterBuildinType(ctx *BuildinTypeContext) {}

// ExitBuildinType is called when production buildinType is exited.
func (s *BaseRosmsgListener) ExitBuildinType(ctx *BuildinTypeContext) {}
