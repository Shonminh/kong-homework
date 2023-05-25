package api

import (
	"context"

	"github.com/Shonminh/kong-homework/apps/linting_rule"
)

type LintingRuleService interface {
	// 创建Linting规则
	CreateLintingRule(ctx context.Context, projectID string, ruleName string, ruleContent string) error

	// 更新Linting规则
	UpdateLintingRule(ctx context.Context, ruleID string, ruleName string, ruleContent string) error

	// 检索Linting规则
	GetLintingRule(ctx context.Context, ruleID string) (*linting_rule.LintingRule, error)
}
