// Code generated from GuardQuery.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GuardQuery
import "github.com/antlr4-go/antlr/v4"

// GuardQueryListener is a complete listener for a parse tree produced by GuardQueryParser.
type GuardQueryListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterTimeStmt is called when entering the timeStmt production.
	EnterTimeStmt(c *TimeStmtContext)

	// EnterAbsTimeStmt is called when entering the absTimeStmt production.
	EnterAbsTimeStmt(c *AbsTimeStmtContext)

	// EnterRelatTimeStmt is called when entering the relatTimeStmt production.
	EnterRelatTimeStmt(c *RelatTimeStmtContext)

	// EnterTimeUnit is called when entering the timeUnit production.
	EnterTimeUnit(c *TimeUnitContext)

	// EnterWhereStmt is called when entering the whereStmt production.
	EnterWhereStmt(c *WhereStmtContext)

	// EnterSourceOrSubquery is called when entering the sourceOrSubquery production.
	EnterSourceOrSubquery(c *SourceOrSubqueryContext)

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// EnterNode is called when entering the node production.
	EnterNode(c *NodeContext)

	// EnterProvider is called when entering the provider production.
	EnterProvider(c *ProviderContext)

	// EnterProviderArg is called when entering the providerArg production.
	EnterProviderArg(c *ProviderArgContext)

	// EnterResultColumn is called when entering the resultColumn production.
	EnterResultColumn(c *ResultColumnContext)

	// EnterResultAlias is called when entering the resultAlias production.
	EnterResultAlias(c *ResultAliasContext)

	// EnterAnyName is called when entering the anyName production.
	EnterAnyName(c *AnyNameContext)

	// EnterBoolExpr is called when entering the boolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterValueExpr is called when entering the valueExpr production.
	EnterValueExpr(c *ValueExprContext)

	// EnterBuildinSource is called when entering the buildinSource production.
	EnterBuildinSource(c *BuildinSourceContext)

	// EnterBuildinFunction is called when entering the buildinFunction production.
	EnterBuildinFunction(c *BuildinFunctionContext)

	// EnterLiteralValue is called when entering the literalValue production.
	EnterLiteralValue(c *LiteralValueContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitTimeStmt is called when exiting the timeStmt production.
	ExitTimeStmt(c *TimeStmtContext)

	// ExitAbsTimeStmt is called when exiting the absTimeStmt production.
	ExitAbsTimeStmt(c *AbsTimeStmtContext)

	// ExitRelatTimeStmt is called when exiting the relatTimeStmt production.
	ExitRelatTimeStmt(c *RelatTimeStmtContext)

	// ExitTimeUnit is called when exiting the timeUnit production.
	ExitTimeUnit(c *TimeUnitContext)

	// ExitWhereStmt is called when exiting the whereStmt production.
	ExitWhereStmt(c *WhereStmtContext)

	// ExitSourceOrSubquery is called when exiting the sourceOrSubquery production.
	ExitSourceOrSubquery(c *SourceOrSubqueryContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitProvider is called when exiting the provider production.
	ExitProvider(c *ProviderContext)

	// ExitProviderArg is called when exiting the providerArg production.
	ExitProviderArg(c *ProviderArgContext)

	// ExitResultColumn is called when exiting the resultColumn production.
	ExitResultColumn(c *ResultColumnContext)

	// ExitResultAlias is called when exiting the resultAlias production.
	ExitResultAlias(c *ResultAliasContext)

	// ExitAnyName is called when exiting the anyName production.
	ExitAnyName(c *AnyNameContext)

	// ExitBoolExpr is called when exiting the boolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitValueExpr is called when exiting the valueExpr production.
	ExitValueExpr(c *ValueExprContext)

	// ExitBuildinSource is called when exiting the buildinSource production.
	ExitBuildinSource(c *BuildinSourceContext)

	// ExitBuildinFunction is called when exiting the buildinFunction production.
	ExitBuildinFunction(c *BuildinFunctionContext)

	// ExitLiteralValue is called when exiting the literalValue production.
	ExitLiteralValue(c *LiteralValueContext)
}
