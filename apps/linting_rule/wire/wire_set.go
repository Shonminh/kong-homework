package linting_rule

import "github.com/google/wire"

var LintingRuleSet = wire.NewSet(
	NewLintingRuleHttpSchema,
	NewLintingRuleService,
	NewLintingRuleRepo,
)
