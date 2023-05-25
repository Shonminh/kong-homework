package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/Shonminh/kong-homework/apps/linting_rule"
)

func (schema *LintingRuleHttpSchema) RegisterSchema(router *gin.Engine) {

	router.Group("/api/linting-rules", schema.AuthMiddleware())
	// 注册创建Linting规则的API接口
	router.POST("/", schema.CreateLintingRuleHandler)

	// 注册更新Linting规则的API接口
	router.PUT("/:ruleID", schema.UpdateLintingRuleHandler)

	// 注册获取Linting规则的API接口
	router.GET("/:ruleID", schema.GetLintingRuleHandler)
}

func (schema *LintingRuleHttpSchema) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		schema.authMiddleware(c)
	}
}

// authMiddleware 是用于验证身份的中间件
func (schema *LintingRuleHttpSchema) authMiddleware(c *gin.Context) {
	// 从请求中获取 token
	token := c.GetHeader("Authorization")

	// 验证 token
	userID, err := schema.AuthenticationService.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.Abort()
		return
	}
	projects, err := schema.ProjectAccessService.GetUserProjects(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "find project failed",
		})
		return
	}

	// 查询rule ID是否在url中存在
	ruleID := c.Param("ruleID")
	projectID := ""
	if len(ruleID) == 0 {
		var lintingRule linting_rule.LintingRule
		_ = c.ShouldBindBodyWith(&lintingRule, binding.JSON) // 保证下次
		projectID = lintingRule.ProjectID

	} else {
		lintingRule, err := schema.LintingRuleService.GetLintingRule(c.Request.Context(), ruleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "find project failed",
			})
			return
		}
		projectID = lintingRule.ProjectID
	}
	hasProjectAccess := false
	for _, v := range projects {
		if v.ID == projectID {
			hasProjectAccess = true
			break
		}
	}
	if !hasProjectAccess {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid access",
		})
		c.Abort()
		return
	}

	// 将验证通过的用户 ID 存储在上下文中
	c.Set(linting_rule.UserID, userID)
	c.Set(linting_rule.ProjectID, projectID)

	// 继续处理请求
	c.Next()
}
