package auth

import "github.com/google/wire"

var AuthenticationSet = wire.NewSet(
	NewUserRepo,
	NewAuthenticationService,
	NewProjectAccessService,
)
