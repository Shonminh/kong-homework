package internal

import (
	"github.com/Shonminh/kong-homework/apps/linting_rule"
	"github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository/api"
)

type LintingRuleServiceImpl struct {
	Repo api.LintingRuleRepo
}

func (impl *LintingRuleServiceImpl) CreateLintingRule(projectID string, ruleName string, ruleContent string) error {
	// TODO implement me
	panic("implement me")
}

func (impl *LintingRuleServiceImpl) UpdateLintingRule(ruleID string, ruleName string, ruleContent string) error {
	// TODO implement me
	panic("implement me")
}

func (impl *LintingRuleServiceImpl) GetLintingRule(ruleID string) (*linting_rule.LintingRule, error) {
	// TODO implement me
	panic("implement me")
}

func (impl *LintingRuleServiceImpl) GetProjectLintingRule(projectID string) (*linting_rule.LintingRule, error) {
	// TODO implement me
	panic("implement me")
}

func (impl *LintingRuleServiceImpl) DeleteLintingRule(ruleID string) error {
	// TODO implement me
	panic("implement me")
}
