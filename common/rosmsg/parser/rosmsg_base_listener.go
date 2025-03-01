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

// EnterMsg_stat is called when production msg_stat is entered.
func (s *BaseRosmsgListener) EnterMsg_stat(ctx *Msg_statContext) {}

// ExitMsg_stat is called when production msg_stat is exited.
func (s *BaseRosmsgListener) ExitMsg_stat(ctx *Msg_statContext) {}

// EnterSrv_stat is called when production srv_stat is entered.
func (s *BaseRosmsgListener) EnterSrv_stat(ctx *Srv_statContext) {}

// ExitSrv_stat is called when production srv_stat is exited.
func (s *BaseRosmsgListener) ExitSrv_stat(ctx *Srv_statContext) {}

// EnterField is called when production field is entered.
func (s *BaseRosmsgListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseRosmsgListener) ExitField(ctx *FieldContext) {}

// EnterField_type is called when production field_type is entered.
func (s *BaseRosmsgListener) EnterField_type(ctx *Field_typeContext) {}

// ExitField_type is called when production field_type is exited.
func (s *BaseRosmsgListener) ExitField_type(ctx *Field_typeContext) {}

// EnterType is called when production type is entered.
func (s *BaseRosmsgListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseRosmsgListener) ExitType(ctx *TypeContext) {}

// EnterCustomed_type is called when production customed_type is entered.
func (s *BaseRosmsgListener) EnterCustomed_type(ctx *Customed_typeContext) {}

// ExitCustomed_type is called when production customed_type is exited.
func (s *BaseRosmsgListener) ExitCustomed_type(ctx *Customed_typeContext) {}

// EnterField_name is called when production field_name is entered.
func (s *BaseRosmsgListener) EnterField_name(ctx *Field_nameContext) {}

// ExitField_name is called when production field_name is exited.
func (s *BaseRosmsgListener) ExitField_name(ctx *Field_nameContext) {}

// EnterBuildin_type is called when production buildin_type is entered.
func (s *BaseRosmsgListener) EnterBuildin_type(ctx *Buildin_typeContext) {}

// ExitBuildin_type is called when production buildin_type is exited.
func (s *BaseRosmsgListener) ExitBuildin_type(ctx *Buildin_typeContext) {}
