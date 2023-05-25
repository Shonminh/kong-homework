package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	api2 "github.com/Shonminh/kong-homework/apps/auth/api"
	"github.com/Shonminh/kong-homework/apps/linting_rule"
	"github.com/Shonminh/kong-homework/apps/linting_rule/api"
	"github.com/Shonminh/kong-homework/pkg/yaml"
)

type LintingRuleHttpSchema struct {
	LintingRuleService    api.LintingRuleService
	AuthenticationService api2.AuthenticationService
	ProjectAccessService  api2.ProjectAccessService
}

// CreateLintingRuleHandler 创建Linting规则的API接口
func (schema *LintingRuleHttpSchema) CreateLintingRuleHandler(c *gin.Context) {

	// 解析请求数据

	lintingRule := getLintingRule(c)
	if lintingRule == nil {
		return
	}
	// 检查是否存在admin角色
	if !schema.hasAdminRole(c) {
		return
	}

	// 调用LintingRuleService的CreateLintingRule方法创建规则
	err := schema.LintingRuleService.CreateLintingRule(c.Request.Context(), lintingRule.ProjectID, lintingRule.RuleName, lintingRule.RuleContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "save linting rule failed, try later",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}

func (schema *LintingRuleHttpSchema) hasAdminRole(c *gin.Context) bool {
	userID := c.GetString(linting_rule.UserID)
	ok, err := schema.AuthenticationService.IsUserAdmin(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "save linting rule failed, try later",
		})
		return true
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "you are not admin user",
		})
		return true
	}
	return false
}

func getLintingRule(c *gin.Context) *linting_rule.LintingRule {
	var lintingRule linting_rule.LintingRule
	if err := c.ShouldBindBodyWith(&lintingRule, binding.JSON); err != nil {
		// 处理解析错误
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bind request failed",
		})
		return nil
	}
	_, err := yaml.VerifyYamlContent(lintingRule.RuleContent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid yaml format",
		})
		return nil
	}
	return &lintingRule
}

// UpdateLintingRuleHandler 更新Linting规则的API接口
func (schema *LintingRuleHttpSchema) UpdateLintingRuleHandler(c *gin.Context) {

	lintingRule := getLintingRule(c)
	if lintingRule == nil {
		return
	}
	if !schema.hasAdminRole(c) {
		return
	}

	// 调用LintingRuleService的UpdateLintingRule方法更新规则
	ruleID := c.Param("ruleID")
	err := schema.LintingRuleService.UpdateLintingRule(c.Request.Context(), ruleID, lintingRule.RuleName, lintingRule.RuleContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "update linting rule failed, try later",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	return
}

// GetLintingRuleHandler 获取Linting规则的API接口
func (schema *LintingRuleHttpSchema) GetLintingRuleHandler(c *gin.Context) {
	// 解析路径参数
	ruleID := c.Param("ruleID")

	// 调用LintingRuleService的GetLintingRule方法获取规则
	lintingRule, err := schema.LintingRuleService.GetLintingRule(c.Request.Context(), ruleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "query linting rule failed, try later",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lintingRule})
}
