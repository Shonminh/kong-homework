package repository

import (
	"context"

	"github.com/Shonminh/bilibee/pkg/db"

	"github.com/Shonminh/kong-homework/apps/project/internal/repository/model"
)

type UserProjectRepository struct {
}

func (r *UserProjectRepository) GetUserProjects(ctx context.Context, userID string) ([]model.UserProjectTab, error) {
	var res []model.UserProjectTab
	if err := db.GetDb(ctx).Where("user_id = ? ", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
