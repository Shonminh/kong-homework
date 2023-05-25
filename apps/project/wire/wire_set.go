package project

import "github.com/google/wire"

var ProjectSet = wire.NewSet(
	NewUserProjectRepository,
	NewUserProjectService,
)
