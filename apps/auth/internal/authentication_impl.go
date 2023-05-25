package internal

import (
	"context"
	"errors"

	"github.com/dgrijalva/jwt-go"

	"github.com/Shonminh/kong-homework/apps/auth/internal/repository/api"
)

// AuthenticationServiceImpl 是基于 JWT 的 TokenVerifier 实现
type AuthenticationServiceImpl struct {
	// secretKey string
	UserRepo api.UserRepo
}

const (
	secretKey = "DBBF3419-148C-489D-92C8-BC8166D95B8C" // TODO implement me  这里应该是设置成配置的会更好
)

// VerifyToken 验证和解析 JWT token
func (impl *AuthenticationServiceImpl) VerifyToken(tokenString string) (string, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		// 返回用于验证的密钥
		return []byte(secretKey), nil
	})

	// 验证 token 是否有效
	if err != nil {
		return "", err
	}

	// 检查 token 是否有效并获取其中的用户信息
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(string)
		return userID, nil
	}

	return "", errors.New("invalid token")
}

func (impl *AuthenticationServiceImpl) IsUserAdmin(ctx context.Context, userId string) (bool, error) {
	if userId == "" {
		return false, nil
	}
	return impl.UserRepo.IsUserAdmin(ctx, userId)
}
