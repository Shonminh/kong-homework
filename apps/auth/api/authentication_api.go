package api

import "context"

// AuthenticationService 是用于身份验证的接口
type AuthenticationService interface {
	VerifyToken(token string) (string, error)
	IsUserAdmin(ctx context.Context, userId string) (bool, error)
}
