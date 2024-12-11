// Code generated from GuardQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GuardQuery
import "github.com/antlr4-go/antlr/v4"

// BaseGuardQueryListener is a complete listener for a parse tree produced by GuardQueryParser.
type BaseGuardQueryListener struct{}

var _ GuardQueryListener = &BaseGuardQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGuardQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGuardQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGuardQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGuardQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseGuardQueryListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseGuardQueryListener) ExitParse(ctx *ParseContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseGuardQueryListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseGuardQueryListener) ExitSelect(ctx *SelectContext) {}

// EnterWhereStmt is called when production whereStmt is entered.
func (s *BaseGuardQueryListener) EnterWhereStmt(ctx *WhereStmtContext) {}

// ExitWhereStmt is called when production whereStmt is exited.
func (s *BaseGuardQueryListener) ExitWhereStmt(ctx *WhereStmtContext) {}

// EnterSourceOrSubquery is called when production sourceOrSubquery is entered.
func (s *BaseGuardQueryListener) EnterSourceOrSubquery(ctx *SourceOrSubqueryContext) {}

// ExitSourceOrSubquery is called when production sourceOrSubquery is exited.
func (s *BaseGuardQueryListener) ExitSourceOrSubquery(ctx *SourceOrSubqueryContext) {}

// EnterSource is called when production source is entered.
func (s *BaseGuardQueryListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production source is exited.
func (s *BaseGuardQueryListener) ExitSource(ctx *SourceContext) {}

// EnterNode is called when production node is entered.
func (s *BaseGuardQueryListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BaseGuardQueryListener) ExitNode(ctx *NodeContext) {}

// EnterProvider is called when production provider is entered.
func (s *BaseGuardQueryListener) EnterProvider(ctx *ProviderContext) {}

// ExitProvider is called when production provider is exited.
func (s *BaseGuardQueryListener) ExitProvider(ctx *ProviderContext) {}

// EnterProviderArg is called when production providerArg is entered.
func (s *BaseGuardQueryListener) EnterProviderArg(ctx *ProviderArgContext) {}

// ExitProviderArg is called when production providerArg is exited.
func (s *BaseGuardQueryListener) ExitProviderArg(ctx *ProviderArgContext) {}

// EnterResultColumn is called when production resultColumn is entered.
func (s *BaseGuardQueryListener) EnterResultColumn(ctx *ResultColumnContext) {}

// ExitResultColumn is called when production resultColumn is exited.
func (s *BaseGuardQueryListener) ExitResultColumn(ctx *ResultColumnContext) {}

// EnterResultAlias is called when production resultAlias is entered.
func (s *BaseGuardQueryListener) EnterResultAlias(ctx *ResultAliasContext) {}

// ExitResultAlias is called when production resultAlias is exited.
func (s *BaseGuardQueryListener) ExitResultAlias(ctx *ResultAliasContext) {}

// EnterAnyName is called when production anyName is entered.
func (s *BaseGuardQueryListener) EnterAnyName(ctx *AnyNameContext) {}

// ExitAnyName is called when production anyName is exited.
func (s *BaseGuardQueryListener) ExitAnyName(ctx *AnyNameContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseGuardQueryListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseGuardQueryListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterValueExpr is called when production valueExpr is entered.
func (s *BaseGuardQueryListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BaseGuardQueryListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterBuildinSource is called when production buildinSource is entered.
func (s *BaseGuardQueryListener) EnterBuildinSource(ctx *BuildinSourceContext) {}

// ExitBuildinSource is called when production buildinSource is exited.
func (s *BaseGuardQueryListener) ExitBuildinSource(ctx *BuildinSourceContext) {}

// EnterBuildinFunction is called when production buildinFunction is entered.
func (s *BaseGuardQueryListener) EnterBuildinFunction(ctx *BuildinFunctionContext) {}

// ExitBuildinFunction is called when production buildinFunction is exited.
func (s *BaseGuardQueryListener) ExitBuildinFunction(ctx *BuildinFunctionContext) {}

// EnterLiteralValue is called when production literalValue is entered.
func (s *BaseGuardQueryListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production literalValue is exited.
func (s *BaseGuardQueryListener) ExitLiteralValue(ctx *LiteralValueContext) {}
