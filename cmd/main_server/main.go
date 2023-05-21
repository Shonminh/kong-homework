package main

import (
	"github.com/Shonminh/bilibee/pkg/logger"
)

func main() {
	app, err := InitMainServerApp()
	if err != nil {
		logger.LogPanicf("err=%v", err)
	}
	app.Register()
	_ = app.RunHttpServer(":8080")
}
