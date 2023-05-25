package api

import "context"

type UserProjectService interface {
	GetUserProjects(ctx context.Context, userID string) ([]Project, error)
}

type Project struct {
	ID string `json:"id"`
	// 其他项目属性
}
