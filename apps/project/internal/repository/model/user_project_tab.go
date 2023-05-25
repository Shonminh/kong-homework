package model

type UserProjectTab struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	UserID     string `gorm:"type:VARCHAR(64);"`
	ProjectID  string `gorm:"not null" json:"project_id"`
	CreateTime uint64 `gorm:"not null" json:"create_time"`
	UpdateTime uint64 `gorm:"not null" json:"update_time"`
}
