// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/Shonminh/kong-homework/apps/auth/wire"
	"github.com/Shonminh/kong-homework/apps/linting_rule/access/http"
	"github.com/Shonminh/kong-homework/apps/linting_rule/wire"
	"github.com/Shonminh/kong-homework/apps/project/wire"
	"github.com/Shonminh/kong-homework/routers"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitMainServerApp() (*MainServerApp, error) {
	gormDB := db.NewDB()
	engine := routers.NewRouters(gormDB)
	lintingRuleRepo := linting_rule.NewLintingRuleRepo()
	lintingRuleService := linting_rule.NewLintingRuleService(lintingRuleRepo)
	userProjectRepo := project.NewUserProjectRepository()
	userProjectService := project.NewUserProjectService(userProjectRepo)
	projectAccessService := auth.NewProjectAccessService(userProjectService)
	userRepo := auth.NewUserRepo()
	authenticationService := auth.NewAuthenticationService(userRepo)
	lintingRuleHttpSchema := linting_rule.NewLintingRuleHttpSchema(lintingRuleService, projectAccessService, authenticationService)
	mainServerApp := &MainServerApp{
		router: engine,
		schema: lintingRuleHttpSchema,
	}
	return mainServerApp, nil
}

// wire.go:

type MainServerApp struct {
	router *gin.Engine
	schema *http.LintingRuleHttpSchema
}

func (app *MainServerApp) Register() {
	app.schema.RegisterSchema(app.router)
}

func (app *MainServerApp) RunHttpServer(addr string) error {
	return app.router.Run(addr)
}
