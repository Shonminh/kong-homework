package api

import (
	"context"

	"github.com/Shonminh/kong-homework/apps/project/api"
)

type ProjectAccessService interface {
	GetUserProjects(ctx context.Context, userID string) ([]api.Project, error)
}
