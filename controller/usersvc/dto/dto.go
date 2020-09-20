package dto

type User struct {
	ID         int64  `json:"id"`
	ExternalID string `json:"external_id"`
	Properties
}

type Properties struct {
	NickName       string `json:"nickname"`
	ProfileImage   string `json:"profile_image"`
	ThumbnailImage string `json:"thumbnail_image"`
}

type UpdateNickNameRequest struct {
	NickName string `form:"nickname"`
}
