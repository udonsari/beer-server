package dto

type Properties struct {
	NickName       string `json:"nickname"`
	ProfileImage   string `json:"profile_image"`
	ThumbnailImage string `json:"thumbnail_image"`
}

type GetUserRequest struct {
	AccessToken string `query:"access_token" validate:"required"`
}
