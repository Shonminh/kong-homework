package linting_rule

import (
	api3 "github.com/Shonminh/kong-homework/apps/auth/api"
	"github.com/Shonminh/kong-homework/apps/linting_rule/access/http"
	"github.com/Shonminh/kong-homework/apps/linting_rule/api"
	"github.com/Shonminh/kong-homework/apps/linting_rule/internal"
	"github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository"
	api2 "github.com/Shonminh/kong-homework/apps/linting_rule/internal/repository/api"
)

func NewLintingRuleHttpSchema(lintingRuleService api.LintingRuleService, accessService api3.ProjectAccessService,
	authenticationService api3.AuthenticationService,
) *http.LintingRuleHttpSchema {

	return &http.LintingRuleHttpSchema{
		LintingRuleService:    lintingRuleService,
		AuthenticationService: authenticationService,
		ProjectAccessService:  accessService,
	}
}

func NewLintingRuleService(Repo api2.LintingRuleRepo) api.LintingRuleService {
	return &internal.LintingRuleServiceImpl{Repo: Repo}
}

func NewLintingRuleRepo() api2.LintingRuleRepo {
	return &repository.LintingRuleRepository{}
}
