//go:build wireinject

package main

import (
	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/Shonminh/kong-homework/routers"
)

type MainServerApp struct {
	router *gin.Engine
}

func (app *MainServerApp) Register() {

}

func (app *MainServerApp) RunHttpServer(addr string) error {
	return app.router.Run(addr)
}

func InitMainServerApp() (*MainServerApp, error) {
	wire.Build(
		// linting_rule.LintingRuleSet,
		routers.NewRouters,
		db.NewDB,
		wire.Struct(new(MainServerApp), "*"),
	)

	return &MainServerApp{}, nil
}
