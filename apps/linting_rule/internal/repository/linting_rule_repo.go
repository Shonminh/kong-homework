package repository

import (
	"context"

	"github.com/Shonminh/bilibee/pkg/db"

	"github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository/model"
)

type LintingRuleRepository struct {
}

func (r *LintingRuleRepository) CreateLintingRule(ctx context.Context, rule *model.LintingRuleTab) error {
	return db.GetDb(ctx).Create(rule).Error
}

func (r *LintingRuleRepository) UpdateLintingRule(ctx context.Context, rule *model.LintingRuleTab) error {
	return db.GetDb(ctx).Save(rule).Error
}

func (r *LintingRuleRepository) GetLintingRuleByID(ctx context.Context, id string) (*model.LintingRuleTab, error) {
	var rule model.LintingRuleTab
	if err := db.GetDb(ctx).First(&rule, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &rule, nil
}

func (r *LintingRuleRepository) GetLintingRulesByProjectID(ctx context.Context, projectID string) ([]*model.LintingRuleTab, error) {
	var rules []*model.LintingRuleTab
	if err := db.GetDb(ctx).Find(&rules, "project_id = ?", projectID).Error; err != nil {
		return nil, err
	}
	return rules, nil
}

func (r *LintingRuleRepository) DeleteLintingRule(ctx context.Context, rule *model.LintingRuleTab) error {
	return db.GetDb(ctx).Delete(rule).Error
}
