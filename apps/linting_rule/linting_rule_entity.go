package linting_rule

type LintingRule struct {
	ID          string `json:"id"`
	ProjectID   string `json:"project_id"`
	RuleName    string `json:"rule_name"`
	RuleContent string `json:"rule_content"`
}
