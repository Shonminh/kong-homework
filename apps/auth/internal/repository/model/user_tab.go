package model

type UserTab struct {
	ID             uint64   `gorm:"type:INT;PRIMARY_KEY;"`
	UserID         string   `gorm:"type:VARCHAR(64);"`
	Name           string   `gorm:"type:VARCHAR(255);"`
	OrganizationID int      `gorm:"type:INT;"`
	RoleType       RoleType `gorm:"type:INT;"`
	CreateTime     uint64   `gorm:"not null" json:"create_time"`
	UpdateTime     uint64   `gorm:"not null" json:"update_time"`
}
type RoleType int

const (
	NonAdminRoleType RoleType = 0
	AdminRoleType    RoleType = 1
)
