//go:build wireinject

package main

import (
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	auth "github.com/Shonminh/kong-homework/apps/auth/wire"
	"github.com/Shonminh/kong-homework/apps/linting_rule/access/http"
	linting_rule "github.com/Shonminh/kong-homework/apps/linting_rule/wire"
	project "github.com/Shonminh/kong-homework/apps/project/wire"
	"github.com/Shonminh/kong-homework/routers"
)

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

func InitMainServerApp() (*MainServerApp, error) {
	wire.Build(
		linting_rule.LintingRuleSet,
		project.ProjectSet,
		auth.AuthenticationSet,
		routers.NewRouters,
		db.NewDB,
		wire.Struct(new(MainServerApp), "*"),
	)

	return &MainServerApp{}, nil
}
