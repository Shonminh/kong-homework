package api

import "context"

type UserRepo interface {
	IsUserAdmin(ctx context.Context, userID string) (bool, error)
}
