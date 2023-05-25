package repository

import (
	"context"

	"github.com/Shonminh/bilibee/pkg/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/Shonminh/kong-homework/apps/auth/internal/repository/model"
)

type UserRepository struct {
}

func (r *UserRepository) IsUserAdmin(ctx context.Context, userID string) (bool, error) {
	var user model.UserTab
	err := db.GetDb(ctx).Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return user.RoleType == model.AdminRoleType, nil
}
