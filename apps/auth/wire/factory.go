package auth

import (
	api2 "github.com/Shonminh/kong-homework/apps/auth/api"
	"github.com/Shonminh/kong-homework/apps/auth/internal"
	"github.com/Shonminh/kong-homework/apps/auth/internal/repository"
	"github.com/Shonminh/kong-homework/apps/auth/internal/repository/api"
	api3 "github.com/Shonminh/kong-homework/apps/project/api"
)

func NewUserRepo() api.UserRepo {
	return &repository.UserRepository{}
}

func NewAuthenticationService(Repo api.UserRepo) api2.AuthenticationService {
	return &internal.AuthenticationServiceImpl{UserRepo: Repo}
}

func NewProjectAccessService(srv api3.UserProjectService) api2.ProjectAccessService {
	return &internal.ProjectAccessServiceImpl{Srv: srv}
}
