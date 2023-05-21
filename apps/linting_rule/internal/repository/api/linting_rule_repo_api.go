package api

import (
	"context"

	"github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository/model"
)

type LintingRuleRepo interface {
	CreateLintingRule(ctx context.Context, rule *model.LintingRuleTab) error
	UpdateLintingRule(ctx context.Context, rule *model.LintingRuleTab) error
	GetLintingRuleByID(ctx context.Context, id string) (*model.LintingRuleTab, error)
	GetLintingRulesByProjectID(ctx context.Context, projectID string) ([]*model.LintingRuleTab, error)
	DeleteLintingRule(ctx context.Context, rule *model.LintingRuleTab) error
}
