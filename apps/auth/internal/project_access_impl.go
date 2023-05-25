package internal

import (
	"context"

	"github.com/Shonminh/kong-homework/apps/project/api"
)

type ProjectAccessServiceImpl struct {
	Srv api.UserProjectService
}

func (impl *ProjectAccessServiceImpl) GetUserProjects(ctx context.Context, userID string) ([]api.Project, error) {
	return impl.Srv.GetUserProjects(ctx, userID)
}
