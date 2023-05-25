package api

import (
	"context"

	"github.com/Shonminh/kong-homework/apps/project/internal/repository/model"
)

type UserProjectRepo interface {
	GetUserProjects(ctx context.Context, userID string) ([]model.UserProjectTab, error)
}
