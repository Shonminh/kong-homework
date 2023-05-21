package http

import "github.com/gin-gonic/gin"

func (schema *LintingRuleHttpSchema) RegisterSchema(router *gin.Engine) {

	// 注册创建Linting规则的API接口
	router.POST("/linting-rules", schema.CreateLintingRuleHandler)

	// 注册更新Linting规则的API接口
	router.PUT("/linting-rules/:ruleID", schema.UpdateLintingRuleHandler)

	// 注册获取Linting规则的API接口
	router.GET("/linting-rules/:ruleID", schema.GetLintingRuleHandler)

}
