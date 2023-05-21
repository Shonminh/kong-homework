package http

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Shonminh/kong-homework/apps/linting_rule"
	"github.com/Shonminh/kong-homework/apps/linting_rule/api"
)

type LintingRuleHttpSchema struct {
	LintingRuleService api.LintingRuleService
}

// 创建Linting规则的API接口
func (schema *LintingRuleHttpSchema) CreateLintingRuleHandler(c *gin.Context) {
	// 解析请求数据
	var lintingRule linting_rule.LintingRule
	if err := c.ShouldBindJSON(&lintingRule); err != nil {
		// 处理解析错误
		// ...
		return
	}

	// 调用LintingRuleService的CreateLintingRule方法创建规则
	err := schema.LintingRuleService.CreateLintingRule(lintingRule.ProjectID, lintingRule.RuleName, lintingRule.RuleContent)
	if err != nil {
		// 处理错误
		// ...
		return
	}

	// 返回成功响应
	// ...
}

// 更新Linting规则的API接口
func (schema *LintingRuleHttpSchema) UpdateLintingRuleHandler(c *gin.Context) {
	// 解析请求数据和路径参数
	ruleID := c.Param("ruleID")
	var lintingRule linting_rule.LintingRule
	if err := c.ShouldBindJSON(&lintingRule); err != nil {
		// 处理解析错误
		// ...
		return
	}

	// 调用LintingRuleService的UpdateLintingRule方法更新规则
	err := schema.LintingRuleService.UpdateLintingRule(ruleID, lintingRule.RuleName, lintingRule.RuleContent)
	if err != nil {
		// 处理错误
		// ...
		return
	}

	// 返回成功响应
	// ...
}

// 获取Linting规则的API接口
func (schema *LintingRuleHttpSchema) GetLintingRuleHandler(c *gin.Context) {
	// 解析路径参数
	ruleID := c.Param("ruleID")

	// 调用LintingRuleService的GetLintingRule方法获取规则
	lintingRule, err := schema.LintingRuleService.GetLintingRule(ruleID)
	// TODO implement me
	fmt.Println(lintingRule)
	if err != nil {
		// 处理错误
		// ...
		return
	}

	// 返回规则响应
	// ...
}
