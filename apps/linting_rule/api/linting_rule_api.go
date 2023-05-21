package api

import "github.com/Shonminh/kong-homework/apps/linting_rule"

type LintingRuleService interface {
	// 创建Linting规则
	CreateLintingRule(projectID string, ruleName string, ruleContent string) error

	// 更新Linting规则
	UpdateLintingRule(ruleID string, ruleName string, ruleContent string) error

	// 检索Linting规则
	GetLintingRule(ruleID string) (*linting_rule.LintingRule, error)

	// 检索项目的Linting规则
	GetProjectLintingRule(projectID string) (*linting_rule.LintingRule, error)

	// 删除Linting规则
	DeleteLintingRule(ruleID string) error
}
