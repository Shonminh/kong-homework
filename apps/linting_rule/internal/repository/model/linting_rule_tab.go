package model

type LintingRuleTab struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	ProjectID   string `gorm:"not null" json:"project_id"`
	RuleName    string `gorm:"not null" json:"rule_name"`
	RuleContent string `gorm:"type:text" json:"rule_content"`
	CreateTime  uint64 `gorm:"not null" json:"create_time"`
	UpdateTime  uint64 `gorm:"not null" json:"update_time"`
}
