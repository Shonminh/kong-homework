package project

import (
	api2 "github.com/Shonminh/kong-homework/apps/project/api"
	"github.com/Shonminh/kong-homework/apps/project/internal"
	"github.com/Shonminh/kong-homework/apps/project/internal/repository"
	"github.com/Shonminh/kong-homework/apps/project/internal/repository/api"
)

func NewUserProjectRepository() api.UserProjectRepo {
	return &repository.UserProjectRepository{}
}

func NewUserProjectService(Repo api.UserProjectRepo) api2.UserProjectService {
	return &internal.UserProjectServiceImpl{Repo: Repo}
}
