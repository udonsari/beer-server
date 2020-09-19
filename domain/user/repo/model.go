package repo

import "time"

type DBUser struct {
	ID             int64
	ExternalID     string
	NickName       string `gorm:"unique"`
	ProfileImage   string
	ThumbnailImage string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (DBUser) TableName() string {
	return "user_info"
}
