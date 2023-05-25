package internal

import (
	"context"

	"github.com/Shonminh/kong-homework/apps/project/api"
	api2 "github.com/Shonminh/kong-homework/apps/project/internal/repository/api"
)

type UserProjectServiceImpl struct {
	Repo api2.UserProjectRepo
}

func (impl *UserProjectServiceImpl) GetUserProjects(ctx context.Context, userID string) ([]api.Project, error) {
	projects, err := impl.Repo.GetUserProjects(ctx, userID)
	if err != nil {
		return nil, err
	}
	var res []api.Project
	for _, row := range projects {
		res = append(res, api.Project{ID: row.ProjectID})
	}
	return res, nil
}
