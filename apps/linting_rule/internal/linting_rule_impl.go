package internal

import (
	"context"
	"errors"

	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/bilibee/pkg/time"

	"github.com/Shonminh/kong-homework/apps/linting_rule"
	"github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository/api"
	"github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository/model"
)

type LintingRuleServiceImpl struct {
	Repo api.LintingRuleRepo
}

func (impl *LintingRuleServiceImpl) CreateLintingRule(ctx context.Context, projectID string, ruleName string, ruleContent string) error {
	err := impl.Repo.CreateLintingRule(ctx, &model.LintingRuleTab{
		ProjectID:   projectID,
		RuleName:    ruleName,
		RuleContent: ruleContent,
		CreateTime:  time.NowUint64(),
		UpdateTime:  time.NowUint64(),
	})
	return err
}

var ErrNotFound = errors.New("query not found")

func (impl *LintingRuleServiceImpl) UpdateLintingRule(ctx context.Context, ruleID string, ruleName string, ruleContent string) error {
	err := db.Transaction(ctx, func(ctx context.Context) error {

		ruleTab, err := impl.Repo.GetLintingRuleByID(ctx, ruleID)
		if err != nil {
			return err
		}
		if ruleTab == nil {
			return ErrNotFound
		}
		ruleTab.RuleName = ruleName
		ruleTab.RuleContent = ruleContent
		if err = impl.Repo.UpdateLintingRule(ctx, ruleTab); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (impl *LintingRuleServiceImpl) GetLintingRule(ctx context.Context, ruleID string) (*linting_rule.LintingRule, error) {
	ruleTab, err := impl.Repo.GetLintingRuleByID(ctx, ruleID)
	if err != nil {
		return nil, err
	}
	return &linting_rule.LintingRule{
		ProjectID:   ruleTab.ProjectID,
		RuleName:    ruleTab.RuleName,
		RuleContent: ruleTab.RuleContent,
	}, nil
}
