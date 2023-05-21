package linting_rule

import (
	"github.com/Shonminh/kong-homework/apps/linting_rule/access/http"
	"github.com/Shonminh/kong-homework/apps/linting_rule/api"
	"github.com/Shonminh/kong-homework/apps/linting_rule/internal"
	"github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository"
	api2 "github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository/api"
)

func NewLintingRuleHttpSchema(service api.LintingRuleService) *http.LintingRuleHttpSchema {
	return &http.LintingRuleHttpSchema{
		LintingRuleService: service,
	}
}

func NewLintingRuleService(Repo api2.LintingRuleRepo) api.LintingRuleService {
	return &internal.LintingRuleServiceImpl{Repo: Repo}
}

func NewLintingRuleRepo() api2.LintingRuleRepo {
	return &repository.LintingRuleRepository{}
}
