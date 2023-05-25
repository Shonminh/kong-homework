package model

type ProjectTab struct {
	ID             uint   `gorm:"primary_key" json:"id"`
	ProjectID      string `gorm:"not null" json:"project_id"`
	OrganizationID int    `gorm:"type:INT;"`
	CreateTime     uint64 `gorm:"not null" json:"create_time"`
	UpdateTime     uint64 `gorm:"not null" json:"update_time"`
}
